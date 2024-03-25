// Package papara provides functions for interacting with the Papara API.
package papara

// PaymentResultNotification is a struct representing a payment result notification from Papara.
// Papara calls the notification URL with this struct when a payment is completed.
// If papara does not call the notification URL, the payment is not considered complete.
type PaymentResultNotification struct {
	MerchantId               string  `json:"merchantId"`
	UserId                   string  `json:"userId"`
	PaymentMethod            int     `json:"paymentMethod"`
	PaymentMethodDescription string  `json:"paymentMethodDescription"`
	ReferenceId              string  `json:"referenceId"`
	OrderDescription         string  `json:"orderDescription"`
	Status                   int     `json:"status"`
	StatusDescription        string  `json:"statusDescription"`
	Amount                   float64 `json:"amount"`
	Fee                      float64 `json:"fee"`
	Currency                 string  `json:"currency"`
	NotificationUrl          string  `json:"notificationUrl"`
	NotificationDone         bool    `json:"notificationDone"`
	RedirectUrl              string  `json:"redirectUrl"`
	MerchantSecretKey        string  `json:"merchantSecretKey"`
	PaymentUrl               string  `json:"paymentUrl"`
	ReturningRedirectUrl     string  `json:"returningRedirectUrl"`
	Id                       string  `json:"id"`
	CreatedAt                string  `json:"createdAt"`
	TurkishNationalId        string  `json:"turkishNationalId"`
}
