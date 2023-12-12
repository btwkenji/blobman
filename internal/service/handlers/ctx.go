package handlers

import (
	"context"
	"github.com/btwkenji/blobman/internal/data"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	blobQCtxKey
	ownerQCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func BlobQ(r *http.Request) data.Blobs {
	return r.Context().Value(blobQCtxKey).(data.Blobs).New()
}

func CtxBlobQ(b data.Blobs) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, blobQCtxKey, b)
	}
}

func OwnerQ(r *http.Request) data.Owners {
	return r.Context().Value(ownerQCtxKey).(data.Owners).New()
}

func CtxOwnerQ(b data.Owners) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, ownerQCtxKey, b)
	}
}
