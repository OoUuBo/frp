package basic

import (
	"fmt"

	"github.com/fatedier/frp/test/e2e/framework"
	"github.com/fatedier/frp/test/e2e/framework/consts"
	"github.com/fatedier/frp/test/e2e/pkg/port"

	. "github.com/onsi/ginkgo"
)

// TODO
// * includes

var _ = Describe("[Feature: Config]", func() {
	f := framework.NewDefaultFramework()

	Describe("Template", func() {
		It("render by env", func() {
			serverConf := consts.DefaultServerConfig
			clientConf := consts.DefaultClientConfig

			portName := port.GenName("TCP")
			serverConf += fmt.Sprintf(`
			token = {{ %s{{ .Envs.FRP_TOKEN }}%s }}
			`, "`", "`")

			clientConf += fmt.Sprintf(`
			token = {{ %s{{ .Envs.FRP_TOKEN }}%s }}

			[tcp]
			type = tcp
			local_port = {{ .%s }}
			remote_port = {{ .%s }}
			`, "`", "`", framework.TCPEchoServerPort, portName)

			f.SetEnvs([]string{"FRP_TOKEN=123"})
			f.RunProcesses([]string{serverConf}, []string{clientConf})

			framework.NewRequestExpect(f).PortName(portName).Ensure()
		})
	})

	Describe("Includes", func() {
		It("split tcp proxies into different files", func() {
			// TODO
		})
	})
})
