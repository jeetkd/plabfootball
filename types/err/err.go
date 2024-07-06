package err

import "fmt"

const (
	BindingFailed = "bind 실패 : "
	ServerErr     = "server 에러 : "
	NoDocument    = "데이터 없음 : "
	NOSQLResult   = "sql: no rows in result set"
	ExistDocument = "데이터 존재"
)

// ErrorMsg 는 커스텀 에러 메시지를 반환합니다.
func ErrorMsg(status string, err error) string {
	return fmt.Sprintf(status+"%s", err.Error())
}
