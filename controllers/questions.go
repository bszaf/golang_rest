package controllers
import (
    "fmt"
    "net/http"
    "encoding/json"
    "strings"
    "strconv"
    "github.com/bszaf/golang_rest/db"
    "github.com/bszaf/golang_rest/model"
    "github.com/bszaf/golang_rest/dto"
)

type Questions struct { Database *db.Db }
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
    resp := dto.QuestionsGetAllRep{}
    for _, question := range q.Database.ListQuestions() {
        anonimized := question.HideValidAnswer()
        resp.Questions = append(resp.Questions, anonimized)
    }
    json.NewEncoder(w).Encode(resp)
}

func (q Questions) handlePostNew(w http.ResponseWriter, r *http.Request) {
    req := dto.QuestionsPostReq{}
    err := json.NewDecoder(r.Body).Decode(&req)
    if err == nil {
        question := model.NewQuestion(req.Text, req.ValidAnswer, req.Answers)
        id, _ := q.Database.PutQuestion(&question)
        json.NewEncoder(w).Encode(dto.QuestionsPostRep{Id: id})
    }
}
