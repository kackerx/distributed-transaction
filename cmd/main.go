package main

import (
	"context"

	"distributed-transaction/mq"
)

type OrderSubmitMsg struct {
	Code string
}

func main() {
	ctx := context.Background()

	err := mq.GetRocketProducer().SendMsgWithTx(ctx, "test", OrderSubmitMsg{Code: "110"})
	if err != nil {
		panic(err)
	}
}
