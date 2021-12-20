package setting

import (
	"github.com/go-ini/ini"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	HTTPAddr     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	RdsDsn  string
	RdbOpts *redis.Options

	SentryDsn string
	LoggerCfg *LogConf
)

type LogConf struct {
	Filename   string // 日志文件路径
	Level      int8   // 日志级别
	MaxSize    int    // 每个日志文件保存的最大尺寸 单位：M
	MaxBackups int    // 日志文件最多保存多少个备份
	MaxAge     int    // 文件最多保存多少天
	Compress   bool   // 是否压缩
}

func Open(filename string) error {
	cfg, err := ini.Load(filename)
	if err != nil {
		return err
	}
	if err := LoadApp(cfg); err != nil {
		return err
	}
	if err := LoadServer(cfg); err != nil {
		return err
	}
	if err := LoadRdb(cfg); err != nil {
		return err
	}
	if err := LoadRds(cfg); err != nil {
		return err
	}
	if err := LoadSentry(cfg); err != nil {
		return err
	}
	if err := LoadLogger(cfg); err != nil {
		return err
	}

	return nil
}

func LoadApp(cfg *ini.File) error {
	return nil
}

func LoadServer(cfg *ini.File) error {
	sec, err := cfg.GetSection("server")
	if err != nil {
		return err
	}
	HTTPAddr = sec.Key("addr").MustString("127.0.0.1:8088")
	ReadTimeout = time.Duration(sec.Key("read_timeout").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("write_timeout").MustInt(59)) * time.Second
	return nil
}

func LoadRds(cfg *ini.File) error {
	sec, err := cfg.GetSection("rds")
	if err != nil {
		return err
	}
	RdsDsn = sec.Key("dsn").MustString("")
	return nil
}

func LoadRdb(cfg *ini.File) error {
	sec, err := cfg.GetSection("rdb")
	if err != nil {
		return err
	}
	RdbOpts = &redis.Options{
		Addr:     sec.Key("address").MustString("127.0.0.1:6379"),
		Password: sec.Key("password").MustString(""),
		DB:       sec.Key("database").MustInt(0),
	}
	return nil
}

func LoadSentry(cfg *ini.File) error {
	SentryDsn = cfg.Section("sentry").Key("dsn").MustString("")
	return nil
}

func LoadLogger(cfg *ini.File) error {
	sec, err := cfg.GetSection("log")
	if err != nil {
		return err
	}
	LoggerCfg = &LogConf{
		Filename:   sec.Key("filename").MustString("./zap.log"),
		Level:      int8(sec.Key("level").MustInt(-1)),
		MaxSize:    sec.Key("max_size").MustInt(200),
		MaxBackups: sec.Key("max_backups").MustInt(100),
		MaxAge:     sec.Key("max_age").MustInt(1),
		Compress:   sec.Key("compress").MustBool(false),
	}
	return nil
}
