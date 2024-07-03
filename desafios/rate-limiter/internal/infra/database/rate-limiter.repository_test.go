package database

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/go-redis/redis"
	"github.com/stretchr/testify/suite"
)

type RateLimitRepositoryTestSuite struct {
	suite.Suite
	localRedis *redis.Client
	repo *RateLimitRepository
	createTestData *RateLimitInput
}

func (suite *RateLimitRepositoryTestSuite) SetupSuite() {
	client := suite.connectDBLocally()
	repo := NewRateLimitRepository(client)
	suite.repo = repo
	suite.createTestData = &RateLimitInput{
		"key",
        10,
    	5000*time.Millisecond,
	}
}

func (suite *RateLimitRepositoryTestSuite) TearDownTest() {
    // Clean up or teardown after tests
    suite.connectDBLocally().FlushAll()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(RateLimitRepositoryTestSuite))
}

func (suite *RateLimitRepositoryTestSuite) Test_GetRateLimitInfo_When_UserNotExist() {
	user, err := suite.repo.Get("key")
	suite.Empty(user)
	suite.Error(err, redis.Nil)

}

func (suite *RateLimitRepositoryTestSuite) Test_GetRateLimitInfo_When_UserExist() {
	suite.repo.Create(suite.createTestData)
	user, err := suite.repo.Get("key")
	suite.Equal(suite.createTestData.Limit, user.Limit)
	suite.NoError(err)
}

func (suite *RateLimitRepositoryTestSuite) Test_CreateRateLimitInfo() { 
	
	suite.repo.Create(suite.createTestData)
	user, err := suite.repo.Get("key")
	suite.Equal(suite.createTestData.Limit, user.Limit)
	suite.NoError(err)
}

func (suite *RateLimitRepositoryTestSuite) Test_UpdateRateLimitInfo() {
	suite.repo.Create(suite.createTestData)
	user, err := suite.repo.Get("key")
	suite.NoError(err)
	suite.Equal(suite.createTestData.Limit, user.Limit)

	suite.repo.Update("key", user, false)
	user, err = suite.repo.Get("key")
	suite.NoError(err)
	suite.False(user.Result)
}

func(suite *RateLimitRepositoryTestSuite) connectDBLocally() *redis.Client{
	host:= "localhost"
	port:= "6379"

	if os.Getenv("DB_HOST") != "" {
		host = os.Getenv("DB_HOST")
	}
	if os.Getenv("DB_PORT") != "" {
		port = os.Getenv("DB_PORT")
	} 

	return redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s",   host, port),
		Password: "", 
		DB: 0, 
    })
}