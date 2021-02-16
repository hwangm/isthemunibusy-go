package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/hwangm/isthemunibusy-go/dal"
	"github.com/hwangm/isthemunibusy-go/mutations"
	"github.com/hwangm/isthemunibusy-go/queries"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	dal.InitDb()
	// Schema
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: queries.GetRootFields()}
	rootMutation := graphql.ObjectConfig{Name: "RootMutation", Fields: mutations.GetRootFields()}
	schemaConfig := graphql.SchemaConfig{
		Query:    graphql.NewObject(rootQuery),
		Mutation: graphql.NewObject(rootMutation),
	}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)
	http.HandleFunc("/websocket", echo)
	http.ListenAndServe(":8080", nil)

}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)

		responseMessage := fmt.Sprintf("You said: %s at %s, so I waited 5 seconds to respond", message, time.Now().String())
		time.Sleep(5 * time.Second)

		err = c.WriteMessage(mt, []byte(responseMessage))
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
