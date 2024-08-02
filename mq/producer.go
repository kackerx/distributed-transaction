package mq

import (
	"context"
	"fmt"
	"sync"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	jsoniter "github.com/json-iterator/go"
)

type Producer interface {
	SendMsgWithTx(ctx context.Context, topic string, msg any)
}

type rocketProducer struct {
	p rocketmq.TransactionProducer

	once sync.Once
}

func (r rocketProducer) SendMsgWithTx(ctx context.Context, topic string, msg any) (err error) {
	if err = r.p.Start(); err != nil {
		return
	}

	msgData, err := jsoniter.Marshal(msg)
	if err != nil {
		return err
	}

	resp, err := r.p.SendMessageInTransaction(ctx, primitive.NewMessage(topic, msgData))
	if err != nil {
		return
	}

	fmt.Printf("send msg success: %s\n", resp.String())

	if err = r.p.Shutdown(); err != nil {
		return
	}

	dp, err := producer.NewDefaultProducer(producer.WithQueueSelector(producer.NewManualQueueSelector()))
	if err != nil {
		return
	}

	dp.SendAsync(ctx, func(ctx context.Context, result *primitive.SendResult, err error) {

	})

	dp.SendSync(ctx, primitive.NewMessage())

	return
}

var rp = &rocketProducer{}

func GetRocketProducer() *rocketProducer {
	sync.OnceFunc(func() {
		rp.p, _ = rocketmq.NewTransactionProducer(
			NewOrderListener(),
			producer.WithNsResolver(primitive.NewPassthroughResolver([]string{"127.0.0.1:9876"})),
			producer.WithRetry(1),
		)
	})()

	return rp
}

// OrderListener 事务监听回调
type OrderListener struct {
}

func NewOrderListener() *OrderListener {
	return &OrderListener{}
}

func (o OrderListener) ExecuteLocalTransaction(message *primitive.Message) primitive.LocalTransactionState {
	return primitive.CommitMessageState
}

func (o OrderListener) CheckLocalTransaction(ext *primitive.MessageExt) primitive.LocalTransactionState {
	return primitive.CommitMessageState
}
