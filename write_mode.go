package gex

const (
	// WriteModeSkip 跳过
	WriteModeSkip writeMode = 1
	// WriteModeOverride 覆盖
	WriteModeOverride writeMode = 2
	// WriteModeError 返回错误
	WriteModeError writeMode = 3
	// WriteModeRename 重命名
	WriteModeRename writeMode = 4
	// WriteModeAppend 追加
	WriteModeAppend writeMode = 5
)

type writeMode int
