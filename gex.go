package gex

// Exec 执行外部命令
func Exec(command string, opts ...option) (code int, err error) {
	return NewCommand(command, opts...).Exec()
}
