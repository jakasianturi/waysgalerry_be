package models

type Project struct {
	Base
	HiredId     int            `json:"-"`
	Hired       Hired          `gorm:"foreignKey:HiredId" json:"hired"`
	Description string         `json:"description" gorm:"type: text"`
	Photos      []ProjectImage `json:"photos" gorm:"foreignkey:ProjectID"`
}
