package mail

import "time"

type Mail struct {
	SendTo      string
	Title       string
	MessageHTML string
	SendAt      time.Time
}

type MailOption func(*Mail)

func NewMail(opts ...MailOption) *Mail {
	m := &Mail{}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

func WithSendTo(sendTo string) MailOption {
	return func(m *Mail) {
		m.SendTo = sendTo
	}
}

func WithTitle(title string) MailOption {
	return func(m *Mail) {
		m.Title = title
	}
}

func WithMessageHTML(messageHTML string) MailOption {
	return func(m *Mail) {
		m.MessageHTML = messageHTML
	}
}

func WithSendAt(sendAt time.Time) MailOption {
	return func(m *Mail) {
		m.SendAt = sendAt
	}
}
