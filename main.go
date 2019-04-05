package main

import (
    "fmt"
    "net/http"
    "github.com/bszaf/golang_rest/controllers"
    "github.com/bszaf/golang_rest/db"
)


var database = db.NewDB()

func main() {
	fmt.Println("Hello, world.")
    mux := http.NewServeMux()
    questionsController := controllers.Questions{Database: &database}
    AnswersController := controllers.Answers{Database: &database}


    mux.Handle("/api/questions/", questionsController)
    mux.Handle("/api/questions", questionsController)
    mux.Handle("/api/answer", AnswersController)

    if err := http.ListenAndServe(":8000", mux); err != nil {
        fmt.Println(err)
    }
}
