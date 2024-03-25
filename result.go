// Package papara provides functions for interacting with the Papara API.
package papara

// AccountResultData is a struct representing account data returned by the Papara API.
type AccountResultData struct {
	ID       string `json:"id"`
	Balances []struct {
		Currency         int     `json:"currency"`
		TotalBalance     float64 `json:"totalBalance"`
		LockedBalance    float64 `json:"lockedBalance"`
		AvailableBalance float64 `json:"availableBalance"`
		Iban             any     `json:"iban"`
		CurrencyInfo     struct {
			CurrencyEnum                         int    `json:"currencyEnum"`
			Symbol                               string `json:"symbol"`
			Code                                 string `json:"code"`
			Number                               int    `json:"number"`
			PreferredDisplayCode                 string `json:"preferredDisplayCode"`
			Name                                 string `json:"name"`
			IsCryptocurrency                     bool   `json:"isCryptocurrency"`
			IsInternationalMoneyTransferCurrency bool   `json:"isInternationalMoneyTransferCurrency"`
			Precision                            int    `json:"precision"`
			IconURL                              string `json:"iconUrl"`
			FlagURL                              string `json:"flagUrl"`
			CurrencyEnumIso                      int    `json:"currencyEnumIso"`
			IsMetalCurrency                      bool   `json:"isMetalCurrency"`
		} `json:"currencyInfo"`
	} `json:"balances"`
	LegalName           string `json:"legalName"`
	IconURL             any    `json:"iconUrl"`
	BrandName           string `json:"brandName"`
	AllowedPaymentTypes []struct {
		PaymentMethod int `json:"paymentMethod"`
	} `json:"allowedPaymentTypes"`
	AllowingGuestCheckout                  bool `json:"allowingGuestCheckout"`
	IsCorporateCardsListOnPaparaAppEnabled bool `json:"isCorporateCardsListOnPaparaAppEnabled"`
	AtmDepositEnabled                      bool `json:"atmDepositEnabled"`
}

// AccountResult is a struct representing the complete account result from the Papara API.
type AccountResult struct {
	Data      *AccountResultData `json:"data"`
	Succeeded bool               `json:"succeeded"`
	Error     *ErrorResult       `json:"error"`
}

// AccountLedgersResult is a struct representing the account ledger result from the Papara API.
type AccountLedgersResult struct {
	Data struct {
		Items []struct {
			Id            int    `json:"id"`
			CreatedAt     string `json:"createdAt"`
			EntryType     int    `json:"entryType"`
			EntryTypeName string `json:"entryTypeName"`
			Amount        int    `json:"amount"`
			Fee           int    `json:"fee"`
			Currency      int    `json:"currency"`
			CurrencyInfo  struct {
				CurrencyEnum         int    `json:"currencyEnum"`
				Symbol               string `json:"symbol"`
				Code                 string `json:"code"`
				PreferredDisplayCode string `json:"preferredDisplayCode"`
				Name                 string `json:"name"`
				IsCryptocurrency     bool   `json:"isCryptocurrency"`
				Precision            int    `json:"precision"`
				IconUrl              string `json:"iconUrl"`
			} `json:"currencyInfo"`
			ResultingBalance           int    `json:"resultingBalance"`
			Description                string `json:"description"`
			MassPaymentId              *int   `json:"massPaymentId,omitempty"`
			CheckoutPaymentId          *int   `json:"checkoutPaymentId,omitempty"`
			CheckoutPaymentReferenceId *int   `json:"checkoutPaymentReferenceId,omitempty"`
		} `json:"items"`
		Page           int `json:"page"`
		PageItemCount  int `json:"pageItemCount"`
		TotalItemCount int `json:"totalItemCount"`
		TotalPageCount int `json:"totalPageCount"`
		PageSkip       int `json:"pageSkip"`
	} `json:"data"`
	Succeeded            bool         `json:"succeeded"`
	Error                *ErrorResult `json:"error"`
	ServiceResultSuccess any          `json:"serviceResultSuccess"`
}

type AccountSettlementResult struct {
	Data struct {
		Count  int `json:"count"`
		Volume int `json:"volume"`
	} `json:"data"`
	Succeeded bool         `json:"succeeded"`
	Error     *ErrorResult `json:"error"`
}

type ValidationResult struct {
	Data struct {
		UserId        string `json:"userId"`
		FirstName     string `json:"firstName"`
		LastName      string `json:"lastName"`
		Tckn          string `json:"tckn"`
		AccountNumber int    `json:"accountNumber"`
	} `json:"data"`
	Succeeded bool         `json:"succeeded"`
	Error     *ErrorResult `json:"error"`
}

type PaymentCreateResult struct {
	Data struct {
		Merchant struct {
			LegalName           string `json:"legalName"`
			BrandName           string `json:"brandName"`
			AllowedPaymentTypes []struct {
				PaymentMethod int `json:"paymentMethod"`
			} `json:"allowedPaymentTypes"`
			Balances []struct {
				Currency         int `json:"currency"`
				TotalBalance     int `json:"totalBalance"`
				LockedBalance    int `json:"lockedBalance"`
				AvailableBalance int `json:"availableBalance"`
			} `json:"balances"`
		} `json:"merchant"`
		UserId                   string `json:"userId"`
		PaymentMethod            int    `json:"paymentMethod"`
		PaymentMethodDescription string `json:"paymentMethodDescription"`
		ReferenceId              string `json:"referenceId"`
		OrderDescription         string `json:"orderDescription"`
		Status                   int    `json:"status"`
		StatusDescription        string `json:"statusDescription"`
		Amount                   int    `json:"amount"`
		Fee                      int    `json:"fee"`
		Currency                 int    `json:"currency"`
		NotificationUrl          string `json:"notificationUrl"`
		NotificationDone         bool   `json:"notificationDone"`
		RedirectUrl              string `json:"redirectUrl"`
		PaymentUrl               string `json:"paymentUrl"`
		MerchantSecretKey        string `json:"merchantSecretKey"`
		ReturningRedirectUrl     string `json:"returningRedirectUrl"`
		TurkishNationalId        int    `json:"turkishNationalId"`
		Id                       string `json:"id"`
		CreatedAt                string `json:"createdAt"`
	} `json:"data"`
	Succeeded bool         `json:"succeeded"`
	Error     *ErrorResult `json:"error"`
}

type PaymentGetResult struct {
	Data struct {
		Merchant struct {
			ID       string `json:"id"`
			Balances []struct {
				Currency         int     `json:"currency"`
				TotalBalance     float64 `json:"totalBalance"`
				LockedBalance    float64 `json:"lockedBalance"`
				AvailableBalance float64 `json:"availableBalance"`
				Iban             any     `json:"iban"`
				CurrencyInfo     struct {
					CurrencyEnum                         int    `json:"currencyEnum"`
					Symbol                               string `json:"symbol"`
					Code                                 string `json:"code"`
					Number                               int    `json:"number"`
					PreferredDisplayCode                 string `json:"preferredDisplayCode"`
					Name                                 string `json:"name"`
					IsCryptocurrency                     bool   `json:"isCryptocurrency"`
					IsInternationalMoneyTransferCurrency bool   `json:"isInternationalMoneyTransferCurrency"`
					Precision                            int    `json:"precision"`
					IconURL                              string `json:"iconUrl"`
					FlagURL                              string `json:"flagUrl"`
					CurrencyEnumIso                      int    `json:"currencyEnumIso"`
					IsMetalCurrency                      bool   `json:"isMetalCurrency"`
				} `json:"currencyInfo"`
			} `json:"balances"`
			LegalName           string `json:"legalName"`
			IconURL             any    `json:"iconUrl"`
			BrandName           string `json:"brandName"`
			AllowedPaymentTypes []struct {
				PaymentMethod int `json:"paymentMethod"`
			} `json:"allowedPaymentTypes"`
			AllowingGuestCheckout                  bool `json:"allowingGuestCheckout"`
			IsCorporateCardsListOnPaparaAppEnabled bool `json:"isCorporateCardsListOnPaparaAppEnabled"`
			AtmDepositEnabled                      bool `json:"atmDepositEnabled"`
		} `json:"merchant"`
		UserName                 any     `json:"userName"`
		RelatedTransactions      []any   `json:"relatedTransactions"`
		TotalRefundedAmount      float64 `json:"totalRefundedAmount"`
		ID                       string  `json:"id"`
		CreatedAt                string  `json:"createdAt"`
		MerchantID               string  `json:"merchantId"`
		UserID                   any     `json:"userId"`
		PaymentMethod            int     `json:"paymentMethod"`
		PaymentMethodDescription string  `json:"paymentMethodDescription"`
		ReferenceID              string  `json:"referenceId"`
		OrderDescription         string  `json:"orderDescription"`
		Status                   int     `json:"status"`
		StatusDescription        string  `json:"statusDescription"`
		Amount                   float64 `json:"amount"`
		Fee                      float64 `json:"fee"`
		Currency                 int     `json:"currency"`
		CurrencyInfo             struct {
			CurrencyEnum                         int    `json:"currencyEnum"`
			Symbol                               string `json:"symbol"`
			Code                                 string `json:"code"`
			Number                               int    `json:"number"`
			PreferredDisplayCode                 string `json:"preferredDisplayCode"`
			Name                                 string `json:"name"`
			IsCryptocurrency                     bool   `json:"isCryptocurrency"`
			IsInternationalMoneyTransferCurrency bool   `json:"isInternationalMoneyTransferCurrency"`
			Precision                            int    `json:"precision"`
			IconURL                              string `json:"iconUrl"`
			FlagURL                              string `json:"flagUrl"`
			CurrencyEnumIso                      int    `json:"currencyEnumIso"`
			IsMetalCurrency                      bool   `json:"isMetalCurrency"`
		} `json:"currencyInfo"`
		NotificationURL       string `json:"notificationUrl"`
		FailNotificationURL   any    `json:"failNotificationUrl"`
		NotificationDone      bool   `json:"notificationDone"`
		PaymentURL            string `json:"paymentUrl"`
		MerchantSecretKey     any    `json:"merchantSecretKey"`
		RemainingRefundAmount any    `json:"remainingRefundAmount"`
		ReturningRedirectURL  string `json:"returningRedirectUrl"`
		ErrorCode             any    `json:"errorCode"`
		ErrorMessage          any    `json:"errorMessage"`
		TurkishNationalID     int    `json:"turkishNationalId"`
	} `json:"data"`
	Succeeded bool         `json:"succeeded"`
	Error     *ErrorResult `json:"error"`
}

type ErrorResult struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type BasicResult struct {
	Succeeded bool         `json:"succeeded"`
	Error     *ErrorResult `json:"error"`
}
