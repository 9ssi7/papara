// Package papara provides functions for interacting with the Papara API.
package papara

// AccountLedgersFilterRequest is a struct representing the request parameters for the AccountLedgers method.
type AccountLedgersFilterRequest struct {
	StartDate     string `json:"startDate"`
	EndDate       string `json:"endDate"`
	Page          int    `json:"page"`
	PageSize      int    `json:"pageSize"`
	EntryType     *int   `json:"entryType,omitempty"`
	AccountNumber *int   `json:"accountNumber,omitempty"`
}

// AccountSettlementRequest is a struct representing the request parameters for the AccountSettlement method.
type AccountSettlementRequest struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	EntryType *int   `json:"entryType,omitempty"`
}

// PaymentCreateRequest is a struct representing the request parameters for the CreatePayment method.
type PaymentCreateRequest struct {
	Amount              float64 `json:"amount"`
	ReferenceId         string  `json:"referenceId"`
	OrderDescription    string  `json:"orderDescription"`
	NotificationUrl     string  `json:"notificationUrl"`
	FailNotificationUrl string  `json:"failNotificationUrl"`
	RedirectUrl         string  `json:"redirectUrl"`
	TurkishNationalId   *string `json:"turkishNationalId,omitempty"`
	Currency            *int    `json:"currency,omitempty"`
}

// PaymentRefundRequest is a struct representing the request parameters for the RefundPayment method.
type PaymentRefundRequest struct {
	PaymentId    string `json:"paymentId"`
	RefundAmount string `json:"refundAmount"`
}
