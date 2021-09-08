package items

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ilyalevyant/bookstore_items-API/src/clients/elasticsearch"
	"github.com/ilyalevyant/bookstore_items-API/src/domain/queries"
	"github.com/ilyalevyant/bookstore_utils-go/rest_errors"
)

const(
	indexItems = "items"
)

func (i *Item) Save() *rest_errors.RestErr{
	result, err := elasticsearch.Client.Index(indexItems, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save item", errors.New("DB error"))
	}
	i.Id = result.Id
	return nil
}

func (i *Item) Get() *rest_errors.RestErr {
	itemId := i.Id
	result, err := elasticsearch.Client.Get(indexItems, i.Id)
	if err != nil {
		return rest_errors.NewInternalServerError(fmt.Sprintf("error when trying to get id %s", i.Id), errors.New("DB error"))
	}
	if !result.Found {
		return rest_errors.NewNotFoundError(fmt.Sprintf("no items found with id %s", i.Id))
	}

	bytes, err := result.Source.MarshalJSON()
	if err != nil{
		return rest_errors.NewInternalServerError("error when trying to parse DB response", errors.New("DB error"))
	}
	if err := json.Unmarshal(bytes, i); err != nil{
		return rest_errors.NewInternalServerError("error when trying to parse DB response", errors.New("DB error"))
	}
	i.Id = itemId
	return nil
}

func (i *Item) Search(query queries.EsQuery) ([]Item, *rest_errors.RestErr){
	result, err := elasticsearch.Client.Search(indexItems, query.Build())
	if err != nil{
		return nil, rest_errors.NewInternalServerError("error when trying to search documents", errors.New("DB error"))
	}
	items := make([]Item, result.TotalHits())
	for i, hit := range result.Hits.Hits{
		bytes, _ := hit.Source.MarshalJSON()
		var item Item
		if err := json.Unmarshal(bytes, &item); err != nil {
			return nil, rest_errors.NewInternalServerError("error when trying to parse response", errors.New("DB error"))
		}
		item.Id = hit.Id
		items[i] = item
	}

	if len(items) == 0 {
		return nil, rest_errors.NewNotFoundError("no items was found matching given criteria")
	}
	return items, nil
}