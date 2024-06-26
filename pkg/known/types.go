/*
SPDX-License-Identifier: Apache-2.0

Copyright Contributors to the Submariner project.

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

package known

import (
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
)

type SyncerConfig struct {
	// LocalRestConfig the REST config used to access the local resources to sync.
	LocalRestConfig *rest.Config

	// LocalClient the client used to access local resources to sync. This is optional and is provided for unit testing
	// in lieu of the LocalRestConfig. If not specified, one is created from the LocalRestConfig.
	LocalClient     dynamic.Interface
	LocalNamespace  string
	LocalClusterID  string
	RemoteNamespace string
}

type Specification struct {
	ClusterID          string
	BootStrapToken     string
	HubSecretNamespace string
	HubSecretName      string
	ShareNamespace     string
	LocalNamespace     string
	HubURL             string
	CIDR               []string
	IsHub              bool
	Endpoint           string
}
