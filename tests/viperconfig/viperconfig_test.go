package viperconfig

import (
	"testing"
	"time"

	"github.com/bang9211/ossicones/implements/viperconfig"
	"github.com/bang9211/ossicones/interfaces/config"

	. "github.com/stretchr/testify/assert"
)

var (
	testBoolVal                   = false
	defaultBoolVal                = true
	testStringVal                 = "test 998877"
	defaultStringVal              = "Hello World!"
	testIntVal            int     = 998877
	defaultIntVal         int     = 112233
	testInt32Val          int32   = 2147483647
	defaultInt32Val       int32   = 112233
	testInt64Val          int64   = 9223372036854775807
	defaultInt64Val       int64   = 112233
	testUintVal           uint    = 998877
	defaultUintVal        uint    = 112233
	testUint32Val         uint32  = 4294967295
	defaultUint32Val      uint32  = 112233
	testUint64Val         uint64  = 18446744073709551615
	defaultUint64Val      uint64  = 112233
	testFloat64Val        float64 = 998.877
	defaultFloat64Val     float64 = 112.233
	testTimeVal                   = time.Date(2021, 9, 15, 15, 31, 48, 123, time.UTC)
	defaultTimeVal                = time.Now()
	testDurationVal               = 9 * time.Minute
	defaultDurationVal            = 12 * time.Second
	testIntSliceVal               = []int{9, 9, 8, 8, 7, 7}
	defaultIntSliceVal            = []int{1, 1, 2, 2, 3, 3}
	testStringSliceVal            = []string{"test1", "test2", "test3"}
	defaultStringSliceVal         = []string{"Hello", "World!", "Nice Day!"}
	testStringMapVal              = map[string]interface{}{
		"test1": "testValue",
		"test2": 999999,
		"test3": false,
	}
	defaultStringMapVal = map[string]interface{}{
		"hello": "World!",
		"Nice":  999999,
		"Day":   true,
	}
	testStringMapStringVal = map[string]string{
		"test1": "testValue1",
		"test2": "testValue2",
		"test3": "testValue3",
	}
	defaultStringMapStringVal = map[string]string{
		"Hello":   "World!",
		"Nice To": "Meet You",
		"Me":      "Too",
	}
	testStringMapSliceVal = map[string][]string{
		"test1": {"testValue1", "testValue2"},
		"test2": {"testValue3", "testValue4"},
		"test3": {"testValue5", "testValue6"},
	}
	defaultStringMapSliceVal = map[string][]string{
		"test1": {"Hello", "World!"},
		"test2": {"Nice", "To"},
		"test3": {"Meet", "You"},
	}
)

func TestImplementConfig(t *testing.T) {
	Implements(t, (*config.Config)(nil), new(viperconfig.ViperConfig),
		"It must implements of interface config.Config")
}

func TestLoadDefault(t *testing.T) {
	cfg, err := initTest()
	NoError(t, err, "Failed to initTest()")
	defer NoError(t, closeTest(cfg), "Failed to closeTest()")

	Equal(t, testBoolVal, cfg.GetBool("test_viper_config_bool_value", defaultBoolVal))
	Equal(t, testStringVal, cfg.GetString("test_viper_config_string_value", defaultStringVal))
	Equal(t, testIntVal, cfg.GetInt("test_viper_config_int_value", defaultIntVal))
	Equal(t, testInt32Val, cfg.GetInt32("test_viper_config_int32_value", defaultInt32Val))
	Equal(t, testInt64Val, cfg.GetInt64("test_viper_config_int64_value", defaultInt64Val))
	Equal(t, testUintVal, cfg.GetUint("test_viper_config_uint_value", defaultUintVal))
	Equal(t, testUint32Val, cfg.GetUint32("test_viper_config_uint32_value", defaultUint32Val))
	Equal(t, testUint64Val, cfg.GetUint64("test_viper_config_uint64_value", defaultUint64Val))
	Equal(t, testFloat64Val, cfg.GetFloat64("test_viper_config_float64_value", defaultFloat64Val))
	Equal(t, testTimeVal, cfg.GetTime("test_viper_config_time_value", defaultTimeVal))
	Equal(t, testDurationVal, cfg.GetDuration("test_viper_config_duration_value", defaultDurationVal))
	// Equal(t, testIntSliceVal, cfg.GetIntSlice("test_viper_config_intslice_value", defaultIntSliceVal)) // doesn't work
	Equal(t, testStringSliceVal, cfg.GetStringSlice("test_viper_config_stringslice_value", defaultStringSliceVal))
	// Equal(t, testStringMapVal, cfg.GetStringMap("test_viper_config_stringmap_value", defaultStringMapVal))
	// Equal(t, testStringMapStringVal, cfg.GetStringMapString("test_viper_config_stringmapstring_value", defaultStringMapStringVal))
	// Equal(t, testStringMapSliceVal, cfg.GetStringMapSlice("test_viper_config_stringmapslice_value", defaultStringMapSliceVal))
}

func TestLoadJSON(t *testing.T) {
	// os.Args[1] = "--config"
	// os.Args[2] = "ossicones.json"
	// cfg, err := initTest()
	// NoError(t, err, "Failed to initTest()")
	// defer NoError(t, closeTest(cfg), "Failed to closeTest()")
}

func TestLoadYAML(t *testing.T) {
	// os.Args[1] = "--config"
	// os.Args[2] = "ossicones.yaml"
	// cfg, err := initTest()
	// NoError(t, err, "Failed to initTest()")
	// defer NoError(t, closeTest(cfg), "Failed to closeTest()")
}

func TestLoadTOML(t *testing.T) {
	// os.Args[1] = "--config"
	// os.Args[2] = "ossicones.toml"
	// cfg, err := initTest()
	// NoError(t, err, "Failed to initTest()")
	// defer NoError(t, closeTest(cfg), "Failed to closeTest()")
}

func TestGetBool(t *testing.T) {

}

func TestGetString(t *testing.T) {

}

func TestGetInt(t *testing.T) {

}

func TestGetInt32(t *testing.T) {

}

func TestGetInt64(t *testing.T) {

}

func TestGetUint(t *testing.T) {

}

func TestGetUint32(t *testing.T) {

}

func TestGetUint64(t *testing.T) {

}

func TestGetFloat64(t *testing.T) {

}

func TestGetTime(t *testing.T) {

}

func TestGetDuration(t *testing.T) {

}

func TestGetIntSlice(t *testing.T) {

}

func TestGetStringSlice(t *testing.T) {

}

func TestGetStringMap(t *testing.T) {

}

func TestGetStringMapString(t *testing.T) {

}

func TestGetStringMapSlice(t *testing.T) {

}

func TestClose(t *testing.T) {

}
