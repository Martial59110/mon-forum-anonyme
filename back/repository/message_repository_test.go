package repository

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Martial59110/mon-forum-anonyme/back/database"
)

func TestGetAllMessages(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock.New() error: %v", err)
	}
	t.Cleanup(func() { _ = db.Close() })
	database.DB = db

	now := time.Date(2026, 1, 10, 12, 0, 0, 0, time.UTC)
	rows := sqlmock.NewRows([]string{"id", "pseudonym", "content", "avatar", "created_at", "updated_at"}).
		AddRow(1, "Alice", "Hello", "https://example.com/a.png", now, now).
		AddRow(2, "Bob", "World", nil, now, now)

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT id, pseudonym, content, avatar, created_at, updated_at 
			  FROM messages 
			  ORDER BY created_at DESC`,
	)).WillReturnRows(rows)

	messages, err := GetAllMessages()
	if err != nil {
		t.Fatalf("GetAllMessages() error: %v", err)
	}

	if len(messages) != 2 {
		t.Fatalf("expected 2 messages, got %d", len(messages))
	}
	if messages[0].ID != 1 || messages[0].Pseudonym != "Alice" || messages[0].Content != "Hello" {
		t.Fatalf("unexpected first message: %+v", messages[0])
	}
	if messages[1].ID != 2 || messages[1].Pseudonym != "Bob" || messages[1].Content != "World" {
		t.Fatalf("unexpected second message: %+v", messages[1])
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet SQL expectations: %v", err)
	}
}

func TestCreateMessage(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock.New() error: %v", err)
	}
	t.Cleanup(func() { _ = db.Close() })
	database.DB = db

	now := time.Date(2026, 1, 10, 12, 0, 0, 0, time.UTC)

	mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO messages (pseudonym, content, avatar, created_at, updated_at) 
			  VALUES ($1, $2, $3, $4, $5) 
			  RETURNING id, pseudonym, content, avatar, created_at, updated_at`,
	)).
		WithArgs("Alice", "Hello", "https://example.com/a.png", sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "pseudonym", "content", "avatar", "created_at", "updated_at"}).
				AddRow(1, "Alice", "Hello", "https://example.com/a.png", now, now),
		)

	msg, err := CreateMessage("Alice", "Hello", "https://example.com/a.png")
	if err != nil {
		t.Fatalf("CreateMessage() error: %v", err)
	}
	if msg == nil {
		t.Fatal("expected non-nil message")
	}
	if msg.ID != 1 || msg.Pseudonym != "Alice" || msg.Content != "Hello" {
		t.Fatalf("unexpected message: %+v", msg)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet SQL expectations: %v", err)
	}
}

