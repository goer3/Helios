package common

// 配置解析入口
type Configuration struct {
	System SystemConfiguration `mapstructure:"system" yaml:"system" json:"system"`
	Log    LogConfiguration    `mapstructure:"log" yaml:"log" json:"log"`
	MySQL  MySQLConfiguration  `mapstructure:"mysql" yaml:"mysql" json:"mysql"`
	Redis  RedisConfiguration  `mapstructure:"redis" yaml:"redis" json:"redis"`
	Auth   AuthConfiguration   `mapstructure:"auth" yaml:"auth" json:"auth"`
}

// 系统配置
type SystemConfiguration struct {
	Mode            string              `mapstructure:"mode" yaml:"mode" json:"mode"`
	Listen          ListenConfiguration `mapstructure:"listen" yaml:"listen" json:"listen"`
	Role            RoleConfiguration   `mapstructure:"role" yaml:"role" json:"role"`
	ShutdownTimeout int                 `mapstructure:"shutdown_timeout" yaml:"shutdown_timeout" json:"shutdown_timeout"`
}

// 监听配置
type ListenConfiguration struct {
	Host string `mapstructure:"host" yaml:"host" json:"host"`
	Port int    `mapstructure:"port" yaml:"port" json:"port"`
}

// 角色配置
type RoleConfiguration struct {
	Leader bool `mapstructure:"leader" yaml:"leader" json:"leader"`
	Worker bool `mapstructure:"worker" yaml:"worker" json:"worker"`
	Web    bool `mapstructure:"web" yaml:"web" json:"web"`
}

// 日志配置
type LogConfiguration struct {
	System BaseLogConfiguration `mapstructure:"system" yaml:"system" json:"system"`
	Access BaseLogConfiguration `mapstructure:"access" yaml:"access" json:"access"`
}

// 基础日志配置
type BaseLogConfiguration struct {
	Enabled   bool                    `mapstructure:"enabled" yaml:"enabled" json:"enabled"`
	Level     string                  `mapstructure:"level" yaml:"level" json:"level"`
	Formatter string                  `mapstructure:"formatter" yaml:"formatter" json:"formatter"`
	Path      string                  `mapstructure:"path" yaml:"path" json:"path"`
	Prefix    string                  `mapstructure:"prefix" yaml:"prefix" json:"prefix"`
	Rolling   LogRollingConfiguration `mapstructure:"rolling" yaml:"rolling" json:"rolling"`
}

// 日志切割配置
type LogRollingConfiguration struct {
	Enabled    bool `mapstructure:"enabled" yaml:"enabled" json:"enabled"`
	MaxSize    int  `mapstructure:"max_size" yaml:"max_size" json:"max_size"`
	MaxBackups int  `mapstructure:"max_backups" yaml:"max_backups" json:"max_backups"`
	MaxAge     int  `mapstructure:"max_age" yaml:"max_age" json:"max_age"`
	Compress   bool `mapstructure:"compress" yaml:"compress" json:"compress"`
}

// MySQL 配置
type MySQLConfiguration struct {
	Host         string `mapstructure:"host" yaml:"host" json:"host"`
	Port         int    `mapstructure:"port" yaml:"port" json:"port"`
	Database     string `mapstructure:"database" yaml:"database" json:"database"`
	Username     string `mapstructure:"username" yaml:"username" json:"username"`
	Password     string `mapstructure:"password" yaml:"password" json:"password"`
	Timeout      int    `mapstructure:"timeout" yaml:"timeout" json:"timeout"`
	Params       string `mapstructure:"params" yaml:"params" json:"params"`
	MaxOpenConns int    `mapstructure:"max_open_conns" yaml:"max_open_conns" json:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns" yaml:"max_idle_conns" json:"max_idle_conns"`
	MaxIdleTime  int    `mapstructure:"max_idle_time" yaml:"max_idle_time" json:"max_idle_time"`
}

// Redis 配置
type RedisConfiguration struct {
	Host         string `mapstructure:"host" yaml:"host" json:"host"`
	Port         int    `mapstructure:"port" yaml:"port" json:"port"`
	Database     int    `mapstructure:"database" yaml:"database" json:"database"`
	Password     string `mapstructure:"password" yaml:"password" json:"password"`
	Timeout      int    `mapstructure:"timeout" yaml:"timeout" json:"timeout"`
	MaxOpenConns int    `mapstructure:"max_open_conns" yaml:"max_open_conns" json:"max_open_conns"`
	MinIdleConns int    `mapstructure:"min_idle_conns" yaml:"min_idle_conns" json:"min_idle_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns" yaml:"max_idle_conns" json:"max_idle_conns"`
	MaxIdleTime  int    `mapstructure:"max_idle_time" yaml:"max_idle_time" json:"max_idle_time"`
}

// 认证配置
type AuthConfiguration struct {
	MultiDevice bool             `mapstructure:"multi_device" yaml:"multi_device" json:"multi_device"`
	JWT         JWTConfiguration `mapstructure:"jwt" yaml:"jwt" json:"jwt"`
}

// JWT 配置
type JWTConfiguration struct {
	Realm   string `mapstructure:"realm" yaml:"realm" json:"realm"`
	Key     string `mapstructure:"key" yaml:"key" json:"key"`
	Timeout int    `mapstructure:"timeout" yaml:"timeout" json:"timeout"`
}
