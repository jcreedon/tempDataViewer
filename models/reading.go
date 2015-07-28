package models

import (
	"database/sql"
	"time"
)

type Reading struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"time"`
	Source    Node
	SourceID  sql.NullInt64 `json:"sourceId"`
	Value     float32       `json:"value"`
}
