package controller

import (
	"hidevops.io/hiboot/pkg/app/web"
	"hidevops.io/hiboot/pkg/log"
	"hidevops.io/hiboot/pkg/utils/io"
	"hidevops.io/mio/console/pkg/aggregate/mocks"
	"hidevops.io/mio/console/pkg/command"
	"net/http"
	"testing"
)

func TestGatewayConfig(t *testing.T) {
	gateway := new(mocks.GatewayConfigAggregate)
	appInfo := newGatewayConfigController(gateway)
	gateway.On("Create", "", "", "", "", "", "").Return(nil, nil)

	app := web.NewTestApp(t, appInfo).
		SetProperty("kube.serviceHost", "test").
		Run(t)
	log.SetLevel(log.DebugLevel)
	log.Println(io.GetWorkDir())
	t.Run("should pass with jwt token", func(t *testing.T) {
		app.Post("/gatewayConfig/create").WithJSON(&command.DeployConfigType{}).Expect().Status(http.StatusOK)
	})
}
