package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

// todo rename all variables
type Transaction struct {
	TransactionId   int64   `json:"transaction_id"`
	SenderId        int64   `json:"sender_id"`
	ReceiverId      int64   `json:"receiver_id"`
	OperationId     int64   `json:"operation_id"`
	TransactionTime string  `json:"transaction_time"`
	Money           float64 `json:"transaction_price"`
}
type TransactionModel struct {
	DB *sql.DB
}
type TransactionHistory struct {
	TransactionId    int64   `json:"transaction_id"`
	SenderId         int64   `json:"sender_id"`
	ReceiverId       int64   `json:"receiver_id"`
	TransactionTime  string  `json:"transaction_time"`
	TransactionPrice float64 `json:"transaction_price"`
	OperationType    string  `json:"operation_type"`
}

type TransactionHistoryModel struct {
	DB *sql.DB
}

type TransactionHistoryResponse struct {
	//AccountId        int64   `json:"transaction_id"`
	SenderId         int64   `json:"sender_id"`
	ReceiverId       int64   `json:"receiver_id"`
	TransactionTime  string  `json:"transaction_time"`
	TransactionPrice float64 `json:"transaction_price"`
	OperationType    string  `json:"operation_type"`
}

type UserReportHistoryResponse struct {
	AccountId    int64   `json:"account_id"`
	ReportTime   string  `json:"report_time"`
	ServicePrice float64 `json:"service_price"`
	ServiceName  string  `json:"service_name"`
}

// todo!!! returns data with paginating feature
func (t TransactionHistoryModel) GetUserTransactionHistory(id int64, filters Filters) ([]TransactionHistoryResponse, Metadata, error) {
	if id < 1 {
		return nil, Metadata{}, errors.New("incorrect id")
	}
	query := fmt.Sprintf(`
			SELECT COUNT(*) OVER(), sender_id, receiver_id, transaction_time, transaction_price, operation.operation_type 
			FROM transaction_history, operation
			WHERE (sender_id = $1 OR receiver_id = $1) AND (operation.operation_id = transaction_history.operation_id)
			ORDER BY %s %s, sender_id ASC
			LIMIT $2 OFFSET $3
		`, filters.sortColumnTransactionQuery(), filters.sortDirection())
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := t.DB.QueryContext(ctx, query, id, filters.limit(), filters.offset())
	if err != nil {
		return nil, Metadata{}, errors.New("No records")
	}

	defer rows.Close()

	totalRecords := 0

	var tranHistory []TransactionHistoryResponse

	for rows.Next() {
		var history TransactionHistoryResponse
		//var str = make([]string, 0)
		err := rows.Scan(&totalRecords, &history.SenderId, &history.ReceiverId,
			&history.TransactionTime, &history.TransactionPrice, &history.OperationType)
		if err != nil {
			return nil, Metadata{}, err
		}
		tranHistory = append(tranHistory, history)
	}

	if err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return tranHistory, metadata, nil
}

func (t TransactionModel) CreateTransaction(senderId, receiverId, operationId int64, transaction_price float64) {

	query := `
		INSERT INTO transaction_history(sender_id, receiver_id, operation_id, transaction_price)
		VALUES ($1, $2, $3, $4)`

	t.DB.QueryRow(query, senderId, receiverId, operationId, transaction_price).Scan()
}
func (t TransactionModel) GetTransactionById(id int64) (*Transaction, error) {
	if id < 1 {
		return nil, errors.New("incorrect id")
	}

	query := `
			SELECT *
			FROM transaction_history
			WHERE transaction_id = $1
		`

	var transaction Transaction
	err := t.DB.QueryRow(query, id).Scan(
		&transaction.TransactionId,
		&transaction.SenderId,
		&transaction.ReceiverId,
		&transaction.OperationId,
		&transaction.TransactionTime,
		&transaction.Money,
	)

	if err != nil {
		return nil, err
	}

	return &transaction, nil
}
