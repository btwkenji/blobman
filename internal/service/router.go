package service

import (
	"github.com/go-chi/chi"
	"github.com/kenjitheman/blobman/internal/config"
	"github.com/kenjitheman/blobman/internal/data/postgres"
	"github.com/kenjitheman/blobman/internal/service/handlers"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router(cfg config.Config) chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxBlobQ(postgres.NewBlobs(cfg.DB())),
			handlers.CtxOwnerQ(postgres.NewOwners(cfg.DB())),
		),
	)
	r.Route("/integrations/blobman", func(r chi.Router) {
		r.Post("/blobs", handlers.CreateBlob)
		r.Get("/blobs", handlers.GetBlobsList)
		r.Get("/blobs/{id}", handlers.GetBlob)
		r.Delete("/blobs/{id}", handlers.DeleteBlob)
	})
	return r
}
