package models

type Minifig struct {
	SetNum         string `json:"set_num"`
	Name           string `json:"name"`
	NumParts       int    `json:"num_parts"`
	SetImgURL      string `json:"set_img_url"`
	SetURL         string `json:"set_url"`
	LastModifiedDT string `json:"last_modified_dt"`
}

type MinifigsResponse struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Results  []Minifig `json:"results"`
}
