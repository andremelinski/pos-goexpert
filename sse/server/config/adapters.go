package config

import (
	"context"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfigMod "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var awsConfig aws.Config
var onceAwsConfig sync.Once

var dynamodbClient *dynamodb.Client
var onceDdbClient sync.Once

func getAwsConfig() aws.Config {
  onceAwsConfig.Do(func() {
    var err error
    awsConfig, err = awsConfigMod.LoadDefaultConfig(context.TODO())
    if err != nil {
      panic(err)
    }
  })

  return awsConfig
}

func GetDynamodbClient() *dynamodb.Client {
  onceDdbClient.Do(func() {
    awsConfig = getAwsConfig()
    dynamodbClient = dynamodb.NewFromConfig(awsConfig, func(opt *dynamodb.Options) {
      opt.Region = awsConfig.Region
    })
  })

  return dynamodbClient
}