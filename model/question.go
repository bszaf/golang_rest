package model

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
    return AnonQuestion{Id: q.Id, Text: q.Text, Answers: append(q.WrongAnswers, q.ValidAnswer)}
}

func (q *Question) setValidAnswer(text string) () {
    q.ValidAnswer = newAnswer(text)
}

func (q *Question) addInvalidAnswer(text string) () {
    q.WrongAnswers = append(q.WrongAnswers, newAnswer(text))
}


