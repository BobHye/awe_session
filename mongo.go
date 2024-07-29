package sessions

import (
	"github.com/gorilla/sessions"
	"github.com/laziercoder/mongostore"
	"go.mongodb.org/mongo-driver/mongo"
)

type AweMongoStore struct {
	*mongostore.MongoStore
}

func (c *AweMongoStore) Options(options AweOptions) {
	c.MongoStore.Options = &sessions.Options{
		Path:     options.Path,
		Domain:   options.Domain,
		MaxAge:   options.MaxAge,
		Secure:   options.Secure,
		HttpOnly: options.HttpOnly,
	}
}

var _ sessions.Store = (*AweMongoStore)(nil)

func NewMongoStore(c *mongo.Collection, maxAge int, ensureTTL bool, keyPairs ...[]byte) Store {
	return &AweMongoStore{mongostore.NewMongoStore(c, maxAge, ensureTTL, keyPairs...)}
}
