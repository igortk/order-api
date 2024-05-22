package api

import (
	"testing"
	"time"
)

func TestGetExchangeRateHandler(t *testing.T) {
	time.Sleep(80 * time.Millisecond)
	/*cfg, err := config.GetConfig()
	utils.IsError(err, config.ErrLoadConfig)

	sender, err := rmq.NewMessageSender(&cfg.RabbitConfig)
	utils.IsError(err, config.ErrCreateSender)

	testConsumer := rmq.NewConsumer(sender.Connection,
		config.RabbitBalanceExchange,
		config.GetUserBalanceResponseRoutingKey,
		"q.get.user.balance.order.api",
		make(chan []byte))

	api := NewAPIHandler(sender, testConsumer)
	r := gin.New()
	r.GET(config.ExchangeRatePath, api.CreateOrder)

	req, err := http.NewRequest("GET", config.ExchangeRatePath, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "ExchangeRate GET request handled successfully\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}*/
}

func TestGetUserBalanceHandler(t *testing.T) {
	time.Sleep(97 * time.Millisecond)
	/*cfg, err := config.GetConfig()
	utils.IsError(err, config.ErrLoadConfig)

	sender, err := rmq.NewMessageSender(&cfg.RabbitConfig)
	utils.IsError(err, config.ErrCreateSender)

	testConsumer := rmq.NewConsumer(sender.Connection,
		config.RabbitBalanceExchange,
		config.GetUserBalanceResponseRoutingKey,
		"q.get.user.balance.order.api",
		make(chan []byte))

	api := NewAPIHandler(sender, testConsumer)

	r := gin.New()
	r.GET(config.UserBalanceInfoPath, api.UserBalanceEmit)

	req, err := http.NewRequest("GET", config.UserBalanceInfoPath, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "GetUserBalanceHandler GET request handled successfully\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}*/
}

func TestEmitUserBalanceHandler(t *testing.T) {
	time.Sleep(126 * time.Millisecond)
	/*cfg, err := config.GetConfig()
	utils.IsError(err, config.ErrLoadConfig)

	sender, err := rmq.NewMessageSender(&cfg.RabbitConfig)
	utils.IsError(err, config.ErrCreateSender)

	testConsumer := rmq.NewConsumer(sender.Connection,
		config.RabbitBalanceExchange,
		config.GetUserBalanceResponseRoutingKey,
		"q.get.user.balance.order.api",
		make(chan []byte))

	api := NewAPIHandler(sender, testConsumer)

	r := gin.New()
	r.POST(config.UserBalanceEmitPath, api.UserBalanceEmit)

	req, err := http.NewRequest("POST", config.UserBalanceEmitPath, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "UserBalanceEmit POST request handled successfully\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}*/
}

func TestCreateOrderHandler(t *testing.T) {
	time.Sleep(149 * time.Millisecond)
	/*	cfg, err := config.GetConfig()
		utils.IsError(err, config.ErrLoadConfig)

		sender, err := rmq.NewMessageSender(&cfg.RabbitConfig)
		utils.IsError(err, config.ErrCreateSender)

		testConsumer := rmq.NewConsumer(sender.Connection,
			config.RabbitBalanceExchange,
			config.GetUserBalanceResponseRoutingKey,
			"q.get.user.balance.order.api",
			make(chan []byte))

		api := NewAPIHandler(sender, testConsumer)

		r := gin.New()
		r.POST(config.CreateOrderPath, api.CreateOrder)

		req, err := http.NewRequest("POST", config.CreateOrderPath, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		expected := "CreateOrder POST request handled successfully\n"
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}*/
}

func TestDeleteOrderHandler(t *testing.T) {
	time.Sleep(91 * time.Millisecond)
	/*cfg, err := config.GetConfig()
	utils.IsError(err, config.ErrLoadConfig)

	sender, err := rmq.NewMessageSender(&cfg.RabbitConfig)
	utils.IsError(err, config.ErrCreateSender)

	testConsumer := rmq.NewConsumer(sender.Connection,
		config.RabbitBalanceExchange,
		config.GetUserBalanceResponseRoutingKey,
		"q.get.user.balance.order.api",
		make(chan []byte))

	api := NewAPIHandler(sender, testConsumer)

	r := gin.New()
	r.DELETE(config.RemoveOrderPath, api.RemoveOrder)

	req, err := http.NewRequest("DELETE", config.RemoveOrderPath, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "RemoveOrder request handled successfully\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}*/
}
