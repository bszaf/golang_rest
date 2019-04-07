package cmd

import (
    "fmt"
    "net/http"
    "github.com/bszaf/golang_rest/controllers"
    "github.com/bszaf/golang_rest/db"
    "github.com/spf13/cobra"
)

var listenPort int

var serverCmd = &cobra.Command{
  Use:   "server",
  Short: "start the HTTP server",
  Long:  ``,
  Run:   server,
}

func init() {
  rootCmd.AddCommand(serverCmd)
  serverCmd.Flags().IntVarP(&listenPort, "port", "p", 8000, "Set HTTP port to listen on")

}

func server(*cobra.Command, []string) {
    var database = db.NewDB()
    fmt.Println("Serving server on port: ", listenPort)
    mux := http.NewServeMux()
    questionsController := controllers.Questions{Database: &database}
    AnswersController := controllers.Answers{Database: &database}


    mux.Handle("/api/questions/", questionsController)
    mux.Handle("/api/questions", questionsController)
    mux.Handle("/api/answer", AnswersController)

    if err := http.ListenAndServe(fmt.Sprintf(":%d", listenPort), mux); err != nil {
        fmt.Println(err)
    }
}
