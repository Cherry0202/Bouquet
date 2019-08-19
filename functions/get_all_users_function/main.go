package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	//"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/guregu/dynamo"
	"log"
)

type User struct {
	User_id string `dynamo:"user_id"`
}

var ginLambda *ginadapter.GinLambda

func init() {
	//dynamo
	db := dynamo.New(session.New(), &aws.Config{
		Region: aws.String("ap-northeast-1"), // "ap-northeast-1"等
	})

	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Gin cold start")
	r := gin.Default()
	r.GET("/bouquet/users", func(c *gin.Context) {
		// user table
		table := db.Table("b_bouquet_users")

		var users []User

		err := table.Scan().All(&users)

		// header
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.Header("Access-Control-Allow-Origin", "*")

		if err != nil {
			log.Println("err")
			//panic(err.Error())
			log.Println(err.Error())

			c.JSON(401, gin.H{
				"message": "could not scan",
			})
		}

		fmt.Println(users)

		for i, _ := range users {
			fmt.Println(users[i])
		}

		log.Println(users)

		for i, _ := range users {
			log.Println(users[i])
		}
		// body
		c.JSON(200, users)
	})

	ginLambda = ginadapter.New(r)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
