package db

import "context"

type Repositoty interface{
	Close()
	InsertMeow(ctx context.Context, meow schema.Meow)
}