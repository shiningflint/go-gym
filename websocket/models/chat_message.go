package models

import (
	"fmt"
	"log"
)

type ChatMessage struct {
	Id        int
	ChatId    int
	Message   string
	CreatedAt string
	UpdatedAt string
}

// AllChatMessages returns all chat messages
// from chat_messages table of one chat
func AllChatMessages(id int) ([]*ChatMessage, error) {
	rows, err := db.Query(
		fmt.Sprintf("SELECT * FROM chat_messages WHERE chat_id = %d", id),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var msgs []*ChatMessage
	for rows.Next() {
		msg := new(ChatMessage)
		err = rows.Scan(&msg.Id, &msg.Message, &msg.CreatedAt, &msg.UpdatedAt, &msg.ChatId)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		msgs = append(msgs, msg)
	}

	return msgs, nil
}
