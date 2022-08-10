package config

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ichtrojan/thoth"
	_ "github.com/joho/godotenv/autoload"
)

func Database() *sql.DB {
	logger, _ := thoth.Init("log")

	// dbname, exist := os.LookupEnv("DB_NAME")

	// if !exist {
	// 	logger.Log(errors.New("DB_NAME not set in .env"))
	// 	log.Fatal("DB_NAME not set in .env")
	// }

	user, exist := os.LookupEnv("DB_USER")

	if !exist {
		logger.Log(errors.New("DB_USER not set in .env"))
		log.Fatal("DB_USER not set in .env")
	}

	pass, exist := os.LookupEnv("DB_PASS")

	if !exist {
		logger.Log(errors.New("DB_PASS not set in .env"))
		log.Fatal("DB_PASS not set in .env")
	}

	host, exist := os.LookupEnv("DB_HOST")

	if !exist {
		logger.Log(errors.New("DB_HOST not set in .env"))
		log.Fatal("DB_HOST not set in .env")
	}

	// dbport, exist := os.LookupEnv("DB_PORT")

	// if !exist {
	// 	logger.Log(errors.New("DB_PORT not set in .env"))
	// 	log.Fatal("DB_PORT not set in .env")
	// }

	credentials := fmt.Sprintf("%s:%s@(%s:3306)/?charset=utf8&parseTime=True", user, pass, host)

	database, err := sql.Open("mysql", credentials)

	if err != nil {
		logger.Log(err)
		log.Fatal(err)
	} else {
		fmt.Println("Database Connection Successful")
	}

	_, err = database.Exec(`CREATE DATABASE IF NOT EXISTS gotask_db`)

	if err != nil {
		fmt.Println(err)
	}

	_, err = database.Exec(`USE gotask_db`)

	if err != nil {
		fmt.Println(err)
	}

	_, err = database.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
		    id INT AUTO_INCREMENT PRIMARY KEY,
		    name VARCHAR(255) NOT NULL,
		    description TEXT NOT NULL,
		    employee VARCHAR(255) NOT NULL,
		    deadline DATETIME NOT NULL,
		    status VARCHAR(50) NOT NULL DEFAULT 'In Progress',
		    completed BOOLEAN DEFAULT FALSE,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
			updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP 
		);
	`)

	if err != nil {
		fmt.Println(err)
	}

	return database
}
