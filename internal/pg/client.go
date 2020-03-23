package pg

import (
	"github.com/go-kit/kit/log"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	defaultMaxConnections = 5
)

type Client struct {
	DB             *sqlx.DB
	logger         log.Logger
	maxConnections int
}

func (c *Client) InitSchema() error {
	_, err := c.DB.Exec(Schema)
	return err
}

func (c *Client) Open(dataSourceName string) error {
	_ = c.logger.Log("level", "debug", "msg", "connecting to db") //Q:не работает, ругается на nil pointer:(
	var err error
	c.DB, err = sqlx.Open("postgres", dataSourceName)
	if err != nil {
		return err
	}
	c.DB.SetMaxIdleConns(c.maxConnections)
	c.DB.SetMaxOpenConns(c.maxConnections)
	_ = c.logger.Log("level", "debug", "msg", "connected to db")
	return err
}

func (c *Client) Close() error {
	return c.DB.Close()
}
