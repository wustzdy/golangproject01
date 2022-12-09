package config

import (
	"golangproject01/record/workspace"
	"gorm.io/gorm"
)

const (
	ENV_UNIT    = "unit_test"
	ENV_DEV     = "dev"
	ENV_TEST    = "test"
	ENV_STAGING = "staging"
	ENV_PRO     = "pro"
)

var (
	DB       *gorm.DB
	WSClient workspace.WSClient
)
