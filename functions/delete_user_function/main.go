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
	User_id string `json:"user_id"`
}

type User struct {
	User_id string `dynamo:"user_id"`
}

var ginLambda *ginadapter.GinLambda

func init() {
	// Credentialsは各自で
	////dynamo
	db := dynamo.New(session.New(), &aws.Config{
		Region: aws.String("ap-northeast-1"), // "ap-northeast-1"等
	})

	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Gin cold start")
	var req Request
	r := gin.Default()
	r.DELETE("/bouquet/user", func(c *gin.Context) {
		c.BindJSON(&req)
		// user table
		table := db.Table("b_bouquet_users")
		table1 := db.Table("b_bouquet_personal_data")
		table2 := db.Table("b_bouquet_weight_log")
		table3 := db.Table("b_bouquet_calendar")

		u := User{
			User_id: req.User_id,
		}
		fmt.Println(u)

		// header
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.Header("Access-Control-Allow-Origin", "*")

		if req.User_id == "" {
			c.JSON(401, gin.H{
				"message": "emailを送ってください",
			})
		}

		if err := table.Delete("user_id", u.User_id).Run(); err != nil {
			fmt.Println("err")

			c.JSON(401, gin.H{
				"message": "Failure(b_bouquet_users)",
				"request": u,
			})
		}

		if err := table1.Delete("user_id", u.User_id).Run(); err != nil {
			fmt.Println("err")

			c.JSON(401, gin.H{
				"message": "Failure(b_bouquet_personal_data)",
				"request": u,
			})
		}

		if err := table2.Delete("user_id", u.User_id).Run(); err != nil {
			fmt.Println("err")

			c.JSON(401, gin.H{
				"message": "Failure(b_bouquet_weight_log)",
				"request": u,
			})
		}

		if err := table3.Delete("user_id", u.User_id).Run(); err != nil {
			fmt.Println("err")

			c.JSON(401, gin.H{
				"message": "Failure(b_bouquet_calendar)",
				"request": u,
			})
		}

		// body
		c.JSON(200, gin.H{
			"message": "Success!",
		})
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
