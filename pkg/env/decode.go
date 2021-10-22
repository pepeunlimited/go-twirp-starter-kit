package env

import (
	"log"
	"os"
	"strconv"
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		log.Printf("ℹ️  For key=[%v], environment variable=[%v]", key, value)
		return value
	}
	log.Printf("⚠️  Invoked fallback for key=[%v], environment variable=[%v]",key, fallback)
	return fallback
}

func Int64Env(key string, fallback int64) int64 {
	env := GetEnv(key, strconv.FormatInt(fallback, 10))
	parsed, err := strconv.ParseInt(env, 10, 64)
	if err != nil {
		log.Panicf("❌  strconv.Atoi: %s @ Int64Env", err)
	}
	return parsed
}

func IntEnv(key string, fallback int) int {
	str := GetEnv(key, strconv.Itoa(fallback))
	parsed, err := strconv.Atoi(str)
	if err != nil {
		log.Panicf("❌  strconv.Atoi: %s @ IntEnv", err)
	}
	return parsed
}