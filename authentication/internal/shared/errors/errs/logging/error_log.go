package logging

import (
	"log"
	"os"

	"github.com/sajad-dev/authservice/authentication/internal/config"
)

func CheckDebugTrueOrNot() {
	if !config.Cfg.Debug {
		file, err := os.OpenFile("../../storage/log/", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		log.SetOutput(file)
	}
}
func ErrLog(err error) {
	log.Println(err)
}
