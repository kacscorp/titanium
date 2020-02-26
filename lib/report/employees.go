package report

// Employee ...
type Employee struct {
	ID                   int64  `json:"id"`
	IdentificationNumber string `json:"identification_number"`
	FirstName            string `json:"first_name"`
	LastName             string `json:"last_name"`
	Age                  int    `json:"age"`
}

// Employees ...
type Employees []Employee
