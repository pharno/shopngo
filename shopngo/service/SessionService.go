package service

import "github.com/martini-contrib/sessions"
import "fmt"

type SessionService struct {
}

func getRedisStore() sessions.RediStore {
	store, err := sessions.NewRediStore(10, "tcp", ":6379", "", []byte("secret-key"))
	if err != nil {
		panic(fmt.Sprintf("cant connect to redis at :6379"))
	}
	fmt.Sprintf("connected to redis at :6379")
	return store
}

var Store = getRedisStore()
