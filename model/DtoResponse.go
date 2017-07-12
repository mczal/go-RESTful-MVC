package model

type UserSimpleDtoResponse struct {
	UserId       string `json:"userId"`
	CreatedDate  int64  `json:"createdDate"`
	UpdatedDate  int64  `json:"updatedDate"`
	Email        string `json:"email"`
	FullName     string `json:"fullName"`
	PhoneNumber  string `json:"phoneNumber"`
	Address      string `json:"address"`
	BirthDate    int64  `json:"birthDate"`
	Point        int    `json:"point"`
	PrestigeRank int    `json:"prestigeRank"`
	Status       string `json:"status"`
	Role         string `json:"role"`
	Qrcode       string `json:"qrcode"`
}

type UserDetailDtoResponse struct {
	UserId              string                   `json:"userId"`
	CreatedDate         int64                    `json:"createdDate"`
	UpdatedDate         int64                    `json:"updatedDate"`
	Email               string                   `json:"email"`
	FullName            string                   `json:"fullName"`
	PhoneNumber         string                   `json:"phoneNumber"`
	Address             string                   `json:"address"`
	BirthDate           int64                    `json:"birthDate"`
	Password            string                   `json:"password"`
	Point               int                      `json:"point"`
	PrestigeRank        int                      `json:"prestigeRank"`
	ConfirmationToken   string                   `json:"confirmationToken"`
	ForgetPasswordToken string                   `json:"forgetPasswordToken"`
	Status              string                   `json:"status"`
	Role                string                   `json:"role"`
	PointStatusLogs     []map[string]interface{} `json:"pointStatusLogs"`
	Qrcode              string                   `json:"qrcode"`
}
