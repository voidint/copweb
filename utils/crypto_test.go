package utils

import (
	"encoding/base64"
	"testing"
)

func TestAes(t *testing.T) {
	// AES-128。key长度：16, 24, 32 bytes 对应 AES-128, AES-192, AES-256
	key := []byte("Sdjftmoil45df#$DGJDI^ku26XXkdDJK")
	result, err := AesEncrypt([]byte("我爱北京天安门"), key)
	if err != nil {
		t.Fatal(err)
	}

	// t.Logf("%s\n", result)
	sEnc := base64.StdEncoding.EncodeToString(result)
	t.Logf("%s\n", sEnc)

	bRaw, err := base64.StdEncoding.DecodeString(sEnc)
	if err != nil {
		t.Fatal(err)
	}
	origData, err := AesDecrypt(bRaw, key)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s\n", origData)
}
