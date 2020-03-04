package pg

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type Client struct {
	DB             *sql.DB
	logger         log.Logger
	maxConnections int
}

func (c *Client) Open(dataSourceName string) error {
	//c.logger.Printf("Trying to connect to PostgreSQL db with params: %v", dataSourceName)
	//fmt.Printf("Trying to connect to PostgreSQL db with params: %v", dataSourceName)
	var err error //Q: если использовать строкой ниже :=, то выдает ",', ';', <assign op>, new line or '}' expected, got ':='"
	c.DB, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		//c.logger.Fatal(err)
		c.logger.Printf("Error %v during connection to PostgreSQL db with params: %v", err, dataSourceName)
	}
	return nil
}

func (c *Client) Close() error {
	err := c.DB.Close()
	if err != nil {
		return error(err)
	}
	return nil
}
