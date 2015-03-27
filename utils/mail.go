package utils

import (
	"errors"
	"fmt"
	"net/smtp"
	"time"
)

var (
	ErrTimeout = errors.New("timeout")
)

// CheckAuth 检查指定邮箱服务商的邮箱账号和密码是否匹配。若匹配成功，返回nil，否则返回error。
func CheckAuth(smtpHost string, smtpPort uint, email, pwd string, timeout time.Duration) error {
	doneChan := make(chan bool)
	errChan := make(chan error)

	// 异步检测邮箱帐号和密码
	go func(host string, port uint, username, userpwd string, doneCh chan<- bool, errCh chan<- error) {
		var closeChan = func() {
			close(errCh)
			close(doneCh)
		}

		client, err := smtp.Dial(fmt.Sprintf("%s:%d", host, port))
		if err != nil {
			errCh <- err
			closeChan()
			return
		}

		err = client.Auth(smtp.PlainAuth("", username, userpwd, host))
		if err != nil {
			errCh <- err
			closeChan()
			return
		}

		doneCh <- true
		closeChan()

	}(smtpHost, smtpPort, email, pwd, doneChan, errChan)

	select {
	case <-time.After(timeout):
		return ErrTimeout
	case <-doneChan:
		return nil
	case err := <-errChan:
		return err
	}

}
