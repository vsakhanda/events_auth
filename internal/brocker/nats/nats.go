package nats

import (
	//"event_auth/internal/authorization"
	"github.com/nats-io/nats.go"
	"os"
)

func NewNatsClient() (nc *nats.Conn, err error) {
	return nats.Connect(os.Getenv("NATS_URL"))
}

//
//func InitNatsSubscriber(cli *nats.Conn, module *authorization.AuthorizationModul (err error) {
//	_, err := cli.Subscribe(UserRegEvent.ToString(), module.RegisterNats)
//	if err != nil {
//		return err
//	})
////	log.Printf("Отримано сесію: %s", string(m.Data))
//
//	_, err := cli.Subscribe(UserAuthEvent.ToString(),  module.AuthorizationNats)
//	if err != nil {
//		return err
//	})
