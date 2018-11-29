// Copyright 2018 John Deng (hi.devops.io@gmail.com).
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// if protoc report command not found error, should install proto and protc-gen-go
// go get -u -v github.com/golang/protobuf/{proto,protoc-gen-go}
//go:generate protoc -I pkg/console/protobuf --go_out=plugins=grpc:pkg/console/protobuf pkg/console/protobuf/buildconfig.proto

package main

import (
	"hidevops.io/hiboot/pkg/app/web"
	_ "hidevops.io/hiboot/pkg/starter/actuator"
	_ "hidevops.io/hiboot/pkg/starter/locale"
	_ "hidevops.io/mio/console/pkg/controller"
	_ "hidevops.io/mio/console/pkg/rpc"
	_ "hidevops.io/mio/console/pkg/service"
)

// main
func main() {
	// create new web application and run it
	web.NewApplication().Run()
}
