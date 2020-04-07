package zuo

import "time"

type Token struct {
	Val    string
	Expiry time.Time
}

func NewToken() *Token {
	t := &Token{Val: "abc", Expiry: time.Now()}
	return t
}
