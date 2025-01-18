package config

type sqlConfig struct {
	MaxIdleConns       int
	MaxOpenConns       int
	MaxIdleTimeSeconds int
	MaxLifeTimeSeconds int
}
