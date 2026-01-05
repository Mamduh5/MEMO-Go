package grpc

import (
	"context"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/golang-jwt/jwt/v5"
)

type JWTInterceptor struct {
	secret []byte
}

func NewJWTInterceptor(secret string) *JWTInterceptor {
	return &JWTInterceptor{
		secret: []byte(secret),
	}
}

func (i *JWTInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "missing metadata")
		}

		authHeader := md.Get("authorization")
		if len(authHeader) == 0 {
			return nil, status.Error(codes.Unauthenticated, "missing authorization header")
		}

		tokenStr := authHeader[0]
		if !strings.HasPrefix(tokenStr, "Bearer ") {
			return nil, status.Error(codes.Unauthenticated, "invalid authorization header")
		}

		tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, status.Error(codes.Unauthenticated, "unexpected signing method")
			}
			return i.secret, nil
		})
		if err != nil || !token.Valid {
			return nil, status.Error(codes.Unauthenticated, "invalid token")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "invalid token claims")
		}

		sub, ok := claims["sub"].(string)
		if !ok || sub == "" {
			return nil, status.Error(codes.Unauthenticated, "missing sub claim")
		}

		ctx = ContextWithUserID(ctx, sub)

		return handler(ctx, req)
	}
}
