/*
Copyright 2019 The Machine Controller Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package vsphere

import (
	"github.com/kubermatic/machine-controller/pkg/cloudprovider/instance"
)

type Server struct {
	name      string
	id        string
	status    instance.Status
	addresses []string
}

func (vsphereServer Server) Name() string {
	return vsphereServer.name
}

func (vsphereServer Server) ID() string {
	return vsphereServer.id
}

func (vsphereServer Server) Addresses() []string {
	return vsphereServer.addresses
}

func (vsphereServer Server) Status() instance.Status {
	return vsphereServer.status
}
