package strategy

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type StrategyInterface interface{
	RateLimitStrategy( ctx context.Context, input *RateLimitInput) (*RateLimiterOutput, error)
}


type RateLimitInput struct {
	Key            string
	Limit    int64
	Duration time.Duration
}

type RateLimiterOutput struct {
	Result    bool
	Limit     int64
	Total     int64
	Remaining int64
	ExpiresAt time.Time
}

type StrategyRateLimit struct{
	Redis *redis.Client
}

func NewStrategyRateLimit(redis *redis.Client) *StrategyRateLimit{
	return &StrategyRateLimit{
redis,
	}
}

func (strl *StrategyRateLimit)RateLimitStrategy( ctx context.Context, input *RateLimitInput) (*RateLimiterOutput, error) {

	result, err := strl.getInfo(input.Key)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			fmt.Println("cria um novo")
			strl.create(input)
			return strl.getInfo(input.Key)
		}else{
			// qualquer erro
			return nil, err
		}
	}
	
	// como chave expira depois de um tempo, ele cria um novo. Entao, soh se preocupa com o limite de chamada no decrease. 
	// quando a chave expirar, quer dizer que usuario ficou 5s sem fazer request e pode fazer tudo de novo
	if result.Total <= 1 || !result.Result {
		// returns -1; if t is after u, it returns +1; if they're the same, it returns 0.
		// se pode tentar de novo, limpa e cria um novo devido ao horario
		expired := result.ExpiresAt.Compare(time.Now())
		if expired == -1 {
			fmt.Println("pode tentar de novo")
			strl.create(input)
		}else{
			fmt.Println("block")
			strl.update(input.Key, result, false)
		}
	}else{
		fmt.Println("decreasing")
		strl.update(input.Key, result, true)
	}
	return strl.getInfo(input.Key)
}

func(strl *StrategyRateLimit) getInfo(key string) (*RateLimiterOutput ,error){
	rtInfo := &RateLimiterOutput{}
	result, err := strl.Redis.Get(key).Result()
	if err != nil{
		return nil, err
	}

	json.Unmarshal([]byte(result), rtInfo)
	return rtInfo, nil
}

func(strl *StrategyRateLimit) create(input *RateLimitInput) {
		obj, _ := json.Marshal(RateLimiterOutput{
		Result: true,
		Limit: input.Limit,
		Total: input.Limit,
		Remaining: input.Duration.Milliseconds(),
		ExpiresAt:  time.Now().Add(input.Duration), 
	})
	strl.Redis.Set(input.Key, obj,  input.Duration)
}

func(strl *StrategyRateLimit) update(key string, input *RateLimiterOutput, result bool) {
		obj, _ := json.Marshal(RateLimiterOutput{
		Result: result,
		Limit: input.Limit,
		Total: input.Total -1,
		Remaining: input.Remaining,
		ExpiresAt:  input.ExpiresAt, 
	})

	exp := time.Duration(input.Remaining) *time.Millisecond
	strl.Redis.Set(key, obj,  exp)
}