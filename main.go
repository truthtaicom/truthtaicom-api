package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

var UserList []User

func init() {
	user1 := User{ID: "1", Username: "user_1"}
	user2 := User{ID: "2", Username: "user_2"}
	user3 := User{ID: "3", Username: "user_3"}
	user4 := User{ID: "4", Username: "user_4"}

	UserList = append(UserList, user1, user2, user3, user4)
}

var userType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"username": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"users": &graphql.Field{
			Type:        graphql.NewList(userType),
			Description: "List of user",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return UserList, nil
			},
		},
	},
})

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"users": &graphql.Field{},
	},
})

// define schema, with our rootQuery and rootMutation
var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

func mainHandler(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

func graphQLHandler(c *gin.Context) {
	query := c.Query("query")
	result := executeQuery(query, schema)

	c.JSON(200, result)
}

func main() {
	g := gin.New()

	g.Use(gin.Logger())
	g.Use(gin.Recovery())

	g.GET("/", mainHandler)
	g.POST("/graphql", graphQLHandler)

	g.Run(":6789")
}
