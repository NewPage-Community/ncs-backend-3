package redis

import (
	"backend/pkg/log"
	"context"
	"crypto/rand"
	"fmt"
	"sync"
	"time"

	goredis "github.com/go-redis/redis/v8"
)

const (
	maxLen            = 100
	idLen             = 6
	comsumerSpawnTime = 1 * time.Second
	retryTime         = 5 * time.Second
	groupExistErr     = "BUSYGROUP Consumer Group name already exists"
	canceledErr       = "context canceled"
)

type Stream struct {
	client   *goredis.Client
	group    string
	comsumer string
	closer   func()
	wg       sync.WaitGroup
}

type StreamCallback func(context.Context, string)

func NewStream(client *goredis.Client, service string) *Stream {
	return &Stream{
		client:   client,
		group:    service,
		comsumer: "",
		wg:       sync.WaitGroup{},
	}
}

func (s *Stream) Publish(ctx context.Context, topic, msg string) (err error) {
	return s.client.XAdd(ctx, &goredis.XAddArgs{
		Stream: topic,
		MaxLen: maxLen,
		Values: []string{"msg", msg},
	}).Err()
}

func (s *Stream) Subscribe(topic string, callback StreamCallback) (err error) {
	ctx := context.Background()

	// Create group
	err = s.client.XGroupCreateMkStream(ctx, topic, s.group, "$").Err()
	if err != nil {
		if err.Error() != groupExistErr {
			return
		}
	}

	// Create comsumer
	if len(s.comsumer) == 0 {
		for i := 0; i < idLen; i++ {
			s.comsumer = getRandomID()
			res, err := s.client.XGroupCreateConsumer(ctx, topic, s.group, s.comsumer).Result()

			// Create successfully
			if res == 1 {
				break
			}

			if err != nil {
				log.Error(err)
			} else {
				// Not things wrong, just keep go on!
				i--
			}
			time.Sleep(comsumerSpawnTime)
		}
	}

	// Loop message from topic
	// Add wait group for waiting block cancel and delete consumer
	s.wg.Add(1)
	go func() {
		for {
			ctx := context.Background()

			// Read message
			ctx, s.closer = context.WithCancel(ctx)
			res, err := s.client.XReadGroup(ctx, &goredis.XReadGroupArgs{
				Group:    s.group,
				Consumer: s.comsumer,
				Streams:  []string{topic, ">"},
				Count:    1,
				Block:    0,
				NoAck:    true,
			}).Result()

			if err != nil {
				// Go to cancel
				if err.Error() == canceledErr {
					_ = s.client.XGroupDelConsumer(context.Background(), topic, s.group, s.comsumer)
					s.wg.Done()
					return
				}
				log.Error("redis XReadGroup:", err)
				time.Sleep(retryTime)
			} else {
				for _, xs := range res {
					for _, ms := range xs.Messages {
						for _, m := range ms.Values {
							v, _ := m.(string)
							s.wg.Add(1)
							go func() {
								callback(ctx, v)
								s.wg.Done()
							}()
						}
					}
				}
			}
		}
	}()

	return
}

func (s *Stream) Close() {
	s.closer()
	s.wg.Wait()
}

func getRandomID() string {
	randBytes := make([]byte, idLen/2)
	_, _ = rand.Read(randBytes)
	return fmt.Sprintf("%x", randBytes)
}
