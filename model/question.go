package model

import "math/rand"
import "time"

type Question struct {
    Id int
    Text string
    ValidAnswer Answer
    WrongAnswers []Answer
}

type AnonQuestion struct {
    Id int
    Text string
    Answers []Answer
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
    a := make([]Answer, len(q.WrongAnswers) + 1)
    copy(a, append(q.WrongAnswers, q.ValidAnswer))
    rand.Seed(time.Now().UnixNano())
    rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
    return AnonQuestion{
        Id: q.Id,
        Text: q.Text,
        Answers: a,
    }
}

func (q *Question) setValidAnswer(text string) () {
    q.ValidAnswer = newAnswer(text)
}

func (q *Question) addInvalidAnswer(text string) () {
    q.WrongAnswers = append(q.WrongAnswers, newAnswer(text))
}


