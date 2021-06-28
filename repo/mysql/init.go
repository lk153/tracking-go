package mysql

import (
	kafkaLib "github.com/lk153/go-lib/kafka"
)

var (
	configPath  *string
	producerLib *kafkaLib.KafkaProducer
)

const (
	PRODUCT_KAFKA_TOPIC = "testing_153_product"
	TASK_KAFKA_TOPIC    = "testing_153_task"
)
