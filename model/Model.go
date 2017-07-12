package model

type User struct {
	UserId              string                   `json:"userId"`
	CreatedDate         int64                    `json:"createdDate"`
	UpdatedDate         int64                    `json:"updatedDate"`
	Email               string                   `json:"email"`
	Password            string                   `json:"password"`
	FullName            string                   `json:"fullName"`
	PhoneNumber         string                   `json:"phoneNumber"`
	Address             string                   `json:"address"`
	BirthDate           int64                    `json:"birthDate"`
	Point               int                      `json:"point"`
	PrestigeRank        int                      `json:"prestigeRank"`
	ConfirmationToken   string                   `json:"confirmationToken"`
	ForgotPasswordToken string                   `json:"forgotPasswordToken"`
	Status              string                   `json:"status"`
	Role                string                   `json:"role"`
	PointStatusLogs     []map[string]interface{} `json:"pointStatusLogs"`
	Qrcode              string                   `json:"qrcode"`
}
