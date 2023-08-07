package bamboogo

import "net/http"

// Client
type Client struct {
	HostURL    string
	Company    string
	HTTPClient *http.Client
}

type User struct {
	ID         int    `json:"id"`
	EmployeeID int    `json:"employeeId"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	Status     string `json:"status"`
	LastLogin  string `json:"lastLogin"`
}

type Users map[string]User
