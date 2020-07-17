package kafka

import (
	"testing"

	"github.com/Shopify/sarama"
)

func Test_LazyClientWithNoConfig(t *testing.T) {
	c := &LazyClient{}
	_, err := c.ListACLs()

	if err == nil {
		t.Fatalf("exepted err, got %v", err)
	}
}

func Test_LazyClientErrors(t *testing.T) {
	c := &LazyClient{}
	_, err := c.ListACLs()
	if err == nil {
		t.Fatalf("exepted err, got %v", err)
	}
	_, err = c.ListACLs()
	if err == nil {
		t.Fatalf("exepted err, got %v", err)
	}
}

func Test_LazyClientWithConfigErors(t *testing.T) {
	config := &Config{
		BootstrapServers: &[]string{"localhost:9000"},
		Timeout:          10,
	}
	c := &LazyClient{
		Config: config,
	}
	err := c.init()

	if err != sarama.ErrOutOfBrokers {
		t.Fatalf("exepted err, got %v", err)
	}
}
