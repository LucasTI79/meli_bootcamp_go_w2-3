package domain

// Product represents an underlying URL with statistics on how it is used.
type Product struct {
	ID             int     `json:"id"`
	Description    string  `json:"description"`
	ExpirationRate int     `json:"expiration_rate"`
	FreezingRate   int     `json:"freezing_rate"`
	Height         float32 `json:"height"`
	Length         float32 `json:"length"`
	Netweight      float32 `json:"netweight"`
	ProductCode    string  `json:"product_code"`
	RecomFreezTemp float32 `json:"recommended_freezing_temperature"`
	Width          float32 `json:"width"`
	ProductTypeID  int     `json:"product_type_id"`
	SellerID       int     `json:"seller_id"`
}

type ProductRequest struct {
	Description    string  `json:"description"`
	ExpirationRate int     `json:"expiration_rate"`
	FreezingRate   int     `json:"freezing_rate"`
	Height         float32 `json:"height"`
	Length         float32 `json:"length"`
	Netweight      float32 `json:"netweight"`
	ProductCode    string  `json:"product_code"`
	RecomFreezTemp float32 `json:"recommended_freezing_temperature"`
	Width          float32 `json:"width"`
	ProductTypeID  int     `json:"product_type_id"`
	SellerID       int     `json:"seller_id"`
}

type ProductResponse struct {
	Data []Product `json:"data"`
}

type ProductResponseById struct {
	Data Product `json:"data"`
}

type RecordByProduct struct {
	ProductID    int    `json:"product_id"`
	Description  string `json:"description"`
	RecordsCount int    `json:"records_count"`
}
