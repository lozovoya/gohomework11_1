package dto

type CardDTO struct {
	Id       int    `json:"id,omitempty"`
	Number   string `json:"number,omitempty"`
	Issuer   string `json:"issuer,omitempty"`
	HolderId int    `json:"holder_id,omitempty"`
	Type     string `json:"type,omitempty"`
	Message  string `json:"message,omitempty"`
}

type MessageDTO struct {
	Message string `json:"message"`
}
