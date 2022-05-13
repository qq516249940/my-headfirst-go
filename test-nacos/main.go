package main

import (
	"fmt"
	"time"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func main() {
	sc := []constant.ServerConfig{{
		IpAddr: "mxxxx.com",
		Port:   8848,
	}}

	cc := constant.ClientConfig{
		NamespaceId:         "a7132ad4-60b7-4acb-ba12-xxxxxxx", // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "log",
		CacheDir:            "cache",
		// RotateTime:          "1h",
		// MaxAge:              3,
		LogLevel: "debug",
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		panic(err)
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "config-debug.yaml",
		Group:  "dev",
	})

	if err != nil {
		panic(err)
	}
	fmt.Println(content) //字符串 - yaml
	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: "config-debug.yaml",
		Group:  "dev",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("配置文件发生了变化...")
			fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
		},
	})

	time.Sleep(300 * time.Second)
}
