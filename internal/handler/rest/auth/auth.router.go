package auth

import (
	guard "antrein/dd-dashboard-auth/application/middleware"
	"antrein/dd-dashboard-auth/internal/usecase/auth"
	validate "antrein/dd-dashboard-auth/internal/utils/validator"
	"antrein/dd-dashboard-auth/model/config"
	"antrein/dd-dashboard-auth/model/dto"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	cfg     *config.Config
	usecase *auth.Usecase
	vld     *validator.Validate
}

func New(cfg *config.Config, usecase *auth.Usecase, vld *validator.Validate) *Router {
	return &Router{
		cfg:     cfg,
		usecase: usecase,
		vld:     vld,
	}
}

func (r *Router) RegisterRoute(app *fiber.App) {
	g := app.Group("/bc/dashboard/auth")
	g.Post("/register", guard.DefaultGuard(r.RegisterTenant))
	g.Post("/login", guard.DefaultGuard(r.LoginTenantAccount))
}

func (r *Router) RegisterTenant(g *guard.GuardContext) error {
	req := dto.CreateTenantRequest{}

	err := g.FiberCtx.BodyParser(&req)
	if err != nil {
		return g.ReturnError(http.StatusBadRequest, "Request tidak sesuai format")
	}

	err = r.vld.StructCtx(g.FiberCtx.Context(), &req)
	if err != nil {
		return g.ReturnError(http.StatusBadRequest, "Request tidak sesuai format")
	}

	err = validate.ValidateCreateAccount(req)
	if err != nil {
		return g.ReturnError(http.StatusBadRequest, err.Error())
	}

	ctx := g.FiberCtx.Context()
	resp, errRes := r.usecase.RegisterNewTenant(ctx, req)
	if errRes != nil {
		return g.ReturnError(errRes.Status, errRes.Error)
	}

	return g.ReturnCreated(resp)
}

func (r *Router) LoginTenantAccount(g *guard.GuardContext) error {
	req := dto.LoginRequest{}

	err := g.FiberCtx.BodyParser(&req)
	if err != nil {
		return g.ReturnError(http.StatusBadRequest, "Request tidak sesuai format")
	}

	err = r.vld.StructCtx(g.FiberCtx.Context(), &req)
	if err != nil {
		return g.ReturnError(http.StatusBadRequest, "Request tidak sesuai format")
	}

	ctx := g.FiberCtx.Context()
	resp, errRes := r.usecase.LoginTenantAccount(ctx, req)
	if errRes != nil {
		return g.ReturnError(errRes.Status, errRes.Error)
	}

	return g.ReturnSuccess(resp)
}
