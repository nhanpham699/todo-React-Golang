package dto

type AuthRequest struct {
	Email    string `json:"email" bson:"email" validate:"required"`
	Password string `json:"password" bson:"password" validate:"required"`
}

type RegisterRequest struct {
	FirstName string `json:"firstname" bson:"firstname"`
	LastName  string `json:"lastname" bson:"lastname"`
	Email     string `json:"email" bson:"email"`
	Password  string `json:"password" bson:"password"`
}

type AuthResponse struct {
	ResultCode string `json:"resultCode" bson:"resultCode`
	Message    string `json:"message" bson:"message"`
	Token      string `json:"token" bson:"token"`
}
