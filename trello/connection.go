package trello

import mytrello "github.com/VojtechVitek/go-trello"

func Conn(key string, token *string) (*mytrello.Client, error) {
	conn, err := mytrello.NewAuthClient(key, token)
	return conn, err
}
