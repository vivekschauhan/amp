package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/appcelerator/amp-kafka-pilot/pilot/api/admin"
	"google.golang.org/grpc"
)

// Kafka singleton
type Kafka struct {
	sarama sarama.Client
	admin  admin.AdminClient
}

// Connect to kafka
func (k *Kafka) Connect(host string) error {
	// Sarama client
	config := sarama.NewConfig()
	config.Version = sarama.V0_10_0_0
	var err error
	k.sarama, err = sarama.NewClient([]string{host}, config)
	if err != nil {
		return err
	}

	// gRPC administration client
	clientConn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		return err
	}
	k.admin = admin.NewAdminClient(clientConn)
	return nil
}

// NewConsumer create a new consumer
func (k *Kafka) NewConsumer() (sarama.Consumer, error) {
	return sarama.NewConsumerFromClient(k.sarama)
}

// Admin return the administration client
func (k *Kafka) Admin() admin.AdminClient {
	return k.admin
}
