package models

import "testing"

func TestAddMailSettings(t *testing.T) {
	sett := &MailSettings{
		Account:      "voidint@126.com",
		Pwd:          "123456",
		Outgoing:     "smtp.126.com",
		OutgoingPort: 25,
		UserId:       "7450b7fb-8145-11e4-beed-005056c00008",
	}
	_, err := AddMailSettings(sett)
	if err != nil {
		t.Fatal(err)
	}
}
