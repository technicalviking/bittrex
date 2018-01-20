package bittrex

import "fmt"

type bittrexError struct {
	location string
	msg      string
}

func (b *bittrexError) Error() string {

	return fmt.Sprintf("Error at location %s: %s", b.location, b.msg)
}

func (c *Client) setError(location string, msg string) {
	if c.err == nil || c.err.location == "" {
		c.err = &bittrexError{
			location,
			msg,
		}
	}
}

func (c *Client) clearError() {
	c.err = nil
}
