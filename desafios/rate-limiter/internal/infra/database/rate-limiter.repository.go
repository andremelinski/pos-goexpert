package database

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
)

type RateLimitInput struct {
	Key            string
	Limit    int64
	Duration time.Duration
}

type RateLimitOutput struct {
	Result    bool
	Limit     int64
	Total     int64
	Remaining int64
	ExpiresAt time.Time
}

type RateLimitRepositoryInterface interface{
	Get(key string) (*RateLimitOutput ,error)
	Create(input *RateLimitInput)
	Update(key string, input *RateLimitOutput, result bool)
}

type RateLimitRepository struct{
	Redis *redis.Client
}

func NewRateLimitRepository(redis *redis.Client) *RateLimitRepository{
	return &RateLimitRepository{
redis,
	}
}

func(strl *RateLimitRepository) Get(key string) (*RateLimitOutput ,error){
	rtInfo := &RateLimitOutput{}
	result, err := strl.Redis.Get(key).Result()
	if err != nil{
		return nil, err
	}

	json.Unmarshal([]byte(result), rtInfo)
	return rtInfo, nil
}

func(strl *RateLimitRepository) Create(input *RateLimitInput) {
		obj, _ := json.Marshal(&RateLimitOutput{
		Result: true,
		Limit: input.Limit,
		Total: input.Limit,
		Remaining: input.Duration.Milliseconds(),
		ExpiresAt:  time.Now().Add(input.Duration), 
	})
	strl.Redis.Set(input.Key, obj,  input.Duration)
}

func(strl *RateLimitRepository) Update(key string, input *RateLimitOutput, result bool) {
		obj, _ := json.Marshal(RateLimitOutput{
		Result: result,
		Limit: input.Limit,
		Total: input.Total -1,
		Remaining: input.Remaining,
		ExpiresAt:  input.ExpiresAt, 
	})

	exp := time.Duration(input.Remaining) *time.Millisecond
	strl.Redis.Set(key, obj,  exp)
}