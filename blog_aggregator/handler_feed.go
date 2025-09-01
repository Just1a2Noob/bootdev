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

	err = handlerFollow(s, command{Name: "follow", Args: []string{url}})
	if err != nil {
		return err
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

	user, _ := get_User(s)

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})

	if err != nil {
		return fmt.Errorf("Creating follow entry failed : %s", err)
	}

	fmt.Printf("%s successfully followed feed: %v", user.Name, feed.Name)
	return nil
}

func handlerFollowing(s *state, cmd command) error {

	user, _ := get_User(s)

	follows, err := s.db.GetFollowForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("Current user is not following any feeds : %s", err)
	}

	for _, follow := range follows {
		feed, err := s.db.FindFeedID(context.Background(), follow.FeedID)
		if err != nil {
			return fmt.Errorf("Finding feed ID unsuccessful : %s", err)
		}
		fmt.Printf("* %s\n", feed.Name)
		fmt.Printf("* %s\n", feed.Url)
		fmt.Println("================")
	}

	return nil
}

func get_User(s *state) (database.User, error) {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return database.User{}, fmt.Errorf("User is not found: %s", err)
	}
	if user.Name == "" {
		return database.User{}, fmt.Errorf("Please login first to follow a feed :%s", err)
	}

	return user, nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* UserID:        %s\n", feed.UserID)
}
