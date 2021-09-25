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
	idLen             = 5
	comsumerSpawnTime = 5 * time.Second
	retryTime         = 5 * time.Second
	groupExistErr     = "BUSYGROUP Consumer Group name already exists"
)

type Stream struct {
	client     *goredis.Client
	comsumer   string
	close      bool
	callbackWG sync.WaitGroup
}

func NewStream(client *goredis.Client) *Stream {
	return &Stream{
		client:     client,
		comsumer:   "",
		close:      false,
		callbackWG: sync.WaitGroup{},
	}
}

func (s *Stream) Publish(ctx context.Context, topic, msg string) (err error) {
	return s.client.XAdd(ctx, &goredis.XAddArgs{
		Stream: topic,
		MaxLen: maxLen,
		Values: []string{"msg", msg},
	}).Err()
}

func (s *Stream) Subscribe(service, topic string, callback func(context.Context, string, error)) (err error) {
	ctx := context.Background()

	// Create group
	err = s.client.XGroupCreateMkStream(ctx, topic, service, "$").Err()
	if err != nil {
		if err.Error() != groupExistErr {
			return
		}
	}

	// Create group
	if len(s.comsumer) == 0 {
		for {
			s.comsumer = getRandomID()
			if s.client.XGroupCreateConsumer(ctx, topic, service, s.comsumer).Val() == 1 {
				break
			}
			time.Sleep(comsumerSpawnTime)
		}
	}

	// Loop message from topic
	go func() {
		for {
			// Check stream close
			if s.close {
				return
			}

			// Read message
			ctx := context.Background()
			res, err := s.client.XReadGroup(ctx, &goredis.XReadGroupArgs{
				Group:    topic,
				Consumer: s.comsumer,
				Streams:  []string{topic},
				Count:    1,
				Block:    0,
				NoAck:    false,
			}).Result()

			if err != nil {
				log.Error("redis XReadGroup:", err)
				time.Sleep(retryTime)
			} else {
				for _, xs := range res {
					for _, ms := range xs.Messages {
						for _, m := range ms.Values {
							v, _ := m.(string)
							s.callbackWG.Add(1)
							go func() {
								callback(ctx, v, nil)
								s.callbackWG.Done()
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
	s.close = true
	s.callbackWG.Wait()
}

func getRandomID() string {
	randBytes := make([]byte, idLen/2)
	_, _ = rand.Read(randBytes)
	return fmt.Sprintf("%x", randBytes)
}
