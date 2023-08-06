package database

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DBConnection struct {
	*sql.DB
}

func ConnectDB() (*sql.DB, error) {
	// connection := os.Getenv("DB_CONNECTION")
	// host := os.Getenv("DB_HOST")
	// database := os.Getenv("DB_DATABASE")
	// username := os.Getenv("DB_USERNAME")
	// password := os.Getenv("DB_PASSWORD")

	// dataSource := fmt.Sprintf("%s:%s@(%s)/%s?parseTime=true", username, password, host, database)

	dataSource := "root:password@tcp(go-mysql)/example?parseTime=true"

	var db *sql.DB
	var err error
	maxRetries := 5
	retryInterval := 1 * time.Second

	for i := 0; i < maxRetries; i++ {
		db, err = sql.Open("mysql", dataSource)
		if err == nil {
			break
		}

		log.Printf("データベースの接続に失敗しました。リトライします... (エラー: %v)", err)

		time.Sleep(retryInterval)
	}

	if err != nil {
		log.Print("リトライ後もデータベースの接続に失敗しました。")
		return nil, err
	}

	return db, nil
}
