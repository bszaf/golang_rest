package controllers
import (
    "fmt"
    "net/http"
    "encoding/json"
    "github.com/bszaf/golang_rest/db"
    "github.com/bszaf/golang_rest/dto"
)

type Answers struct { Database *db.Db }

func (q Answers) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "POST":
        req := dto.AnswerReq{}
        err := json.NewDecoder(r.Body).Decode(&req)
        if err == nil {
            fmt.Println(req)
            score := calculateScore(q.Database, req)
            ranking := compareToOthers(q.Database, score)
            q.Database.AppendScore(score)
            json.NewEncoder(w).Encode(dto.AnswerRep{Score: score, Ranking: ranking})
        } else {
            resp := make(map[string]error)
            resp["error"] = err
            json.NewEncoder(w).Encode(resp)

        }
    }
}

func calculateScore(d *db.Db, req dto.AnswerReq) (int) {
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

func compareToOthers(d *db.Db, score int) (float32) {
    total := 0
    worse := 0
    for _, s := range d.GetScores() {
        total++
        if score > s {
            worse++
        }
    }
    return float32(worse)/float32(total)
}
