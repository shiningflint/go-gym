package models

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type ChatMessage struct {
	Id         int `json:"-"`
	ChatId     int `json:"-"`
	Message    string
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
	TimeString string
	NickName   string
	Color      string
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
		err = rows.Scan(
			&msg.Id,
			&msg.Message,
			&msg.CreatedAt,
			&msg.UpdatedAt,
			&msg.ChatId,
			&msg.NickName,
			&msg.Color,
		)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		msg.TimeString = msg.CreatedAt.Format("15:04")
		msgs = append(msgs, msg)
	}

	return msgs, nil
}

func SaveChatMessage(jsonString []byte) (*ChatMessage, error) {
	var data map[string]interface{}
	msg := new(ChatMessage)
	if err := json.Unmarshal(jsonString, &data); err != nil {
		return nil, err
	}
	sqs := `
	INSERT INTO chat_messages (message, chat_id, nickname, color)
	VALUES ($1, $2, $3, $4)
	RETURNING *`
	err := db.QueryRow(
		sqs,
		data["message"].(string),
		1,
		data["nickname"].(string),
		data["color"].(string),
	).
		Scan(
			&msg.Id,
			&msg.Message,
			&msg.CreatedAt,
			&msg.UpdatedAt,
			&msg.ChatId,
			&msg.NickName,
			&msg.Color,
		)
	if err != nil {
		return nil, err
	}
	msg.TimeString = msg.CreatedAt.Format("15:04")
	return msg, nil
}
