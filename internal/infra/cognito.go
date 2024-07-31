package infra

import (
	"context"
	"errors"
	"time"

	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"go.uber.org/zap"
)

func NewCognito() Cognito {
	return Cognito{}
}

type AWSCognitoJWK struct {
	Keys []struct {
		Kid string `json:"kid"`
		Alg string `json:"alg"`
		Kty string `json:"kty"`
		E   string `json:"e"`
		N   string `json:"n"`
		Use string `json:"use"`
	} `json:"keys"`
}

type AWSCognitoAccessTokenPayload struct {
	Sub           string   `json:"sub"`
	DeviceKey     string   `json:"device_key"`
	CognitoGroups []string `json:"cognito:groups"`
	Iss           string   `json:"iss"`
	Version       int      `json:"version"`
	ClientID      string   `json:"client_id"`
	OriginJti     string   `json:"origin_jti"`
	EventID       string   `json:"event_id"`
	TokenUse      string   `json:"token_use"`
	Scope         string   `json:"scope"`
	AuthTime      int      `json:"auth_time"`
	Exp           int      `json:"exp"`
	Iat           int      `json:"iat"`
	Jti           string   `json:"jti"`
	Username      string   `json:"username"`
}

type Cognito struct {
	jwksURL       string
	publicKeysURL string
	jwks          AWSCognitoJWK
	jwkCache      *jwk.Cache
	logger        *zap.Logger
	client_id     string
	iss           string
	config        CognitoConfig
}

func (c *Cognito) VerifyToken(ctx context.Context, rawtoken []byte) error {
	_, err := c.jwkCache.Refresh(ctx, c.publicKeysURL)
	if err != nil {
		c.logger.Error("failed to refresh aws jwks", zap.Error(err))
		return err
	}
	keyset, err := c.jwkCache.Get(ctx, c.publicKeysURL)
	if err != nil {
		c.logger.Error("failed to retrieve aws jwks", zap.Error(err))
		return err
	}
	token, err := jwt.Parse(rawtoken, jwt.WithKeySet(keyset), jwt.WithValidate(true))
	if err != nil {
		c.logger.Error("failed to parse token", zap.Error(err))
		return err
	}
	clientID, _ := token.Get("client_id")
	if clientID != c.client_id {
		return errors.New("invalid access token")
	}
	iss, _ := token.Get("iss")
	if iss != c.config.IssuerURL() {
		return errors.New("invalid access token")
	}
	return nil
}

func (c *Cognito) LoadJWK() error {
	ca := jwk.NewCache(context.Background())
	ca.Register(c.jwksURL, jwk.WithMinRefreshInterval(15*time.Minute))
	c.jwkCache = ca
	return nil
}
