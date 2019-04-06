package controllers
import (
    "fmt"
    "net/http"
    "encoding/json"
    "github.com/bszaf/golang_rest/db"
)

type Answers struct { Database *db.Db }

type questionAnswer struct {
    QuestionId int
    AnswerId string
}
type answerReq struct { Answers []questionAnswer }
type answerRep struct { Score int }

func (q Answers) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "POST":
        req := answerReq{}
        err := json.NewDecoder(r.Body).Decode(&req)
        if err == nil {
            fmt.Println(req)
            score := calculateScore(q.Database, req)
            json.NewEncoder(w).Encode(answerRep{Score: score})
        } else {
            resp := make(map[string]error)
            resp["error"] = err
            json.NewEncoder(w).Encode(resp)

        }
    }
}

func calculateScore(d *db.Db, req answerReq) (int) {
    goodAnswers := 0
    for _, v := range req.Answers {
        if q, err := d.GetQuestion(v.QuestionId); err == nil {
            if q.ValidAnswer.Id == v.AnswerId {
                goodAnswers++
            }
        }
    }
    return goodAnswers
}
