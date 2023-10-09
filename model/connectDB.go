package model

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type connectMySQL struct {
	conn *sql.DB
}

func (cm *connectMySQL) Conn() *sql.DB {
	return cm.conn
}

var lock = &sync.Mutex{}
var instance *connectMySQL

func InstanceDB() *connectMySQL {

	godotenv.Load(".env")

	var user string = os.Getenv("DB_USER")
	var pass string = os.Getenv("DB_PASS")
	var host string = os.Getenv("DB_HOST")
	var port string = os.Getenv("DB_PORT")
	var name string = os.Getenv("DB_NAME")

	if instance == nil {
		lock.Lock()
		defer lock.Unlock()

		if instance == nil {

			connStr := "host=" + host + " port=" + port + " user=" + user + " " + "password=" + pass + " dbname=" + name
			conn, err := sql.Open("postgres", connStr)
			if err != nil {
				fmt.Println(err)
			}

			instance = &connectMySQL{
				conn: conn,
			}
		}
	}

	return instance
}
