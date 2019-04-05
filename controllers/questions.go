package controllers
import (
    "fmt"
    "net/http"
    "encoding/json"
    "strings"
    "strconv"
    "github.com/bszaf/golang_rest/db"
)

type Questions struct { Database *db.Db }
type questionsPostReq struct { Text string; ValidAnswer string; Answers []string }
type questionsPostRep struct { Id int }

func (q Questions) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    uriSegments := strings.Split(r.URL.Path, "/")
    switch r.Method {
    case "GET":
        if len(uriSegments) == 4 {
            q.handleGetSingle(uriSegments[3], w, r)
        } else if len(uriSegments) == 3 {
            q.handleGetMany(w, r)
        } else {
            fmt.Println("unhandled method")
        }
    case "POST":
        if len(uriSegments) == 3 {
            q.handlePostNew(w, r)
        } else {
            fmt.Println("unhandled")
        }
    default:
        fmt.Println("unhandled method")
    }
}

func (q Questions) handleGetSingle(idString string, w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(idString)
    resp := make(map[string]error)
    if err != nil {
        resp["error"] = err
        json.NewEncoder(w).Encode(resp)
        return
    }
    question, err := q.Database.GetQuestion(id)
    if err != nil {
        resp["error"] = err
        json.NewEncoder(w).Encode(resp)
        return
    }
    anonQuestion := question.HideValidAnswer()
    json.NewEncoder(w).Encode(anonQuestion)
}

func (q Questions) handleGetMany(w http.ResponseWriter, r *http.Request) {
    return
}

func (q Questions) handlePostNew(w http.ResponseWriter, r *http.Request) {
    req := questionsPostReq{}
    err := json.NewDecoder(r.Body).Decode(&req)
    if err == nil {
        fmt.Println(req)
        question := db.NewQuestion(req.Text, req.ValidAnswer, req.Answers)
        id, _ := q.Database.PutQuestion(&question)
        json.NewEncoder(w).Encode(questionsPostRep{Id: id})
    }
}
