package mail

import (
	"fmt"
	"net/smtp"
	"sync"

	"github.com/sajad-dev/authservice/authentication/internal/config"
	"github.com/sajad-dev/authservice/authentication/internal/shared/errors/errs/logging"
	"github.com/sajad-dev/authservice/authentication/internal/shared/job"
)

var (
	smtpConn smtp.Auth
	once     sync.Once
)

func _connectSMTP() {
	smtpConn = smtp.PlainAuth("", config.Cfg.Mail.Email, config.Cfg.Mail.Password, config.Cfg.Mail.Host)
}

func Send(params ...any) {
	once.Do(func() {
		_connectSMTP()
	})

	err := smtp.SendMail(fmt.Sprintf("%s:%d", config.Cfg.Mail.Host, config.Cfg.Mail.Port), smtpConn, config.Cfg.Mail.Email, []string{params[0].(string)}, []byte(params[2].(string)))
	if err != nil {
		logging.ErrLog(err)
	}
}

func (m *Mail) AddJob(jb job.Worker) error {
	jb.Add(job.NewTask(
		job.WithParams(m.SendTo, m.Title, m.MessageHTML, m.SendAt),
		job.WithRun(Send),
		job.WithRunAt(m.SendAt),
	))
	return nil
}
