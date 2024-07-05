package err

import "fmt"

const (
	BindingFailed = "bind 실패 : "
	ServerErr     = "server 에러 : "
	NoDocument    = "데이터 없음 : "
	NOSQLResult   = "sql: no rows in result set"
	ExistDocument = "데이터 존재"
)

func ErrorMsg(status string, err error) string {
	return fmt.Sprintf(status+"%s", err.Error())
}
