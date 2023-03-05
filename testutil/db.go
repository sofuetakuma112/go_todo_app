package testutil

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// テスト用のMySQLへのコネクションを確立して返す
func OpenDBForTest(t *testing.T) *sqlx.DB {
	t.Helper()

	port := 33306
	if _, defined := os.LookupEnv("CI"); defined {
		// CI環境
		port = 3306
	}

	db, err := sql.Open("mysql", fmt.Sprintf("todo:todo@tcp(127.0.0.1:%d)/todo?parseTime=true", port))
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() { _ = db.Close() })

	// NewDbは、既存の*sql.DBに対する新しいsqlx DBラッパーを返します。
	// 名前付きクエリをサポートするために、
	// 元のデータベースのdriverNameが必要です。
	return sqlx.NewDb(db, "mysql")
}
