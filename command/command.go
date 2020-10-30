package command

import "log"

type CMD interface {
	Run(cmd []byte) error
}

type Command struct {
}

func New() *Command {
	return &Command{}
}

func (c *Command) Run(cmd []byte) error {
	log.Printf("%s\n", cmd)
	return nil
}
