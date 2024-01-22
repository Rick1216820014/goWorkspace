package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gvb_server/core"
)

var client *elastic.Client

func EsConnect() *elastic.Client {
	var err error
	sniffOpt := elastic.SetSniff(false)
	host := "http://192.168.31.124:9200/"

	c, err := elastic.NewClient(
		elastic.SetURL(host),
		sniffOpt,
		elastic.SetBasicAuth("", ""),
	)
	if err != nil {
		logrus.Fatalf("es连接失败%s", err.Error())
	}
	return c
}

func init() {
	core.InitConf()
	core.InitLogger()
	client = EsConnect()
}

type DemoModel struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	UserID   string `json:"user_id"`
	CreateAt string `json:"create_at"`
}

func (DemoModel) Index() string {

	return "demo_index"
}

// 使用Elasticsearch客户端的Index方法创建一个索引请求。
// 通过调用Index方法的Index函数指定索引的名称（使用data.Index()函数返回的名称）。
// 使用BodyJson方法将data作为JSON格式的数据添加到索引请求中。
// 执行索引请求，并获取索引响应indexResponse和可能出现的错误err。
// 如果有错误，记录错误并将其返回。
// 输出索引响应的详细信息。
// 将索引响应中的Id赋值给data.ID，表示索引操作成功。
// 返回nil表示索引操作成功完成。
func Create(data *DemoModel) error {
	indexResponse, err := client.Index().Index(data.Index()).BodyJson(data).Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return err
	}
	logrus.Infof("%#v", indexResponse)
	data.ID = indexResponse.Id
	return nil
}

func FindList(key string, page int, limit int) (demoList []DemoModel, count int) {
	boolSearch := elastic.NewBoolQuery()
	from := page

	// 如果提供了关键字，添加一个匹配查询以根据"title"字段匹配关键字
	if key != "" {
		boolSearch.Must(
			//在Elasticsearch中，Bool查询可以由多个子句组成，包括 must、should、must_not 和 filter 等。在这里，Must 是其中的一个子句，它表示查询结果必须满足该子句中指定的条件。
			//具体到这段代码，boolSearch.Must 被用于将一个匹配查询条件添加到boolSearch 布尔查询中。匹配查询条件用于在 title 字段上查找与给定关键字 key 匹配的文档。
			elastic.NewMatchQuery("title", key),
		)
	}

	// 如果limit为0，则设置默认限制为10
	if limit == 0 {
		limit = 10
	}

	// 如果from为0，则设置默认页码为1
	if from == 0 {
		from = 1
	}

	// 执行搜索查询
	res, err := client. // 执行搜索操作的Elasticsearch客户端对象
				Search(DemoModel{}.Index()). // 指定搜索的索引，使用`DemoModel{}.Index()`返回的索引名称
				Query(boolSearch).           // 指定搜索的查询条件，使用之前构建的`boolSearch`布尔查询对象
				From((from - 1) * limit).    // 指定搜索结果的起始位置（与分页有关）
				Size(limit).                 // 指定搜索结果的最大数量（与分页有关）
				Do(context.Background())     // 执行搜索操作，使用上下文对象
	if err != nil {
		logrus.Error(err.Error())
		return
	}

	// 获取搜索结果的总条数
	count = int(res.Hits.TotalHits.Value)

	// 解析每个命中的文档为DemoModel对象
	for _, hit := range res.Hits.Hits {
		var demo DemoModel
		data, err := hit.Source.MarshalJSON()
		if err != nil {
			logrus.Error(err.Error())
			continue
		}
		err = json.Unmarshal(data, &demo)
		if err != nil {
			logrus.Error(err)
			continue
		}
		demo.ID = hit.Id
		demoList = append(demoList, demo)
	}

	return demoList, count
}

func main() {
	//DemoModel{}.CreateIndex()
	//Create(&DemoModel{Title: "golang_es操作测试1", UserID: "2", CreateAt: time.Now().Format("2006-01-02 15:04:05")})
	list, count := FindList("", 1, 10)
	fmt.Println(list, count)
}
