package guard

import (
	"antrein/dd-dashboard-auth/model/config"
	"antrein/dd-dashboard-auth/model/dto"
	"antrein/dd-dashboard-auth/model/entity"
	"fmt"
	"net/http"

	jwtware "github.com/gofiber/contrib/jwt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type GuardContext struct {
	FiberCtx *fiber.Ctx
}

type AuthGuardContext struct {
	FiberCtx *fiber.Ctx
	Claims   entity.JWTClaim
}

var statusErrorMap = map[int]string{
	400: "Bad Request",
	401: "Unauthorized",
	403: "Forbidden",
	404: "Not Found",
	500: "Internal Server Error",
	503: "Service Unavailable",
}

func (g *GuardContext) ReturnError(
	status int,
	err string,
	detail ...any,
) error {
	if status == http.StatusInternalServerError {
		fmt.Println("error", err)
	}

	response := dto.DefaultResponse{
		Status:  status,
		Message: statusErrorMap[status],
		Error:   err,
	}

	if len(detail) > 0 {
		response.Data = detail[0]
	}

	return g.FiberCtx.Status(status).JSON(response)
}

func (g *GuardContext) ReturnSuccess(
	data interface{},
) error {
	return g.FiberCtx.Status(http.StatusOK).JSON(dto.DefaultResponse{
		Status:  http.StatusOK,
		Message: "OK",
		Data:    data,
	})
}

func (g *GuardContext) ReturnCreated(
	data interface{},
) error {
	return g.FiberCtx.Status(http.StatusCreated).JSON(dto.DefaultResponse{
		Status:  http.StatusCreated,
		Message: "Created",
		Data:    data,
	})
}

func (g *GuardContext) ReturnHTML(htmlContent string) error {
	g.FiberCtx.Set("Content-Type", "text/html; charset=utf-8")
	return g.FiberCtx.Status(http.StatusOK).SendString(htmlContent)
}

func (g *AuthGuardContext) ReturnError(
	status int,
	err string,
	detail ...any,
) error {
	if status == http.StatusInternalServerError {
		fmt.Println("error", err)
	}

	response := dto.DefaultResponse{
		Status:  status,
		Message: statusErrorMap[status],
		Error:   err,
	}

	if len(detail) > 0 {
		response.Data = detail[0]
	}

	return g.FiberCtx.Status(status).JSON(response)
}

func (g *AuthGuardContext) ReturnSuccess(
	data interface{},
) error {
	return g.FiberCtx.Status(http.StatusOK).JSON(dto.DefaultResponse{
		Status:  http.StatusOK,
		Message: "OK",
		Data:    data,
	})
}

func (g *AuthGuardContext) ReturnCreated(
	data interface{},
) error {
	return g.FiberCtx.Status(http.StatusCreated).JSON(dto.DefaultResponse{
		Status:  http.StatusCreated,
		Message: "Created",
		Data:    data,
	})
}

func (g *AuthGuardContext) ReturnHTML(htmlContent string) error {
	g.FiberCtx.Set("Content-Type", "text/html; charset=utf-8")
	return g.FiberCtx.Status(http.StatusOK).SendString(htmlContent)
}

func DefaultGuard(handlerFunc func(g *GuardContext) error) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		guardCtx := GuardContext{
			FiberCtx: ctx,
		}
		return handlerFunc(&guardCtx)
	}
}

func AuthGuard(cfg *config.Config, handlerFunc func(g *AuthGuardContext) error) []fiber.Handler {
	handlers := []fiber.Handler{
		jwtware.New(jwtware.Config{
			SigningKey: jwtware.SigningKey{
				Key: []byte(cfg.Secrets.JWTSecret),
			},
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				return c.Status(fiber.StatusUnauthorized).JSON(dto.DefaultResponse{
					Status:  http.StatusUnauthorized,
					Message: statusErrorMap[http.StatusUnauthorized],
					Error:   "Tidak ada autentikasi",
				})
			},
		}),
		func(ctx *fiber.Ctx) error {
			user := ctx.Locals("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)
			expireAt, err := claims.GetExpirationTime()
			if err != nil {
				return ctx.Status(http.StatusUnauthorized).JSON(dto.DefaultResponse{
					Status:  http.StatusUnauthorized,
					Message: statusErrorMap[http.StatusUnauthorized],
					Error:   "Sesi anda telah berakhir",
				})
			}
			issuedAt, err := claims.GetIssuedAt()
			if err != nil {
				return ctx.Status(http.StatusUnauthorized).JSON(dto.DefaultResponse{
					Status:  http.StatusUnauthorized,
					Message: statusErrorMap[http.StatusUnauthorized],
					Error:   "Sesi anda telah berakhir",
				})
			}
			userID, ok := claims["user_id"].(string)
			if !ok {
				return ctx.Status(http.StatusUnauthorized).JSON(dto.DefaultResponse{
					Status:  http.StatusUnauthorized,
					Message: statusErrorMap[http.StatusUnauthorized],
					Error:   "Tidak terautentikasi",
				})
			}
			ety := entity.JWTClaim{
				UserID: userID,
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: expireAt,
					IssuedAt:  issuedAt,
				},
			}
			authGuardCtx := AuthGuardContext{
				FiberCtx: ctx,
				Claims:   ety,
			}
			return handlerFunc(&authGuardCtx)
		},
	}
	return handlers
}
