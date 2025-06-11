package models

type Series struct {
	SeriesID            uint    `json:"series_id"`
	Name                string  `json:"name"`
	Description         string  `json:"description"`
	ParentID            *uint   `json:"parent_series_id"`
	Parent              *Series `json:"parent,omitempty"`
	RebrickableID       int     `json:"id"`
	RebrickableParentID uint    `json:"parent_id"`
}

type SeriesResponse struct {
	Count   int      `json:"count"`
	Next    string   `json:"next"`
	Results []Series `json:"results"`
}
