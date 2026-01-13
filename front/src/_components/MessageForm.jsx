import React, { useState } from 'react';
import PropTypes from 'prop-types';
import './MessageForm.css';

// Without a reverse-proxy, the browser must reach the API port directly.
// You can override with VITE_API_BASE_URL, otherwise default to :8080 on same host.
const API_BASE_URL =
  import.meta.env.VITE_API_BASE_URL ||
  `${window.location.protocol}//${window.location.hostname}:8080`;

function MessageForm({ onMessagePosted }) {
  const [formData, setFormData] = useState({
    pseudonym: '',
    content: '',
    avatar: '',
  });
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [error, setError] = useState(null);
  const [success, setSuccess] = useState(null);

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: value,
    }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    // Validation
    if (!formData.pseudonym.trim()) {
      setError('Le pseudonyme est requis');
      return;
    }

    if (!formData.content.trim()) {
      setError('Le contenu du message est requis');
      return;
    }

    if (formData.pseudonym.length > 50) {
      setError('Le pseudonyme ne peut pas dépasser 50 caractères');
      return;
    }

    if (formData.content.length > 1000) {
      setError('Le message ne peut pas dépasser 1000 caractères');
      return;
    }

    try {
      setIsSubmitting(true);
      setError(null);

      const response = await fetch(`${API_BASE_URL}/api/messages`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          pseudonym: formData.pseudonym.trim(),
          content: formData.content.trim(),
          avatar: formData.avatar.trim() || null,
        }),
      });

      if (!response.ok) {
        throw new Error("Erreur lors de l'envoi du message");
      }

      const newMessage = await response.json();

      // Show success message
      setSuccess('✅ Message posté avec succès !');

      // Reset form
      setFormData({
        pseudonym: '',
        content: '',
        avatar: '',
      });

      // Clear success message after 3 seconds
      setTimeout(() => setSuccess(null), 3000);

      // Notify parent component
      if (onMessagePosted) {
        onMessagePosted(newMessage);
      }
    } catch (err) {
      setError(err.message);
      console.error('Error posting message:', err);
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <div className="message-form animate-fade-in-up">
      <h2>✨ Poster un message</h2>
      <form onSubmit={handleSubmit} noValidate>
        <div className="form-group">
          <label htmlFor="pseudonym">👤 Pseudonyme *</label>
          <input
            type="text"
            id="pseudonym"
            name="pseudonym"
            value={formData.pseudonym}
            onChange={handleInputChange}
            placeholder="Votre pseudonyme anonyme"
            maxLength="50"
            required
          />
        </div>

        <div className="form-group">
          <label htmlFor="avatar">🖼️ Avatar (URL optionnel)</label>
          <input
            type="url"
            id="avatar"
            name="avatar"
            value={formData.avatar}
            onChange={handleInputChange}
            placeholder="https://example.com/avatar.jpg"
          />
        </div>

        <div className="form-group">
          <label htmlFor="content">💬 Message *</label>
          <textarea
            id="content"
            name="content"
            value={formData.content}
            onChange={handleInputChange}
            placeholder="Écrivez votre message ici..."
            maxLength="1000"
            rows="4"
            required
          />
          <div className="char-count">
            {formData.content.length}/1000 caractères
          </div>
        </div>

        {error && (
          <div className="error-message animate-bounce-in">{error}</div>
        )}

        {success && (
          <div className="success-message animate-bounce-in">{success}</div>
        )}

        <button
          type="submit"
          disabled={isSubmitting}
          className={`submit-button ${isSubmitting ? 'loading' : ''}`}
        >
          {isSubmitting ? '📤 Envoi en cours...' : '🚀 Poster le message'}
        </button>
      </form>
    </div>
  );
}

MessageForm.propTypes = {
  onMessagePosted: PropTypes.func,
};

export default MessageForm;
