package db

import "github.com/bszaf/golang_rest/model"

// Database model
type Db struct {
    questions map[int]model.Question
    scores []int
    questionsSeq int
    channel chan dbRequest
}

// Request and Reply models
type dbRequest interface {isDbRequest()}

type getQuestionReq struct {
    id int
    replyChannel chan getQuestionRep
}
type getQuestionRep struct {
    val *model.Question
    err error
}

type listQuestionsReq struct {
    replyChannel chan listQuestionsRep
}
type listQuestionsRep struct {
    val []*model.Question
}

type putQuestionReq struct {
    replyChannel chan putQuestionRep
    question *model.Question
}
type putQuestionRep struct {
    val int
    err error
}

type getScoresReq struct {
    replyChannel chan getScoresRep
}

type getScoresRep struct {
    val []int
}

type appendScoreReq struct {
    replyChannel chan appendScoreRep
    score int
}

type appendScoreRep struct{}

// Add Requests to interface dbRequest
func (_ getQuestionReq) isDbRequest() {}
func (_ listQuestionsReq) isDbRequest() {}
func (_ putQuestionReq) isDbRequest() {}
func (_ getScoresReq) isDbRequest() {}
func (_ appendScoreReq) isDbRequest() {}

type dbError struct {
    Msg string
}

func (d dbError) Error() string {
    return d.Msg
}

// API functions
func NewDB() (Db) {
    db := Db{
        questions: make(map[int]model.Question),
        questionsSeq: 0,
        scores: make([]int, 10),
        channel: make(chan dbRequest, 10)}
    go dbLoop(db)
    return db
}

func (d Db) GetQuestion(id int) (*model.Question, error) {
    replyChannel := make(chan getQuestionRep)
    req := getQuestionReq{id: id, replyChannel: replyChannel}
    d.channel <- req
    rep := <- replyChannel
    close(replyChannel)
    return rep.val, rep.err
}

func (d Db) ListQuestions() ([]*model.Question) {
    replyChannel := make(chan listQuestionsRep)
    req := listQuestionsReq{replyChannel: replyChannel}
    d.channel <- req
    rep := <- replyChannel
    close(replyChannel)
    return rep.val
}
func (d Db) PutQuestion(q *model.Question) (int, error) {
    replyChannel := make(chan putQuestionRep)
    req := putQuestionReq{question: q, replyChannel: replyChannel}
    d.channel <- req
    rep := <- replyChannel
    close(replyChannel)
    return rep.val, rep.err
}

func (d Db) GetScores() []int {
    replyChannel := make(chan getScoresRep)
    req := getScoresReq{replyChannel: replyChannel}
    d.channel <- req
    rep := <- replyChannel
    close(replyChannel)
    return rep.val
}

func (d Db) AppendScore(s int) {
    replyChannel := make(chan appendScoreRep)
    req := appendScoreReq{replyChannel: replyChannel, score: s}
    d.channel <- req
    <- replyChannel
    close(replyChannel)
    return
}

// Internal functions

func dbLoop(db Db) {
    req := <-db.channel
    switch r := req.(type) {
    case getQuestionReq:
        val, err := db.getQuestion(r.id)
        r.replyChannel <- getQuestionRep{val: val, err: err}
    case listQuestionsReq:
        r.replyChannel <- listQuestionsRep{val: db.listQuestions()}
    case putQuestionReq:
        val, err := db.putQuestion(r.question)
        r.replyChannel <- putQuestionRep{val: val, err: err}
    case getScoresReq:
        r.replyChannel <- getScoresRep{val: db.getScores()}
    case appendScoreReq:
        db.appendScore(r.score)
        r.replyChannel <- appendScoreRep{}
    }
    dbLoop(db)
}

func (d *Db) getQuestion(id int) (*model.Question, error) {
    if id > d.questionsSeq - 1 {
        return nil, dbError{Msg: "not existing"}
    } else {
        q := d.questions[id]
        return &q, nil
    }
}

func (d *Db) listQuestions() ([]*model.Question) {
    all := make([]*model.Question, 0, len(d.questions))
    for _, v := range d.questions {
        all = append(all, &v)
    }
    return all
}

func (d *Db) putQuestion(q *model.Question) (int, error) {
    newId := d.questionsSeq
    q.Id = newId
    d.questionsSeq += 1
    d.questions[newId] = *q
    return newId, nil
}

func (d *Db) getScores() []int {
    return d.scores
}

func (d *Db) appendScore(s int) {
    d.scores = append(d.scores, s)
}
