package drivers

import (
	"context"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/nats-io/nats.go"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Postgres (GORM)
func ConnectPostgres(dsn string) (*gorm.DB, error) {
    return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

// Redis
func ConnectRedis(url string) (*redis.Client, error) {
    opt, err := redis.ParseURL(url)
    if err != nil {
        return nil, err
    }
    client := redis.NewClient(opt)
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()
    if err := client.Ping(ctx).Err(); err != nil {
        return nil, err
    }
    return client, nil
}

// NATS
func ConnectNATS(url string) (*nats.Conn, error) {
    return nats.Connect(url)
}

// MQTT
func ConnectMQTT(broker string, clientID string) (mqtt.Client, error) {
    opts := mqtt.NewClientOptions().AddBroker(broker).SetClientID(clientID)
    client := mqtt.NewClient(opts)
    token := client.Connect()
    if token.Wait() && token.Error() != nil {
        return nil, token.Error()
    }
    return client, nil
} 