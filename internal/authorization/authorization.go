package authorization

import (
	"context"
	"event_auth/internal/brocker/nats/subjects"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/nats-io/nats.go"
	"time"
)

type AuthorizationModul struct {
	cashe *redis.Client
	nats  *nats.Conn
}

func NewAuthorizationModule(casheCli *redis.Client, natsCli *nats.Conn) *AuthorizationModul {
	return &AuthorizationModul{
		cashe: casheCli,
		nats:  natsCli,
	}
}

func (am *AuthorizationModul) InitNatsSubscribers() (err error) {
	_, err = am.nats.Subscribe(
		subjects.UserRegEvent.ToString(),
		am.RegisterNats,
	)
	if err != nil {
		return err
	}
	_, err = am.nats.Subscribe(subjects.UserAuthEvent.ToString(), am.AuthorizationNats)
	if err != nil {
		return err
	}

	return

}

func (am *AuthorizationModul) RegisterNats(m *nats.Msg) {
	fmt.Printf("Registration Nats called: %s\n", string(m.Data))

	ctx := context.Background()
	if err := am.set(ctx, "key", "value", time.Hour*24*7); err != nil {
		m.Respond([]byte("Registration error set: " + err.Error()))
	}
	m.Respond([]byte("Successfully registered"))

}

func (am *AuthorizationModul) AuthorizationNats(m *nats.Msg) {
	fmt.Printf("Authorization Nats called: %s\n", string(m.Data))

	ctx := context.Background()
	val, err := am.get(ctx, "key")
	if err != nil {
		m.Respond([]byte("Error get " + err.Error()))
	}
	m.Respond([]byte(val))
}

func (am *AuthorizationModul) set(ctx context.Context, key string, value any, expiration ...time.Duration) error {
	var exp time.Duration
	if len(expiration) > 0 {
		exp = expiration[0]
	}

	return am.cashe.Set(ctx, key, value, exp).Err()
}

func (am *AuthorizationModul) get(ctx context.Context, key string) (string, error) {
	return am.cashe.Get(ctx, key).Result()
}

func (am *AuthorizationModul) del(ctx context.Context, key string) error {
	return am.cashe.Del(ctx, key).Err()
}
