package config_client

import (
	"github.com/golang/mock/gomock"
	"github.com/peggypig/nacos-go/clients/nacos_client"
	"github.com/peggypig/nacos-go/common/constant"
	"github.com/peggypig/nacos-go/common/http_agent"
	"github.com/peggypig/nacos-go/vo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

/**
*
* @description :
*
* @author : codezhang
*
* @create : 2019-01-10 15:55
**/

func TestMockIConfigClient_GetConfig(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := http_agent.NewMockIHttpAgent(ctrl)
	mockINacosClient := nacos_client.NewMockINacosClient(ctrl)

	mockINacosClient.EXPECT().SetHttpAgent(gomock.Eq(mockIHttpAgent)).Times(1).Return(nil)

	mockINacosClient.EXPECT().SetClientConfig(gomock.Eq(constant.ClientConfig{
		TimeoutMs:      10 * 1000,
		ListenInterval: 10 * 1000,
	}))

	mockINacosClient.EXPECT().SetServerConfig(gomock.Eq([]constant.ServerConfig{{
		IpAddr:      "console.nacos.io",
		ContextPath: "/nacos",
		Port:        80,
	}}))

	mockINacosClient.EXPECT().GetClientConfig().Times(1).Return(constant.ClientConfig{
		TimeoutMs:      10 * 1000,
		ListenInterval: 10 * 1000,
	}, nil)

	mockINacosClient.EXPECT().GetServerConfig().Times(1).Return([]constant.ServerConfig{{
		IpAddr:      "console.nacos.io",
		ContextPath: "/nacos",
		Port:        80,
	}}, nil)

	mockINacosClient.EXPECT().GetHttpAgent().Times(1).Return(mockIHttpAgent, nil)

	mockIHttpAgent.EXPECT().Get(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/cs/configs"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.AssignableToTypeOf(uint64(10*1000)),
		gomock.Eq(map[string]string{
			"dataId": "TEST",
			"group":  "TEST",
		})).
		Times(1).
		Return(http_agent.FakeHttpResponse(200, `MOCK RESULT`), nil)

	client := ConfigClient{}
	client.INacosClient = mockINacosClient
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(constant.ClientConfig{
		TimeoutMs:      10 * 1000,
		ListenInterval: 10 * 1000,
	})
	_ = client.SetServerConfig([]constant.ServerConfig{{
		IpAddr:      "console.nacos.io",
		Port:        80,
		ContextPath: "/nacos",
	}})
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: "TEST",
		Group:  "TEST",
	})
	assert.Equal(t, nil, err)
	assert.Equal(t, "MOCK RESULT", content)
}

func TestMockIConfigClient_PublishConfig(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := http_agent.NewMockIHttpAgent(ctrl)
	mockINacosClient := nacos_client.NewMockINacosClient(ctrl)

	mockINacosClient.EXPECT().SetHttpAgent(gomock.Eq(mockIHttpAgent)).Times(1).Return(nil)

	mockINacosClient.EXPECT().SetClientConfig(gomock.Eq(constant.ClientConfig{
		TimeoutMs:      10 * 1000,
		ListenInterval: 10 * 1000,
	}))

	mockINacosClient.EXPECT().SetServerConfig(gomock.Eq([]constant.ServerConfig{{
		IpAddr:      "console.nacos.io",
		Port:        80,
		ContextPath: "/nacos",
	}}))

	mockINacosClient.EXPECT().GetClientConfig().Times(1).Return(constant.ClientConfig{
		TimeoutMs:      10 * 1000,
		ListenInterval: 10 * 1000,
	}, nil)

	mockINacosClient.EXPECT().GetServerConfig().Times(1).Return([]constant.ServerConfig{{
		IpAddr:      "console.nacos.io",
		Port:        80,
		ContextPath: "/nacos",
	}}, nil)

	mockINacosClient.EXPECT().GetHttpAgent().Times(1).Return(mockIHttpAgent, nil)

	mockIHttpAgent.EXPECT().Post(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/cs/configs"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.AssignableToTypeOf(uint64(10*1000)),
		gomock.Eq(map[string]string{
			"dataId":  "TEST",
			"group":   "TEST",
			"content": "test",
		})).
		Times(1).
		Return(http_agent.FakeHttpResponse(200, `true`), nil)

	client := ConfigClient{}
	client.INacosClient = mockINacosClient
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(constant.ClientConfig{
		TimeoutMs:      10 * 1000,
		ListenInterval: 10 * 1000,
	})
	_ = client.SetServerConfig([]constant.ServerConfig{{
		IpAddr:      "console.nacos.io",
		Port:        80,
		ContextPath: "/nacos",
	}})
	content, err := client.PublishConfig(vo.ConfigParam{
		DataId:  "TEST",
		Group:   "TEST",
		Content: "test",
	})
	assert.Equal(t, nil, err)
	assert.Equal(t, true, content)
}

func TestMockIConfigClientMockRecorder_DeleteConfig(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := http_agent.NewMockIHttpAgent(ctrl)
	mockINacosClient := nacos_client.NewMockINacosClient(ctrl)

	mockINacosClient.EXPECT().SetHttpAgent(gomock.Eq(mockIHttpAgent)).Times(1).Return(nil)

	mockINacosClient.EXPECT().SetClientConfig(gomock.Eq(constant.ClientConfig{
		TimeoutMs:      10 * 1000,
		ListenInterval: 10 * 1000,
	}))

	mockINacosClient.EXPECT().SetServerConfig(gomock.Eq([]constant.ServerConfig{{
		IpAddr:      "console.nacos.io",
		Port:        80,
		ContextPath: "/nacos",
	}}))

	mockINacosClient.EXPECT().GetClientConfig().Times(1).Return(constant.ClientConfig{
		TimeoutMs:      10 * 1000,
		ListenInterval: 10 * 1000,
	}, nil)

	mockINacosClient.EXPECT().GetServerConfig().Times(1).Return([]constant.ServerConfig{{
		IpAddr:      "console.nacos.io",
		Port:        80,
		ContextPath: "/nacos",
	}}, nil)

	mockINacosClient.EXPECT().GetHttpAgent().Times(1).Return(mockIHttpAgent, nil)

	mockIHttpAgent.EXPECT().Delete(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/cs/configs"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.AssignableToTypeOf(uint64(10*1000)),
		gomock.Eq(map[string]string{
			"dataId": "TEST",
			"group":  "TEST",
		})).
		Times(1).
		Return(http_agent.FakeHttpResponse(200, `true`), nil)

	client := ConfigClient{}
	client.INacosClient = mockINacosClient
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(constant.ClientConfig{
		TimeoutMs:      10 * 1000,
		ListenInterval: 10 * 1000,
	})
	_ = client.SetServerConfig([]constant.ServerConfig{{
		IpAddr:      "console.nacos.io",
		Port:        80,
		ContextPath: "/nacos",
	}})
	content, err := client.DeleteConfig(vo.ConfigParam{
		DataId: "TEST",
		Group:  "TEST",
	})
	assert.Equal(t, nil, err)
	assert.Equal(t, true, content)
}

func TestMockIConfigClientMockRecorder_GetConfigContent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := http_agent.NewMockIHttpAgent(ctrl)
	mockINacosClient := nacos_client.NewMockINacosClient(ctrl)

	mockINacosClient.EXPECT().SetHttpAgent(gomock.Eq(mockIHttpAgent)).Times(1).Return(nil)

	mockINacosClient.EXPECT().SetClientConfig(gomock.Eq(constant.ClientConfig{
		TimeoutMs:      10 * 1000,
		ListenInterval: 10 * 1000,
	}))

	mockINacosClient.EXPECT().SetServerConfig(gomock.Eq([]constant.ServerConfig{{
		IpAddr:      "console.nacos.io",
		Port:        80,
		ContextPath: "/nacos",
	}}))

	mockINacosClient.EXPECT().GetClientConfig().Times(1).Return(constant.ClientConfig{
		TimeoutMs:      10 * 1000,
		ListenInterval: 10 * 1000,
	}, nil)

	mockINacosClient.EXPECT().GetServerConfig().Times(1).Return([]constant.ServerConfig{{
		IpAddr:      "console.nacos.io",
		Port:        80,
		ContextPath: "/nacos",
	}}, nil)

	mockINacosClient.EXPECT().GetHttpAgent().Times(1).Return(mockIHttpAgent, nil)

	mockIHttpAgent.EXPECT().Get(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/cs/configs"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.AssignableToTypeOf(uint64(10*1000)),
		gomock.Eq(map[string]string{
			"dataId": "TEST",
			"group":  "TEST",
		})).
		Times(1).
		Return(http_agent.FakeHttpResponse(200, `MOCK RESULT`), nil)

	client := ConfigClient{}
	client.INacosClient = mockINacosClient
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(constant.ClientConfig{
		TimeoutMs:      10 * 1000,
		ListenInterval: 10 * 1000,
	})
	_ = client.SetServerConfig([]constant.ServerConfig{{
		IpAddr:      "console.nacos.io",
		Port:        80,
		ContextPath: "/nacos",
	}})
	content, err := client.GetConfigContent("TEST", "TEST")
	assert.Equal(t, nil, err)
	assert.Equal(t, "MOCK RESULT", content)
}

func TestMockIConfigClient_ListenConfig(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := http_agent.NewMockIHttpAgent(ctrl)
	mockINacosClient := nacos_client.NewMockINacosClient(ctrl)

	mockINacosClient.EXPECT().SetHttpAgent(gomock.Eq(mockIHttpAgent)).Times(1).Return(nil)

	mockINacosClient.EXPECT().SetClientConfig(gomock.Eq(constant.ClientConfig{
		TimeoutMs:      10 * 1000,
		ListenInterval: 10 * 1000,
		BeatInterval:   1 * 1000,
	})).Times(1).Return(nil)

	mockINacosClient.EXPECT().SetServerConfig(gomock.Eq([]constant.ServerConfig{{
		IpAddr:      "console.nacos.io",
		Port:        80,
		ContextPath: "/nacos",
	}})).Times(1).Return(nil)

	mockINacosClient.EXPECT().GetClientConfig().AnyTimes().Return(constant.ClientConfig{
		TimeoutMs:      10 * 1000,
		ListenInterval: 10 * 1000,
	}, nil)

	mockINacosClient.EXPECT().GetServerConfig().AnyTimes().Return([]constant.ServerConfig{{
		IpAddr:      "console.nacos.io",
		Port:        80,
		ContextPath: "/nacos",
	}}, nil)

	mockINacosClient.EXPECT().GetHttpAgent().AnyTimes().Return(mockIHttpAgent, nil)

	mockIHttpAgent.EXPECT().Post(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/cs/configs/listener"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.AssignableToTypeOf(uint64(10*1000)),
		gomock.Eq(map[string]string{
			"Listening-Configs": "TEST" + constant.SPLIT_CONFIG_INNER + "TEST" + constant.SPLIT_CONFIG_INNER +
				constant.SPLIT_CONFIG_INNER + constant.SPLIT_CONFIG,
		})).AnyTimes().
		Return(http_agent.FakeHttpResponse(200, ``), nil)

	client := ConfigClient{}
	client.INacosClient = mockINacosClient
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(constant.ClientConfig{
		TimeoutMs:      10 * 1000,
		ListenInterval: 10 * 1000,
		BeatInterval:   1 * 1000,
	})
	_ = client.SetServerConfig([]constant.ServerConfig{{
		IpAddr:      "console.nacos.io",
		Port:        80,
		ContextPath: "/nacos",
	}})
	err := client.ListenConfig([]vo.ConfigParam{{
		DataId: "TEST",
		Group:  "TEST",
	}})
	time.Sleep(21 * time.Second)
	assert.Equal(t, nil, err)
}

func TestMockIConfigClient_StopListenConfig(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()
	mockIHttpAgent := http_agent.NewMockIHttpAgent(ctrl)
	mockINacosClient := nacos_client.NewMockINacosClient(ctrl)

	mockINacosClient.EXPECT().SetHttpAgent(gomock.Eq(mockIHttpAgent)).Times(1).Return(nil)

	mockINacosClient.EXPECT().SetClientConfig(gomock.Eq(constant.ClientConfig{
		TimeoutMs:      10 * 1000,
		ListenInterval: 10 * 1000,
	})).Times(1).Return(nil)

	mockINacosClient.EXPECT().SetServerConfig(gomock.Eq([]constant.ServerConfig{{
		IpAddr:      "console.nacos.io",
		Port:        80,
		ContextPath: "/nacos",
	}})).Times(1).Return(nil)

	mockINacosClient.EXPECT().GetClientConfig().AnyTimes().Return(constant.ClientConfig{
		TimeoutMs:      10 * 1000,
		ListenInterval: 10 * 1000,
	}, nil)

	mockINacosClient.EXPECT().GetServerConfig().AnyTimes().Return([]constant.ServerConfig{{
		IpAddr:      "console.nacos.io",
		Port:        80,
		ContextPath: "/nacos",
	}}, nil)

	mockINacosClient.EXPECT().GetHttpAgent().AnyTimes().Return(mockIHttpAgent, nil)

	mockIHttpAgent.EXPECT().Post(
		gomock.Eq("http://console.nacos.io:80/nacos/v1/cs/configs/listener"),
		gomock.AssignableToTypeOf(http.Header{}),
		gomock.AssignableToTypeOf(uint64(10*1000)),
		gomock.Eq(map[string]string{
			"Listening-Configs": "TEST" + constant.SPLIT_CONFIG_INNER + "TEST" + constant.SPLIT_CONFIG_INNER +
				constant.SPLIT_CONFIG_INNER + constant.SPLIT_CONFIG,
		})).AnyTimes().
		Return(http_agent.FakeHttpResponse(200, ``), nil)

	client := ConfigClient{}
	client.INacosClient = mockINacosClient
	_ = client.SetHttpAgent(mockIHttpAgent)
	_ = client.SetClientConfig(constant.ClientConfig{
		TimeoutMs:      10 * 1000,
		ListenInterval: 10 * 1000,
	})
	_ = client.SetServerConfig([]constant.ServerConfig{{
		IpAddr:      "console.nacos.io",
		Port:        80,
		ContextPath: "/nacos",
	}})
	err := client.ListenConfig([]vo.ConfigParam{{
		DataId: "TEST",
		Group:  "TEST",
	}})
	go func() {
		time.Sleep(11 * time.Second)
		client.StopListenConfig()
	}()
	time.Sleep(21 * time.Second)
	assert.Equal(t, nil, err)
}
