package jwt

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

var JwtKey = []byte(goDotEnvVariable("JWTKey"))
var CSRFKey = []byte(goDotEnvVariable("CSRFKey"))
