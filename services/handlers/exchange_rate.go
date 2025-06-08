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

type ExchangeRateHandler struct {
	sender                  *rmq.MessageSender
	getExchangeRateConsumer *rmq.Consumer
	req                     *proto.GetExchangeRateRequest
}

func NewExchangeRateHandler(sender *rmq.MessageSender, getExchangeRateConsumer *rmq.Consumer) *ExchangeRateHandler {
	return &ExchangeRateHandler{
		sender:                  sender,
		getExchangeRateConsumer: getExchangeRateConsumer,
		req:                     &proto.GetExchangeRateRequest{},
	}
}

func (h *ExchangeRateHandler) Action(ctx *gin.Context) {
	log.Printf("getting req...")

	err := ctx.BindJSON(h.req)
	utils.IsError(err, config.ErrBindJson)
	log.Printf(fmt.Sprintf("req was received\nreq Id: %s", h.req.Id))

	h.sender.Send(config.RabbitExchangeRateExchange, config.GetExchangeRateRequestRoutingKey, h.req)

	respBytes := h.getExchangeRateConsumer.GetMessageByCondition(h.condition, 60)

	resp := &proto.GetExchangeRateResponse{}
	err = gitProto.Unmarshal(respBytes, resp)

	ctx.IndentedJSON(http.StatusOK, resp)
	log.Printf(fmt.Sprintf("resp was publish\nresponse: %s", resp))
}
func (h *ExchangeRateHandler) condition(message []byte) bool {
	resp := &proto.GetExchangeRateResponse{}
	err := gitProto.Unmarshal(message, resp)
	utils.IsError(err, "failed unmarshal message")

	return h.req.Id == resp.Id
}
