package gorm

import "gorm.io/gen"

type Query interface {
	// SELECT * FROM @@table WHERE id = @id LIMIT 1
	GetByID(id uint64) gen.T
}
