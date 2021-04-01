// Package grpc implements functions to expose project service endpoint using GRPC protocol.
package grpc

import (
	"context"

	"github.com/decentralized-cloud/edge-cluster/models"
	"github.com/go-kit/kit/endpoint"
)

func (service *transportService) createAuthMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			// token, err := grpc.ParseAndVerifyToken(ctx, service.jwksURL, true)
			// if err != nil {
			// 	return nil, err
			// }

			//			parsedToken := models.ParsedToken{Email: token.PrivateClaims()["email"].(string)}
			parsedToken := models.ParsedToken{Email: "morteza.alizadeh@gmail.com"}
			ctx = context.WithValue(ctx, models.ContextKeyParsedToken, parsedToken)

			return next(ctx, request)
		}
	}
}
