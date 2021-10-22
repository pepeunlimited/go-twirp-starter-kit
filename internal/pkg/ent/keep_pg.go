package ent


/*
import (
	"database/sql"
	"github.com/pepeunlimited/microservice-kit/misc"
	"github.com/pepeunlimited/microservice-kit/sqlz"
	"log"
	"strconv"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const databaseUser string 		= "root"
const databasePassword string 	= "root"
const databaseHost string 		= "localhost"
const databasePort string 		= "5432"
const databaseName string 		= "places_db"
const databaseSSLMode string 	= "false"

func NewEntClient() *Client {
	username 	:= misc.GetEnv(sqlz.PostgresqlUser, databaseUser)
	password 	:= misc.GetEnv(sqlz.PostgresqlPassword, databasePassword)
	host 		:= misc.GetEnv(sqlz.PostgresqlHost, databaseHost)
	port, _ 	:= strconv.Atoi(misc.GetEnv(sqlz.PostgresqlPort, databasePort))
	database 	:= misc.GetEnv(sqlz.PostgresqlDatabase, databaseName)
	sslmode,_ 	:= strconv.ParseBool(misc.GetEnv(sqlz.PostgresqlSslMode, databaseSSLMode))
	uri 		:= sqlz.PostgreSQLURI(username, password, host, port, database, sslmode)
	db, err := sql.Open("pgx", uri)
	if err != nil {
		log.Fatal(err)
	}
	drv := entsql.OpenDB(dialect.Postgres, db)
	return NewClient(Driver(drv))
}

// sometimes is needed to expose the DB for repository
func (c *Client) DB() *sql.DB {
	return c.driver.(*entsql.Driver).DB()
}
 */