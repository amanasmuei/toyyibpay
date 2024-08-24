package models

// PaymentRequest represents the payload to create a new bill in ToyyibPay.
type PaymentRequest struct {
	UserSecretKey   string `json:"userSecretKey"`
	CategoryCode    string `json:"categoryCode"`
	BillName        string `json:"billName"`
	BillAmount      int    `json:"billAmount"`
	BillDescription string `json:"billDescription"`
	// Add other necessary fields
}

// PaymentResponse represents the response from ToyyibPay after creating a bill.
type PaymentResponse struct {
	Status   string `json:"status"`
	BillCode string `json:"billCode"`
	// Add other necessary fields
}
