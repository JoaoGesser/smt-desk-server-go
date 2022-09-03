package cfg

import (
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"log"
)

var (
	ElasticSearchClient = "elastic-search"
)

func StartElasticSearch(ctx context.Context) context.Context {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
		Username: "elastic",
		Password: "changeme",
	}

	clientElasticSearch, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the cliente %s", err)
	}

	res, err := clientElasticSearch.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	fmt.Println(res)
	defer res.Body.Close()
	return context.WithValue(ctx, ElasticSearchClient, clientElasticSearch)
}
