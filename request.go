package bri

type CreateVaRequest struct {
	InstitutionCode string `json:"institutionCode"`
	BrivaNo         string `json:"brivaNo"`
	CustCode        string `json:"custCode"`
	Nama            string `json:"nama"`
	Amount          string `json:"amount"`
	Keterangan      string `json:"keterangan"`
	ExpiredDate     string `json:"expiredDate"`
}
