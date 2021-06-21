package server

import (
	"context"
	"encoding/json"
	"fmt"

	"factory/exam/repo"
	"factory/exam/services"

	kafkaLib "github.com/lk153/go-lib/kafka"
)

const (
	BOOTSTRAP_SERVERS = "bootstrap.servers"
	SASL_MECHANISMS   = "sasl.mechanisms"
	SECURITY_PROTOCOL = "security.protocol"
	SASL_USERNAME     = "sasl.username"
	SASL_PASSWORD     = "sasl.password"
)

//KafkaConsumer ...
type KafkaConsumer struct {
	cache services.ProductServiceInterface
	repo  repo.ProductRepoInterface
}

//NewKafkaConsumer ...
func NewKafkaConsumer(
	cache services.ProductServiceInterface,
	repo repo.ProductRepoInterface,
) (*KafkaConsumer, error) {
	return &KafkaConsumer{
		cache: cache,
		repo:  repo,
	}, nil
}

//Close ...
func (kc *KafkaConsumer) Close() error {
	return nil
}

//Start ...
func (kc *KafkaConsumer) Start() error {
	consumerOutput := make(chan []byte, 1)
	go func() {
		kafkaLib.Start(consumerOutput)
	}()

	go func() {
		for item := range consumerOutput {
			fmt.Printf("Consumer Set Cache: %v\n", string(item))
			productModel := &repo.ProductModel{}
			err := json.Unmarshal([]byte(item), productModel)
			if err != nil {
				fmt.Printf("Unmarshal consumed product has error: %v", err)
			}

			kc.cache.GetProduct(context.Background(), int(productModel.ID))
		}
	}()

	return nil
}
