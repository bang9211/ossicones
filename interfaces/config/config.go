package config

import "time"

type Config interface {
	// Load reads and sets config.
	Load() error
	// GetBool gets boolean value of the key if present, Otherwise, it gets defaultVal.
	GetBool(key string, defaultVal bool) bool
	// GetString gets string value of the key if present, Otherwise, it gets defaultVal.
	GetString(key string, defaultVal string) string
	// GetInt gets int value of the key if present, Otherwise, it gets defaultVal.
	GetInt(key string, defaultVal int) int
	// GetBoGetInt32ol gets int32 value of the key if present, Otherwise, it gets defaultVal.
	GetInt32(key string, defaultVal int32) int32
	// GetInt64 gets int64 value of the key if present, Otherwise, it gets defaultVal.
	GetInt64(key string, defaultVal int64) int64
	// GetUint gets uint value of the key if present, Otherwise, it gets defaultVal.
	GetUint(key string, defaultVal uint) uint
	// GetUint32 gets uint32 value of the key if present, Otherwise, it gets defaultVal.
	GetUint32(key string, defaultVal uint32) uint32
	// GetUint64 gets uint64 value of the key if present, Otherwise, it gets defaultVal.
	GetUint64(key string, defaultVal uint64) uint64
	// GetFloat64 gets float64 value of the key if present, Otherwise, it gets defaultVal.
	GetFloat64(key string, defaultVal float64) float64
	// GetTime gets time.Time value of the key if present, Otherwise, it gets defaultVal.
	GetTime(key string, defaultVal time.Time) time.Time
	// GetDuration gets time.Duration value of the key if present, Otherwise, it gets defaultVal.
	GetDuration(key string, defaultVal time.Duration) time.Duration
	// GetIntSlice gets []int value of the key if present, Otherwise, it gets defaultVal.
	GetIntSlice(key string, defaultVal []int) []int
	// GetStringSlice gets []string value of the key if present, Otherwise, it gets defaultVal.
	GetStringSlice(key string, defaultVal []string) []string
	// GetStringMap gets map[string]interface{} value of the key if present, Otherwise, it gets defaultVal.
	GetStringMap(key string, defaultVal map[string]interface{}) map[string]interface{}
	// GetStringMapString gets map[string]string value of the key if present, Otherwise, it gets defaultVal.
	GetStringMapString(key string, defaultVal map[string]string) map[string]string
	// GetStringMapSlice gets map[string][]string value of the key if present, Otherwise, it gets defaultVal.
	GetStringMapSlice(key string, defaultVal map[string][]string) map[string][]string
	// Close closes config.
	Close() error
}
