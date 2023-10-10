package cache

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/go-redis/redis/v8"
	"github.com/smartystreets/goconvey/convey"
)

func TestRedis(t *testing.T) {
	convey.Convey("TestIsAdult", t, func() {
		f1 := gomonkey.ApplyFunc(DefaultRedisOption, func() redis.Options {
			return redis.Options{
				Addr:        fmt.Sprintf("%s:%d", "192.168.97.224", 6379),
				DB:          0,
				DialTimeout: 10 * time.Millisecond,
				Password:    "fdKMLz6LAPSDPRHF",
			}
		})
		defer f1.Reset()
		convey.Convey("empty", func() {
			ctx := context.Background()
			err := InitRedis(ctx)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestRedisSetNX(t *testing.T) {
	convey.Convey("TestIsAdult", t, func() {
		f1 := gomonkey.ApplyFunc(DefaultRedisOption, func() redis.Options {
			return redis.Options{
				Addr:        fmt.Sprintf("%s:%d", "192.168.97.224", 6379),
				DB:          0,
				DialTimeout: 10 * time.Millisecond,
				Password:    "fdKMLz6LAPSDPRHF",
			}
		})
		defer f1.Reset()
		convey.Convey("empty", func() {
			ctx := context.Background()
			err := InitRedis(ctx)
			convey.So(err, convey.ShouldBeNil)
			b, err := redisClient.SetNX(ctx, "testkey:1", "abcdef", time.Minute*30).Result()
			println("%v %s", b, err)
			b2, err2 := redisClient.SetNX(ctx, "testkey:1", "aaa", time.Minute*30).Result()
			println("%v %s", b2, err2)

		})
	})
}
