package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Just1a2Noob/bootdev/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		Name:      name,
		Url:       url,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}

	fmt.Println("Feed created successfully:")
	printFeed(feed)
	fmt.Println()
	fmt.Println("=====================================")

	return nil
}

func handlerListFeeds(s *state, cmd command) error {
	feeds, err := s.db.ListFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		printFeed(feed)
		fmt.Println()
		fmt.Println("=====================================")
	}
	return nil
}

func handlerFollow(s *state, cmd command) error {
	url := cmd.Args[0]

	feed, err := s.db.FindFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("Feed URL is not found in database: %s", err)
	}

	follows, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    feed.UserID,
		FeedID:    feed.ID,
	})

	if err != nil {
		return fmt.Errorf("Creating follow entry failed : %s", err)
	}

	printFollow(follows)

	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* UserID:        %s\n", feed.UserID)
}

func printFollow(follow database.CreateFeedFollowRow) {
	fmt.Printf("* ID:            %s\n", follow.ID)
	fmt.Printf("* Created:       %v\n", follow.CreatedAt)
	fmt.Printf("* Updated:       %v\n", follow.UpdatedAt)
	fmt.Printf("* UserID:        %s\n", follow.UserID)
	fmt.Printf("* FeedID:        %s\n", follow.FeedID)
}
