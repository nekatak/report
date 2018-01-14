// Package db is responsible for retrieving test data from psql
// Test Data from suade looked like this:
// id   | 1
// data | {"organization":"Dunder Mifflin","reported_at":"2015-04-21","created_at":"2015-04-22",
// "inventory":[
// {"name":"paper","price":"2.00"},{"name":"stapler","price":"5.00"},
// {"name":"printer","price":"125.00"},{"name":"ink","price":"3000.00"}
// ]}
// id   | 2
// data | {"organization":"Skynet Papercorp","reported_at":"2015-04-22","created_at":"2015-04-23",
// "inventory":[{"name":"paper","price":"4.00"}]}
// with the supplied psql connection string pointing to Azure cloud psql host
// (suadecandidate.northeurope.cloudapp.azure.com has address 52.169.204.248).

package db

import (
	"encoding/json"
	"fmt"
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/astaxie/beego/logs"

	"report/generator"
	"log"
)

type Conf struct {
	Host string
	User string
	Password string
	Database string
	SSLMode string
}

type Client struct {
	db *sql.DB
}

type ClientInterface interface {
	GetData() (generator.Data, error)
}

//// Setting both clients to fill interface
var _ ClientInterface = &Client{}
//var _ ClientInterface = &MockClient{}

// Connect to the psql db and return client with sql.DB in it
// on which you perform queries
func Connect(conf Conf) (*Client, error){
	connectionString := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s",
					conf.User, conf.Password, conf.Host,
					conf.Database, conf.SSLMode)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		logs.Error("Could not connect to database", err.Error())
		return nil, err
	}

	var client = &Client{
		db: db,
	}

	return client, err
}

// GetData returns only id 1 every time as the data struct will become complicated
// and there was no requirement to show all data in db.
func (c *Client) GetData() (d generator.Data, err error) {
	query := `select data from reports where id = 1`

	var str string
	err = c.db.QueryRow(query).Scan(&str)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	err = json.Unmarshal([]byte(str), &d)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	return
}
