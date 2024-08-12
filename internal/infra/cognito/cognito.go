package cognito

import (
	"algvisual/internal/entities"
	"algvisual/internal/infra/config"
	"algvisual/internal/shared"
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"go.uber.org/zap"
)

func NewCognito(c *config.AppConfig, log *zap.Logger) *Cognito {
	return &Cognito{publicKeysURL: c.Cognito.IssuerURL(), logger: log, config: c.Cognito}
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
	publicKeysURL string
	jwkCache      *jwk.Cache
	logger        *zap.Logger
	config        config.CognitoConfig
}

func (c *Cognito) VerifyToken(ctx context.Context, rawtoken []byte) (*entities.UserSession, error) {
	_, err := c.jwkCache.Refresh(ctx, c.publicKeysURL)
	if err != nil {
		c.logger.Error("failed to refresh aws jwks", zap.Error(err))
		return nil, shared.WrapWithAppError(err, "", err.Error())
	}
	keyset, err := c.jwkCache.Get(ctx, c.publicKeysURL)
	if err != nil {
		c.logger.Error("failed to retrieve aws jwks", zap.Error(err))
		return nil, shared.WrapWithAppError(err, "", err.Error())
	}
	token, err := jwt.Parse(rawtoken, jwt.WithKeySet(keyset), jwt.WithValidate(true))
	if err != nil {
		c.logger.Error("failed to parse token", zap.Error(err))
		return nil, shared.WrapWithAppError(err, "", err.Error())
	}
	clientID, _ := token.Get("client_id")
	if clientID != c.config.ClientID {
		return nil, errors.New("invalid access token: client id does not match")
	}
	// iss, _ := token.Get("iss")
	// if iss != c.config.IssuerURL() {
	// 	return errors.New("invalid access token: issuer does not match")
	// }
	username, _ := token.Get("username")
	// companyID, _ := token.Get("custom:company_id")

	var companyID int64
	user, err := c.GetUser(c.createSession(), string(rawtoken))
	if err != nil {
		c.logger.Error("failed to get cognito user", zap.Error(err))
	} else {
		for idx := range user.UserAttributes {
			if user.UserAttributes[idx].Name != nil && *user.UserAttributes[idx].Name == "custom:company_id" && user.UserAttributes[idx].Value != nil {
				companyID, err = strconv.ParseInt(*user.UserAttributes[idx].Value, 10, 64)
				if err != nil {
					c.logger.Error("failed to parse company id", zap.Error(err))
				}
			}
		}
	}
	return &entities.UserSession{
		Username:  username.(string),
		CompanyID: companyID,
	}, nil
}

func (c *Cognito) LoadJWK() error {
	ca := jwk.NewCache(context.Background())
	err := ca.Register(c.publicKeysURL, jwk.WithMinRefreshInterval(15*time.Minute))
	if err != nil {
		return err
	}
	c.jwkCache = ca
	return nil
}

func (c *Cognito) createSession() *session.Session {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(c.config.Region), // Replace with your region
	})
	if err != nil {
		fmt.Println("Failed to create session,", err)
		return nil
	}
	return sess
}

func (c *Cognito) GetUser(
	sess *session.Session,
	accessToken string,
) (*cognitoidentityprovider.GetUserOutput, error) {
	svc := cognitoidentityprovider.New(sess)

	input := &cognitoidentityprovider.GetUserInput{
		AccessToken: aws.String(accessToken),
	}

	result, err := svc.GetUser(input)
	if err != nil {
		fmt.Println("Failed to get user,", err)
		return nil, err
	}
	return result, nil
}
