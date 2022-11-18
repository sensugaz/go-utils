package required

import (
	"log"
	"os"
)

func GetEnv(env string) string {
	e := os.Getenv(env)
	if len(e) == 0 {
		log.Panicf("[GetEnv]: unable to get value from %s", env)
	}

	return e
}
