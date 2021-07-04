package tests

import (
	"database/sql"

	"bitbucket.org/terirem/simpleserver/pkg/storage"
)

func MakeTestStorage() error {
	st, err := storage.NewSQLiteStorage("testdb.sqlite")
	if err != nil {
		return err
	}

	conn, err := st.GetConnect()
	if err != nil {
		return err
	}

	defer conn.Close()

	migrate(conn)
	fillTestData(conn)

	return nil
}

func migrate(conn *sql.Conn) error {

	return nil
}

func fillTestData(db *sql.Conn) error {

	return nil
}
