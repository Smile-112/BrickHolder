package models

type Minifig struct {
	SetNum         string `gorm:"primaryKey" json:"set_num"`
	Name           string `gorm:"size:255;not null" json:"name"`
	NumParts       int    `json:"num_parts"`
	SetImgURL      string `json:"set_img_url"`
	SetURL         string `json:"set_url"`
	LastModifiedDT string `json:"last_modified_dt"`
}
