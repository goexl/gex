package builder

import (
	"fmt"

	"github.com/goexl/gex/internal/builder/internal"
	"github.com/goexl/gex/internal/internal/collector"
	"github.com/goexl/gex/internal/internal/constant"
	"github.com/goexl/gex/internal/param"
)

type Strings struct {
	*internal.Collect

	command *Command
	collect *Collect
	params  *param.Collect
	target  *[]string
}

func NewStrings(command *Command, collect *Collect, target *[]string) (strings *Strings) {
	strings = new(Strings)
	strings.params = param.NewCollect()
	strings.Collect = internal.NewCollect(strings.params)
	strings.command = command
	strings.collect = collect
	strings.target = target

	return
}

func (s *Strings) Build() (collect *Collect) {
	s.command.params.Collectors[fmt.Sprintf(constant.Point, s)] = collector.NewStrings(s.target, s.params)
	s.command.params.Async = true
	collect = s.collect

	return
}
