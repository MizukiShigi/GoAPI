package Utilities

import (
	"os"
	"github.com/joho/godotenv"
	"database/sql"
	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	// 環境変数の取得
	err := godotenv.Load()
	CheckErr(err)

	db, err := sql.Open("postgres", "host=" + os.Getenv("DB_HOST") + " port=5432" + " user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASS") + " dbname=" + os.Getenv("DB_NAME") + " sslmode=disable")
	CheckErr(err)

	return db
}

func CheckErr(err error) {
    if err != nil {
        panic(err)
    }
}