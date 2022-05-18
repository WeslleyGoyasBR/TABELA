package main

import (
	"context"
	"fmt"

	_ "github.com/jackc/pgx/v4"
)

var err error

func CriaTabela() error {
	sql := "CREATE TABLE clientes(cod uuid, nome varchar(25), telefone varchar)"
	_, err = conn.Exec(context.Background(), sql)
	if err != nil {
		fmt.Errorf("falha ao criar tabela %v", err)
		return err

	}
	return nil

}
