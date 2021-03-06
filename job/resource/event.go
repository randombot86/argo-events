/*
Copyright 2018 BlackRock, Inc.

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

package resource

import (
	"encoding/json"
	"time"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	"github.com/argoproj/argo-events/job"
)

type event struct {
	job.AbstractEvent
	resource  *resource
	obj       *unstructured.Unstructured
	timestamp time.Time
}

func (e *event) GetID() string {
	return e.resource.AbstractSignal.GetID()
}

func (e *event) GetSource() string {
	return e.resource.Resource.Namespace + ":" + e.resource.Resource.Group + "/" + e.resource.Resource.Version + "/" + e.resource.Resource.Kind
}

func (e *event) GetBody() []byte {
	data, _ := json.Marshal(e.obj)
	return data
}

func (e *event) GetTimestamp() time.Time {
	return e.timestamp
}

func (e *event) GetSignal() job.Signal {
	return e.resource
}
