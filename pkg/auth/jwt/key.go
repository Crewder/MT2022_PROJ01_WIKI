package jwt

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GoDotEnvVariable(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

var JwtKey = []byte(GoDotEnvVariable("JWTKey"))
var CSRFKey = GoDotEnvVariable("CSRFKey")
