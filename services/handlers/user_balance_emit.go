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

type UserBalanceEmitHandler struct {
	sender                  *rmq.MessageSender
	emitUserBalanceConsumer *rmq.Consumer
	req                     *proto.EmmitBalanceByUserIdRequest
}

func NewUserBalanceEmitHandler(sender *rmq.MessageSender, emitUserBalanceConsumer *rmq.Consumer) *UserBalanceEmitHandler {
	return &UserBalanceEmitHandler{
		sender:                  sender,
		emitUserBalanceConsumer: emitUserBalanceConsumer,
		req:                     &proto.EmmitBalanceByUserIdRequest{},
	}
}

func (h *UserBalanceEmitHandler) Action(ctx *gin.Context) {
	log.Printf("getting req (UserBalanceEmit)...")

	err := ctx.BindJSON(h.req)
	utils.IsError(err, config.ErrBindJson)
	log.Printf(fmt.Sprintf("req was received\nrequest: %s", h.req))

	h.sender.SendEmmitUserBalanceRequest(h.req)
	respBalance := &proto.UserBalance{}

	respBytes := h.emitUserBalanceConsumer.GetMessageByCondition(h.condition, 60)

	err = gitProto.Unmarshal(respBytes, respBalance)
	if err != nil {
		ctx.IndentedJSON(http.StatusUnprocessableEntity, respBalance)
	} else {
		ctx.IndentedJSON(http.StatusOK, respBalance)
	}

	log.Printf(fmt.Sprintf("resp was publish\nresponse: %s", respBalance))
}
func (h *UserBalanceEmitHandler) condition(message []byte) bool {
	response := &proto.UserBalance{}
	err := gitProto.Unmarshal(message, response)
	utils.IsError(err, "failed unmarshal message")

	return h.req.UserId == response.UserId
}
