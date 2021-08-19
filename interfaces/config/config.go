package config

import "time"

type Config interface {
	//Load reads and sets config
	Load() error
	GetBool(key string, defaultVal bool) bool
	GetString(key string, defaultVal string) string
	GetInt(key string, defaultVal int) int
	GetInt32(key string, defaultVal int32) int32
	GetInt64(key string, defaultVal int64) int64
	GetUint(key string, defaultVal uint) uint
	GetUint32(key string, defaultVal uint32) uint32
	GetUint64(key string, defaultVal uint64) uint64
	GetFloat64(key string, defaultVal float64) float64
	GetTime(key string, defaultVal time.Time) time.Time
	GetDuration(key string, defaultVal time.Duration) time.Duration
	GetIntSlice(key string, defaultVal []int) []int
	GetStringSlice(key string, defaultVal []string) []string
	GetStringMap(key string, defaultVal map[string]interface{}) map[string]interface{}
	GetStringMapString(key string, defaultVal map[string]string) map[string]string
	GetStringMapSlice(key string, defaultVal map[string][]string) map[string][]string
}
