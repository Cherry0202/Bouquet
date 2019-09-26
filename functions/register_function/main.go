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
	"time"
)

type Request struct {
	User_id   string `json:"user_id"`
	Password  string `json:"password"`
	User_name string `json:"user_name"`
}

type User struct {
	User_id     string    `dynamo:"user_id"`
	Password    string    `dynamo:"password"`
	User_name   string    `dynamo:"user_name"`
	CreatedTime time.Time `dynamo:"created_time"`
}

type Personal struct {
	User_id string `dynamo:"user_id"`
}

type Log struct {
	User_id string `dynamo:"user_id"`
}

type Calender struct {
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
	var req Request
	r := gin.Default()
	r.POST("/bouquet/user", func(c *gin.Context) {
		c.BindJSON(&req)
		// user table
		table := db.Table("bouquet_users")
		table1 := db.Table("bouquet_personal_data")
		table2 := db.Table("bouquet_weight_log")
		table3 := db.Table("bouquet_calendar")

		jst, _ := time.LoadLocation("Asia/Tokyo")

		u := User{
			User_id:     req.User_id,
			Password:    req.Password,
			User_name:   req.User_name,
			CreatedTime: time.Now().In(jst),
		}

		p := Personal{
			User_id: req.User_id,
		}

		l := Log{
			User_id: req.User_id,
		}

		cc := Calender{
			User_id: req.User_id,
		}

		fmt.Println(u)

		// header
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.Header("Access-Control-Allow-Origin", "*")

		//　以下並列処理で書き直す

		if req.User_id == "" {
			c.JSON(401, gin.H{
				"message": "emailを入力してください",
			})
		}
		//user
		if err := table.Put(u).If("attribute_not_exists(user_id)").Run(); err != nil {
			fmt.Println("err")
			//panic(err.Error())
			c.JSON(401, gin.H{
				"message": "認証エラー" +
					"既に存在しているIDです" +
					"(b_bouquet_users)",
			})
		}

		//Personal
		if err := table1.Put(p).If("attribute_not_exists(user_id)").Run(); err != nil {
			fmt.Println("err")
			//panic(err.Error())
			c.JSON(401, gin.H{
				"message": "認証エラー" +
					"既に存在しているIDです" +
					"(b_bouquet_personal_data)",
			})
		}

		//Log
		if err := table2.Put(l).If("attribute_not_exists(user_id)").Run(); err != nil {
			fmt.Println("err")
			//panic(err.Error())
			c.JSON(401, gin.H{
				"message": "認証エラー" +
					"既に存在しているIDです" +
					"(b_bouquet_weight_log)",
			})
		}

		//Log
		if err := table3.Put(cc).If("attribute_not_exists(user_id)").Run(); err != nil {
			fmt.Println("err")
			//panic(err.Error())
			c.JSON(401, gin.H{
				"message": "認証エラー" +
					"既に存在しているIDです" +
					"(b_bouquet_weight_log)",
			})
		}

		//後にtokenを渡す処理に変更
		// body
		c.JSON(200, gin.H{
			"user_id":   req.User_id,
			"user_name": req.User_name,
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
