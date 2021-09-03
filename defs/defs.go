package defs

import "sync"

var (
	GlobalConfig *config
	onceConfig sync.Once
	)

type config struct {
	Port string
	XuidLength int
	ResultMap sync.Map
}

func GenerateGlobalConfig() error {
	config := config{}

	onceConfig.Do(func() {
		GlobalConfig = &config
	})
	return nil
}
