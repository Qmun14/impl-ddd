package kedai

import (
	"time"

	"github.com/google/uuid"
)

// Transaction adalah value object karena tidak memiliki pengenal/identifier dan tidak dapat diubah (Imutable)
type Transaction struct {
	amount    int
	from      uuid.UUID
	to        uuid.UUID
	createdAt time.Time
}
