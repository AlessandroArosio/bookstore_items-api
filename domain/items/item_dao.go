package items

import (
	"errors"
	"github.com/alessandroarosio/bookstore-utils-go/rest_errors"
	"github.com/alessandroarosio/bookstore_items-api/clients/elasticsearch"
)

const (
	indexItems = "items"
)

func (i *Item) Save() rest_errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save item", errors.New("database error"))
	}
	i.Id = result.Id
	return nil
}
