package main

import (
	"context"
	"fmt"
	"smt-desk-server/app/data-base/elastic-search/cfg"
	"smt-desk-server/app/data-base/elastic-search/users"
	user "smt-desk-server/pkg/dto/user"
)

func main() {
	ctx := context.Background()
	ctx = cfg.StartElasticSearch(ctx)
	fmt.Println(ctx.Value(cfg.ElasticSearchClient))
	userIndex := user.UserDTO{
		Name: "João",
	}
	id, _ := users.IndexUser(ctx, userIndex)
	fmt.Printf("Generated ID %s", id)
	users.GetUser(ctx, id)

}
