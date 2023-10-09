package model

import (
	"database/sql"
	"fmt"
	"sync"

	// _ "github.com/go-sql-driver/mysql"
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

// var user string = "root"
// var pass string = ""
// var host string = "localhost"
// var port string = "3306"
// var database string = "librarydb"

var user string = "fl0user"
var pass string = "8NcmbZuOk0zq"
var host string = "ep-patient-waterfall-23004259.us-east-2.aws.neon.fl0.io"
var port string = "5432"
var database string = "librarydb"

func InstanceDB() *connectMySQL {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()

		if instance == nil {

			connStr := "host=" + host + " port=" + port + " user=%" + user + " " + "password=" + pass + " dbname=" + database + " sslmode=enable"
			conn, err := sql.Open("postgres", connStr)
			if err != nil {
				fmt.Println(err)
			}

			// strConnect := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + database
			// conn, err := sql.Open("mysql", strConnect)
			// if err != nil {
			// 	fmt.Println(err)
			// }

			instance = &connectMySQL{
				conn: conn,
			}
		}
	}

	return instance
}
