package banners

import (
	"fmt"
	"time"

	model "github.com/M0rdovorot/effective_mobile/internal/model"
	"github.com/gomodule/redigo/redis"
)

const (
	maxIdles   = 10
	timeOutSec = 240
	TTL = 300
)

type BannerRedisStorage struct {
	// redisConn redis.Conn
	redisPool *redis.Pool
}

func CreateBannerRedisStorage(pool *redis.Pool) CashRepository {
	return &BannerRedisStorage{
		redisPool: pool,
	}
}

func NewPool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     maxIdles,
		IdleTimeout: timeOutSec * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", addr) },
	}
}

func (storage *BannerRedisStorage) GetUserBanner(tagId int, featureId int) (string, bool, error) {
	conn := storage.redisPool.Get()
	defer conn.Close()

	content, err := redis.Bytes(conn.Do("GET", fmt.Sprintf("%d:%d:content", featureId, tagId)))
	if err != nil {
		return "", false, err
	}

	isActive, err := redis.Bool(conn.Do("GET", fmt.Sprintf("%d:%d:isActive", featureId, tagId)))
	if err != nil {
		return "", false, err
	}

	return string(content), isActive, nil
}

func (storage *BannerRedisStorage) CreateBanner(banner model.Banner, tagId int) (error) {
	conn := storage.redisPool.Get()
	defer conn.Close()

	_, err := redis.String(conn.Do("SET", fmt.Sprintf("%d:%d:content", banner.FeatureId, tagId), banner.JSONContent, "EX", TTL))
	if err != nil {
		return err
	}
	_, err = redis.String(conn.Do("SET", fmt.Sprintf("%d:%d:isActive", banner.FeatureId, tagId), banner.IsActive, "EX", TTL))
	if err != nil {
		return err
	}

	return nil
}
