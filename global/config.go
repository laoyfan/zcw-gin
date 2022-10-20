package global

// 配置文件结构体

type Config struct {
	App   App     `yaml:"app"`
	Zap   Zap     `yaml:"zap"`
	Mysql []Mysql `yaml:"mysql"`
	Redis []Redis `yaml:"rides"`
}

// 系统配置

type App struct {
	Env      string
	Mode     string
	Debug    bool
	Port     string
	Limit    float64
	Language string
}

// 日志配置

type Zap struct {
	Director      string
	Level         string
	MaxAge        int
	Format        string
	StackTraceKey string
	EncodeLevel   string
	Prefix        string
	LogInConsole  bool
	ShowLine      bool
}

// mysql 配置

type Mysql struct {
	Name         string
	Disable      bool
	Type         string
	Node         []Node
	MaxIdleConns int
	MaxOpenConns int
	LogLevel     string
	Log          bool
}

// mysql 节点配置

type Node struct {
	Path     string
	Port     string
	Database string
	Username string
	Password string
	Config   string
	Role     bool
}

// redis 配置

type Redis struct {
	Name     string
	Disable  bool
	Addr     string
	Port     string
	Password string
	Db       int
}
