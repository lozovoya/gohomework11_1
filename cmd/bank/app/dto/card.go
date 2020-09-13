package dto

type CardDTO struct {
	Id       int    `json:"id"`
	Number   string `json:"number"`
	Issuer   string `json:"issuer"`
	HolderId int    `json:"holder_id"`
	Type     string `json:"type"`
	Message  string `json:"message"`
}

type MessageDTO struct {
	Message string `json:"message"`
}
