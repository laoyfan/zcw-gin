package global

// Config 配置文件结构体
type Config struct {
	App   App     `yaml:"app"`   // 系统
	Zap   Zap     `yaml:"zap"`   // 日志
	Jwt   Jwt     `yaml:"jwt"`   // Jwt
	Cors  Cors    `yaml:"cors"`  // 跨域
	Mysql []Mysql `yaml:"mysql"` // 数据库
	Redis []Redis `yaml:"rides"` // Redis
}

// App 系统配置
type App struct {
	Mode     string  `yaml:"mode"`     // 环境
	Port     string  `yaml:"port"`     // 端口
	Limit    float64 `yaml:"limit"`    // 限流
	Language string  `yaml:"language"` // 语言
}

// Jwt token
type Jwt struct {
	SigningKey  string `yaml:"signingKey"`  // jwt签名
	ExpiresTime int64  `yaml:"expiresTime"` // 过期时间
	BufferTime  int64  `yaml:"bufferTime"`  // 缓冲时间
	Issuer      string `yaml:"issuer"`      // 签发者
}

// Cors 跨域配置
type Cors struct {
	AllowOrigins     []string `yaml:"allowOrigins"`     // 允许跨域origin
	AllowMethods     string   `yaml:"allowMethods"`     // 方法
	AllowHeaders     string   `yaml:"allowHeaders"`     // 请求头
	ExposeHeaders    string   `yaml:"exposeHeaders"`    //
	AllowCredentials string   `yaml:"allowCredentials"` //
	MaxAge           string   `yaml:"maxAge"`           //
}

// Zap 日志配置
type Zap struct {
	Director      string `yaml:"director"`      // 配置文件目录
	Level         string `yaml:"level"`         // 日志登记
	MaxAge        int    `yaml:"maxAge"`        //
	Format        string `yaml:"format"`        //
	StackTraceKey string `yaml:"stackTraceKey"` //
	EncodeLevel   string `yaml:"encodeLevel"`   //
	Prefix        string `yaml:"prefix"`        //
	LogInConsole  bool   `yaml:"logInConsole"`  //
	ShowLine      bool   `yaml:"showLine"`      //
}

// Mysql 配置
type Mysql struct {
	Name         string `yaml:"name"`         // 名称
	Disable      bool   `yaml:"disable"`      // 启用
	Type         string `yaml:"type"`         // 类型
	Node         []Node `yaml:"node"`         // 节点
	MaxIdleConns int    `yaml:"maxIdleConns"` // 空闲最大连接数
	MaxOpenConns int    `yaml:"maxOpenConns"` // 打开最大连接数据
	LogLevel     string `yaml:"logLevel"`     //
	Log          bool   `yaml:"log"`          //
}

// Node mysql 节点配置
type Node struct {
	Path     string `yaml:"path"`     // 地址
	Port     string `yaml:"port"`     // 端口
	Database string `yaml:"database"` // 库名
	Username string `yaml:"username"` // 用户名
	Password string `yaml:"password"` // 密码
	Config   string `yaml:"config"`   // 高级配置
	Role     bool   `yaml:"role"`     // 主从
}

// Redis 配置
type Redis struct {
	Name     string `yaml:"name"`     // 名称
	Disable  bool   `yaml:"disable"`  // 启用
	Addr     string `yaml:"addr"`     // 地址
	Port     string `yaml:"port"`     // 端口
	Password string `yaml:"password"` // 密码
	Db       int    `yaml:"db"`       // 库
}
