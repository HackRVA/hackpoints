package auth

import (
	"context"
	"fmt"
	"hackpoints/models"
	"net/http"
	"time"

	"github.com/shaj13/go-guardian/v2/auth"
	"github.com/shaj13/go-guardian/v2/auth/strategies/basic"
	"github.com/shaj13/go-guardian/v2/auth/strategies/jwt"
	"github.com/shaj13/go-guardian/v2/auth/strategies/union"
	"github.com/shaj13/libcache"
	_ "github.com/shaj13/libcache/fifo"
	log "github.com/sirupsen/logrus"
)

type AuthProvider interface {
	User(*http.Request) auth.Info
	IssueAccessToken(info auth.Info) (string, error)
	AuthMiddleware(next http.Handler) http.Handler
}

type UserStore interface {
	SignIn(username, password string) error
	GetMemberByEmail(email string) (models.Member, error)
}

type AuthServer struct {
	Expires     time.Time
	keeper      jwt.SecretsKeeper
	JWTInterval time.Duration
	strategy    union.Union
	userStore   UserStore
}

func (a *AuthServer) IssueAccessToken(user auth.Info) (string, error) {
	println("in auth info", user)
	return jwt.IssueAccessToken(user, a.keeper, jwt.SetExpDuration(time.Hour*a.JWTInterval))
}

func (a *AuthServer) User(r *http.Request) auth.Info {
	println("in auth user")
	return auth.User(r)
}

func Setup(userStore UserStore) *AuthServer {
	cache := libcache.FIFO.New(0)
	cache.SetTTL(time.Minute * 5)
	cache.RegisterOnExpired(func(key, _ interface{}) {
		cache.Peek(key)
	})

	server := &AuthServer{
		userStore: userStore,
		keeper: jwt.StaticSecret{
			ID:        "secret-id",
			Secret:    []byte("secret"), // change this
			Algorithm: jwt.HS256,
		},
		JWTInterval: 8,
	}

	basicStrategy := basic.NewCached(server.getValidator(), cache)
	jwtStrategy := jwt.New(cache, server.keeper)
	server.strategy = union.New(jwtStrategy, basicStrategy)

	return server
}

func (a AuthServer) getValidator() basic.AuthenticateFunc {
	validator := func(ctx context.Context, r *http.Request, userName, password string) (auth.Info, error) {
		err := a.userStore.SignIn(userName, password)
		if err != nil {
			log.Errorf("error signing in: %s", err)
			return nil, fmt.Errorf("invalid credentials")
		}
		// If we reach this point, that means the users password was correct, and that they are authorized
		// we could attach some of their privledges to this return val I think
		return auth.NewDefaultUser(userName, userName, []string{}, nil), nil
	}
	return validator
}

func (a *AuthServer) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, user, err := a.strategy.AuthenticateRequest(r)
		if err != nil {
			log.Error("whoa error", err)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		r = auth.RequestWithUser(user, r)
		next.ServeHTTP(w, r)
	})
}
