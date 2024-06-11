package pkg

import (
	"context"
	"database/sql"
	"fmt"
)

// toda vez que queremos usar transacao usa o *sql.Tx.
//Ele deve retornar o repo conectado a transaction. Como nao sabemos o type od repo, voltamos interface{} que eh do tipo any
type RepositoryFactoryInterface func(tx *sql.Tx) interface{}



type UowInterface interface{
	Register(name string, fc RepositoryFactoryInterface)
	GetRepository(ctx context.Context,name string) (interface{}, error)
	// funcao eh a funcao que faz tudo ser executado
	Do(ctx context.Context, fn func(uow *UOW) error)  error
	CommitOrRollBack() error
	Rollback() error
	Unregister(name string) error
}

type UOW struct{
	Db           *sql.DB
	Tx           *sql.Tx
	Repositories map[string]RepositoryFactoryInterface
} 

func NewUOW(ctx context.Context, db *sql.DB) *UOW{
	return &UOW{
		Db: db,
		Repositories: make(map[string]RepositoryFactoryInterface),
	}
} 

func (u *UOW) Register(name string, fc RepositoryFactoryInterface) {
	// {"nome do repo" : funcao que retorna o Repo que sera usada no tx}
	u.Repositories[name] = fc
}

func (u *UOW) UnRegister(name string) {
	delete(u.Repositories, name)
}

func (u *UOW) Do(ctx context.Context, fn func(uow *UOW) error)  error {
	if u.Tx != nil {
		return fmt.Errorf("transaction already started")
	}
	
	tx, err := u.Db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	u.Tx =  tx
	// agora fica a cargo dos repos de chamarem db.New() que utiliza tudo o que tem no struct UOW 
	err = fn(u)
	if err != nil {
	if errRb := tx.Rollback(); errRb != nil {
		return fmt.Errorf("error on rollback: %v, original error: %w", errRb, err)
	}
	return err
	}

	return u.CommitOrRollBack()
}

func (u *UOW) CommitOrRollBack() error {
	err := u.Tx.Commit()
	if err != nil {
	if errRb := u.Tx.Rollback(); errRb != nil {
		return fmt.Errorf("error on rollback: %v, original error: %w", errRb, err)
	}
	return err
	}
	return nil
}

func (u *UOW) GetRepository(ctx context.Context,name string) (interface{}, error){
	// se eu pegar o repo mas tx ainda nao foi inicializada, inicia aqui. Pq? pq as vezes vc nao quer utilizar o Do(), as  vezes vc soh quer pegar o repo e fazer algo
	if u.Tx == nil {
		tx, err := u.Db.BeginTx(ctx, nil)

		if err != nil {
			return nil, err
		}

		u.Tx = tx
	}
	repo := u.Repositories[name](u.Tx)

	return repo, nil
}