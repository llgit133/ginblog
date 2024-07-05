package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	ERROR_USER_USED      = 1001
	ERROR_PASSWORD_WRONG = 1002
	ERROR_USER_NOT_EXIST = 1003

	ERROR_TOKEN_EXIST      = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007

	ERROR_CATEGORY_NOT_EXIST = 2001

	ERROR_ARTICLE_NOT_EXIST = 3001
	ERROR_ARTICLE_EXIST     = 3002
	ERROR_ARTICLE_DELETE    = 3003
)

// 1000开头为用户模块的错误
// 2000开头为分类模块的错误
// 3000开头为文章模块的错误

var codeMsg = map[int]string{
	SUCCESS:                  "OK",
	ERROR:                    "FAIL",
	ERROR_USER_NOT_EXIST:     "用户不存在",
	ERROR_PASSWORD_WRONG:     "密码错误",
	ERROR_USER_USED:          "用户已存在",
	ERROR_TOKEN_EXIST:        "TOKEN已存在",
	ERROR_TOKEN_RUNTIME:      "TOKEN已过期",
	ERROR_TOKEN_WRONG:        "TOKEN不正确",
	ERROR_TOKEN_TYPE_WRONG:   "TOKEN格式错误",
	ERROR_CATEGORY_NOT_EXIST: "分类不存在",
	ERROR_ARTICLE_NOT_EXIST:  "文章不存在",
	ERROR_ARTICLE_EXIST:      "文章已存在",
	ERROR_ARTICLE_DELETE:     "文章删除失败",
}

func GetErrMsg(code int) string {
	return codeMsg[code]

}

//https://www.bilibili.com/video/BV1oA411e7kM/?spm_id_from=333.788&vd_source=b88b8b3483f8d9fb78e502aa8b359b50
