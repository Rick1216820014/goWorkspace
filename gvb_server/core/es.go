package core

import (
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gvb_server/global"
	"net/http"
)

func EsConnect() *elastic.Client {
	var err error
	sniffOpt := elastic.SetSniff(false)

	//2024年1月24日，邪门的bug:
	//不知道哪里设置了全局代理，只要连接就会走代理，我明明已经关了代理了，还是会走（可见什么地方给配置上了）
	//es配置文件中只有一行http.host: 0.0.0.0
	//使用es head可以访问es,es是没问题的
	//bug如下：
	//es连接失败health check timeout: Head "http://192.168.31.124:9200": proxyconnect tcp: dial tcp 127.0.0.1:1080:
	//connectex: No connection could be made because the target machine actively refused it.: no Elasticsearch node available

	//gpt老师帮忙后防止了该bug出现，虽然我仍然不知道这个bug是怎么出现的：
	//在创建 elastic.NewClient 时，我们使用自定义的 http.Client，在传输中设置了 Proxy 为 nil，这将覆盖任何可能的全局代理设置，确保请求不经过代理。

	c, err := elastic.NewClient(
		elastic.SetURL(global.Config.ES.Url()),
		sniffOpt,
		elastic.SetHttpClient(&http.Client{
			Transport: &http.Transport{
				Proxy: nil,
			},
		}),
		elastic.SetBasicAuth(global.Config.ES.User, global.Config.ES.Password),
	)
	if err != nil {
		logrus.Fatalf("es连接失败%s", err.Error())
	}
	return c
}
