import React, { useState, useEffect } from 'react';
import './MessageList.css';

const API_BASE_URL = 'http://localhost:8080';

function MessageList() {
  const [messages, setMessages] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    fetchMessages();
  }, []);

  const fetchMessages = async () => {
    try {
      setLoading(true);
      const response = await fetch(`${API_BASE_URL}/api/messages`);
      if (!response.ok) {
        throw new Error('Failed to fetch messages');
      }
      const data = await response.json();
      // Ensure data is an array
      setMessages(Array.isArray(data) ? data : []);
      setError(null);
    } catch (err) {
      setError(err.message);
      console.error('Error fetching messages:', err);
      setMessages([]); // Reset to empty array on error
    } finally {
      setLoading(false);
    }
  };

  const formatDate = (dateString) => {
    const date = new Date(dateString);
    return date.toLocaleString('fr-FR', {
      year: 'numeric',
      month: 'short',
      day: 'numeric',
      hour: '2-digit',
      minute: '2-digit',
    });
  };

  if (loading) {
    return (
      <div className="message-list-loading animate-pulse">
        ⏳ Chargement des messages...
      </div>
    );
  }

  if (error) {
    return (
      <div className="message-list-error animate-bounce-in">
        🚫 Erreur lors du chargement: {error}
        <button onClick={fetchMessages} className="retry-button">
          🔄 Réessayer
        </button>
      </div>
    );
  }

  return (
    <div className="message-list animate-fade-in-up">
      <h2>📝 Messages du forum</h2>
      {!Array.isArray(messages) || messages.length === 0 ? (
        <p className="no-messages animate-bounce-in">
          📭 Aucun message pour le moment. Soyez le premier à poster ! ✨
        </p>
      ) : (
        <div className="messages-container">
          {messages.map((message) => (
            <div key={message.id} className="message-card animate-fade-in">
              <div className="message-header">
                <span className="pseudonym">{message.pseudonym}</span>
                <span className="timestamp">
                  {formatDate(message.created_at)}
                </span>
              </div>
              {message.avatar && (
                <div className="avatar">
                  <img
                    src={message.avatar}
                    alt={`Avatar de ${message.pseudonym}`}
                  />
                </div>
              )}
              <div className="message-content">{message.content}</div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
}

export default MessageList;
