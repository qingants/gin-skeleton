// 1. 全局logger
// 2. 支持日志分割
// 3. 接管os.StdOut
// 4. 可读性强的时间信息
// 5. Fatal或者Error打印调用栈
// 6. 显示文件和行号信息
// 7. 开发模式输出到os.Stdout, goland可以超链跳转
// 8. 可以支持单独日志对象

package log

import (
	"github.com/qingants/gin-skeleton/pkg/utils"
	"os"

	"github.com/natefinch/lumberjack"

	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	once       sync.Once
	logger     *zap.Logger
	globalUndo func()
	stdUndo    func()
)

func Open(filename string, level zapcore.Level, maxSize, maxBackups, maxAge int, compress bool) {
	once.Do(func() {
		caller := zap.AddCaller()
		development := zap.Development()
		stacktrace := zap.AddStacktrace(zap.ErrorLevel)

		var core zapcore.Core
		// 本地开发输出到控制台即可
		if utils.IsLocal() {
			core = initConsoleCore(level)
		} else {
			core = initFileCore(filename, level, maxSize, maxBackups, maxAge, compress)
		}
		logger = zap.New(core, caller, development, stacktrace)
		sCore := &SentryCore{
			LevelEnabler: zap.ErrorLevel,
			withFields:   []zapcore.Field{},
		}
		logger = logger.WithOptions(zap.WrapCore(func(core zapcore.Core) zapcore.Core {
			return zapcore.NewTee(core, sCore)
		}))
		// 替换zap的全局logger
		globalUndo = zap.ReplaceGlobals(logger)
		// 重定向log.Stdout
		stdUndo, _ = zap.RedirectStdLogAt(logger, zap.InfoLevel)
	})
}

func Close() {
	stdUndo()
	globalUndo()
	if err := logger.Sync(); err != nil {
		panic(err)
	}
}

func initFileCore(filename string, level zapcore.Level, maxSize, maxBackups, maxAge int, compress bool) zapcore.Core {
	writer := lumberjack.Logger{
		Filename:   filename,   // 日志文件路径
		MaxSize:    maxSize,    // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: maxBackups, // 日志文件最多保存多少个备份
		MaxAge:     maxAge,     // 文件最多保存多少天
		Compress:   compress,   // 是否压缩
	}
	encoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	return zapcore.NewCore(encoder, zapcore.AddSync(&writer), zap.NewAtomicLevelAt(level))
}

func initConsoleCore(level zapcore.Level) zapcore.Core {
	// mutex保护os.Stdout
	syncer := zapcore.Lock(os.Stdout)

	// 默认选择development配置
	conf := zap.NewDevelopmentEncoderConfig()

	// 开发模式开启全路径，方便跳转
	conf.EncodeCaller = zapcore.FullCallerEncoder

	// 彩色终端
	conf.EncodeLevel = zapcore.CapitalColorLevelEncoder

	encoder := zapcore.NewConsoleEncoder(conf)
	return zapcore.NewCore(encoder, syncer, zap.NewAtomicLevelAt(level))
}
