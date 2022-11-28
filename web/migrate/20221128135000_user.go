package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"golangproject01/web/models"
	"gorm.io/gorm"
)

func init() {
	type User struct {
		SceneSource models.SceneSource `gorm:"type:tinyint(1) NOT NULL DEFAULT 1;index"`
	}
	MyMigration = append(MyMigration, &gormigrate.Migration{
		ID: "20221128135000",
		Migrate: func(tx *gorm.DB) error {

			if err := tx.AutoMigrate(&User{}); err != nil {
				return err
			} else {
				return nil
			}
		},
		Rollback: func(tx *gorm.DB) error {
			if err := tx.Migrator().DropColumn(&User{}, "scene_source"); err != nil {
				return err
			} else {
				return nil
			}
		},
	})
}
