package subscriber
import (
	"cloud.google.com/go/pubsub"
	"fmt"
	"golang.org/x/net/context"
	"log"
	"os"
	"sync"
	"time"
)
type Subscriber struct {
	SubscriptionName string
	TopicName        string
	ProjectId        string
	mux              sync.Mutex
}
func (sub *Subscriber) PullMessages(c chan string) {
	log.Println("Go subscriber is listening to events")
	os.Setenv("PUBSUB_EMULATOR_HOST", "localhost:8085")
	ctx := context.Background()
	pubsubClient, err := pubsub.NewClient(ctx, sub.ProjectId)
	if err != nil {
		log.Fatalf("Could not create pubsub Client: %v", err)
	}
	topic := createTopicIfNotExists(pubsubClient, sub.TopicName)
	subscription := createSubscriptionIfNotExists(pubsubClient, topic, sub.SubscriptionName)
	err = subscription.Receive(context.Background(), func(ctx context.Context, msg *pubsub.Message) {
		fmt.Printf("Got message: %q\n", string(msg.Data))
		c <- string(msg.Data) // Sending message to channel
		msg.Ack()		
	})
	if err != nil {
		log.Fatalf("Error during Receive: %v", err)
	}
}

func createTopicIfNotExists(client *pubsub.Client, topicName string) *pubsub.Topic {
	ctx := context.Background()
	t := client.Topic(topicName)
	ok, err := t.Exists(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if ok {
		return t
	}
	// creating
	t, err = client.CreateTopic(ctx, topicName)
	if err != nil {
		log.Fatalf("Failed to create the topic: %v", err)
	}
	return t
}

func createSubscriptionIfNotExists(client *pubsub.Client, topic *pubsub.Topic, name string) *pubsub.Subscription {
	ctx := context.Background()
	sub := client.Subscription(name)
	ok, err := sub.Exists(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if ok {
		return sub
	}
	// creating
	sub, err = client.CreateSubscription(ctx, name, pubsub.SubscriptionConfig{
		Topic:       topic,
		AckDeadline: 10 * time.Second,
	})
	return sub
}