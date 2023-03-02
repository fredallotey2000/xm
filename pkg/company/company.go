package company

type CompanyType string

const (
	Corporation        CompanyType = "Corporation"
	NonProfit          CompanyType = "NonProfit"
	Cooperative        CompanyType = "Cooperative"
	SoleProprietorship CompanyType = "Sole Proprietorship"
)

type Company struct {
	ID                string      `json:"id"`
	Name              string      `json:"name" binding:"required"`
	Description       string      `json:"description" binding:"max=15"`
	AmountOfEmployees int         `json:"amountOfEmployees" binding:"required"`
	Registered        bool        `json:"registered" binding:"required"`
	Type              CompanyType `json:"type" binding:"required"`
}

// type Company struct {
// 	ID                string      `json:"id" binding:"required"`
// 	Name              string      `json:"name" binding:"required"`
// 	Description       string      `json:"description"`
// 	AmountOfEmployees int         `json:"amountOfEmployees" binding:"required"`
// 	Registered        bool        `json:"registered" binding:"required"`
// 	Type              CompanyType `json:"type" binding:"required"`
// }
