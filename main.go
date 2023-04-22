package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jackc/pgx/v5"
)

func processTable(ctx context.Context, db *pgx.Conn, tableName string) error {
	tableColumns, err := getTableColumns(ctx, db, tableName)
	if err != nil {
		return fmt.Errorf("could not get table columns: %w", err)
	}
	updateSQL := generateUpdateSQL(tableName, tableColumns)
	log.Printf("%s", updateSQL)
	return nil
}

func generateUpdateSQL(tableName string, tableColumns []tableColumn) string {
	var exprs []string
	for _, c := range tableColumns {
		updateSQL := generateUpdateSQLForColumn(c)
		if updateSQL != "" {
			exprs = append(exprs, updateSQL)
		}
	}
	return fmt.Sprintf("UPDATE %s SET\n%s\n;", tableName, strings.Join(exprs, ",\n"))
}

func generateUpdateSQLForColumn(c tableColumn) string {
	if c.DataType == "character varying" {
		return fmt.Sprintf("%s = 'REDACTED' /* %s */", c.ColumnName, c.DataType)
	}
	if c.DataType == "text" {
		return fmt.Sprintf("%s = 'REDACTED' /* %s */", c.ColumnName, c.DataType)
	}
	return ""
}

func getTableNames(ctx context.Context, db *pgx.Conn) ([]string, error) {
	q, err := db.Query(ctx, `SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'`)
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	defer q.Close()
	rows, err := pgx.CollectRows(q, pgx.RowTo[string])
	if err != nil {
		return nil, fmt.Errorf("collect: %w", err)
	}
	return rows, nil
}

type tableColumn struct {
	ColumnName string
	DataType   string
}

func getTableColumns(ctx context.Context, db *pgx.Conn, tableName string) ([]tableColumn, error) {
	q, err := db.Query(ctx,
		`SELECT column_name, data_type FROM information_schema.columns WHERE table_schema = 'public' AND table_name = $1`,
		tableName,
	)
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	defer q.Close()
	tableColumns, err := pgx.CollectRows(q, pgx.RowToStructByPos[tableColumn])
	if err != nil {
		return nil, fmt.Errorf("collect: %w", err)
	}
	return tableColumns, nil
}

func run(ctx context.Context) error {
	db, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return fmt.Errorf("connect error: %w", err)
	}
	defer db.Close(ctx)

	tableNames, err := getTableNames(ctx, db)
	if err != nil {
		return fmt.Errorf("could not get table names: %w", err)
	}
	for _, tableName := range tableNames {
		if err := processTable(ctx, db, tableName); err != nil {
			return fmt.Errorf("could not mask the table: %w", err)
		}
	}
	return nil
}

func main() {
	log.SetFlags(log.Lmicroseconds | log.Lshortfile)
	if err := run(context.Background()); err != nil {
		log.Fatalf("error: %s", err)
	}
}
