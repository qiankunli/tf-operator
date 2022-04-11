package pytorch

import (
	"fmt"
	"strconv"
	"strings"

	pytorchv1 "github.com/kubeflow/training-operator/pkg/apis/pytorch/v1"
	corev1 "k8s.io/api/core/v1"
	volcanov1beta1 "volcano.sh/apis/pkg/apis/scheduling/v1beta1"
)

func setPodLabel(obj interface{}, podTemplateSpec *corev1.PodTemplateSpec, rtype, index string) error {
	pytorchjob, ok := obj.(*pytorchv1.PyTorchJob)
	if !ok {
		return fmt.Errorf("%+v is not a type of PyTorchJob", obj)
	}
	if len(podTemplateSpec.Labels) == 0 {
		podTemplateSpec.Labels = make(map[string]string)
	}
	if pytorchjob.Spec.PyTorchReplicaSpecs[pytorchv1.PyTorchReplicaTypeMaster] != nil {
		if rtype == strings.ToLower(string(pytorchv1.PyTorchReplicaTypeMaster)) {
			podTemplateSpec.Labels[volcanov1beta1.PodPreemptable] = "false"
		} else {
			podTemplateSpec.Labels[volcanov1beta1.PodPreemptable] = "true"
		}
	} else {
		// If the master is null, then we need to set the volcano.sh/preemptable = false to make sure that work0 can not be preempted
		rank, err := strconv.Atoi(index)
		if err != nil {
			return err
		}
		if rank == 0 {
			podTemplateSpec.Labels[volcanov1beta1.PodPreemptable] = "false"
		} else {
			podTemplateSpec.Labels[volcanov1beta1.PodPreemptable] = "true"
		}
	}
	return nil
}
