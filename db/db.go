package db

import "github.com/bszaf/golang_rest/model"

type Db struct {
    Questions map[int]model.Question
    QuestionsSeq int
}

type dbError struct {
    Msg string
}

func (d dbError) Error() string {
    return d.Msg
}

func NewDB() (Db) {
    return Db{ Questions: make(map[int]model.Question), QuestionsSeq: 0}
}


func (d *Db) GetQuestion(id int) (*model.Question, error) {
    if id > d.QuestionsSeq - 1 {
        return nil, dbError{Msg: "not existing"}
    } else {
        q := d.Questions[id]
        return &q, nil
    }
}

func (d *Db) ListQuestions() ([]*model.Question) {
    all := make([]*model.Question, 0, len(d.Questions))
    for _, v := range d.Questions {
        all = append(all, &v)
    }
    return all
}

func (d *Db) PutQuestion(q *model.Question) (int, error) {
    newId := d.QuestionsSeq
    q.Id = newId
    d.QuestionsSeq += 1
    d.Questions[newId] = *q
    return newId, nil
}
