package persistence

import (
	"context"
)

// Repository (data structure)
type Repository struct {
	Context context.Context
}

// Interface (repository contract)
type Interface interface {
	SetCollectionName(collection string)
	Find(key string, item interface{})
	List(item interface{})
	Create(item interface{}) string
	Update(key string, item interface{})
	Delete(key string)
}
