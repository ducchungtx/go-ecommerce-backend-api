package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	kafka "github.com/segmentio/kafka-go"
)

var (
	kafkaProducer *kafka.Writer
)

const (
	kafkaURL   = "localhost:19092"
	kafkaTopic = "user_topic_vip"
)

// For Producer
func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

// For Consumer
func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers, // []string{"localhost:9092", "localhost:9093"},
		GroupID:        groupID,
		Topic:          topic,
		MinBytes:       10e3,              // 10KB
		MaxBytes:       10e6,              // 10MB
		CommitInterval: time.Second,       // flushes commits to Kafka every second
		StartOffset:    kafka.FirstOffset, // kafka.LastOffset, kafka.FirstOffset, kafka.LastOffset, kafka.FirstOffset, kafka.FirstOffsetWithLeader
	})
}

type StockInfo struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

// Buy, Sell Stock
func newStock(message string, stockType string) *StockInfo {
	s := StockInfo{}
	s.Message = message
	s.Type = stockType

	return &s
}

func actionStock(c *gin.Context) {
	s := newStock(c.Query("msg"), c.Query("type"))
	body := make(map[string]interface{})
	body["action"] = "action"
	body["info"] = s
	jsonbody, _ := json.Marshal(body)
	// tao msg
	msg := kafka.Message{
		Key:   []byte("action"),
		Value: []byte(jsonbody),
	}
	// viet message by producer
	err := kafkaProducer.WriteMessages(context.Background(), msg)
	if err != nil {
		c.JSON(200, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"err": "",
		"msg": "action successfully",
	})
}

// Consumer hong mua ATC
func RegisterConsumerATC(id int) {
	// group consumer??
	kafkaGroupId := fmt.Sprintf("consumer-group-%d", id) //consumer-group-
	reader := getKafkaReader(kafkaURL, kafkaTopic, kafkaGroupId)
	defer reader.Close()

	fmt.Printf("Consumer(%d) Hong Phien ATC::\n", id)
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("Consumer (%d) error: %v\n", id, err)
		}
		fmt.Printf("Consumer (%d), hong topic:%v, partition:%v, offset:%v,time%d %s = %s\n", id, m.Topic, m.Partition, m.Offset, m.Time.Unix(), string(m.Key), string(m.Value))
	}
}

func main() {
	r := gin.Default()
	kafkaProducer = getKafkaWriter(kafkaURL, kafkaTopic)
	defer kafkaProducer.Close()
	r.POST("action/stock", actionStock)
	// dang ky 2 user de mau stock trong atc (1) (2)
	go RegisterConsumerATC(1)
	go RegisterConsumerATC(2)
	go RegisterConsumerATC(3)
	r.Run(":8999")
}
