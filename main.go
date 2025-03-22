package main

import (
	"fmt"
	"github.com/S4mkiel/finance-backend/adapter"
	"github.com/S4mkiel/finance-backend/application"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	if os.Getenv("ENV") != "production" {
		loadConfig()
	}

	fx.New(
		adapter.Module,
		application.Module,
	).Run()
}

func loadConfig() {
	_, b, _, _ := runtime.Caller(0)

	basepath := filepath.Dir(b)

	err := godotenv.Load(fmt.Sprintf("%v/.env", basepath))
	if err != nil {
		panic(err)
	}
}
