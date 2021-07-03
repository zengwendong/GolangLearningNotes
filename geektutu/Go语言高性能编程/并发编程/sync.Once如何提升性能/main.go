/*
sync.Once 的使用场景
sync.Once 是 Go 标准库提供的使函数只执行一次的实现，常应用于单例模式，例如初始化配置、保持数据库连接等。

sync.Once 的原理
type Once struct {
	done uint32 // 标记是否已初始化
	m    Mutex  // 互斥锁
}
首先：保证变量仅被初始化一次，需要有个标志来判断变量是否已初始化过，若没有则需要初始化。
第二：线程安全，支持并发，无疑需要互斥锁来实现
 */


package main

import (
	"log"
	"strconv"
	"sync"
	"time"
)

type Config struct {
	Server string
	Port   int64
}

var (
	once   sync.Once
	config *Config
)

func ReadConfig() *Config {
	once.Do(func() {
		var err error
		//config = &Config{Server: os.Getenv("TT_SERVER_URL")}
		//config.Port, err = strconv.ParseInt(os.Getenv("TT_PORT"), 10, 0)
		// test
		config = &Config{Server: "127.0.0.1"}
		config.Port, err = strconv.ParseInt("8088", 10, 0)
		if err != nil {
			config.Port = 8080 // default port
		}
		log.Println("init config")
	})
	return config
}

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			_ = ReadConfig()
		}()
	}
	log.Printf("%s:%d starting...\n", config.Server, config.Port)
	time.Sleep(time.Second)
}

/*
2021/07/03 20:22:43 init config
2021/07/03 20:22:43 127.0.0.1:8088 starting...
 */