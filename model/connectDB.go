package model

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

type connectMySQL struct {
	conn *sql.DB
}

func (cm *connectMySQL) Conn() *sql.DB {
	return cm.conn
}

var lock = &sync.Mutex{}
var instance *connectMySQL

var user string = "root"
var pass string = ""
var host string = "localhost"
var port string = "3306"
var database string = "librarydb"

func InstanceDB() *connectMySQL {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()

		if instance == nil {
			strConnect := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + database
			conn, err := sql.Open("mysql", strConnect)
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
