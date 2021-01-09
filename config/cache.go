package config

import (
	. "gitee.com/pangxianfei/frame/config"
)

func init() {
	cache := make(map[string]interface{})

	cache["default"] = Env("CACHE_DRIVER", "memory")
	cache["stores"] = map[string]interface{}{
		"memory": map[string]interface{}{
			"driver":                    "memory",
			"default_expiration_minute": 5,
			"cleanup_interval_minute":   5,
			"prefix":                    Env("APP_NAME", "tmaic").(string) + "_cache_",
		},
		"redis": map[string]interface{}{
			"driver":     "redis",
			"connection": "cache",
		},
	}

	Add("cache", cache)
}
