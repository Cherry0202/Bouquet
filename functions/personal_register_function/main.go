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
	User_id     string `json:"user_id"`
	Height      int    `json:"height"`
	Weight      int    `json:"weight"`
	Position    string `json:"position"`
	Wedding_day string `json:"wedding_day"`
	Goal_weight int    `json:"goal_weight"`
}

type Personal struct {
	User_id     string    `dynamo:"user_id"`
	Height      int       `dynamo:"height"`
	Weight      int       `dynamo:"weight"`
	Position    string    `dynamo:"position"`
	Wedding_day string    `dynamo:"wedding_day"`
	Goal_weight int       `dynamo:"goal_weight"`
	CreatedTime time.Time `dynamo:"created_time"`
}

type Log struct {
	User_id     string `dynamo:"user_id"`
	Base_weight int    `dynamo:"base_weight"`
}

type Calender struct {
	User_id     string    `dynamo:"user_id"`
	Wedding_day string    `dynamo:"wedding_day"`
	CreatedTime time.Time `dynamo:"created_time"`
}

type Response struct {
	User_id     string `json:"user_id"`
	Height      int    `json:"height"`
	Weight      int    `json:"weight"`
	Position    string `json:"position"`
	Wedding_day string `json:"wedding_day"`
	Goal_weight int    `json:"goal_weight"`
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
	r.POST("/bouquet/user/personal", func(c *gin.Context) {
		c.BindJSON(&req)
		// tables
		table_p := db.Table("b_bouquet_personal_data")
		table_log := db.Table("b_bouquet_weight_log")
		table_calender := db.Table("b_bouquet_calendar")

		jst, _ := time.LoadLocation("Asia/Tokyo")

		p := Personal{
			User_id:     req.User_id,
			Height:      req.Height,
			Weight:      req.Weight,
			Position:    req.Position,
			Goal_weight: req.Goal_weight,
			Wedding_day: req.Wedding_day,
			CreatedTime: time.Now().In(jst),
		}

		l := Log{
			User_id:     req.User_id,
			Base_weight: req.Weight,
		}

		cc := Calender{
			User_id:     req.User_id,
			Wedding_day: req.Wedding_day,
			CreatedTime: time.Now().In(jst),
		}

		fmt.Println("#######")
		fmt.Println(req.Wedding_day)

		// header
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.Header("Access-Control-Allow-Origin", "*")

		//Personal
		if err := table_p.Put(p).Run(); err != nil {
			fmt.Println("err")
			//panic(err.Error())
			c.JSON(401, gin.H{
				"message": "認証エラー(b_bouquet_personal_data)",
			})
		}

		//calender
		if err := table_calender.Put(cc).Run(); err != nil {
			fmt.Println("err")
			//panic(err.Error())
			c.JSON(401, gin.H{
				"message": "認証エラー(b_bouquet_calendar)",
			})
		}

		//Log
		if err := table_log.Put(l).Run(); err != nil {
			fmt.Println("err")
			//panic(err.Error())
			c.JSON(401, gin.H{
				"message": "認証エラー(b_bouquet_weight_log)",
			})
		}

		// body
		c.JSON(200, p)
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
