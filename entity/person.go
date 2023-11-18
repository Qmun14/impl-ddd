// Package entities ini menampung semua entitas yang dibagikan di seluruh sub domain
package entity

import "github.com/google/uuid"

// Person adalah entitas yang mewakili seseorang di semua domain
type Person struct {
	// ID adalah pengidentifikasi entitas itu
	ID   uuid.UUID
	Name string
	Age  int
}
