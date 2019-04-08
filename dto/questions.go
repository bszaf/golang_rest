package dto

import "github.com/bszaf/golang_rest/model"

type QuestionGetRep struct { model.Question }

type QuestionsPostReq struct { Text string; ValidAnswer string; Answers []string }
type QuestionsPostRep struct { Id int }

type QuestionsGetAllRep struct { Questions []model.AnonQuestion }


