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

func TestServiceConfig(t *testing.T) {
	serviceConfigAggregate := new(mocks.ServiceConfigAggregate)
	service := newServiceConfigController(serviceConfigAggregate)

	app := web.NewTestApp(t, service).
		SetProperty("kube.serviceHost", "test").
		Run(t)
	log.SetLevel(log.DebugLevel)
	log.Println(io.GetWorkDir())
	serviceConfigAggregate.On("Create", "", "", "", "", "", "").Return(nil, nil)
	t.Run("should pass with jwt token", func(t *testing.T) {
		app.Post("/serviceConfig/create").WithJSON(&command.DeployConfigType{}).Expect().Status(http.StatusOK)
	})

}
