package models

import "time"

// FileMetadata represents metadata for an uploaded file
type FileMetadata struct {
	ID         string         `json:"id" bson:"_id"`
	Filename   string         `json:"filename" bson:"filename"`
	Size       int64          `json:"size" bson:"size"`
	Checksum   string         `json:"checksum" bson:"checksum"`
	StorageURL string         `json:"storage_url" bson:"storage_url"`
	UploadedBy string         `json:"uploaded_by" bson:"uploaded_by"`
	UploadedAt time.Time      `json:"uploaded_at" bson:"uploaded_at"`
	Status     string         `json:"status" bson:"status"` // "uploading", "completed"
	Metadata   map[string]any `json:"metadata" bson:"metadata"`
}
