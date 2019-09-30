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

var ginLambda *ginadapter.GinLambda

type User_name struct {
	User_name string `dynamo:"user_name" json:"user_name"`
}

type Personal struct {
	Height      int    `dynamo:"height"`
	Weight      int    `dynamo:"weight"`
	Goal_Weight int    `dynamo:"goal_weight"`
	Position    string `dynamo:"position"`
	Wedding_day string `dynamo:"wedding_day"`
}

//type Wedding_day struct {
//	Wedding_day string `dynamo:"wedding_day"`
//}

type Nail_and_extetiton struct {
	Title string `json:"title"`
	Start string `json:"start"`
	End   string `json:"end"`
}

type Nail_consideration struct {
	Title string `json:"title"`
	Start string `json:"start"`
	End   string `json:"end"`
}

type Bridal_beauty_treatment_salon struct {
	Title string `json:"title"`
	Start string `json:"start"`
	End   string `json:"end"`
}

type Whitening struct {
	Title string `json:"title"`
	Start string `json:"start"`
	End   string `json:"end"`
}

type Salon_consideration struct {
	Title string `json:"title"`
	Start string `json:"start"`
	End   string `json:"end"`
}

//{
//	"nail_and_extetiton": {
//		"title": ""
//	},
//	"test": {
//		"a": b
//	}
//}

type Response struct {
	User_name   string `json:"user_name"`
	Height      int    `json:"height"`
	Goal_weight int    `json:"goal_weight"`
	//Weight             int    `json:"weight"`
	Until_goal_weight  int    `json:"until_goal_weight"`
	Position           string `json:"position"`
	Counting_days      int    `json:"counting_days"`
	Nail_and_extetiton struct {
		Title string `json:"title"`
		Start string `json:"start"`
		End   string `json:"end"`
	}
	Nail_consideration struct {
		Title string `json:"title"`
		Start string `json:"start"`
		End   string `json:"end"`
	}
	Bridal_beauty_treatment_salon struct {
		Title string `json:"title"`
		Start string `json:"start"`
		End   string `json:"end"`
	}
	Whitening struct {
		Title string `json:"title"`
		Start string `json:"start"`
		End   string `json:"end"`
	}
	Salon_consideration struct {
		Title string `json:"title"`
		Start string `json:"start"`
		End   string `json:"end"`
	}
}

func init() {

	db := dynamo.New(session.New(), &aws.Config{
		Region: aws.String("ap-northeast-1"), // "ap-northeast-1"等
	})
	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Gin cold start")
	var name User_name
	var personal Personal
	//var weight GoalWeight
	//var wedding Wedding_day
	r := gin.Default()
	//var req Request
	r.GET("/bouquet/calendar", func(c *gin.Context) {
		// header
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.Header("Access-Control-Allow-Origin", "*")

		qs := c.Query("user_id")

		fmt.Println("hogehogoe")
		fmt.Println(qs)

		table := db.Table("bouquet_users")
		table_p := db.Table("bouquet_personal_data")
		//table_log := db.Table("b_bouquet_weight_log")
		//table_c := db.Table("bouquet_calendar")

		if qs == "" {
			c.JSON(401, gin.H{
				"message": "user_idを送ってください",
				"Request": qs,
			})
		} else if err := table.Get("user_id", qs).One(&name); err != nil {
			fmt.Println("err")
			c.JSON(401, gin.H{
				"message": "不正なuser_idです",
				"request": qs,
			})
		} else if err := table_p.Get("user_id", qs).One(&personal); err != nil {
			fmt.Println("err")
			c.JSON(401, gin.H{
				"message": "不正なuser_idです",
				"request": qs,
			})
		} else {
			//現在の日付を取得
			jst, _ := time.LoadLocation("Asia/Tokyo")
			today := time.Now().In(jst)
			log.Print(today, "現在時刻")
			//bSix := t.AddDate(0, 6, 0)
			//bThree := t.AddDate(0, 3, 0)
			//bOne := t.AddDate(0, 1, 0)
			//fmt.Println("aaaaaaaa")

			str := personal.Wedding_day
			const layout = "2006-01-02"
			t, _ := time.Parse(layout, str)
			log.Print("↓")
			log.Println(t)
			fmt.Println(t.Local())

			//残り日数計算
			duration := t.Sub(today)
			hours0 := int(duration.Hours())
			fmt.Println(hours0)
			days := hours0 / 24

			//目標体重まであと~
			until_goal_weight := personal.Weight - personal.Goal_Weight
			//user_nameを返す
			res := Response{
				User_name:         name.User_name,
				Until_goal_weight: until_goal_weight,
				Counting_days:     days,
				Goal_weight:       personal.Goal_Weight,
			}
			c.JSON(200, res)
		}
		//else

		//if err := table_log.Get("user_id", qs).One(&weight); err != nil {
		//	fmt.Println("err")
		//	c.JSON(401, gin.H{
		//		"message": "不正なuser_idです",
		//		"request": qs,
		//	})
		//}

		/*		if err := table_c.Get("user_id", qs).One(&wedding); err != nil {
				fmt.Println("err")
				c.JSON(401, gin.H{
					"message": "不正なuser_idです",
					"request": qs,
				})
		*/
		//}
		//else {
		//
		//str := wedding.Wedding_day
		////fmt.Println(str)
		////str := "2019-01-01"
		//layout := "2006-01-02"
		//t, _ := time.Parse(layout, str)
		//fmt.Println("#########################################")
		//fmt.Println(str) // => "2003-04-18 00:00:00 +0000 UTC"
		//fmt.Println(t)
		//fmt.Println(t.Local())
		//
		//today := time.Now()
		//bSix := t.AddDate(0, 6, 0)
		//bThree := t.AddDate(0, 3, 0)
		//bOne := t.AddDate(0, 1, 0)
		//fmt.Println("aaaaaaaa")
		//
		////残り日数計算
		//duration := t.Sub(today)
		//hours0 := int(duration.Hours())
		//fmt.Println(hours0)
		//days := hours0 / 24
		//
		//if days < 0 {
		//	c.JSON(401, gin.H{
		//		"Error": "結婚日時がおかしいです。",
		//	})
		//} else if days == 0 {
		//	return
		//} else {
		//	days = days + 1
		//}
		//fmt.Println(days)
		//
		//if !bSix.After(today) {
		//	fmt.Printf("６ヶ月以上")
		//	nailStart := t.AddDate(0, 0, -7)
		//	nailEnd := t.AddDate(0, 0, -1)
		//	nsStart := t.AddDate(0, -1, 0)
		//	nsEnd := t.AddDate(0, 0, -8)
		//	beStart := t.AddDate(0, -3, 0)
		//	beEnd := t.AddDate(0, -2, 0)
		//	wStart := t.AddDate(0, -4, 0)
		//	wEnd := t.AddDate(0, -3, -1)
		//	scStart := today.AddDate(0, 0, 1)
		//	scEnd := today.AddDate(0, 1, 0)
		//	res := Response{
		//		User_name:     name.User_name,
		//		Height:        height.Height,
		//		Goal_weight:   height.Goal_Weight,
		//		Weight:        height.Weight,
		//		Counting_days: days,
		//		Position:      height.Position,
		//		Nail_and_extetiton: Nail_and_extetiton{
		//			Title: "ネイル＆マツエク",
		//			Start: nailStart.String(),
		//			End:   nailEnd.String(),
		//		},
		//		Nail_consideration: Nail_consideration{
		//			Title: "ネイルサロンの検討",
		//			Start: nsStart.String(),
		//			End:   nsEnd.String(),
		//		},
		//		Bridal_beauty_treatment_salon: Bridal_beauty_treatment_salon{
		//			Title: "ブライダルエステ",
		//			Start: beStart.String(),
		//			End:   beEnd.String(),
		//		},
		//		Whitening: Whitening{
		//			Title: "ホワイトニング",
		//			Start: wStart.String(),
		//			End:   wEnd.String(),
		//		},
		//		Salon_consideration: Salon_consideration{
		//			Title: "エステサロンの検討",
		//			Start: scStart.String(),
		//			End:   scEnd.String(),
		//		},
		//	}
		//	c.JSON(200, res)
		//
		//} else if !today.After(bSix) || !bThree.After(today) {
		//	fmt.Printf("6ヶ月未満３ヶ月以上")
		//	nailStart := t.AddDate(0, 0, -7)
		//	nailEnd := t.AddDate(0, 0, -1)
		//	nsStart := t.AddDate(0, -1, 0)
		//	nsEnd := t.AddDate(0, 0, -8)
		//	beStart := today.AddDate(0, 0, 1)
		//	beEnd := t.AddDate(0, -2, 0)
		//	res := Response{
		//		User_name:     name.User_name,
		//		Height:        height.Height,
		//		Goal_weight:   height.Goal_Weight,
		//		Weight:        height.Weight,
		//		Counting_days: days,
		//		Position:      height.Position,
		//		Nail_and_extetiton: Nail_and_extetiton{
		//			Title: "ネイル＆マツエク",
		//			Start: nailStart.String(),
		//			End:   nailEnd.String(),
		//		},
		//		Nail_consideration: Nail_consideration{
		//			Title: "ネイルサロンの検討",
		//			Start: nsStart.String(),
		//			End:   nsEnd.String(),
		//		},
		//		Bridal_beauty_treatment_salon: Bridal_beauty_treatment_salon{
		//			Title: "ブライダルエステ",
		//			Start: beStart.String(),
		//			End:   beEnd.String(),
		//		},
		//	}
		//	c.JSON(200, res)
		//} else if !today.After(bThree) || !bOne.After(today) {
		//	fmt.Printf("3ヶ月未満1ヶ月以上")
		//	nailStart := t.AddDate(0, 0, -7)
		//	nailEnd := t.AddDate(0, 0, -1)
		//	nsStart := today.AddDate(0, 0, 1)
		//	nsEnd := t.AddDate(0, 0, -8)
		//	res := Response{
		//		User_name:     name.User_name,
		//		Height:        height.Height,
		//		Goal_weight:   height.Goal_Weight,
		//		Weight:        height.Weight,
		//		Counting_days: days,
		//		Position:      height.Position,
		//		Nail_and_extetiton: Nail_and_extetiton{
		//			Title: "ネイル＆マツエク",
		//			Start: nailStart.String(),
		//			End:   nailEnd.String(),
		//		},
		//		Nail_consideration: Nail_consideration{
		//			Title: "ネイルサロンの検討",
		//			Start: nsStart.String(),
		//			End:   nsEnd.String(),
		//		},
		//	}
		//	c.JSON(200, res)
		//} else if !today.After(bOne) {
		//	fmt.Println("1ヶ月未満だ")
		//	nailStart := t.AddDate(0, 0, -7)
		//	nailEnd := t.AddDate(0, 0, -1)
		//	nsStart := today.AddDate(0, 0, 1)
		//	nsEnd := today.AddDate(0, 0, 8)
		//	res := Response{
		//		User_name:     name.User_name,
		//		Height:        height.Height,
		//		Goal_weight:   height.Goal_Weight,
		//		Weight:        height.Weight,
		//		Counting_days: days,
		//		Position:      height.Position,
		//		Nail_and_extetiton: Nail_and_extetiton{
		//			Title: "ネイル＆マツエク",
		//			Start: nailStart.String(),
		//			End:   nailEnd.String(),
		//		},
		//		Nail_consideration: Nail_consideration{
		//			Title: "ネイルサロンの検討",
		//			Start: nsStart.String(),
		//			End:   nsEnd.String(),
		//		},
		//	}
		//	c.JSON(200, res)
		//}
		//
		// }
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
