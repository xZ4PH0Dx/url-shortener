package pg

import (
	"database/sql"
	"github.com/go-kit/kit/log"
	_ "github.com/lib/pq"
)

const (
	defaultMaxConnections = 5
)

type Client struct {
	DB             *sql.DB
	logger         log.Logger
	maxConnections int
}

func (c *Client) InitSchema() error {
	_, err := c.DB.Exec(Schema)
	return err
}

func (c *Client) Open(dataSourceName string) error {
	//c.logger.Log("level", "debug", "msg", "connecting to db")//Q:не работает, ругается на nil pointer:(
	var err error
	c.DB, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		return err
	}
	c.DB.SetMaxIdleConns(c.maxConnections)
	c.DB.SetMaxOpenConns(c.maxConnections)
	//c.logger.Log("level", "debug", "msg", "connected to db")
	return err
}

func (c *Client) Close() error {
	return c.DB.Close()
}
