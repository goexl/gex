package gex

// Exec 执行外部命令
func Exec(command string, opts ...option) (code int, err error) {
	var cmd *_command
	if cmd, err = newCommand(command, opts...); nil != err {
		return
	}
	code, err = cmd.Exec()

	return
}
