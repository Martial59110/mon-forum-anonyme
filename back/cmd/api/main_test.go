package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Martial59110/mon-forum-anonyme/back/database"
)

func TestHandleHealth(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/health", nil)
	rr := httptest.NewRecorder()

	handleHealth(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rr.Code)
	}

	var payload map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &payload); err != nil {
		t.Fatalf("invalid json response: %v", err)
	}
	if payload["status"] != "ok" {
		t.Fatalf("expected status=ok, got %q", payload["status"])
	}
}

func TestHandleMessagesGET(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock.New() error: %v", err)
	}
	t.Cleanup(func() { _ = db.Close() })
	database.DB = db

	now := time.Date(2026, 1, 10, 12, 0, 0, 0, time.UTC)
	rows := sqlmock.NewRows([]string{"id", "pseudonym", "content", "avatar", "created_at", "updated_at"}).
		AddRow(1, "Alice", "Hello", "https://example.com/a.png", now, now)

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT id, pseudonym, content, avatar, created_at, updated_at 
			  FROM messages 
			  ORDER BY created_at DESC`,
	)).WillReturnRows(rows)

	req := httptest.NewRequest(http.MethodGet, "/api/messages", nil)
	rr := httptest.NewRecorder()

	handleMessages(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rr.Code)
	}

	var messages []map[string]any
	if err := json.Unmarshal(rr.Body.Bytes(), &messages); err != nil {
		t.Fatalf("invalid json response: %v", err)
	}
	if len(messages) != 1 {
		t.Fatalf("expected 1 message, got %d", len(messages))
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet SQL expectations: %v", err)
	}
}

func TestHandleMessagesPOSTValidation(t *testing.T) {
	body := bytes.NewBufferString(`{"pseudonym":"","content":""}`)
	req := httptest.NewRequest(http.MethodPost, "/api/messages", body)
	rr := httptest.NewRecorder()

	handleMessages(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, rr.Code)
	}
}

func TestHandleMessagesPOSTSuccess(t *testing.T) {
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

	body := bytes.NewBufferString(`{"pseudonym":"Alice","content":"Hello","avatar":"https://example.com/a.png"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/messages", body)
	rr := httptest.NewRecorder()

	handleMessages(rr, req)

	if rr.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, rr.Code)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet SQL expectations: %v", err)
	}
}
