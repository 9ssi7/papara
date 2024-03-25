## Papara Go API Client

This library makes it easy to interact with the Papara API.

### Supported APIs

- [x] Account API
- [x] User Verification API 
- [x] Pay with Papara API 
- [ ] Pay with Vpos API 
- [ ] Pay with QR API 
- [ ] Pay with Link API 
- [ ] Mass Payment API

### Installation

```go
go get github.com/9ssi7/papara
```

### Usage

```go
import (
	"fmt"
	"github.com/9ssi7/papara"
)

func main() {
	// Set your API key and URL
	apiKey := "YOUR_API_KEY"
	apiUrl := "YOUR_API_URL"

	// Create a new Papara client
	client := papara.New(papara.Config{
		ApiKey: apiKey,
		ApiUrl: apiUrl,
	})

	// Get account information
	account, err := client.Account()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Account Information:")
	fmt.Println("  ID:", account.Data.ID)
	fmt.Println("  Balance:", account.Data.Balances[0].TotalBalance)

	// Create a payment
	payment, err := client.CreatePayment(papara.PaymentCreateRequest{
		Amount:              100.0,
		ReferenceId:         "ORDER-1234",
		OrderDescription:    "Order #1234",
		NotificationUrl:     "https://your-website.com/callback",
		FailNotificationUrl: "https://your-website.com/fail-callback",
		RedirectUrl:         "https://your-website.com/success",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Payment Created:")
	fmt.Println("  ID:", payment.Data.ID)
	fmt.Println("  Payment URL:", payment.Data.PaymentUrl)
}
```

### Contributing

If you want to contribute to this library, you are welcome to send pull requests or open issues on Github.