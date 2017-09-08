package routes

import (
	"github.com/sh3rp/databox/auth"
	"github.com/sh3rp/databox/db"
	"github.com/sh3rp/databox/search"
)

type RouterBase struct {
	DB         db.BoxDB
	Search     search.SearchEngine
	Auth       auth.Authenticator
	TokenStore auth.TokenStore
}

func NewRouterBase(db db.BoxDB, search search.SearchEngine, auth auth.Authenticator, tokenStore auth.TokenStore) *RouterBase {
	return &RouterBase{
		DB:         db,
		Search:     search,
		Auth:       auth,
		TokenStore: tokenStore,
	}
}
