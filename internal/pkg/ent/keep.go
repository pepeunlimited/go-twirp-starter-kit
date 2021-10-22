package ent

/*
import (
	"database/sql"
	"entgo.io/ent/dialect"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pepeunlimited/microservice-kit/misc"
	"github.com/pepeunlimited/microservice-kit/sqlz"
	"log"
	"strconv"

	entsql "entgo.io/ent/dialect/sql"
)

const databaseUser string 		= "root"
const databasePassword string 	= "root"
const databaseHost string 		= "localhost"
const databasePort string 		= "3306"
const databaseName string 		= "helloWorld"

func NewEntClient() *Client {
	username 	:= misc.GetEnv(sqlz.MysqlUser, databaseUser)
	password 	:= misc.GetEnv(sqlz.MysqlRootPassword, databasePassword)
	host 		:= misc.GetEnv(sqlz.MysqlHost, databaseHost)
	port, _ 	:= strconv.Atoi(misc.GetEnv(sqlz.MysqlPort, databasePort))
	database 	:= misc.GetEnv(sqlz.MysqlDatabase, databaseName)
	uri 		:= sqlz.MySQLURI(username, password, host, port, database)
	client, err := Open(dialect.MySQL, uri)
	if err != nil {
		log.Panic("failed to open MySQL connection err: " + err.Error())
	}
	return client
}

// sometimes is needed to expose the DB for repository
func (c *Client) DB() *sql.DB {
	return c.driver.(*entsql.Driver).DB()
}
*/