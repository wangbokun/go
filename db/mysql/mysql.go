package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"time"
)

type MySQLConfig struct {
	Addr     string `yaml:"addr"`
	User     string `yaml:"user"`
	Pass     string `yaml:"pass"`
	Database string `yaml:"database"`
}

var engine *sql.DB

func MySQLInit(addr, user, pass, database string) error {
	var dsn string
	var err error
	if user == "" || pass == "" {
		dsn = fmt.Sprintf("%s/%s?charset=utf8", addr, database)
	} else {
		dsn = fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8", user, pass, addr, database)
	}
	fmt.Println(dsn)
	engine, err = sql.Open("mysql", dsn)
	engine.SetConnMaxLifetime(6 * time.Hour)
	engine.SetMaxOpenConns(200)
	engine.SetMaxIdleConns(100)
	return err
}

func Insert(table string, row map[string]interface{}) (sql.Result, error) {
	if err := engine.Ping(); err != nil {
		return nil, err
	}
	var keyList []string
	var valueList []interface{}
	var markList []string
	for k, v := range row {
		keyList = append(keyList, k)
		valueList = append(valueList, v)
		markList = append(markList, "?")
	}
	sql := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, strings.Join(keyList, ", "), strings.Join(markList, ", "))
	fmt.Println(sql, fmt.Sprintf("%#v", valueList))
	stmt, err := engine.Prepare(sql)
	if err != nil {
		return nil, err
	}
	return stmt.Exec(valueList...)
}

func Update(table string, row map[string]interface{}, selector string, value interface{}) (sql.Result, error) {
	if err := engine.Ping(); err != nil {
		return nil, err
	}
	var updateList []string
	var valueList []interface{}
	for k, v := range row {
		updateList = append(updateList, k+"=?")
		valueList = append(valueList, v)
	}
	sql := fmt.Sprintf("UPDATE %s SET %s WHERE %s=?", table, strings.Join(updateList, ", "), selector)
	fmt.Println(sql, fmt.Sprintf("%#v", valueList))
	valueList = append(valueList, value)
	stmt, err := engine.Prepare(sql)
	if err != nil {
		return nil, err
	}
	return stmt.Exec(valueList...)
}

func MySQLDelete(table string, selector string, value interface{}) (sql.Result, error) {
	err := engine.Ping()
	if err != nil {
		return nil, err
	}
	sql := fmt.Sprintf("DELETE FROM %s WHERE %s=?", table, selector)
	fmt.Println(sql, fmt.Sprintf("%#v", value))
	stmt, err := engine.Prepare(sql)
	if err != nil {
		return nil, err
	}
	return stmt.Exec(value)
}
