package company

type CompanyType string

const (
	Corporation        CompanyType = "Corporation"
	NonProfit          CompanyType = "NonProfit"
	Cooperative        CompanyType = "Cooperative"
	SoleProprietorship CompanyType = "Sole Proprietorship"
)

// type Company struct {
// 	ID                string      `json:"id"`
// 	Name              string      `json:"name" validate:"required"`
// 	Description       string      `json:"description" validate:"max=15"`
// 	AmountOfEmployees int         `json:"amountOfEmployees" validate:"required"`
// 	Registered        bool        `json:"registered" bindvalidateing:"required"`
// 	Type              CompanyType `json:"type" validate:"required"`
// }

type Company struct {
	ID                string `json:"id"`
	Name              string `json:"name" validate:"required"`
	Description       string `json:"description" validate:"omitempty,max=15"`
	AmountOfEmployees int    `json:"amountOfEmployees" validate:"required,numeric"`
	Registered        bool   `json:"registered" validate:"required,boolean"`
	Type              string `json:"type" validate:"required"`
}

type CompanyPatch struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	Description       string `json:"description" validate:"omitempty,max=15"`
	AmountOfEmployees int    `json:"amountOfEmployees" validate:"numeric"`
	Registered        bool   `json:"registered" validate:"boolean"`
	Type              string `json:"type"`
}

// type Company struct {
// 	ID                string      `json:"id" binding:"required"`
// 	Name              string      `json:"name" binding:"required"`
// 	Description       string      `json:"description"`
// 	AmountOfEmployees int         `json:"amountOfEmployees" binding:"required"`
// 	Registered        bool        `json:"registered" binding:"required"`
// 	Type              CompanyType `json:"type" binding:"required"`
// }
