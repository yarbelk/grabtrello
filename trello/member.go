package trello

import (
	mytrello "github.com/VojtechVitek/go-trello"
)

func Member(username, key string, token *string) (*mytrello.Member, error) {
	conn, err := Conn(key, token)
	if err != nil {
		return nil, err
	}
	return conn.Member(username)
}
