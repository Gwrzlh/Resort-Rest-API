package config
 
 import(
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
 )

 func ConnectDB() *sql.DB {
	dbUser := "user"
	dbPassword := ""
	dbHost := "localhost"
	dbPort := "3306"
	dbName := "dapresto"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Database Connection Error", err)
	}

	fmt.Println("Database Connected Successfully")
	return db

 }