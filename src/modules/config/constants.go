package config

// 系统环境变量
const (
	SysMasterHost  = "MASTER_HOST"  // 管理端主机地址与端口
	SysLocalHost   = "LOCAL_HOST"   // 本机暴露给主机或其它服务的主机地址与端口
	SysServiceInfo = "SERVICE_INFO" // 应用信息；将原封不动的发给管理端
)

// 应用配置文件
const (
	AppLogFile   = "logger.filename" // 日志文件(包括路径)
	AppLogFormat = "logger.format"   // 日志格式
	AppLogLevel  = "logger.level"    // 日志可记录的最大级别
	AppCatchDir  = "catch.dir"       // 缓存目录
)
