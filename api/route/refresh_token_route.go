package route

import (
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/altafino/go-backend-clean-architecture-chi/api/controller"
	"github.com/altafino/go-backend-clean-architecture-chi/bootstrap"
	"github.com/altafino/go-backend-clean-architecture-chi/domain"
	"github.com/altafino/go-backend-clean-architecture-chi/mongo"
	"github.com/altafino/go-backend-clean-architecture-chi/repository"
	"github.com/altafino/go-backend-clean-architecture-chi/usecase"
)

func NewRefreshTokenRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, router chi.Router) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	rtc := &controller.RefreshTokenController{
		RefreshTokenUsecase: usecase.NewRefreshTokenUsecase(ur, timeout),
		Env:                 env,
	}
	router.Post("/refresh", rtc.RefreshToken)
}
