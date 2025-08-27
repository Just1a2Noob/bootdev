package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Just1a2Noob/bootdev/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}
	fmt.Printf("Feed: %+v\n", feed)
	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	username := s.ptr_to_cfg.CurrentUserName

	name := cmd.Args[0]
	url := cmd.Args[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		// TODO: Fix this, it should have a function that takes a username and spits out its ID from database
		UserID: handleGetUsers(s, cmd),
	})
	if err != nil {
		return err
	}

	// TODO: After fixing the above paramater print out its output
	return nil
}
