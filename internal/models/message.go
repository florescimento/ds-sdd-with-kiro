package models

import "time"

// MessageStatus represents the delivery status of a message
type MessageStatus string

const (
	MessageStatusSent      MessageStatus = "SENT"
	MessageStatusDelivered MessageStatus = "DELIVERED"
	MessageStatusRead      MessageStatus = "READ"
)

// MessagePayloadType defines the type of message content
type MessagePayloadType string

const (
	MessagePayloadTypeText MessagePayloadType = "text"
	MessagePayloadTypeFile MessagePayloadType = "file"
)

// Message represents a chat message
type Message struct {
	ID             string               `json:"id" bson:"_id"` // message_id (UUID)
	ConversationID string               `json:"conversation_id" bson:"conversation_id"`
	From           string               `json:"from" bson:"from"` // username
	To             []string             `json:"to" bson:"to"`     // usernames
	Payload        MessagePayload       `json:"payload" bson:"payload"`
	Channels       []string             `json:"channels" bson:"channels"`
	Status         MessageStatus        `json:"status" bson:"status"`
	StatusHistory  []StatusHistoryEntry `json:"status_history" bson:"status_history"`
	SequenceNumber int64                `json:"sequence_number" bson:"sequence_number"`
	CreatedAt      time.Time            `json:"created_at" bson:"created_at"`
	Metadata       map[string]any       `json:"metadata" bson:"metadata"`
}

// MessagePayload contains the actual message content
type MessagePayload struct {
	Type   MessagePayloadType `json:"type" bson:"type"`
	Text   string             `json:"text,omitempty" bson:"text,omitempty"`
	FileID string             `json:"file_id,omitempty" bson:"file_id,omitempty"`
}

// StatusHistoryEntry tracks status changes
type StatusHistoryEntry struct {
	Status    MessageStatus `json:"status" bson:"status"`
	Timestamp time.Time     `json:"timestamp" bson:"timestamp"`
}
