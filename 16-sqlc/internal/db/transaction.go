package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Transaction struct {
	dbConn *sql.DB
}

func NewTransaction(dbConn *sql.DB) *Transaction {
return &Transaction{
		dbConn:  dbConn,
	}
}

func (t *Transaction)callTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := t.dbConn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	// iniciou um novo *Queries (onde possui todos os metodos) injentando uma transacao. ou seja, agr o que ocorrer aqui eh dentro da transaction
	q := New(tx)

	// executa a funcao anonima vinda de q (que nao as queries que serao executadas que utilizam transaction)
	err = fn(q)
	if err != nil {
		if errRb := tx.Rollback(); errRb != nil {
			return fmt.Errorf("error on rollback: %v, original error: %w", errRb, err)
		}
		return err
	}
	return tx.Commit()
}