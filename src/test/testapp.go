// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package main

import (
	"time"

	"github.com/yyyar/gobetween/config"
	"github.com/yyyar/gobetween/core"
	"github.com/yyyar/gobetween/launch"
	"github.com/yyyar/gobetween/manager"
)

func main() {
	// 	launch.Launch(`
	// [servers]
	//   [servers.sample]
	//   protocol = "tcp"
	//   bind = "127.0.0.1:200"

	//     [servers.sample.discovery]
	//     kind = "static"
	// 	static_list = ["127.0.0.1:8990"]
	// `)
	launch.Launch(config.Config{Api: config.ApiConfig{Enabled: true, Bind: "127.0.0.1:8889"}})
	manager.Create("foo", config.Server{
		Protocol: "tcp",
		Bind:     "127.0.0.1:199",
		Discovery: &config.DiscoveryConfig{
			Kind:                  "static",
			StaticDiscoveryConfig: &config.StaticDiscoveryConfig{StaticList: []string{}},
			// StaticDiscoveryConfig: &config.StaticDiscoveryConfig{StaticList: []string{"127.0.0.1:8990"}},
		},
		Healthcheck: &config.HealthcheckConfig{Kind: "ping",
			Interval: "2s",
			Timeout:  "500ms",
		},
	})
	time.Sleep(time.Second * 5)
	manager.Modify("foo", &[]core.Backend{core.Backend{Target: core.Target{Host: "127.0.0.1", Port: "8990"}, Stats: core.BackendStats{Live: true} }})
	time.Sleep(9000 * time.Second)

}
