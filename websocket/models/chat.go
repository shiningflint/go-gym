package models

import (
	"fmt"
	"log"
)

type Chat struct {
	Id          int
	ChannelName string
	CreatedAt   string
	UpdatedAt   string
}

func GetChat(id int) (*Chat, error) {
	chat := new(Chat)
	row := db.QueryRow(fmt.Sprintf("SELECT * FROM chats WHERE id = %d", id))
	err := row.Scan(&chat.Id, &chat.ChannelName, &chat.CreatedAt, &chat.UpdatedAt)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return chat, nil
}
