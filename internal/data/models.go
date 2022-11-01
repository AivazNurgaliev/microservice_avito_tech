package data

import "database/sql"

type Models struct {
	Account            AccountModel
	Report             ReportModel
	TransactionHistory TransactionHistoryModel
	Transaction        TransactionModel
	Service            ServiceModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Account:            AccountModel{DB: db},
		Report:             ReportModel{DB: db},
		TransactionHistory: TransactionHistoryModel{DB: db},
		Transaction:        TransactionModel{DB: db},
		Service:            ServiceModel{DB: db},
	}
}
