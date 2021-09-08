package elasticsearch

import (
	"context"
	"fmt"
	"github.com/ilyalevyant/bookstore_users-api/logger"
	"github.com/olivere/elastic"
	"time"
)
var (
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	setClient(c *elastic.Client)
	Index(string, interface{}) (*elastic.IndexResponse, error)
	Get(string, string) (*elastic.GetResult, error)
	Search(string, elastic.Query) (*elastic.SearchResult, error)
}

type esClient struct {
	client *elastic.Client
}

func Init(){
	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetHealthcheckInterval(10*time.Second),
		//elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		//elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
		//elastic.SetHeaders(http.Header{
		//	"X-Caller-Id": []string{"..."},
		//}),
	)
	if err != nil {
		panic(err)
	}
	Client.setClient(client)
}

func (c *esClient) setClient(client *elastic.Client){
	c.client = client
}

func (c *esClient) Index(index string, doc interface{}) (*elastic.IndexResponse, error){
	ctx := context.Background()
	result, err := c.client.Index().Index(index).BodyJson(doc).Do(ctx)
	if err != nil{
		logger.Error("error when trying to imdex document in es", err)
		return nil, err
	}
	return result, nil
}

func (c *esClient) Get(index string, id string) (*elastic.GetResult, error) {
	ctx := context.Background()
	result, err := c.client.Get().Index(index).Id(id).Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to get id %s", id), err)
		return nil, err
	}
	return result, nil
}

func (c *esClient) Search(index string, query elastic.Query) (*elastic.SearchResult, error){
	ctx := context.Background()
	result, err := c.client.Search(index).Query(query).Do(ctx)
	if err != nil{
		logger.Error(fmt.Sprintf("error when trying to search documents in index %s", index), err)
		return nil, err
	}
	return result, nil
}