// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createAccountStmt, err = db.PrepareContext(ctx, createAccount); err != nil {
		return nil, fmt.Errorf("error preparing query CreateAccount: %w", err)
	}
	if q.createCategoryStmt, err = db.PrepareContext(ctx, createCategory); err != nil {
		return nil, fmt.Errorf("error preparing query CreateCategory: %w", err)
	}
	if q.createUserStmt, err = db.PrepareContext(ctx, createUser); err != nil {
		return nil, fmt.Errorf("error preparing query CreateUser: %w", err)
	}
	if q.deleteAccountStmt, err = db.PrepareContext(ctx, deleteAccount); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteAccount: %w", err)
	}
	if q.deleteCategoryStmt, err = db.PrepareContext(ctx, deleteCategory); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteCategory: %w", err)
	}
	if q.getAccountStmt, err = db.PrepareContext(ctx, getAccount); err != nil {
		return nil, fmt.Errorf("error preparing query GetAccount: %w", err)
	}
	if q.getAccountsStmt, err = db.PrepareContext(ctx, getAccounts); err != nil {
		return nil, fmt.Errorf("error preparing query GetAccounts: %w", err)
	}
	if q.getAccountsGraphStmt, err = db.PrepareContext(ctx, getAccountsGraph); err != nil {
		return nil, fmt.Errorf("error preparing query GetAccountsGraph: %w", err)
	}
	if q.getAccountsReportsStmt, err = db.PrepareContext(ctx, getAccountsReports); err != nil {
		return nil, fmt.Errorf("error preparing query GetAccountsReports: %w", err)
	}
	if q.getCategoriesStmt, err = db.PrepareContext(ctx, getCategories); err != nil {
		return nil, fmt.Errorf("error preparing query GetCategories: %w", err)
	}
	if q.getCategoryStmt, err = db.PrepareContext(ctx, getCategory); err != nil {
		return nil, fmt.Errorf("error preparing query GetCategory: %w", err)
	}
	if q.getUserStmt, err = db.PrepareContext(ctx, getUser); err != nil {
		return nil, fmt.Errorf("error preparing query GetUser: %w", err)
	}
	if q.getUserByIdStmt, err = db.PrepareContext(ctx, getUserById); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserById: %w", err)
	}
	if q.updateAccountsStmt, err = db.PrepareContext(ctx, updateAccounts); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateAccounts: %w", err)
	}
	if q.updateCategoriesStmt, err = db.PrepareContext(ctx, updateCategories); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateCategories: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createAccountStmt != nil {
		if cerr := q.createAccountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createAccountStmt: %w", cerr)
		}
	}
	if q.createCategoryStmt != nil {
		if cerr := q.createCategoryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createCategoryStmt: %w", cerr)
		}
	}
	if q.createUserStmt != nil {
		if cerr := q.createUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createUserStmt: %w", cerr)
		}
	}
	if q.deleteAccountStmt != nil {
		if cerr := q.deleteAccountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteAccountStmt: %w", cerr)
		}
	}
	if q.deleteCategoryStmt != nil {
		if cerr := q.deleteCategoryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteCategoryStmt: %w", cerr)
		}
	}
	if q.getAccountStmt != nil {
		if cerr := q.getAccountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAccountStmt: %w", cerr)
		}
	}
	if q.getAccountsStmt != nil {
		if cerr := q.getAccountsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAccountsStmt: %w", cerr)
		}
	}
	if q.getAccountsGraphStmt != nil {
		if cerr := q.getAccountsGraphStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAccountsGraphStmt: %w", cerr)
		}
	}
	if q.getAccountsReportsStmt != nil {
		if cerr := q.getAccountsReportsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAccountsReportsStmt: %w", cerr)
		}
	}
	if q.getCategoriesStmt != nil {
		if cerr := q.getCategoriesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getCategoriesStmt: %w", cerr)
		}
	}
	if q.getCategoryStmt != nil {
		if cerr := q.getCategoryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getCategoryStmt: %w", cerr)
		}
	}
	if q.getUserStmt != nil {
		if cerr := q.getUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserStmt: %w", cerr)
		}
	}
	if q.getUserByIdStmt != nil {
		if cerr := q.getUserByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserByIdStmt: %w", cerr)
		}
	}
	if q.updateAccountsStmt != nil {
		if cerr := q.updateAccountsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateAccountsStmt: %w", cerr)
		}
	}
	if q.updateCategoriesStmt != nil {
		if cerr := q.updateCategoriesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateCategoriesStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                     DBTX
	tx                     *sql.Tx
	createAccountStmt      *sql.Stmt
	createCategoryStmt     *sql.Stmt
	createUserStmt         *sql.Stmt
	deleteAccountStmt      *sql.Stmt
	deleteCategoryStmt     *sql.Stmt
	getAccountStmt         *sql.Stmt
	getAccountsStmt        *sql.Stmt
	getAccountsGraphStmt   *sql.Stmt
	getAccountsReportsStmt *sql.Stmt
	getCategoriesStmt      *sql.Stmt
	getCategoryStmt        *sql.Stmt
	getUserStmt            *sql.Stmt
	getUserByIdStmt        *sql.Stmt
	updateAccountsStmt     *sql.Stmt
	updateCategoriesStmt   *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                     tx,
		tx:                     tx,
		createAccountStmt:      q.createAccountStmt,
		createCategoryStmt:     q.createCategoryStmt,
		createUserStmt:         q.createUserStmt,
		deleteAccountStmt:      q.deleteAccountStmt,
		deleteCategoryStmt:     q.deleteCategoryStmt,
		getAccountStmt:         q.getAccountStmt,
		getAccountsStmt:        q.getAccountsStmt,
		getAccountsGraphStmt:   q.getAccountsGraphStmt,
		getAccountsReportsStmt: q.getAccountsReportsStmt,
		getCategoriesStmt:      q.getCategoriesStmt,
		getCategoryStmt:        q.getCategoryStmt,
		getUserStmt:            q.getUserStmt,
		getUserByIdStmt:        q.getUserByIdStmt,
		updateAccountsStmt:     q.updateAccountsStmt,
		updateCategoriesStmt:   q.updateCategoriesStmt,
	}
}
