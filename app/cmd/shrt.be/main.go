package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/pr0xity/go-url-short/app/internal/model"
	"github.com/pr0xity/go-url-short/app/internal/server"
)

func main() {
	godotenv.Load("../../../env/.env")
	app_host, app_port, db_host, db_port, username, password, database, timezone :=
		os.Getenv("APP_HOST"),
		os.Getenv("APP_PORT"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USERNAME"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DATABASE"),
		os.Getenv("POSTGRES_TIMEZONE")

	model.Setup(db_host, db_port, username, password, database, timezone)
	server.SetupAndListen(app_host, app_port)
}
