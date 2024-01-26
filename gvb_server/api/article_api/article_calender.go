package article_api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"time"
)

type CalendarResponse struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

type BucketsType struct {
	Buckets []struct {
		KeyAsString string `json:"key_as_string"`
		Key         int64  `json:"key"`
		DocCount    int    `json:"doc_count"`
	} `json:"buckets"`
}

var DateCount = map[string]int{}

func (ArticleApi) ArticleCalenderView(c *gin.Context) {
	//时间聚合
	agg := elastic.NewDateHistogramAggregation().Field("created_at").CalendarInterval("day")
	//时间段搜索
	//从今天开始到去年的今天
	formate := "2006-01-02 15:04:05"
	now := time.Now()
	lastYear := now.AddDate(-1, 0, 0)
	//lastYear := now.Add(-2 * time.Hour)
	//lt是小于，gt是大于
	query := elastic.NewRangeQuery("created_at").Gte(lastYear.Format(formate)).Lte(now.Format(formate))

	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(query).
		Aggregation("calendar", agg).
		Size(0).
		Do(context.Background())
	fmt.Println(result)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("查询失败", c)
		return
	}

	var data BucketsType
	_ = json.Unmarshal(result.Aggregations["calendar"], &data)
	var resList = make([]CalendarResponse, 0)

	for _, bucket := range data.Buckets {
		Time, _ := time.Parse(formate, bucket.KeyAsString)
		DateCount[Time.Format("2006-01-02")] = bucket.DocCount
		//resList = append(resList, CalendarResponse{
		//	Date:  Time.Format("2006-01-02"),
		//	Count: bucket.DocCount,
		//})
	}
	days := int(now.Sub(lastYear).Hours() / 24)
	for i := 0; i < days; i++ {
		day := lastYear.AddDate(0, 0, i).Format("2006-01-02")
		//day := lastYear.AddDate(0, 0, i).Format("2006-01-02")
		count, _ := DateCount[day]
		resList = append(resList, CalendarResponse{
			Date:  day,
			Count: count,
		})

	}
	res.OkWithData(resList, c)

}
