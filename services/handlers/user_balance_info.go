package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	gitProto "github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"net/http"
	"order-api/config"
	"order-api/dto/proto"
	"order-api/services/rmq"
	"order-api/utils"
)

type UserBalanceInfoHandler struct {
	sender                 *rmq.MessageSender
	getUserBalanceConsumer *rmq.Consumer
	req                    *proto.GetBalanceByUserIdRequest
}

func NewUserBalanceInfoHandler(sender *rmq.MessageSender, getUserBalanceConsumer *rmq.Consumer) *UserBalanceInfoHandler {
	return &UserBalanceInfoHandler{
		sender:                 sender,
		getUserBalanceConsumer: getUserBalanceConsumer,
		req:                    &proto.GetBalanceByUserIdRequest{},
	}
}

func (h *UserBalanceInfoHandler) Action(ctx *gin.Context) {
	log.Printf("getting req (UserBalanceInfo)...")
	h.req = &proto.GetBalanceByUserIdRequest{
		Id:     uuid.New().String(),
		UserId: ctx.Param("id"),
	}

	h.sender.Send(config.RabbitBalanceExchange, config.GetUserBalanceRequestRoutingKey, h.req)
	respBalance := &proto.GetBalanceByUserIdResponse{}

	respBytes := h.getUserBalanceConsumer.GetMessageByCondition(h.condition, 60)

	err := gitProto.Unmarshal(respBytes, respBalance)
	if err != nil {
		ctx.IndentedJSON(http.StatusUnprocessableEntity, respBalance)
	} else {
		ctx.IndentedJSON(http.StatusOK, respBalance)
	}
	log.Printf(fmt.Sprintf("resp was publish\nresponse: %s", respBalance))
}
func (h *UserBalanceInfoHandler) condition(message []byte) bool {
	response := &proto.GetBalanceByUserIdResponse{}
	err := gitProto.Unmarshal(message, response)
	utils.IsError(err, "failed unmarshal message")

	return h.req.Id == response.Id
}
