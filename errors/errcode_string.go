// Code generated by "stringer -type errCode -linecomment"; DO NOT EDIT.

package errors

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ServerError-10101]
	_ = x[TooManyRequests-10102]
	_ = x[ParamBindError-10103]
	_ = x[MySQLNoQueryError-20505]
	_ = x[MySQLNoFieldError-20506]
}

const (
	_errCode_name_0 = "服务器内部错误请求太多参数信息有误"
	_errCode_name_1 = "查询数据库无结果查询数据库字段不存在"
)

var (
	_errCode_index_0 = [...]uint8{0, 21, 33, 51}
	_errCode_index_1 = [...]uint8{0, 24, 54}
)

func (i errCode) String() string {
	switch {
	case 10101 <= i && i <= 10103:
		i -= 10101
		return _errCode_name_0[_errCode_index_0[i]:_errCode_index_0[i+1]]
	case 20505 <= i && i <= 20506:
		i -= 20505
		return _errCode_name_1[_errCode_index_1[i]:_errCode_index_1[i+1]]
	default:
		return "errCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
