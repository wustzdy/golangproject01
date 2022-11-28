package migration

import (
	"fmt"

	"github.com/go-gormigrate/gormigrate/v2"
)

var MyMigration []*gormigrate.Migration

func init() {
	fmt.Println("starting migration")
}
