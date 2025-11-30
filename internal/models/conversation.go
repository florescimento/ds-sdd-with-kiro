package models

import "time"

// ConversationType defines the type of conversation
type ConversationType string

const (
	ConversationTypePrivate ConversationType = "private"
	ConversationTypeGroup   ConversationType = "group"
)

// Conversation represents a chat conversation
type Conversation struct {
	ID            string           `json:"id" bson:"_id"`
	Type          ConversationType `json:"type" bson:"type"`
	Members       []string         `json:"members" bson:"members"` // usernames
	CreatedAt     time.Time        `json:"created_at" bson:"created_at"`
	LastMessageAt time.Time        `json:"last_message_at" bson:"last_message_at"`
	Metadata      map[string]any   `json:"metadata" bson:"metadata"`
}
