package config

// "github.com/Grs2080w/grp_server/core/config"

import (
	"log"
	"os"
	"strconv"
	"github.com/joho/godotenv"
)


func GetValueEnv(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar .env", err)
	}
	return os.Getenv(key)
}

func GetValueEnvInt(key string) int {

	value := GetValueEnv(key)
	envInt, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("Error on convert %s to int: %v", value, err)
	}

	return envInt
}

