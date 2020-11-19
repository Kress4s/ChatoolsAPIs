package constant

const (
	EXPEIRE_ACCOUNT_CODE   = 2000 //用户账户登录过期状态码
	EXPEIRE_WXACCOUNT_CODE = 2001 //微信登录状态过期状态码(暂且不用)
	TOKEN_IS_NIL           = 2002 //token为空
)

//header
const (
	H_AUTHORIZATION = "Authorization" //
	H_WXID          = "wx_id"         //微信ID
	H_UUID          = "UUID"          //变化的
	H_TOKEN_KEY     = "WX_TOKEN"      // 给协议顶层的token
)

// params
const (
	P_UUID  = "uuid"
	P_WXID  = "wx_id"
	P_TOKEN = "token"
)

// session
const (
	S_WX_ID   = "WX_ID"
	S_UUID    = "uuid"
	S_ACCOUNT = "account"
)

const (
	BASE_URL        = "http://49.234.86.244:8080" //微信对接baselink
	GRPC_BASE_URL   = "49.234.86.244:8081"        // 消息监控的地址
	GRPC_RUN_SERVER = "139.159.237.126:9001"
)

// login
const (
	LOGIN_AUTO_URL   = BASE_URL + "/login/auto"
	LOGIN_AWAKE_URL  = BASE_URL + "/login/awake"
	LOGIN_CHECK_URL  = BASE_URL + "/login/check"
	LOGIN_HEART_URL  = BASE_URL + "/login/heartbeat"
	LOGIN_INIT_URL   = BASE_URL + "/login/init"
	LOGIN_LOGOUT_URL = BASE_URL + "/login/logout"
	LOGIN_QRCODE_URL = BASE_URL + "/login/qr_code"
)

// wxuser
const (
	WXUSER_PROFILE_URL = BASE_URL + "/user/profile"
)

//contact
const (
	CONTACT_ACCEPT_URL     = BASE_URL + "/contact/accept"
	CONTACT_ADD_URL        = BASE_URL + "/contact/add"
	CONTACT_BATCH_URL      = BASE_URL + "/contact/batch"
	CONTACT_LIST_URL       = BASE_URL + "/contact/list/all"
	CONTACT_GROUP_LIST_URL = BASE_URL + "/contact/list/group"
	CONTACT_SEARCH_URL     = BASE_URL + "/contact/search"
)

// group
const (
	GROUP_ACCEPT_URL      = BASE_URL + "/group/accept"
	GROUP_CREATE_URL      = BASE_URL + "/group/create"
	GROUP_DETAIL_URL      = BASE_URL + "/group/get/detail"
	GROUP_INFO_URL        = BASE_URL + "/group/get/info"
	GROUP_MEMBERS_URL     = BASE_URL + "/group/get/members"
	GROUP_ADD_MEMBERS_URL = BASE_URL + "/group/members/add"
	GROUP_DEL_MEMBERS_URL = BASE_URL + "/group/members/delete"
	// GROUP_ADD_MEMBERS_URL  = BASE_URL + "/group/members/invite"
	GROUP_QUIT_URL         = BASE_URL + "/group/quit"
	GROUP_SET_ANNOUNCE_URL = BASE_URL + "/group/set/announcement"
)

// label
const (
	LABEL_ADD_URL         = BASE_URL + "/label/add"
	LABEL_DELETE_URL      = BASE_URL + "/label/delete"
	LABEL_LIST_URL        = BASE_URL + "/label/list"
	LABEL_UPDATE_URL      = BASE_URL + "/label/update"
	LABEL_LIST_UPDATE_URL = BASE_URL + "/label/update/list"
)

// Send Message
const (
	SEND_TEXT  = BASE_URL + "/message/send/text"
	SEND_IMAGE = BASE_URL + "/message/send/image"
	SEND_VIDEO = BASE_URL + "/message/send/video"
)

// sns
const (
	SNS_LIST_ME      = BASE_URL + "/sns/list/me"
	SNS_LIST_FRIEND  = BASE_URL + "/sns/list/friend"
	SNS_SEND_TEXT    = BASE_URL + "/sns/send/text"
	SNS_SEND_IMAGE   = BASE_URL + "/sns/send/image"
	SNS_SEND_VIDEO   = BASE_URL + "/sns/send/video"
	SNS_UPLOAD_IMAGE = BASE_URL + "/sns/upload/image"
	SNS_UPLOAD_VIDEO = BASE_URL + "/sns/upload/video"
)

// message info
const (
	TEXT_TYPE_MESSAGE = iota
	IMAGE_TYPE_MESSAGE
	VIDEO_TYPE_MESSAGE
	CARD_TYPE_MESSAGE
	EMOJI_TYPE_MESSAGE
	SMALL_PROGRAM_TYPE_MESSAGE
)

// message source
const (
	CONTACT_MESSAGE = iota
	GROUP_MESSAGE
	PUBLIC_MESSAGE
	SYSTEM_MESSAGE
)

// resource type
const (
	SOURCE_TEXT = iota
	SOURCE_IMAGE
	SOURCE_VOICE
	SOURCE_VIDEO
	SOURCE_FILE
	SOURCE_LINK
	SOURCE_APP
	SOURCE_EMOJI
)

// 认证方式
const (
	AUTH_CODE = "authCode"
	AUTH_PWD  = "authPwd"
)

// 定时任务状态
const (
	UN_SEND    = "unSend"     // 未发送
	SENDED     = "sended"     // 已发送
	FAILEDSEND = "sendFailed" // 发送失败
)

const (
	MESSAGE_TASK      = "messagetask"
	ANNOUNCEMENT_TASK = "announcementask"
)
