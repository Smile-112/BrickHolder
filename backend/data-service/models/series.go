package models

type Series struct {
	SeriesID            uint    `gorm:"primaryKey" json:"series_id"`
	Name                string  `gorm:"size:255;not null" json:"name"`
	Description         string  `json:"description"`
	ParentID            *uint   `json:"parent_series_id"` // nullable внешний ключ
	Parent              *Series `gorm:"foreignKey:ParentID;references:SeriesID" json:"parent,omitempty"`
	RebrickableID       int     `gorm:"unique;not null" json:"id"`
	RebrickableParentID uint    `gorm:"not null" json:"parent_id"`
}
