package schema

import (
    "context"
    "gorm.io/gorm"
)

type contextKey string

const dbContextKey = contextKey("db")

func WithDB(ctx context.Context, db *gorm.DB) context.Context {
    return context.WithValue(ctx, dbContextKey, db)
}

func GetDB(ctx context.Context) *gorm.DB {
    if db, ok := ctx.Value(dbContextKey).(*gorm.DB); ok {
        return db
    }
    return nil
}
