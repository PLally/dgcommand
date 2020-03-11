package permissions

import "github.com/jinzhu/gorm"

type PermissionsModel struct {
	gorm.Model
	Name string
	Default bool
}

type PermValueModel struct {
	gorm.Model
	Snowflake string

}

