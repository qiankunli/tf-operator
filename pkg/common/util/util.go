// Copyright 2021 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License

package util

import (
	commonv1 "github.com/kubeflow/common/pkg/apis/common/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ObjectFilterFunction func(obj metav1.Object) bool

// ConvertServiceList convert service list to service point list
func ConvertServiceList(list []corev1.Service) []*corev1.Service {
	if list == nil {
		return nil
	}
	ret := make([]*corev1.Service, 0, len(list))
	for i := range list {
		ret = append(ret, &list[i])
	}
	return ret
}

// ConvertPodList convert pod list to pod pointer list
func ConvertPodList(list []corev1.Pod) []*corev1.Pod {
	if list == nil {
		return nil
	}
	ret := make([]*corev1.Pod, 0, len(list))
	for i := range list {
		ret = append(ret, &list[i])
	}
	return ret
}

// ConvertPodListWithFilter converts pod list to pod pointer list with ObjectFilterFunction
func ConvertPodListWithFilter(list []corev1.Pod, pass ObjectFilterFunction) []*corev1.Pod {
	if list == nil {
		return nil
	}
	ret := make([]*corev1.Pod, 0, len(list))
	for i := range list {
		obj := &list[i]
		if pass != nil {
			if pass(obj) {
				ret = append(ret, obj)
			}
		} else {
			ret = append(ret, obj)
		}
	}
	return ret
}

func GetReplicaTypes(specs map[commonv1.ReplicaType]*commonv1.ReplicaSpec) []commonv1.ReplicaType {
	keys := make([]commonv1.ReplicaType, 0, len(specs))
	for k := range specs {
		keys = append(keys, k)
	}
	return keys
}

// GetContainerExitCode gets the container exit code from the given pod.
func GetContainerExitCode(pod *corev1.Pod, name string) int32 {
	var exitCode int32 = 0xbeef // magic number
	for _, status := range pod.Status.ContainerStatuses {
		state := status.State
		if status.Name == name && state.Terminated != nil {
			exitCode = state.Terminated.ExitCode
		}
	}
	return exitCode
}
