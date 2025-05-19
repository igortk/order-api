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

type UserOrdersHandler struct {
	sender                *rmq.MessageSender
	getUserOrdersConsumer *rmq.Consumer
	req                   *proto.GetUserOrdersRequest
}

func NewUserOrdersHandler(sender *rmq.MessageSender, getUserOrdersConsumer *rmq.Consumer) *UserOrdersHandler {
	return &UserOrdersHandler{
		sender:                sender,
		getUserOrdersConsumer: getUserOrdersConsumer,
		req:                   &proto.GetUserOrdersRequest{},
	}
}

func (h *UserOrdersHandler) Action(ctx *gin.Context) {
	log.Printf("getting req...")

	err := ctx.BindJSON(h.req)
	utils.IsError(err, config.ErrBindJson)
	log.Printf(fmt.Sprintf("req was received\nrequest: %s", h.req))

	h.sender.SendGetUserOrdersRequest(h.req)

	respBytes := h.getUserOrdersConsumer.GetMessageByCondition(h.condition, 60)

	resp := &proto.GetUserOrdersResponse{}
	err = gitProto.Unmarshal(respBytes, resp)

	ctx.IndentedJSON(http.StatusOK, resp)
	log.Printf(fmt.Sprintf("resp was publish\nresponse: %s", resp))
}
func (h *UserOrdersHandler) condition(message []byte) bool {
	resp := &proto.GetUserOrdersResponse{}
	err := gitProto.Unmarshal(message, resp)
	utils.IsError(err, "failed unmarshal message")

	return h.req.Id == resp.Id
}
