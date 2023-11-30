package kedai

import "github.com/google/uuid"

// Item adalah entitas yang mewakili sebuah item di semua domain
type Item struct {
	// ID adalah pengidentifikasi entitas itu
	ID          uuid.UUID
	Name        string
	Description string
}
