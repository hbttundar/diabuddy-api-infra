package repository

import (
	"context"
	"database/sql"
	baseErrors "errors"
	"fmt"
	"github.com/hbttundar/diabuddy-api-infra/database"
	"github.com/hbttundar/diabuddy-errors"
	"log"
	"math"
	"reflect"
)

type Pagination struct {
	Data     interface{} `json:"data"`
	Page     int         `json:"page"`
	PerPage  int         `json:"per_page"`
	Total    int         `json:"total"`
	LastPage int         `json:"last_page"`
	HasNext  bool        `json:"has_next"`
	HasPrev  bool        `json:"has_prev"`
}

type BaseRepository struct {
	Connection database.Connection
}

func NewBaseRepository(connection database.Connection) *BaseRepository {
	repository := &BaseRepository{
		Connection: connection,
	}
	return repository
}

// ExecContext handles execution with transaction support
func (br *BaseRepository) ExecContext(ctx context.Context, tx *sql.Tx, query string, args ...interface{}) (sql.Result, errors.ApiErrors) {
	var (
		result sql.Result
		err    error
	)

	if tx != nil {
		result, err = tx.ExecContext(ctx, query, args...)
	} else {
		result, err = br.Connection.DB().ExecContext(ctx, query, args...)
		log.Printf("error in executing exec context:%v", err)
	}

	if err != nil {
		return nil, errors.NewApiError(errors.InternalServerErrorType, "SQL execution error", errors.WithInternalError(err))
	}
	return result, nil
}

// QueryRowContext handles querying a single row with transaction support
func (br *BaseRepository) QueryRowContext(ctx context.Context, tx *sql.Tx, query string, args ...interface{}) *sql.Row {
	if tx != nil {
		return tx.QueryRowContext(ctx, query, args...)
	} else {
		return br.Connection.DB().QueryRowContext(ctx, query, args...)
	}
}

// QueryContext handles querying multiple rows with context
func (br *BaseRepository) QueryContext(ctx context.Context, tx *sql.Tx, query string, args ...interface{}) (*sql.Rows, errors.ApiErrors) {
	var (
		rows *sql.Rows
		err  error
	)

	if tx != nil {
		rows, err = tx.QueryContext(ctx, query, args...)
	} else {
		rows, err = br.Connection.DB().QueryContext(ctx, query, args...)
	}

	if err != nil {
		log.Printf("error in executing query context: %v", err)
		return nil, errors.NewApiError(errors.InternalServerErrorType, "SQL query error", errors.WithInternalError(err))
	}
	return rows, nil
}

// ScanRow scans a row into the provided attributes and checks for errors
func (br *BaseRepository) ScanRow(row *sql.Row, attributes ...any) errors.ApiErrors {
	if err := row.Scan(attributes...); err != nil {
		if baseErrors.Is(err, sql.ErrNoRows) {
			return errors.NewApiError(errors.NotFoundErrorType, "no record found", errors.WithInternalError(err))
		}
		log.Printf("error in executing scan row :%v", err)
		return errors.NewApiError(errors.InternalServerErrorType, "error scanning row", errors.WithInternalError(err))
	}
	return nil
}

func (br *BaseRepository) ScanRows(
	rows *sql.Rows,
	results interface{},
	scanTarget func() (interface{}, []interface{}),
) errors.ApiErrors {
	defer rows.Close()

	rv := reflect.ValueOf(results)
	if rv.Kind() != reflect.Ptr || rv.Elem().Kind() != reflect.Slice {
		return errors.NewApiError(
			errors.InternalServerErrorType,
			"ScanRows: results must be a pointer to a slice",
			nil,
		)
	}
	sliceVal := rv.Elem()
	elemType := sliceVal.Type().Elem()

	for rows.Next() {
		obj, scanFields := scanTarget()
		if err := rows.Scan(scanFields...); err != nil {
			return errors.NewApiError(
				errors.InternalServerErrorType,
				"error scanning row",
				errors.WithInternalError(err),
			)
		}

		v := reflect.ValueOf(obj)
		if !v.Type().AssignableTo(elemType) {
			return errors.NewApiError(
				errors.InternalServerErrorType,
				fmt.Sprintf("ScanRows: scanTarget returned %s, but expected %s",
					v.Type(), elemType,
				),
				nil,
			)
		}

		sliceVal.Set(reflect.Append(sliceVal, v))
	}

	if err := rows.Err(); err != nil {
		return errors.NewApiError(
			errors.InternalServerErrorType,
			"error after row iteration",
			errors.WithInternalError(err),
		)
	}
	return nil
}

// ParseResult parses an executed query result and checks for errors
func (br *BaseRepository) ParseResult(result sql.Result, operationType string) errors.ApiErrors {
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.NewApiError(errors.InternalServerErrorType, "error retrieving rows affected", errors.WithInternalError(err))
	}
	if rowsAffected == 0 {
		return errors.NewApiError(errors.NotFoundErrorType, fmt.Sprintf("no records found to %s", operationType))
	}
	return nil
}

// Paginate runs a COUNT(*) query, then your SELECT...LIMIT/OFFSET,
// and scans into the slice pointed at by ‘resultsPtr’ via ScanRows.
// - countQuery & countArgs:    e.g. "SELECT COUNT(*) FROM users", nil
// - dataQuery & dataArgs:      e.g. "SELECT id,… FROM users ORDER BY … LIMIT $1 OFFSET $2", []interface{}{perPage, offset}
// - scanTarget:                factory to make one row object + its []interface{} of field-ptrs
// - resultsPtr:                pointer to a slice (e.g. *[]*User)
func (br *BaseRepository) Paginate(ctx context.Context, tx *sql.Tx, page, perPage int, countQuery string, countArgs []interface{}, dataQuery string, dataArgs []interface{}, scanTarget func() (interface{}, []interface{}), results interface{}) (*Pagination, errors.ApiErrors) {
	row := br.QueryRowContext(ctx, tx, countQuery, countArgs...)
	var total int
	if err := br.ScanRow(row, &total); err != nil {
		return nil, err
	}

	offset := (page - 1) * perPage
	args := append(dataArgs, perPage, offset)
	rows, err := br.QueryContext(ctx, tx, dataQuery, args...)
	if err != nil {
		return nil, err
	}
	if err = br.ScanRows(rows, results, scanTarget); err != nil {
		return nil, err
	}

	lastPage := int(math.Ceil(float64(total) / float64(perPage)))
	return &Pagination{
		Data:     results,
		Page:     page,
		PerPage:  perPage,
		Total:    total,
		LastPage: lastPage,
		HasNext:  page < lastPage,
		HasPrev:  page > 1,
	}, nil
}

func (br *BaseRepository) Close() errors.ApiErrors {
	if br.Connection != nil {
		return br.Connection.Close()
	}
	return nil
}
