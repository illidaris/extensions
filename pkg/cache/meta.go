package cache

import "fmt"

func BuildRedisKey(b BaseCacheKey, name, key string) string {
	return fmt.Sprintf("%s:%s:%s", b, name, key)
}
