package rmq

import (
	"github.com/streadway/amqp"
	"sync"
	"time"
)

type MessageRouter struct {
	incomingChan chan amqp.Delivery
	subscribers  []*routerSubscriber
	mu           sync.Mutex
}

type routerSubscriber struct {
	cond Condition
	ch   chan []byte
	once bool
}

func NewMessageRouting(incoming chan amqp.Delivery) *MessageRouter {
	router := &MessageRouter{
		incomingChan: incoming,
	}
	go router.start()
	return router
}

func (mc *MessageRouter) start() {
	for msg := range mc.incomingChan {
		mc.mu.Lock()
		for i := 0; i < len(mc.subscribers); i++ {
			sub := mc.subscribers[i]
			if sub.cond(msg.Body) {
				sub.ch <- msg.Body
				if sub.once {
					mc.subscribers = append(mc.subscribers[:i], mc.subscribers[i+1:]...)
					i--
				}
			}
		}
		mc.mu.Unlock()
	}
}

func (mc *MessageRouter) Subscribe(cond Condition, once bool, timeout time.Duration) []byte {
	ch := make(chan []byte, 1)
	sub := &routerSubscriber{
		cond: cond,
		ch:   ch,
		once: once,
	}

	mc.mu.Lock()
	mc.subscribers = append(mc.subscribers, sub)
	mc.mu.Unlock()

	select {
	case msg := <-sub.ch:
		return msg
	case <-time.After(timeout * time.Second):
		return nil
	}
}

type MessageContainer struct {
	messages map[any][]byte
}
