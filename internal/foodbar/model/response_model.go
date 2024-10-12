package model

type Response struct {
	Success bool   `json:"success,omitempty"`
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

type HalalResponse struct {
	Message    string            `json:"message,omitempty"`
	EntityData []HalalEntityData `json:"entityData,omitempty"`
}

type HalalEntityData struct {
	Score      float32                  `json:"score,omitempty"`
	Label      string                   `json:"label,omitempty"`
	StatsScore string                   `json:"statsScore,omitempty"`
	Atribute   HalalEntityDataAttribute `json:"atribute,omitempty"`
}

type HalalEntityDataAttribute struct {
	Code            string `json:"code,omitempty"`
	Certificate     string `json:"certificate,omitempty"`
	Label           string `json:"label,omitempty"`
	FoodProductId   string `json:"foodproductId,omitempty"`
	HasManufacturer string `json:"hasManufacturer,omitempty"`
}
