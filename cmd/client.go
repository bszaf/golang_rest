package cmd

import (
    "fmt"
    "io/ioutil"
    "github.com/spf13/cobra"
    "net/http"
)

var serverAddr string

var clientCmd = &cobra.Command{
  Use:   "client",
  Short: "connect to the HTTP server",
  Long:  ``,
  Run:   client,
}

func init() {
  rootCmd.AddCommand(clientCmd)
  serverCmd.Flags().StringVarP(&serverAddr, "server", "s", "http://localhost:8000", "Server to connect with")

}

func client(*cobra.Command, []string) {
    fmt.Println("Connecting to server:", serverAddr)
    resp, err := http.Get(serverAddr)
    if err != nil {
	// handle error
        fmt.Println("Error", err)
    } else {
        defer resp.Body.Close()
        body, _ := ioutil.ReadAll(resp.Body)
        fmt.Println(body)
    }
}
