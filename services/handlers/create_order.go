package handlers

import (
	"github.com/gin-gonic/gin"
	gitProto "github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
	"net/http"
	"order-api/config"
	"order-api/dto/proto"
	"order-api/services/rmq"
	"order-api/utils"
)

type Event = proto.OrderUpdateEvent

type Handler interface {
	Action(ctx *gin.Context)
}

type CreateOrderHandler struct {
	sender             *rmq.MessageSender
	orderEventConsumer *rmq.Consumer
	req                *proto.CreateOrderRequest
}

func NewCreateOrderHandler(sender *rmq.MessageSender, orderEventConsumer *rmq.Consumer) *CreateOrderHandler {
	return &CreateOrderHandler{
		sender:             sender,
		orderEventConsumer: orderEventConsumer,
		req:                &proto.CreateOrderRequest{},
	}
}

func (h *CreateOrderHandler) Action(ctx *gin.Context) {
	log.Printf("getting req...")

	err := ctx.BindJSON(h.req)
	utils.IsError(err, config.ErrBindJson)
	log.Printf("req was received request: %s", h.req)

	h.sender.Send(config.RabbitOrderExchange, config.CreateOrderRequestRoutingKey, h.req)

	event, ok := h.getUpdateEvent()

	if ok || event.Error == nil {
		ctx.IndentedJSON(http.StatusOK, event)
	} else {
		ctx.IndentedJSON(http.StatusUnprocessableEntity, event)
	}

	log.Printf("resp was publish resp/req id: %s", event.Id)
}

func (h *CreateOrderHandler) getUpdateEvent() (*Event, bool) {
	event := &Event{Id: h.req.Id}
	evErr := &proto.Error{
		Code:    422,
		Message: "timeout updated order",
	}

	respBytes := h.orderEventConsumer.GetMessageByCondition(h.condition, 4)
	if respBytes == nil {
		event.Error = evErr
		return event, false
	}

	err := gitProto.Unmarshal(respBytes, event)
	if err != nil {
		evErr.Message = "problem parse event"
	}

	return event, true
}

func (h *CreateOrderHandler) condition(message []byte) bool {
	event := &Event{}
	err := gitProto.Unmarshal(message, event)
	utils.IsError(err, "failed unmarshal message")

	return h.req.OrderId == event.Order.OrderId
}
