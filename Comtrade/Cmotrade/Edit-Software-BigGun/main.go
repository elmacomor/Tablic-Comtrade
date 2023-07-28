package main

import (
	"github.com/gin-gonic/gin"
	"main.go/api"
	"main.go/initializers"
	// "encoding/json"
	// "log"
	// "net/http"
	// "github.com/gorilla/mux"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {

	r := gin.Default()
	api.InitializeHandlers(r)
	r.GET("/cards", api.NewDeckHandler)
	r.GET("/cards/:userid/:deckid", api.ShowPlayerCards)
	r.GET("/takecardsfromtable/:deckId/:handPile/:takenPile", api.TakeCardsFromTable)
	r.GET("/throwCard/:cardCode/:deckId/:playerPile", api.ThrowCardHandler)
	r.Run()
	// router := mux.NewRouter()

	// router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	// 	response := map[string]string{
	// 		"message": "Hello Docker!",
	// 	}
	// 	json.NewEncoder(rw).Encode(response)
	// })

	// log.Println("Server is running!")
	// http.ListenAndServe(":4000", router)
}
