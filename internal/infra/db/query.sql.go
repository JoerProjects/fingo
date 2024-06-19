// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package accounting_db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const addTransaction = `-- name: AddTransaction :exec
INSERT INTO accounting_transaction (
  id,
  project_id,
  transaction_type_id,
  account_number,
  memo,
  posting_date,
  created_by,
  created_at
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
`

type AddTransactionParams struct {
	ID                string
	ProjectID         pgtype.Text
	TransactionTypeID int32
	AccountNumber     pgtype.Text
	Memo              pgtype.Text
	PostingDate       pgtype.Timestamptz
	CreatedBy         string
	CreatedAt         pgtype.Timestamp
}

func (q *Queries) AddTransaction(ctx context.Context, arg AddTransactionParams) error {
	_, err := q.db.Exec(ctx, addTransaction,
		arg.ID,
		arg.ProjectID,
		arg.TransactionTypeID,
		arg.AccountNumber,
		arg.Memo,
		arg.PostingDate,
		arg.CreatedBy,
		arg.CreatedAt,
	)
	return err
}

const getTransaction = `-- name: GetTransaction :one
SELECT id, project_id, transaction_type_id, account_number, memo, posting_date, updated_by, updated_at, created_by, created_at FROM accounting_transaction WHERE id = $1
`

func (q *Queries) GetTransaction(ctx context.Context, id string) (AccountingTransaction, error) {
	row := q.db.QueryRow(ctx, getTransaction, id)
	var i AccountingTransaction
	err := row.Scan(
		&i.ID,
		&i.ProjectID,
		&i.TransactionTypeID,
		&i.AccountNumber,
		&i.Memo,
		&i.PostingDate,
		&i.UpdatedBy,
		&i.UpdatedAt,
		&i.CreatedBy,
		&i.CreatedAt,
	)
	return i, err
}

const getTransactions = `-- name: GetTransactions :many
SELECT id, project_id, transaction_type_id, account_number, memo, posting_date, updated_by, updated_at, created_by, created_at FROM accounting_transaction
`

func (q *Queries) GetTransactions(ctx context.Context) ([]AccountingTransaction, error) {
	rows, err := q.db.Query(ctx, getTransactions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AccountingTransaction
	for rows.Next() {
		var i AccountingTransaction
		if err := rows.Scan(
			&i.ID,
			&i.ProjectID,
			&i.TransactionTypeID,
			&i.AccountNumber,
			&i.Memo,
			&i.PostingDate,
			&i.UpdatedBy,
			&i.UpdatedAt,
			&i.CreatedBy,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}