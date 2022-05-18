package main

import (
	"context"
	"fmt"
	"os"

	pgx "github.com/jackc/pgx/v4"
)

func connectDb() (*pgx.Conn, error) {
	connstr := os.Getenv("DATATESTE_URL")
	if len(connstr) == 0 {
		err := fmt.Errorf("URL indisponivel ou inexistente")
		return nil, err
	}
	conn, err := pgx.Connect(context.Background(), connstr)
	if err != nil {
		err = fmt.Errorf("sem acessso ao banco de dados [%v]: %v", connstr, err)
		return nil, err

	}
	return conn, nil
}

var conn *pgx.Conn

func init() {
	var err error
	conn, err = connectDb()
	if err != nil {
		fmt.Fprintf(os.Stderr, "impossivel estabelecer conexão %v\n", err)
		os.Exit(1)
	}
}

func main() {
	defer conn.Close(context.Background())

	err := CriaTabela()
	if err != nil {
		fmt.Fprintf(os.Stderr, "impossivel criar a tabela", err)
		os.Exit(2)
	}
}
