package utils

import "os"

const (
	POKER = "POKER"

	LOCAL       = "LOCAL"
	DEVELOPMENT = "DEVELOPMENT"
	TESTING     = "TESTING"
	STAGING     = "STAGING"
	PRODUCTION  = "PRODUCTION"
)

func GetEnvironment() string {
	env := os.Getenv(POKER)
	// 不设置默认就是本地环境
	if env == "" {
		return LOCAL
	}
	return env
}

func IsLocal() bool {
	return GetEnvironment() == LOCAL
}

func IsDevelopment() bool {
	return GetEnvironment() == DEVELOPMENT
}

func EnableDevelopment() {
	enableEnv(DEVELOPMENT)
}

func IsTesting() bool {
	return GetEnvironment() == TESTING
}

func EnableTesting() {
	enableEnv(TESTING)
}

func IsStaging() bool {
	return GetEnvironment() == STAGING
}

func EnableStaging() {
	enableEnv(STAGING)
}

func IsProduction() bool {
	return GetEnvironment() == PRODUCTION
}

func EnableProduction() {
	enableEnv(PRODUCTION)
}

func enableEnv(v string) {
	if err := os.Setenv(POKER, v); err != nil {
		panic(err)
	}
}
