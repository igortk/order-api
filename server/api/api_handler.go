package api

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
	"time"
)

type Handler struct {
	MessageSender           *rmq.MessageSender
	GetUserBalanceConsumer  *rmq.Consumer
	EmitUserBalanceConsumer *rmq.Consumer
	OrderEventConsumer      *rmq.Consumer
	GetUserOrdersConsumer   *rmq.Consumer
	GetExchangeRateConsumer *rmq.Consumer
}

func NewAPIHandler(messageSender *rmq.MessageSender, getUserBalanceConsumer, emitUserBalanceConsumer, orderEventConsumer, getUserOrdersConsumer, getExchangeRateConsumer *rmq.Consumer) *Handler {
	return &Handler{
		MessageSender:           messageSender,
		GetUserBalanceConsumer:  getUserBalanceConsumer,
		EmitUserBalanceConsumer: emitUserBalanceConsumer,
		OrderEventConsumer:      orderEventConsumer,
		GetUserOrdersConsumer:   getUserOrdersConsumer,
		GetExchangeRateConsumer: getExchangeRateConsumer,
	}
}

func (h *Handler) CreateOrder(ctx *gin.Context) {
	log.Printf("getting req...")

	req := &proto.CreateOrderRequest{}
	err := ctx.BindJSON(&req)
	utils.IsError(err, config.ErrBindJson)
	log.Printf(fmt.Sprintf("req was received\nrequest: %s", req))

	h.MessageSender.SendCreateOrderRequest(req)

	respBytes := h.OrderEventConsumer.GetMessageByCondition(func(message []byte) bool {
		event := &proto.OrderUpdateEvent{}
		err := gitProto.Unmarshal(message, event)
		utils.IsError(err, "failed unmarshal message")

		return req.OrderId == event.Order.OrderId
	}, 60)

	event := &proto.OrderUpdateEvent{}
	err = gitProto.Unmarshal(respBytes, event)

	if err != nil || event.Error != nil {
		ctx.IndentedJSON(http.StatusUnprocessableEntity, event)
	} else {
		ctx.IndentedJSON(http.StatusOK, event)
	}

	log.Printf(fmt.Sprintf("resp was publish\nresponse: %s", event))
}

func (h *Handler) RemoveOrder(ctx *gin.Context) {
	log.Printf("getting req...")

	req := &proto.RemoveOrderRequest{}
	err := ctx.BindJSON(&req)
	utils.IsError(err, config.ErrBindJson)
	log.Printf(fmt.Sprintf("req was received\nrequest: %s", req))

	h.MessageSender.SendRemoveOrderRequest(req)

	respBytes := h.OrderEventConsumer.GetMessageByCondition(func(message []byte) bool {
		event := &proto.OrderUpdateEvent{}
		err := gitProto.Unmarshal(message, event)
		utils.IsError(err, "failed unmarshal message")

		return req.OrderId == event.Order.OrderId && event.Order.Status == proto.OrderStatus_ORDER_STATUS_REMOVED
	}, 60)

	event := &proto.OrderUpdateEvent{}
	err = gitProto.Unmarshal(respBytes, event)

	if err != nil || event.Error != nil {
		ctx.IndentedJSON(http.StatusUnprocessableEntity, event)
	} else {
		ctx.IndentedJSON(http.StatusOK, event)
	}

	log.Printf(fmt.Sprintf("resp was publish\nresponse: %s", event))
}

func (h *Handler) UserBalanceEmit(ctx *gin.Context) {
	log.Printf("getting req (UserBalanceEmit)...")

	req := &proto.EmmitBalanceByUserIdRequest{}
	err := ctx.BindJSON(&req)
	utils.IsError(err, config.ErrBindJson)
	log.Printf(fmt.Sprintf("req was received\nrequest: %s", req))

	h.MessageSender.SendEmmitUserBalanceRequest(req)
	respBalance := &proto.UserBalance{}

	respBytes := h.EmitUserBalanceConsumer.GetMessageByCondition(func(message []byte) bool {
		response := &proto.UserBalance{}
		err := gitProto.Unmarshal(message, response)
		utils.IsError(err, "failed unmarshal message")

		return req.UserId == response.UserId
	}, 60)

	err = gitProto.Unmarshal(respBytes, respBalance)
	if err != nil {
		ctx.IndentedJSON(http.StatusUnprocessableEntity, respBalance)
	} else {
		ctx.IndentedJSON(http.StatusOK, respBalance)
	}

	log.Printf(fmt.Sprintf("resp was publish\nresponse: %s", respBalance))
}

func (h *Handler) UserBalanceInfo(ctx *gin.Context) {
	log.Printf("getting req (UserBalanceInfo)...")
	req := &proto.GetBalanceByUserIdRequest{
		Id:     uuid.New().String(),
		UserId: ctx.Param("id"),
	}

	h.MessageSender.SendGetUserBalanceRequest(req)
	respBalance := &proto.GetBalanceByUserIdResponse{}

	respBytes := h.GetUserBalanceConsumer.GetMessageByCondition(func(message []byte) bool {
		response := &proto.GetBalanceByUserIdResponse{}
		err := gitProto.Unmarshal(message, response)
		utils.IsError(err, "failed unmarshal message")

		return req.Id == response.Id
	}, 60)

	err := gitProto.Unmarshal(respBytes, respBalance)
	if err != nil {
		ctx.IndentedJSON(http.StatusUnprocessableEntity, respBalance)
	} else {
		ctx.IndentedJSON(http.StatusOK, respBalance)
	}
	log.Printf(fmt.Sprintf("resp was publish\nresponse: %s", respBalance))
}

func (h *Handler) ExchangeRate(ctx *gin.Context) {
	log.Printf("getting req...")

	req := &proto.GetExchangeRateRequest{}
	err := ctx.BindJSON(&req)
	utils.IsError(err, config.ErrBindJson)
	log.Printf(fmt.Sprintf("req was received\nrequest: %s", req))

	h.MessageSender.SendGetExchangeRateRequest(req)

	respBytes := h.GetExchangeRateConsumer.GetMessageByCondition(func(message []byte) bool {
		resp := &proto.GetExchangeRateResponse{}
		err := gitProto.Unmarshal(message, resp)
		utils.IsError(err, "failed unmarshal message")

		return req.Id == resp.Id
	}, 60)

	resp := &proto.GetExchangeRateResponse{}
	err = gitProto.Unmarshal(respBytes, resp)

	ctx.IndentedJSON(http.StatusOK, resp)
	log.Printf(fmt.Sprintf("resp was publish\nresponse: %s", resp))
}

func (h *Handler) UserOrders(ctx *gin.Context) {
	log.Printf("getting req...")

	req := &proto.GetUserOrdersRequest{}
	err := ctx.BindJSON(&req)
	utils.IsError(err, config.ErrBindJson)
	log.Printf(fmt.Sprintf("req was received\nrequest: %s", req))

	h.MessageSender.SendGetUserOrdersRequest(req)

	respBytes := h.GetUserOrdersConsumer.GetMessageByCondition(func(message []byte) bool {
		resp := &proto.GetUserOrdersResponse{}
		err := gitProto.Unmarshal(message, resp)
		utils.IsError(err, "failed unmarshal message")

		return req.Id == resp.Id
	}, 60)

	resp := &proto.GetUserOrdersResponse{}
	err = gitProto.Unmarshal(respBytes, resp)

	ctx.IndentedJSON(http.StatusOK, resp)
	log.Printf(fmt.Sprintf("resp was publish\nresponse: %s", resp))
}

func (h *Handler) test(respBalance *proto.GetBalanceByUserIdResponse, req *proto.GetBalanceByUserIdRequest) []byte {
	for respBalance.Id != req.Id {
		select {
		case body := <-h.GetUserBalanceConsumer.MessageChan:
			err := gitProto.Unmarshal(body, respBalance)
			utils.IsError(err, "failed unmarshal message")
			log.Printf("resp456464: %s", respBalance)
			return body
		case <-time.After(10 * time.Second):
			log.Error("response wasn't received")
			return nil
		}
	}
	return nil
}

func cond(message []byte, req *proto.GetBalanceByUserIdRequest) bool {
	response := &proto.GetBalanceByUserIdResponse{}
	err := gitProto.Unmarshal(message, response)
	utils.IsError(err, "failed unmarshal message")
	log.Printf(fmt.Sprintf("resp feom bs5\nresponse: %s", response))
	return req.Id != response.Id
}
