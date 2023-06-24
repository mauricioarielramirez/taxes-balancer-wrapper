package models

import "time"

type SnapshotProduct struct {
	Product Product
	Data    time.Time
}
