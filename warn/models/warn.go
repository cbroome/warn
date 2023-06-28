package models

import (
	"warn/warn/structs"

	"gorm.io/gorm"
)

type WarnNoticeModel struct {
	gorm.Model
	structs.WarnNotice
}

