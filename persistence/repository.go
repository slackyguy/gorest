package persistence

import (
	"context"

	"github.com/slackyguy/gorest/base"
)

// Repository (data structure)
type Repository struct {
	AppSettings *base.AppSettings
	Context     context.Context
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
