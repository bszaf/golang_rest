package cmd

import (
    "fmt"
    "net/http"
    "encoding/json"
    "bytes"
    "github.com/spf13/cobra"
    "github.com/bszaf/golang_rest/dto"
)

// variables for parameters
var serverAddr string
var questionBody string
var questionValidAnswer string
var questionAnswers []string

var clientCmd = &cobra.Command{
    Use:   "client",
    Short: "connect to the HTTP server",
    Long:  ``,
}

var listCmd = &cobra.Command{
    Use: "list",
    Short: "List all available questions",
    Long: ``,
    Run: list,
}

var postCmd = &cobra.Command{
    Use: "post",
    Short: "Pushes new question to the server",
    Long: ``,
    Run: post,
}

var quizCmd = &cobra.Command{
    Use: "quiz",
    Short: "Starts a quiz from the server",
    Long: ``,
    Run: quiz,
}

func init() {
    rootCmd.AddCommand(clientCmd)
    clientCmd.PersistentFlags().StringVarP(&serverAddr, "server", "s", "http://localhost:8000", "Server to connect with")

    clientCmd.AddCommand(listCmd)
    clientCmd.AddCommand(postCmd)
    clientCmd.AddCommand(quizCmd)

    postCmd.Flags().StringVarP(&questionBody, "body", "b", "", "Text of question")
    postCmd.MarkFlagRequired("body")
    postCmd.Flags().StringVarP(&questionValidAnswer, "valid", "v", "", "Text of valid answer")
    postCmd.MarkFlagRequired("valid")
    postCmd.Flags().StringArrayVarP(&questionAnswers, "wrong", "w", []string{}, "Text of wrong answers")
    postCmd.MarkFlagRequired("wrong")


}

func list(*cobra.Command, []string) {
    questions := getQuestions()
    for j, q := range questions.Questions {
        fmt.Println(j, "Text: ", q.Text)
        for i, a := range q.Answers {
            str := fmt.Sprintf("\t %d) %s", i, a.Text)
            fmt.Println(str)
        }
    }
}

func post(*cobra.Command, []string) {
    req := dto.QuestionsPostReq{
        Text: questionBody,
        ValidAnswer: questionValidAnswer,
        Answers: questionAnswers,
    }
    reply := postQuestion(req)
    fmt.Println("Question id =", reply.Id)
}

func quiz(*cobra.Command, []string) {
    questions := getQuestions()
    answers := dto.AnswerReq{}
    for _, q := range questions.Questions {
        fmt.Println("Question:", q.Text)
        for i, a := range q.Answers {
            str := fmt.Sprintf("\t %d) %s", i, a.Text)
            fmt.Println(str)
        }
        selectedIndex := scanAnswerNumber(len(q.Answers))
        selected := q.Answers[selectedIndex]
        answer := dto.QuestionAnswer{
            QuestionId: q.Id,
            AnswerId: selected.Id,
        }
        answers.Answers = append(answers.Answers, answer)
    }
    reply := postAnswers(answers)
    fmt.Println("Total valid answers:", reply.Score)
    str := fmt.Sprintf("You where better than: %.2f %%", reply.Ranking*100)
    fmt.Println(str)
}

// API related functions

func postAnswers(a dto.AnswerReq) *dto.AnswerRep {
    URL := fmt.Sprintf("%s/api/answer", serverAddr)
    fmt.Println("Post answers to :", URL)
    buff := new(bytes.Buffer)
    json.NewEncoder(buff).Encode(a)
    resp, err := http.Post(URL, "application/json", buff)
    if err != nil {
        fmt.Println("Error", err)
        return nil
    }
    if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
        answersReply := dto.AnswerRep{}
        json.NewDecoder(resp.Body).Decode(&answersReply)
        return &answersReply
    } else {
        return nil
    }
}

func postQuestion(q dto.QuestionsPostReq) *dto.QuestionsPostRep {
    URL := fmt.Sprintf("%s/api/questions", serverAddr)
    fmt.Println("Push question to :", URL)
    buff := new(bytes.Buffer)
    json.NewEncoder(buff).Encode(q)
    resp, err := http.Post(URL, "application/json", buff)
    if err != nil {
        fmt.Println("Error", err)
        return nil
    }
    if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
        questionReply := dto.QuestionsPostRep{}
        json.NewDecoder(resp.Body).Decode(&questionReply)
        return &questionReply
    } else {
        return nil
    }
}

func getQuestions() *dto.QuestionsGetAllRep {
    URL := fmt.Sprintf("%s/api/questions", serverAddr)
    fmt.Println("Connecting to server:", URL)
    resp, err := http.Get(URL)
    if err != nil {
        // handle error
        fmt.Println("Error", err)
        return nil
    }
    if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
        questions := dto.QuestionsGetAllRep{}
        json.NewDecoder(resp.Body).Decode(&questions)
        return &questions
    } else {
        return nil
    }
}

// helpers

func scanAnswerNumber(max int) int {
    var i int
    fmt.Print("> ")
    _, err := fmt.Scanf("%d", &i)
    if err != nil || i < 0 || i >= max {
        fmt.Println("Please specify valid number")
        return scanAnswerNumber(max)
    } else {
        return i
    }
}
