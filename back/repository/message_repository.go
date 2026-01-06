package repository

import (
	"database/sql"
	"time"

	"github.com/Martial59110/mon-forum-anonyme/back/database"
	"github.com/Martial59110/mon-forum-anonyme/back/models"
)

func GetAllMessages() ([]models.Message, error) {
	query := `SELECT id, pseudonym, content, avatar, created_at, updated_at 
			  FROM messages 
			  ORDER BY created_at DESC`

	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var msg models.Message
		err := rows.Scan(&msg.ID, &msg.Pseudonym, &msg.Content, &msg.Avatar, &msg.CreatedAt, &msg.UpdatedAt)
		if err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}

	return messages, rows.Err()
}

func CreateMessage(pseudonym, content, avatar string) (*models.Message, error) {
	query := `INSERT INTO messages (pseudonym, content, avatar, created_at, updated_at) 
			  VALUES ($1, $2, $3, $4, $5) 
			  RETURNING id, pseudonym, content, avatar, created_at, updated_at`

	now := time.Now()
	var msg models.Message

	err := database.DB.QueryRow(query, pseudonym, content, avatar, now, now).Scan(
		&msg.ID, &msg.Pseudonym, &msg.Content, &msg.Avatar, &msg.CreatedAt, &msg.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &msg, nil
}

func GetMessageByID(id int) (*models.Message, error) {
	query := `SELECT id, pseudonym, content, avatar, created_at, updated_at 
			  FROM messages 
			  WHERE id = $1`

	var msg models.Message
	err := database.DB.QueryRow(query, id).Scan(
		&msg.ID, &msg.Pseudonym, &msg.Content, &msg.Avatar, &msg.CreatedAt, &msg.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &msg, nil
}
