package entities

type UserSession struct {
	Username  string `json:"username,omitempty"`
	CompanyID int64  `json:"company_id,omitempty"`
}
