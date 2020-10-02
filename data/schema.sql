-- type ChatMessage struct {
--   Timestamp  time.Time `json:"timestamp"`
--   ServerType string    `json:"server_type,omitempty"`
--   ServerID   string    `json:"server_id,omitempty"`
--   Server     string    `json:"server,omitempty"`
--   ChannelID  string    `json:"channel_id,omitempty"`
--   Channel    string    `json:"channel,omitempty"`
--   UserID     string    `json:"user_id"`
--   User       string    `json:"user"`
--   Message    string    `json:"message"`
-- }

CREATE TABLE IF NOT EXISTS logs (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  server_type VARCHAR(256) NOT NULL,
  server_id VARCHAR(256) NOT NULL,
  server VARCHAR(256) NOT NULL,
  channel_id VARCHAR(256) NOT NULL,
  channel VARCHAR(256) NOT NULL,
  user_id VARCHAR(256) NOT NULL,
  user VARCHAR(256) NOT NULL,
  message TEXT NOT NULL,
  timestamp DATETIME,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
