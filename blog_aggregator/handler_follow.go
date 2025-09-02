package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Just1a2Noob/bootdev/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	url := cmd.Args[0]

	feed, err := s.db.FindFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("Feed URL is not found in database: %s", err)
	}

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

func handlerFollowing(s *state, cmd command, user database.User) error {

	follows, err := s.db.GetFollowForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("Current user is not following any feeds : %s", err)
	}

	for _, follow := range follows {
		feed, err := s.db.FindFeedID(context.Background(), follow.FeedID)
		if err != nil {
			return fmt.Errorf("Finding feed ID unsuccessful : %s", err)
		}
		fmt.Println("================")
		fmt.Printf("* %s\n", feed.Name)
		fmt.Printf("* %s\n", feed.Url)
		fmt.Println("================")
	}

	return nil
}

func handlerUnfollowing(s *state, cmd command, user database.User) error {
	url := cmd.Args[0]

	feed, err := s.db.FindFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("Invalid URL : %s", err)
	}

	err = s.db.DeleteFollow(context.Background(), database.DeleteFollowParams{
		FeedID: feed.ID,
		UserID: user.ID,
	})
	if err != nil {
		return fmt.Errorf("Unfollowing feed unsuccessful: %s", err)
	}

	fmt.Print("Successfully unfollowed feed")
	return nil
}
