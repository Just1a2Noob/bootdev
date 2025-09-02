package main

import (
	"context"
	"fmt"

	"github.com/Just1a2Noob/bootdev/blog_aggregator/internal/database"
)

func handlerLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {

	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("Error with user. Not logged in or not in database: %v", err)
		}
		return handler(s, cmd, user)
	}
}
