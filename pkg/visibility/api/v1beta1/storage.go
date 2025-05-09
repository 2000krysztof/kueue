/*
Copyright The Kubernetes Authors.

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

package v1beta1

import (
	"k8s.io/apiserver/pkg/registry/rest"

	"sigs.k8s.io/kueue/pkg/queue"
)

func NewStorage(mgr *queue.Manager) map[string]rest.Storage {
	return map[string]rest.Storage{
		"clusterqueues":                  NewCqREST(),
		"clusterqueues/pendingworkloads": NewPendingWorkloadsInCqREST(mgr),
		"localqueues":                    NewLqREST(),
		"localqueues/pendingworkloads":   NewPendingWorkloadsInLqREST(mgr),
	}
}
