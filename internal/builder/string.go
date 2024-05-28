package builder

import (
	"fmt"

	"github.com/goexl/gex/internal/builder/internal"
	"github.com/goexl/gex/internal/internal/collector"
	"github.com/goexl/gex/internal/internal/constant"
	"github.com/goexl/gex/internal/param"
)

type String struct {
	*internal.Collect

	command *Command
	collect *Collect
	params  *param.Collect
	target  *string
}

func NewString(command *Command, collect *Collect, target *string) (str *String) {
	str = new(String)
	str.params = param.NewCollect()
	str.Collect = internal.NewCollect(str.params)
	str.command = command
	str.collect = collect
	str.target = target

	return
}

func (s *String) Build() (collect *Collect) {
	s.command.params.Collectors[fmt.Sprintf(constant.Point, s)] = collector.NewString(s.target, s.params)
	s.command.params.Async = true
	collect = s.collect

	return
}
