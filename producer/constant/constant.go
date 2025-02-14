package constant

const (
	CON_API_CALL_BACK  = ""
	CON_KAKFA_ADRRESS1 = "localhost:9092"
	CON_KAKFA_TOPICS   = ""
	CON_KAKFA_GROUPID  = ""
)

type CONSTANT_CODE int32 // 对前端响应的状态码

const (
	CONSTANT_CODE_PARAMETER_ERROR CONSTANT_CODE = 10000

	CONSTANT_CODE_NAMEPASS_ERROR        = 10001
	CONSTANT_CODE_SERVER_ERROR          = 10002
	CONSTANT_CODE_TOKEN_EXPIRS          = 10003
	CONSTANT_CODE_NEED_LOGIN            = 10004
	CONSTANT_CODE_USER_INVALIDE         = 10005
	CONSTANT_CODE_USER_ALREADY_EXIST    = 10006
	CONSTANT_CODE_ROLE_LIMIT            = 10007
	CONSTANT_CODE_ACCOUNT_ALREADY_EXIST = 10008

	CONSTANT_CODE_REQUEST_SUCCESS = 20000
	CONSTANT_CODE_LOGIN_SUCCESS   = 20001

	CONSTANT_CODE_OPERATE_DELETE_FAILED   = 30000
	CONSTANT_CODE_OPERATE_UPDATE_FAILED   = 30001
	CONSTANT_CODE_OPERATE_CREATE_FAILED   = 30002
	CONSTANT_CODE_OPERATE_DELETE_SUCCESS  = 30003
	CONSTANT_CODE_OPERATE_UPDATE_SUCCESS  = 30004
	CONSTANT_CODE_OPERATE_CREATE_SUCCESS  = 30005
	CONSTANT_CODE_OPERATE_INTEGRAL_FAILED = 30006
)
