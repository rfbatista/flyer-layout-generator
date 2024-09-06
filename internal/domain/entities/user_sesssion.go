package entities

type UserSession struct {
	UserID    int32  `json:"user_id"`
	Username  string `json:"username"`
	CompanyID int64  `json:"company_id"`
}
