package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"go/codec"
	"github.com/wangbokun/go/log"
	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

func genSQL(cmd string, v ...interface{}) string {
	return fmt.Sprintf(strings.ReplaceAll(cmd, "?", "%v"), v...)
}

// MySQL mysql
type MySQL struct {
	opts Options
	// Log  core.Logger
	db   *sql.DB
}
// New file config
func New(opts ...Option) *MySQL {
	options := NewOptions(opts...)
	return &MySQL{
		opts: options,
		// Log:  log.DefaultStdLog(),
	}
}

// Init init
func (my *MySQL) Init(opts ...Option) {
	for _, o := range opts {
		o(&my.opts)
	}
}

// SetLog set log
// func (my *MySQL) SetLog(log core.Logger) {
// 	my.Log = log
// }

// LoadConfig loadconfig
func (my *MySQL) LoadConfig(v interface{}) error {
	return codec.NewJSONCodec().Format(&my.opts, v)
}

// Connect connect
func (my *MySQL) Connect() error {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8", my.opts.Username, my.opts.Password, my.opts.Hostname, my.opts.Port, my.opts.Database)
	my.db, err = sql.Open("mysql", dsn)
	if err == nil {
		my.db.SetConnMaxLifetime(time.Duration(my.opts.MaxLifetime) * time.Second) // 1h
		// 避免 Invalid Connection， 设置成一样的
		my.db.SetMaxOpenConns(my.opts.MaxConns)
		my.db.SetMaxIdleConns(my.opts.MaxConns)
	}
	return err
}

// Exist is exist
func (my *MySQL) Exist(ctx context.Context, table string, values map[string]interface{}) (bool, error) {
	var exists bool
	var (
		keyList   []string
		valueList []interface{}
	)
	for key, value := range values {
		keyList = append(keyList, fmt.Sprintf("`%s` = ?", key))
		valueList = append(valueList, value)
	}
	cmd := fmt.Sprintf("SELECT EXISTS(SELECT * FROM %s WHERE %s)", table, strings.Join(keyList, " AND "))
	err := my.db.QueryRow(cmd, valueList...).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return false, fmt.Errorf("exists err: %v, cmd=%s", err, fmt.Sprintf(strings.ReplaceAll(cmd, "?", "%s"), valueList...))
	}
	return exists, nil
}

// Put insert
func (my *MySQL) Put(ctx context.Context, table string, v interface{}) (sql.Result, error) {
	var (
		keyList   []string
		valueList []interface{}
		markList  []string
	)

	data := make(map[string]interface{})
	if err := codec.NewJSONCodec().Format(&data, v); err != nil {
		return nil, err
	}

	for key, value := range data {
		keyList = append(keyList, "`"+key+"`")
		valueList = append(valueList, value)
		markList = append(markList, "?")
	}

	cmd := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, strings.Join(keyList, ", "), strings.Join(markList, ", "))
	return my.Exec(ctx, cmd, valueList...)
}

// Update update
func (my *MySQL) Update(ctx context.Context, table string, v interface{}, cond map[string]interface{}) (sql.Result, error) {
	data := make(map[string]interface{})
	if err := codec.NewJSONCodec().Format(&data, v); err != nil {
		return nil, err
	}
	var (
		keyList   []string
		condList  []string
		valueList []interface{}
	)
	for key, value := range data {
		keyList = append(keyList, fmt.Sprintf("`%s` = ?", key))
		valueList = append(valueList, value)
	}
	for key, value := range cond {
		condList = append(condList, fmt.Sprintf("`%s` = ?", key))
		valueList = append(valueList, value)

	}
	cmd := fmt.Sprintf("UPDATE %s SET %s WHERE %s", table, strings.Join(keyList, ","), strings.Join(condList, " AND "))
	return my.Exec(ctx, cmd, valueList...)
}

// Upset update if exist or insert, true is insert, false is update
func (my *MySQL) Upset(ctx context.Context, table string, v interface{}, cond map[string]interface{}) (bool, error) {
	result, err := my.Update(ctx, table, v, cond)
	if err != nil {
		return false, err
	}
	if n, err := result.RowsAffected(); err != nil || n != 0 {
		return false, err
	}

	_, err = my.Put(ctx, table, v)
	return true, err
}

// Get get data
func (my *MySQL) Get(ctx context.Context, cmd string, handle func(row map[string]interface{}) error) error {
	rows, err := my.db.QueryContext(ctx, cmd)
	if err != nil {
		return err
	}
	defer rows.Close()

	cols, err := rows.ColumnTypes()
	if err != nil {
		return err
	}

	columnLen := len(cols)
	values := make([]interface{}, columnLen)
	scanValues := make([]interface{}, columnLen)
	for i := range scanValues {
		scanValues[i] = &values[i]
	}
	for rows.Next() {
		if err = rows.Scan(scanValues...); err != nil {
			return err
		}
		result := make(map[string]interface{}, columnLen)
		for i := 0; i < columnLen; i++ {
			switch cols[i].ScanType().Name() { //字段类型
			case "RawBytes", "NullTime":
				if values[i] != nil {
					result[cols[i].Name()] = string(values[i].([]byte))
				}
			case "NullInt64", "int", "int8", "int32", "int64","uint32":
				if values[i] != nil {
					result[cols[i].Name()], err = strconv.Atoi(string(values[i].([]byte)))
					if err != nil {
						return err
					}
				}
			default:
				if values[i] != nil {
					result[cols[i].Name()] = values[i]
				}
				log.Warn("type unkonw, name=%v, kind=%v, value=%v", cols[i].Name(), cols[i].ScanType().Kind(), values[i])
			}
		}
		if err := handle(result); err != nil {
			return err
		}
	}
	return rows.Err()
}

// Ping ping
func (my *MySQL) Ping() error {
	var err error
	for i := 0; i < my.opts.PingAttempts; i++ {
		if err = my.db.Ping(); err == nil {
			return nil
		}
		log.Warn("Ping fail #%d: %v", i+1, err)
	}
	return err
}

// Exec exec
func (my *MySQL) Exec(ctx context.Context, cmd string, args ...interface{}) (sql.Result, error) {
	stmt, err := my.db.PrepareContext(ctx, cmd)
	if err != nil {
		return nil, fmt.Errorf("exec err=%v, sql=%v", err, genSQL(cmd, args...))
	}
	defer stmt.Close()
	return stmt.ExecContext(ctx, args...)
}

// Stat stats
func (my *MySQL) Stat() string {
	return fmt.Sprintf("%#v", my.db.Stats())
}

// Close close connect
func (my *MySQL) Close() error {
	return my.db.Close()
}
