package controllers
import (
    "net/http"
    "encoding/json"
    "github.com/bszaf/golang_rest/db"
)

type Answers struct { Database *db.Db }

func (q Answers) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    resp := make(map[string]string)
    resp["msg"] = "not implemented"
    json.NewEncoder(w).Encode(resp)
}
