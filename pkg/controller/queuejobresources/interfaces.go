/*
Copyright 2019, 2021 The Multi-Cluster App Dispatcher Authors.

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

package queuejobresources

import (
	qjobv1 "github.com/project-codeflare/multi-cluster-app-dispatcher/pkg/apis/controller/v1beta1"
)

// Interface is an abstract interface for queue job resource management.
type Interface interface {
	UpdateQueueJobStatus(queuejob *qjobv1.AppWrapper) error
	// TODO: Add to calculate more accurate partial deployments while job is being realized
	//	GetAggregatedResourcesByPhase(phase v1.PodPhase, queuejob *qjobv1.AppWrapper) *clusterstateapi.Resource
	Run(stopCh <-chan struct{})
}
