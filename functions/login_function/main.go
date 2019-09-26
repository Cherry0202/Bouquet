package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/guregu/dynamo"
	"log"
)

type Request struct {
	User_id  string `json:"user_id"`
	Password string `json:"password"`
}

type User struct {
	User_id  string `dynamo:"user_id"`
	Password string `dynamo:"password"`
}

var ginLambda *ginadapter.GinLambda

func init() {

	////dynamo
	db := dynamo.New(session.New(), &aws.Config{
		Region: aws.String("ap-northeast-1"), // "ap-northeast-1"等
	})

	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Gin cold start")
	var req Request
	r := gin.Default()
	r.POST("/bouquet/user/login", func(c *gin.Context) {
		c.BindJSON(&req)
		// user table
		table := db.Table("bouquet_users")

		u := User{
			User_id:  req.User_id,
			Password: req.Password,
		}
		fmt.Println(u)

		// header
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.Header("Access-Control-Allow-Origin", "*")

		if req.User_id == "" {
			c.JSON(401, gin.H{
				"message": "emailを入力してください",
				"Request": u,
			})
		}

		if err := table.Get("user_id", u.User_id).
			Filter("'password' = ?", u.Password).
			One(&u); err != nil {
			fmt.Println("err")

			c.JSON(401, gin.H{
				"message": "認証エラー",
				"Request": u,
			})
		} else {
			// body
			c.JSON(200, gin.H{
				"token":   "mti2019",
				"user_id": req.User_id,
			})
		}
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
