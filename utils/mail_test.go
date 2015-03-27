package utils

import (
	"testing"
	"time"
)

func TestCheckAuth(t *testing.T) {
	host := "smtp.126.com"
	port := 25
	email := "voidint2@126.com"
	pwd := ""

	err := CheckAuth(host, uint(port), email, pwd, time.Second*5)
	if err != nil {
		t.Log(err)
	}

}
