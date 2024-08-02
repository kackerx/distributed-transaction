package mq

import (
	"context"
	"testing"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

type Consumer interface {
	SubTopic() string
	Handle(ctx context.Context, msg []byte) (err error)
}

type rocketConsumer struct {
	c rocketmq.PushConsumer
}

func newRocketConsumer(c rocketmq.PushConsumer) *rocketConsumer {
	return &rocketConsumer{c: c}
}

func (r *rocketConsumer) SubTopic() string {
	return "test"
}

func (r *rocketConsumer) Handle(ctx context.Context, msg []byte) (err error) {
	r.c.Subscribe(r.SubTopic(), consumer.MessageSelector{}, func(ctx context.Context, msgs ...*primitive.MessageExt) (res consumer.ConsumeResult, err error) {
		return consumer.ConsumeSuccess, nil
	})

	return
}

func TestName(t *testing.T) {
	c, err := consumer.NewPushConsumer(
		consumer.WithGroupName(""),
		consumer.WithNameServer([]string{""}),
	)
	if err != nil {
		return
	}

	

}
