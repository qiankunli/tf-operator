//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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
// limitations under the License.

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	commonv1 "github.com/kubeflow/common/pkg/apis/common/v1"
	"k8s.io/api/autoscaling/v2beta2"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticPolicy) DeepCopyInto(out *ElasticPolicy) {
	*out = *in
	if in.MinReplicas != nil {
		in, out := &in.MinReplicas, &out.MinReplicas
		*out = new(int32)
		**out = **in
	}
	if in.MaxReplicas != nil {
		in, out := &in.MaxReplicas, &out.MaxReplicas
		*out = new(int32)
		**out = **in
	}
	if in.RDZVBackend != nil {
		in, out := &in.RDZVBackend, &out.RDZVBackend
		*out = new(RDZVBackend)
		**out = **in
	}
	if in.RDZVPort != nil {
		in, out := &in.RDZVPort, &out.RDZVPort
		*out = new(int32)
		**out = **in
	}
	if in.RDZVHost != nil {
		in, out := &in.RDZVHost, &out.RDZVHost
		*out = new(string)
		**out = **in
	}
	if in.RDZVID != nil {
		in, out := &in.RDZVID, &out.RDZVID
		*out = new(string)
		**out = **in
	}
	if in.RDZVConf != nil {
		in, out := &in.RDZVConf, &out.RDZVConf
		*out = make([]RDZVConf, len(*in))
		copy(*out, *in)
	}
	if in.Standalone != nil {
		in, out := &in.Standalone, &out.Standalone
		*out = new(bool)
		**out = **in
	}
	if in.NProcPerNode != nil {
		in, out := &in.NProcPerNode, &out.NProcPerNode
		*out = new(int32)
		**out = **in
	}
	if in.MaxRestarts != nil {
		in, out := &in.MaxRestarts, &out.MaxRestarts
		*out = new(int32)
		**out = **in
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = make([]v2beta2.MetricSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SuccessPolicy != nil {
		in, out := &in.SuccessPolicy, &out.SuccessPolicy
		*out = new(SuccessPolicy)
		**out = **in
	}
	if in.FailurePolicy != nil {
		in, out := &in.FailurePolicy, &out.FailurePolicy
		*out = new(FailurePolicy)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticPolicy.
func (in *ElasticPolicy) DeepCopy() *ElasticPolicy {
	if in == nil {
		return nil
	}
	out := new(ElasticPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PyTorchJob) DeepCopyInto(out *PyTorchJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PyTorchJob.
func (in *PyTorchJob) DeepCopy() *PyTorchJob {
	if in == nil {
		return nil
	}
	out := new(PyTorchJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PyTorchJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PyTorchJobList) DeepCopyInto(out *PyTorchJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PyTorchJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PyTorchJobList.
func (in *PyTorchJobList) DeepCopy() *PyTorchJobList {
	if in == nil {
		return nil
	}
	out := new(PyTorchJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PyTorchJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PyTorchJobSpec) DeepCopyInto(out *PyTorchJobSpec) {
	*out = *in
	in.RunPolicy.DeepCopyInto(&out.RunPolicy)
	if in.ElasticPolicy != nil {
		in, out := &in.ElasticPolicy, &out.ElasticPolicy
		*out = new(ElasticPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.PyTorchReplicaSpecs != nil {
		in, out := &in.PyTorchReplicaSpecs, &out.PyTorchReplicaSpecs
		*out = make(map[commonv1.ReplicaType]*commonv1.ReplicaSpec, len(*in))
		for key, val := range *in {
			var outVal *commonv1.ReplicaSpec
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(commonv1.ReplicaSpec)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PyTorchJobSpec.
func (in *PyTorchJobSpec) DeepCopy() *PyTorchJobSpec {
	if in == nil {
		return nil
	}
	out := new(PyTorchJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDZVConf) DeepCopyInto(out *RDZVConf) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDZVConf.
func (in *RDZVConf) DeepCopy() *RDZVConf {
	if in == nil {
		return nil
	}
	out := new(RDZVConf)
	in.DeepCopyInto(out)
	return out
}
