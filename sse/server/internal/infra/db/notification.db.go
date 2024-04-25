package db

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)


type NotificationDB struct {
	dbConfig *dynamodb.Client
}


func NewNotificationDB(db *dynamodb.Client) *NotificationDB{
	return &NotificationDB{
		db,
	}
}


// func (db *NotificationDB)GetNotification() (interface{}, error) {
// 	result, err := db.dbConfig.GetItem(context.TODO(), &dynamodb.GetItemInput{
// 		Key: GetKey(), TableName: aws.String("notification"),
// 	})

// 	if err!= nil {
// 		return nil, err
// 	}

// 	fmt.Println(result)
// 	return result, nil
// }

func (db *NotificationDB) GetAll() ([]map[string]interface{}, error) {
	var output []map[string]interface{}
	var response *dynamodb.ExecuteStatementOutput
	var err error
	var nextToken *string
	for moreData := true; moreData; {
		response, err = db.dbConfig.ExecuteStatement(context.TODO(), &dynamodb.ExecuteStatementInput{
			Statement: aws.String(
				fmt.Sprintf("SELECT title, info.rating FROM \"%v\"", "notification")),
			Limit:     aws.Int32(100),
			NextToken: nextToken,
		})
		if err != nil {
			log.Printf("Couldn't get movies. Here's why: %v\n", err)
			moreData = false
		} else {
			var pageOutput []map[string]interface{}
			err = attributevalue.UnmarshalListOfMaps(response.Items, &pageOutput)
			if err != nil {
				log.Printf("Couldn't unmarshal response. Here's why: %v\n", err)
			} else {
				log.Printf("Got a page of length %v.\n", len(response.Items))
				output = append(output, pageOutput...)
			}
			nextToken = response.NextToken
			moreData = nextToken != nil
		}
	}
	return output, err
}

// GetKey returns the composite primary key of the movie in a format that can be
// sent to DynamoDB.
func GetKey() map[string]types.AttributeValue {
	title, err := attributevalue.Marshal("test")
	if err != nil {
		panic(err)
	}
	return map[string]types.AttributeValue{"title": title}
}