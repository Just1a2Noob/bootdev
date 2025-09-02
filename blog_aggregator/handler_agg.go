package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Just1a2Noob/bootdev/blog_aggregator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	time_between_reqs, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("Time argument syntax is invalid : %s", err)
	}

	ticker := time.NewTicker(time_between_reqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) error {
	fetch, err := s.db.GetNextFeedtoFetch(context.Background())
	if err != nil {
		return fmt.Errorf("Error fetching feed from database : %s", err)
	}

	_, err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		UpdatedAt:     time.Now().UTC(),
		LastFetchedAt: sql.NullTime{Time: time.Now(), Valid: true},
		ID:            fetch.ID,
	})

	if err != nil {
		return fmt.Errorf("Cannot update feed fetch row : %s", err)
	}

	feed, err := fetchFeed(context.Background(), fetch.Url)
	if err != nil {
		return fmt.Errorf("Couldn't not get the fetch results : %v", err)
	}

	for _, title := range feed.Channel.Title {
		println(title)
	}

	return nil
}
