package constants

const (
	EMAIL    = "email"
	USERNAME = "username"
	SMS   = "sms"
)

const (
	EMAILREGEX    = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	PHONEREGEX    = `^\+[1-9]\d{1,14}$`
	USERNAMEREGEX = `^[a-zA-Z0-9\-]{3,20}$`
)
