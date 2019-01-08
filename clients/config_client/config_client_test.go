package config_client

import (
	"nacos-go/common/constant"
	"nacos-go/vo"
	"testing"
	"time"
)

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-01-08 12:05
**/

func TestConfigClient_GetConfig(t *testing.T) {
	client := ConfigClient{
		ServerConfigs: []constant.ServerConfig{constant.ServerConfig{
			IpAddr: "10.0.0.8",
			Port:   8848,
		}},
	}
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: "TEST",
		Group:  "DEFAULT_GROUP",
	})
	if err != nil {
		t.Error(err)
	} else {
		t.Log(content)
	}
}

func TestConfigClient_PublishConfig(t *testing.T) {
	client := ConfigClient{
		ServerConfigs: []constant.ServerConfig{constant.ServerConfig{
			IpAddr: "10.0.0.8",
			Port:   8848,
		}},
	}
	content, err := client.PublishConfig(vo.ConfigParam{
		DataId:  "TEST",
		Group:   "DEFAULT_GROUP",
		Content: "aaa",
		Tenant:  "bbb",
	})
	if err != nil {
		t.Error(err)
	} else {
		t.Log(content)
	}
}

func TestConfigClient_DeleteConfig(t *testing.T) {
	client := ConfigClient{
		ServerConfigs: []constant.ServerConfig{constant.ServerConfig{
			IpAddr: "10.0.0.8",
			Port:   8848,
		}},
	}
	content, err := client.DeleteConfig(vo.ConfigParam{
		DataId: "TEST",
		Group:  "DEFAULT_GROUP",
		Tenant: "bbb",
	})
	if err != nil {
		t.Error(err)
	} else {
		t.Log(content)
	}
}

func TestConfigClient_ListenConfig(t *testing.T) {
	client := ConfigClient{
		ServerConfigs: []constant.ServerConfig{constant.ServerConfig{
			IpAddr: "10.0.0.8",
			Port:   8848,
		}},
		ClientConfig: constant.ClientConfig{
			TimeoutMs:      30 * 1000,
			ListenInterval: 10 * 1000,
		},
	}
	err := client.ListenConfig([]vo.ConfigParam{{
		DataId: "TEST",
		Group:  "TEST",
		Content: "2019-01-08 09:57:34",
	},{
		DataId: "TEST",
		Group:  "DEFAULT_GROUP",
	}})
	if err != nil {
		t.Error(err)
	}
	time.Sleep(100 * time.Second)
}

func TestConfigClient_GetConfigContent(t *testing.T) {
	client := ConfigClient{
		ServerConfigs: []constant.ServerConfig{constant.ServerConfig{
			IpAddr: "10.0.0.8",
			Port:   8848,
		}},
		ClientConfig: constant.ClientConfig{
			TimeoutMs:      30 * 1000,
			ListenInterval: 10 * 1000,
		},
	}
	content , err := client.GetConfigContent("TEST","TEST1")
	if err != nil {
		t.Error(err)
	}else {
		t.Log(content)
	}
}