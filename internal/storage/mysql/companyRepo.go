package mysql

import (
	"context"
	"database/sql"
	cmp "xm/pkg/company"
)

type companyRepo struct {
	db *sql.DB
}

//New instance of the company repo
func NewCompanyRepo(d *sql.DB) cmp.Repository {
	return &companyRepo{
		db: d,
	}
}

//Add a new company
func (f *companyRepo) CreateCompany(ctx context.Context, comp cmp.Company) (string, error) {
	stmt, _ := f.db.Prepare("INSERT INTO companies VALUES(?, ?, ?, ?, ?, ?)")
	//execute
	res, err := stmt.ExecContext(ctx, comp.ID, comp.Name, comp.Description, comp.AmountOfEmployees, comp.Registered, comp.Type)
	if err != nil {
		return "", err
	}
	rows, err := res.RowsAffected()
	if err != nil || rows < 1 {
		return "", err
	}
	return comp.ID, nil

}

//update a company info
func (f *companyRepo) UpdateCompany(ctx context.Context, comp cmp.Company, compId string) (string, error) {

	stmt, _ := f.db.Prepare("UPDATE companies SET name=?,description=?,amountOfEmployees=?,registered=?,type=? WHERE id=?")
	res, err := stmt.ExecContext(ctx, comp.Name, comp.Description, comp.AmountOfEmployees, comp.Registered, comp.Type, compId)
	if err != nil {
		return "", err
	}
	rows, err := res.RowsAffected()
	if err != nil || rows < 1 {
		return "", err
	}
	return compId, nil

}

//retrieve all drones
func (f *companyRepo) GetCompany(ctx context.Context, companyId string) (*cmp.Company, error) {
	var comp = cmp.Company{}
	SQL := "SELECT id,name,description,amountOfEmployees,registered,type FROM companies d WHERE id=?"
	row := f.db.QueryRowContext(ctx, SQL, companyId)
	err := row.Scan(&comp.ID, &comp.Name, &comp.Description, &comp.AmountOfEmployees, &comp.Registered, &comp.Type)
	if err != nil {
		return nil, err
	}
	return &comp, nil
}

//remove a company
func (f *companyRepo) DeleteCompany(ctx context.Context, companyId string) (string, error) {
	stmt, _ := f.db.Prepare("DELETE FROM companies d WHERE d.id = ?")
	res, err_ := stmt.ExecContext(ctx, companyId)
	if err_ != nil {
		return "", err_
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return "", err_
	}
	if rows > 0 {
		return companyId, nil
	} else {
		return "", nil
	}
}
