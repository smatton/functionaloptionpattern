package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"strconv"
)

type Client struct {
	*http.Client
	host string
	port string
}

type clientOption func(c *Client) error

func newClientWithOptions(c *Client, opts ...clientOption) error {

	for _, option := range opts {
		err := option(c)
		if err != nil {
			return err
		}
	}
	return nil
}

func SetTLS(tls *tls.Config) clientOption {
	return func(c *Client) error {

		return nil
	}
}

func Host(host string) clientOption {
	return func(c *Client) error {
		c.host = host
		return nil
	}
}

func Port(port int) clientOption {
	return func(c *Client) error {
		c.port = strconv.Itoa(port)
		return nil
	}
}

func NewClient(opts ...clientOption) (*Client, error) {
	c := &Client{}

	// Set default port to OS randomly assign a port
	Port(0)(c)

	err := newClientWithOptions(c, opts...)
	if err != nil {
		return c, err
	}

	return c, nil

}
func main() {

	conn, err := NewClient(Host("localhost"))
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Print(conn.host, ":", conn.port)
}
