package dto

type QuestionAnswer struct {
    QuestionId int
    AnswerId string
}
type AnswerReq struct { Answers []QuestionAnswer }
type AnswerRep struct { Ranking float32; Score int }
