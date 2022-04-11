package dao

import (
	"encoding/json"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"gopkg.in/ini.v1"
)

func loadNacos(file *ini.File) {
	NacosAddr = file.Section("nacos").Key("NacosAddr").String()
	NacosPort, _ = file.Section("nacos").Key("NacosPort").Int()
}

//nacos link
func linkNacos() {
	//serverConfig  http://121.4.81.22:8848/nacos
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(
			NacosAddr,
			uint64(NacosPort),
			constant.WithScheme("http"),
			constant.WithContextPath("/nacos")),
	}
	//clientConfig
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId(""),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
	)

	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		panic(err)
	}

	//获取配置
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: "nacos_daily",
		Group:  "daily",
	})
	err = json.Unmarshal([]byte(content), &Res)
	if err != nil {
		panic(err)
	}
}
