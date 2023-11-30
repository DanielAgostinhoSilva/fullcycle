package database

import (
	"database/sql"
	"math"
)

type Pageable struct {
	PageSize      int
	PageNumber    int
	TotalPages    int
	TotalElements int
}

func QueryWithPagination(db *sql.DB, query string, page int, pageSize int) (*sql.Rows, *Pageable, error) {
	var totalItems int
	err := db.QueryRow("SELECT COUNT(*) FROM (" + query + ") AS count").Scan(&totalItems)
	if err != nil {
		return nil, nil, err
	}

	offset := (page - 1) * pageSize
	limit := pageSize
	stmt, err := db.Prepare(query + " LIMIT ? OFFSET ?")
	if err != nil {
		return nil, nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(limit, offset)
	if err != nil {
		return nil, nil, err
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(pageSize)))
	return rows, &Pageable{
		PageSize:      pageSize,
		PageNumber:    page,
		TotalPages:    totalPages,
		TotalElements: totalItems,
	}, nil
}
