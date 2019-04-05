package db

import (
    "crypto/sha1"
    "encoding/hex"
)

type Db struct {
    Questions map[int]Question
    QuestionsSeq int
}

type Question struct {
    Text string
    ValidAnswer Answer
    WrongAnswers []Answer
}

type AnonQuestion struct {
    Text string
    Answers []Answer
}

type Answer struct {
    Text string
    Id string
}

type dbError struct {
    Msg string
}

func (d dbError) Error() string {
    return d.Msg
}

func NewDB() (Db) {
    return Db{ Questions: make(map[int]Question), QuestionsSeq: 0}
}


func (d *Db) GetQuestion(id int) (*Question, error) {
    if id > d.QuestionsSeq - 1 {
        return nil, dbError{Msg: "not existing"}
    } else {
        q := d.Questions[id]
        return &q, nil
    }
}

func (d *Db) PutQuestion(q *Question) (int, error) {
    newId := d.QuestionsSeq
    d.QuestionsSeq += 1
    d.Questions[newId] = *q
    return newId, nil
}

func NewQuestion(q_str string, v_str string, nv_str []string) (Question) {
    q := Question{Text: q_str}
    q.setValidAnswer(v_str)
    for _, str := range nv_str {
        q.addInvalidAnswer(str)
    }
    return q
}

func (q Question) HideValidAnswer() (AnonQuestion) {
    return AnonQuestion{Text: q.Text, Answers: append(q.WrongAnswers, q.ValidAnswer)}
}

func (q *Question) setValidAnswer(text string) () {
    q.ValidAnswer = newAnswer(text)
}

func (q *Question) addInvalidAnswer(text string) () {
    q.WrongAnswers = append(q.WrongAnswers, newAnswer(text))
}

func newAnswer(text string) (Answer) {
    idByte := sha1.Sum([]byte(text))
    idString := hex.EncodeToString(idByte[:])
    return Answer{Text: text, Id: idString}
}


