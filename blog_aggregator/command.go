package main

import (
	"errors"

	"github.com/Just1a2Noob/bootdev/blog_aggregator/internal/config"
)

type state struct {
	ptr_to_cfg *config.Config
}

type command struct {
	Name string
	Args []string
}

type commands struct {
	commands map[string]func(*state, command) error
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("Username is required")
	}

	err := s.ptr_to_cfg.SetUser(cmd.Args[0])
	if err != nil {
		return err
	}

	return nil
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
