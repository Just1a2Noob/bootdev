package main

import (
	"errors"

	"github.com/Just1a2Noob/bootdev/blog_aggregator/internal/config"
	"github.com/Just1a2Noob/bootdev/blog_aggregator/internal/database"
)

type state struct {
	db         *database.Queries
	ptr_to_cfg *config.Config
}

type command struct {
	Name string
	Args []string
}

type commands struct {
	commands map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	function, exists := c.commands[cmd.Name]
	if exists == false {
		return errors.New("Command does not exist")
	}

	err := function(s, cmd)
	if err != nil {
		return err
	}

	return function(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) error {
	_, exists := c.commands[name]
	if exists {
		return errors.New("Command already exists")
	}

	c.commands[name] = f

	return nil
}
