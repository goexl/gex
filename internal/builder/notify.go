package builder

import (
	"github.com/goexl/gex/internal/internal/core"
)

type Notify struct {
	command   *Command
	notifiers []core.Notifier
}

func NewNotify(command *Command) *Notify {
	return &Notify{
		command:   command,
		notifiers: make([]core.Notifier, 0),
	}
}

func (n *Notify) By(by core.Notifier) (notify *Notify) {
	n.notifiers = append(n.notifiers, by)
	notify = n

	return
}

func (n *Notify) Build() (command *Command) {
	n.command.params.Notifiers = append(n.command.params.Notifiers, n.notifiers...)
	command = n.command

	return
}
