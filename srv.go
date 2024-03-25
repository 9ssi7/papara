package papara

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/goccy/go-json"
)

type Service interface {
	Account() (*AccountResult, error)
	AccountLedgers(req AccountLedgersFilterRequest) (*AccountLedgersResult, error)
	AccountSettlement(req AccountSettlementRequest) (*AccountSettlementResult, error)
	ValidateWithAccountNumber(accountNumber int) (*ValidationResult, error)
	ValidateWithPhoneNumber(phoneNumber string) (*ValidationResult, error)
	ValidateWithTckn(tckn string) (*ValidationResult, error)

	CreatePayment(req PaymentCreateRequest) (*PaymentCreateResult, error)
	GetPayment(paymentId string) (*PaymentGetResult, error)
	CancelPayment(paymentId string) (*BasicResult, error)
	RefundPayment(req PaymentRefundRequest) (*BasicResult, error)
}

type service struct {
	apiKey string
	apiUrl string
}

func New(cnf Config) Service {
	return &service{
		apiKey: cnf.ApiKey,
		apiUrl: cnf.ApiUrl,
	}
}

func (s *service) Account() (*AccountResult, error) {
	url := s.makeUrl("/account")
	headers := s.getApiKeyHeader()
	return runCommand[any, AccountResult]("GET", url, nil, headers)
}

func (s *service) AccountLedgers(req AccountLedgersFilterRequest) (*AccountLedgersResult, error) {
	url := s.makeUrl("/account/ledgers")
	headers := s.getApiKeyHeader()
	return runCommand[AccountLedgersFilterRequest, AccountLedgersResult]("POST", url, &req, headers)
}

func (s *service) AccountSettlement(req AccountSettlementRequest) (*AccountSettlementResult, error) {
	url := s.makeUrl("/account/settlement")
	headers := s.getApiKeyHeader()
	return runCommand[AccountSettlementRequest, AccountSettlementResult]("POST", url, &req, headers)
}

func (s *service) ValidateWithAccountNumber(accountNumber int) (*ValidationResult, error) {
	url := s.makeUrl("/validation/accountNumber?accountNumber=" + strconv.Itoa(accountNumber))
	headers := s.getApiKeyHeader()
	return runCommand[any, ValidationResult]("GET", url, nil, headers)
}

func (s *service) ValidateWithPhoneNumber(phoneNumber string) (*ValidationResult, error) {
	url := s.makeUrl("/validation/phoneNumber?phoneNumber=" + phoneNumber)
	headers := s.getApiKeyHeader()
	return runCommand[any, ValidationResult]("GET", url, nil, headers)
}

func (s *service) ValidateWithTckn(tckn string) (*ValidationResult, error) {
	url := s.makeUrl("/validation/tckn?tckn=" + tckn)
	headers := s.getApiKeyHeader()
	return runCommand[any, ValidationResult]("GET", url, nil, headers)
}

func (s *service) CreatePayment(req PaymentCreateRequest) (*PaymentCreateResult, error) {
	url := s.makeUrl("/payments")
	headers := s.getApiKeyHeader()
	return runCommand[PaymentCreateRequest, PaymentCreateResult]("POST", url, &req, headers)
}

func (s *service) GetPayment(paymentId string) (*PaymentGetResult, error) {
	url := s.makeUrl("/payments?id=" + paymentId)
	headers := s.getApiKeyHeader()
	return runCommand[any, PaymentGetResult]("GET", url, nil, headers)
}

func (s *service) CancelPayment(paymentId string) (*BasicResult, error) {
	url := s.makeUrl("/payments?id=" + paymentId)
	headers := s.getApiKeyHeader()
	return runCommand[any, BasicResult]("PUT", url, nil, headers)
}

func (s *service) RefundPayment(req PaymentRefundRequest) (*BasicResult, error) {
	url := s.makeUrl("/payments/refund")
	headers := s.getApiKeyHeader()
	return runCommand[PaymentRefundRequest, BasicResult]("POST", url, &req, headers)
}

func (s *service) makeUrl(path string) string {
	return s.apiUrl + path
}

func (s *service) getApiKeyHeader() map[string]string {
	return map[string]string{
		"ApiKey": s.apiKey,
	}
}

func runCommand[Req interface{}, Res any](method string, url string, req *Req, headers ...map[string]string) (*Res, error) {
	var body []byte
	if req != nil {
		b, err := json.Marshal(req)
		if err != nil {
			return nil, err
		}
		body = b
	}
	client := &http.Client{}
	r, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Accept", "application/json")
	if len(headers) > 0 {
		for k, v := range headers[0] {
			r.Header.Set(k, v)
		}
	}
	resp, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var res Res
	if err := json.Unmarshal(b, &res); err != nil {
		fmt.Println("error unmarshalling response", err)
		return nil, err
	}
	return &res, nil
}
