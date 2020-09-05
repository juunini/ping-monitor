package lib

import "log"

func Logging(message ...interface{}) {
	log.Println(message...)
}