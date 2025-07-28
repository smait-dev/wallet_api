package repository

import (
	"github.com/jmoiron/sqlx"
	"wallet_api/internal/entity"
	"wallet_api/internal/errors"
)

// ITransactionRepository интерфейс для работы с транзакциями.
type ITransactionRepository interface {
	// Send выполненяет отправку средств между кошельками.
	Send(fromAddr string, toAddr string, amount float64) (err error)
	// GetLastTransactions возвращает записи транзакции в количестве count.
	GetLastTransactions(count int) ([]entity.Transaction, error)
}

// TransactionRepository реализует ITransactionRepository.
type TransactionRepository struct {
	db *sqlx.DB
}

// NewTransactionRepository возвращает новый экземпляр TransactionRepository.
func NewTransactionRepository(db *sqlx.DB) *TransactionRepository {
	return &TransactionRepository{db}
}

func (tr *TransactionRepository) Send(fromAddr, toAddr string, amount float64) (err error) {
	tx, err := tr.db.Beginx()
	if err != nil {
		return errors.ErrDatabase
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	if err = tr.debit(tx, fromAddr, amount); err != nil {
		return err
	}

	if err = tr.credit(tx, toAddr, amount); err != nil {
		return err
	}

	_, err = tx.Exec(`
		INSERT INTO transactions (from_wallet_id, to_wallet_id, amount)
		VALUES (
			(SELECT id FROM wallets WHERE address = $1),
			(SELECT id FROM wallets WHERE address = $2),
			$3
		)
	`, fromAddr, toAddr, amount)
	if err != nil {
		return errors.ErrTransactionFailed
	}

	if err = tx.Commit(); err != nil {
		return errors.ErrTransactionFailed
	}

	return nil
}

// debit уменьшает баланс указанного кошелька на заданную сумму(в рамках транзакции).
func (tr *TransactionRepository) debit(tx *sqlx.Tx, address string, amount float64) error {
	res, err := tx.Exec(`
		UPDATE wallets SET balance = balance - $1
		WHERE address = $2 AND balance >= $1
	`, amount, address)
	if err != nil {
		return errors.ErrDebitFailed
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return errors.ErrInsufficientFunds
	}
	return nil
}

// credit увеличивает баланс указанного кошелька на заданную сумму(в рамках транзакции).
func (tr *TransactionRepository) credit(tx *sqlx.Tx, address string, amount float64) error {
	_, err := tx.Exec(`
		UPDATE wallets SET balance = balance + $1
		WHERE address = $2
	`, amount, address)
	if err != nil {
		return errors.ErrCreditFailed
	}
	return nil
}

func (tr *TransactionRepository) GetLastTransactions(count int) ([]entity.Transaction, error) {
	var transactions []entity.Transaction
	err := tr.db.Select(&transactions, `
		SELECT
			w1.address AS from,
			w2.address AS to,
			t.amount
		FROM transactions t
		JOIN wallets w1 ON t.from_wallet_id = w1.id
		JOIN wallets w2 ON t.to_wallet_id = w2.id
		ORDER BY t.created_at DESC
		LIMIT $1
	`, count)
	return transactions, err
}
