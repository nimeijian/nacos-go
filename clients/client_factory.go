package clients

import (
	"errors"
	"nacos-go/clients/config_client"
	"nacos-go/clients/service_client"
	"nacos-go/common/constant"
)

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-01-07 15:13
**/

// 创建配置相关的客户端
func CreateConfigClient(properties map[string]interface{}) (iClient config_client.IConfigClient, err error) {
	client := config_client.ConfigClient{}
	// 设置 timeoutMs
	if err == nil {
		if timeoutMsTmp, exist := properties[constant.KEY_TIMEOUT_MS]; exist {
			if timeoutMs, ok := timeoutMsTmp.(uint64); ok {
				client.ClientConfig = constant.ClientConfig{
					TimeoutMs: timeoutMs,
				}
			} else {
				err = errors.New("invalid timeoutMs value,should be uint64")
			}
		} else {
			client.ClientConfig = constant.ClientConfig{
				TimeoutMs: 10 * 1000,
			}
		}
	}
	// 设置 ListenInterval
	if err == nil {
		if listenIntervalTmp, exist := properties[constant.KEY_LISTEN_INTERVAL]; exist {
			if listenInterval, ok := listenIntervalTmp.(uint64); ok {
				client.ClientConfig.ListenInterval = listenInterval
			} else {
				err = errors.New("invalid listenInterval value,should be uint64")
			}
		} else {
			client.ClientConfig.ListenInterval = 30 * 1000
		}
	}
	// 设置 serverConfig
	if err == nil {
		if serverConfigTmp, exist := properties[constant.KEY_SERVER_CONFIG]; exist {
			if serverConfig, ok := serverConfigTmp.(constant.ServerConfig); ok {
				if len(serverConfig.IpAddr) <= 0 || serverConfig.Port <= 0 || serverConfig.Port > 65535 {
					err = errors.New("invalid server config")
				} else {
					client.ServerConfigs = append(client.ServerConfigs, serverConfig)
				}
			}
		} else {
			err = errors.New("server config not found in properties")
		}
	}
	if err == nil {
		iClient = &client
	}
	return
}

// 创建服务发现相关的客户端
func CreateServiceClient() (client service_client.ServiceClient) {
	client = service_client.ServiceClient{}
	return
}
