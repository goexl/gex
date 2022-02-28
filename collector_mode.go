package exec

const (
	// CollectorModeCache 缓存
	CollectorModeCache CollectorMode = 1
	// CollectorModeDirect 直接输出
	CollectorModeDirect CollectorMode = 2
)

// CollectorMode 收集模式
type CollectorMode uint8
