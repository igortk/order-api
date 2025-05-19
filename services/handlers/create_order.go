package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	gitProto "github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
	"net/http"
	"order-api/config"
	"order-api/dto/proto"
	"order-api/services/rmq"
	"order-api/utils"
)

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
	log.Printf(fmt.Sprintf("req was received\nrequest: %s", h.req))

	h.sender.SendCreateOrderRequest(h.req)

	respBytes := h.orderEventConsumer.GetMessageByCondition(h.condition, 60)

	event := &proto.OrderUpdateEvent{}
	err = gitProto.Unmarshal(respBytes, event)

	if err != nil || event.Error != nil {
		ctx.IndentedJSON(http.StatusUnprocessableEntity, event)
	} else {
		ctx.IndentedJSON(http.StatusOK, event)
	}

	log.Printf(fmt.Sprintf("resp was publish\nresponse: %s", event))
}

func (h *CreateOrderHandler) condition(message []byte) bool {
	event := &proto.OrderUpdateEvent{}
	err := gitProto.Unmarshal(message, event)
	utils.IsError(err, "failed unmarshal message")

	return h.req.OrderId == event.Order.OrderId
}
