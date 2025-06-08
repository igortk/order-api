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

type RemoveOrderHandler struct {
	sender             *rmq.MessageSender
	orderEventConsumer *rmq.Consumer
	req                *proto.RemoveOrderRequest
}

func NewRemoveOrderHandler(sender *rmq.MessageSender, orderEventConsumer *rmq.Consumer) *RemoveOrderHandler {
	return &RemoveOrderHandler{
		sender:             sender,
		orderEventConsumer: orderEventConsumer,
		req:                &proto.RemoveOrderRequest{},
	}
}

func (h *RemoveOrderHandler) Action(ctx *gin.Context) {
	log.Printf("getting req...")

	err := ctx.BindJSON(h.req)
	utils.IsError(err, config.ErrBindJson)
	log.Printf(fmt.Sprintf("req was received\nrequest: %s", h.req))

	h.sender.Send(config.RabbitOrderExchange, config.RemoveOrderRequestRoutingKey, h.req)

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
func (h *RemoveOrderHandler) condition(message []byte) bool {
	event := &proto.OrderUpdateEvent{}
	err := gitProto.Unmarshal(message, event)
	utils.IsError(err, "failed unmarshal message")

	return h.req.OrderId == event.Order.OrderId && event.Order.Status == proto.OrderStatus_ORDER_STATUS_REMOVED
}
