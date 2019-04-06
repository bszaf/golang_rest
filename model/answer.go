package model

import (
    "crypto/sha1"
    "encoding/hex"
)

type Answer struct {
    Text string
    Id string
}

func newAnswer(text string) (Answer) {
    idByte := sha1.Sum([]byte(text))
    idString := hex.EncodeToString(idByte[:])
    return Answer{Text: text, Id: idString}
}
