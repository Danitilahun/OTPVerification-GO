package model

type OTPData struct {
	phoneNumber string  `json: "phoneNumber, omitempty" validate:"required`;
}
type VerifyOTPData struct {
	User *OTPData  `json: "user, omitempty" validate:"required`;
	Code string  `json: "code, omitempty" validate:"required`;
}