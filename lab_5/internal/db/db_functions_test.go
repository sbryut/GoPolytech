package db

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

type TestTableRow struct {
	names       []string
	errExpected error
}

var testTable = []TestTableRow{
	{names: []string{"Sofia", "Nikita", "Marina", "Olga", "Artemiy"}, errExpected: nil},
	{names: nil, errExpected: sql.ErrNoRows},
	{names: nil, errExpected: fmt.Errorf("SQL error")},
}

func mockDbRows(names []string) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"name"})
	for _, name := range names {
		rows.AddRow(name)
	}
	return rows
}

func TestGetNames(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%v' was not expected when opening a mock database connection", err)
	}
	defer mockDB.Close()

	dbService := DBService{DB: mockDB}

	for i, row := range testTable {
		mock.ExpectQuery("SELECT name FROM users").WillReturnRows(mockDbRows(row.names)).WillReturnError(row.errExpected)
		names, err := dbService.GetNames()

		if row.errExpected != nil {
			require.ErrorIs(t, err, row.errExpected, "row: %d, expected error: %v, actual error: %v", i, row.errExpected, err)
			require.Nil(t, names, "row: %d, names must be nil", i)
			continue
		}

		require.NoError(t, err, "row: %d, error must be nil", i)
		require.Equal(t, row.names, names, "row: %d, expected names: %v, actual names: %v", i, row.names, names)
	}
}

type TestSelectUniqueValuesRow struct {
	columnName  string
	tableName   string
	mockRows    []string
	errExpected error
}

var testSelectUniqueValuesTable = []TestSelectUniqueValuesRow{
	{columnName: "department", tableName: "staff", mockRows: []string{"IT", "HR", "Sales"}, errExpected: nil},
	{columnName: "department", tableName: "staff", mockRows: nil, errExpected: sql.ErrNoRows},
	{columnName: "department", tableName: "staff", mockRows: nil, errExpected: fmt.Errorf("SQL error")},
}

func mockSelectUniqueValuesRows(rows []string) *sqlmock.Rows {
	mockRows := sqlmock.NewRows([]string{"value"})
	for _, row := range rows {
		mockRows.AddRow(row)
	}
	return mockRows
}

func TestSelectUniqueValues(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%v' was not expected when opening a mock database connection", err)
	}
	defer mockDB.Close()

	dbService := DBService{DB: mockDB}

	for i, row := range testSelectUniqueValuesTable {
		query := "SELECT DISTINCT " + row.columnName + " FROM " + row.tableName
		mock.ExpectQuery(query).WillReturnRows(mockSelectUniqueValuesRows(row.mockRows)).WillReturnError(row.errExpected)

		values, err := dbService.SelectUniqueValues(row.columnName, row.tableName)

		if row.errExpected != nil {
			require.ErrorIs(t, err, row.errExpected, "row: %d, expected error: %v, actual error: %v", i, row.errExpected, err)
			require.Nil(t, values, "row: %d, values must be nil", i)
			continue
		}

		require.NoError(t, err, "row: %d, error must be nil", i)
		require.Equal(t, row.mockRows, values, "row: %d, expected values: %v, actual values: %v", i, row.mockRows, values)
	}
}
