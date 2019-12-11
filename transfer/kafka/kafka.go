package kafka

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Kafka struct {
	opts     Options
	consumer kafka.Consumer
}

// New file config
func New(opts ...Option) *Kafka {
	options := NewOptions(opts...)
	return &Kafka{
		opts: options,
	}
}

// Init init
func (k *Kafka) Init(opts ...Option) {
	for _, o := range opts {
		o(&k.opts)
	}
}

// Connect connect
func (k *Kafka) Connect() error {
	var err error
	k.consumer, err = kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":    strings.Join(k.opts.Brokers, ","),
		"group.id":             k.opts.Group,
		"session.timeout.ms":   6000,
		"default.topic.config": kafka.ConfigMap{"auto.offset.reset": k.opts.Offset},
	})
	return err
}

func (k *Kafka) Subscribe() error {

	if err := k.consumer.SubscribeTopics(topics, nil); err != nil {
		return err
	}

	defer k.consumer.Close()
	for run == true {
			ev := c.Poll(100)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				fmt.Printf("=> %s: %s\n",
					e.TopicPartition, string(e.Key))
			case kafka.PartitionEOF:
				fmt.Printf("%% Reached %v\n", e)
			case kafka.Error:
				fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
				run = false
			default:
				fmt.Printf("Ignored %v\n", e)
			}
	}
	return nil
}
