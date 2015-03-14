package models

import (
	"testing"
	"time"
)

func TestGetContactMsgPage(t *testing.T) {
	cond := &ContactMessage{
		Name:  "voidint@126.com",
		Email: "voidint@126.com",
	}
	page, err := GetContactMsgPage(cond, 1, 5)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v\n", page)
}

func TestSearchContactMsgPage(t *testing.T) {
	cond := &ContactMessage{}
	begin := time.Now()
	end := time.Now()
	op := LogicalOp("OR")
	page, err := SearchContactMsgPage(cond, begin, end, op, 1, 10)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v\n", page)
}
