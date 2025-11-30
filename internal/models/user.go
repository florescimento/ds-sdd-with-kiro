package models

import "time"

// User represents a platform user
type User struct {
	Username        string           `json:"username" bson:"username"`
	PasswordHash    string           `json:"-" bson:"password_hash"`
	CreatedAt       time.Time        `json:"created_at" bson:"created_at"`
	ChannelMappings []ChannelMapping `json:"channel_mappings" bson:"channel_mappings"`
}

// ChannelMapping represents a user's connection to an external channel
type ChannelMapping struct {
	Channel    string `json:"channel" bson:"channel"`       // "whatsapp", "instagram", etc
	ExternalID string `json:"external_id" bson:"external_id"`
	Verified   bool   `json:"verified" bson:"verified"`
}
