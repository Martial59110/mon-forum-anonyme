package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Martial59110/mon-forum-anonyme/back/database"
	"github.com/Martial59110/mon-forum-anonyme/back/repository"
)

func initDBWithRetry(maxRetries int, retryDelay time.Duration) (*sql.DB, error) {
	for i := 0; i < maxRetries; i++ {
		db, err := database.InitDB()
		if err == nil {
			return db, nil
		}

		log.Printf("Tentative %d/%d - Erreur de connexion à la base de données: %v", i+1, maxRetries, err)

		if i < maxRetries-1 {
			log.Printf("Nouvelle tentative dans %v...", retryDelay)
			time.Sleep(retryDelay)
		}
	}

	return database.InitDB()
}

func main() {

	log.Println("Connexion à la base de données...")
	db, err := initDBWithRetry(60, 2*time.Second)
	if err != nil {
		log.Fatal("Erreur de connexion à la base de données après plusieurs tentatives:", err)
	}
	defer db.Close()
	log.Println("Connexion à la base de données établie avec succès")

	if err := database.CreateTable(); err != nil {
		log.Fatal("Erreur lors de la création de la table:", err)
	}
	log.Println("Table des messages créée/vérifiée")

	http.HandleFunc("/api/messages", handleMessages)
	http.HandleFunc("/api/health", handleHealth)

	port := getEnv("API_PORT", "8080")
	fmt.Printf("Serveur démarré sur le port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func handleMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	switch r.Method {
	case "GET":
		messages, err := repository.GetAllMessages()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(messages)

	case "POST":
		var req struct {
			Pseudonym string `json:"pseudonym"`
			Content   string `json:"content"`
			Avatar    string `json:"avatar"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		if req.Pseudonym == "" || req.Content == "" {
			http.Error(w, "Pseudonym and content are required", http.StatusBadRequest)
			return
		}

		message, err := repository.CreateMessage(req.Pseudonym, req.Content, req.Avatar)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(message)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
