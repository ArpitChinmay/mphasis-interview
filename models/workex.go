package models

import "time"

type Workex struct {
	WorkexId    int       `json:"workexId"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"emdDate"`
	CompanyName string    `json:"company"`
	Role        string    `json:"role"`
	Description string    `json:"description"`
}
