package users

import (
	"bytes"
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"io"
	"smt-desk-server/app/data-base/elastic-search/cfg"

	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"smt-desk-server/pkg/dto/user"
)

var (
//client *elasticsearch.Client
)

func IndexUser(ctx context.Context, user user.UserDTO) (generatedId string, err error) {
	requestBytes, err := json.Marshal(user)
	if err != nil {
		return "", err
	}

	cfgDoc := esapi.IndexRequest{
		Index:   "users",
		Body:    bytes.NewReader(requestBytes),
		Refresh: "true",
	}
	res, err := cfgDoc.Do(context.Background(), getClient(ctx))

	fmt.Println(res)
	b, _ := io.ReadAll(res.Body)
	defer res.Body.Close()

	var result map[string]interface{}
	err = json.Unmarshal(b, &result)
	if err != nil {
		fmt.Println("Error when parsin JSON returned")
	}
	return result["_id"].(string), nil
}

func GetUser(ctx context.Context, id string) error {
	cfg := esapi.GetRequest{
		Index:      "users",
		DocumentID: id,
	}

	res, err := cfg.Do(context.Background(), getClient(ctx))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	var user user.UserDTO
	err = json.Unmarshal(b, &user)
	if err != nil {
		return err
	}

	var results map[string]interface{}
	json.Unmarshal(b, &results)
	fmt.Println(results["_source"].(map[string]interface{}))
	return nil
}

func getClient(ctx context.Context) *elasticsearch.Client {
	return ctx.Value(cfg.ElasticSearchClient).(*elasticsearch.Client)
}
