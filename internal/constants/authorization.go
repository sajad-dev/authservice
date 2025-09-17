package constants

type Permission string

const (
	SUPERUSER     Permission = "superuser"
	CHAT_SUPPOERT Permission = "chatsupport"
	TECH_SUPPORT  Permission = "techsupport"
	MONITOR       Permission = "monitor"
)
