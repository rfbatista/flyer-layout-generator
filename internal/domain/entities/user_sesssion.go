package entities

type UserSession struct {
	UserID    int32  `json:"user_id,omitempty"`
	Username  string `json:"username,omitempty"`
	CompanyID int64  `json:"company_id,omitempty"`
}
