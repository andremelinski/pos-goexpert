package strategy

import (
	"errors"
	"fmt"
	"time"

	"github.com/andremelinski/pos-goexpert/desafios/rate-limiter/internal/infra/database"
	"github.com/go-redis/redis"
)

type StrategyInterface interface{
	RateLimitStrategy(input *database.RateLimitInput) (*database.RateLimitOutput, error)
}

type StrategyRateLimit struct{
	rateLimitRepository database.RateLimitRepositoryInterface
}

func NewStrategyRateLimit(rateLimitRepository database.RateLimitRepositoryInterface) *StrategyRateLimit{
	return &StrategyRateLimit{
rateLimitRepository,
	}
}

func (strl *StrategyRateLimit)RateLimitStrategy(input *database.RateLimitInput) (*database.RateLimitOutput, error) {
	result, err := strl.rateLimitRepository.Get(input.Key)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			fmt.Println("cria um novo")
			strl.rateLimitRepository.Create(input)
			return strl.rateLimitRepository.Get(input.Key)
		}else{
			// qualquer erro
			return nil, err
		}
	}
	
	// como chave expira depois de um tempo, ele cria um novo. Entao, soh se preocupa com o limite de chamada no decrease. 
	// quando a chave expirar, quer dizer que usuario ficou 5s sem fazer request e pode fazer tudo de novo
	if result.Total <= 1 {
		// returns -1; if t is after u, it returns +1; if they're the same, it returns 0.
		// se pode tentar de novo, limpa e cria um novo devido ao horario
		expired := result.ExpiresAt.Compare(time.Now())
		if expired == -1 {
			fmt.Println("pode tentar de novo")
			strl.rateLimitRepository.Create(input)
		}else{
			if result.Result{
				fmt.Println("getting blocked")
				strl.rateLimitRepository.Update(input.Key, result, false)
			}
			fmt.Println("block")
		}
	}else{
		fmt.Println("decreasing")
		strl.rateLimitRepository.Update(input.Key, result, true)
	}
	return strl.rateLimitRepository.Get(input.Key)
}

