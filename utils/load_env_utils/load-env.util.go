package load_env_utils

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	envFilePath := filepath.Join(basepath, "../../.env")

	if err := godotenv.Load(envFilePath); err != nil {
		log.Fatal("Error Loading .env File", err)
	}
}
