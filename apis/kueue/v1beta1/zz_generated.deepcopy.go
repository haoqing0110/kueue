//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 The Kubernetes Authors.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Admission) DeepCopyInto(out *Admission) {
	*out = *in
	if in.PodSetAssignments != nil {
		in, out := &in.PodSetAssignments, &out.PodSetAssignments
		*out = make([]PodSetAssignment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Admission.
func (in *Admission) DeepCopy() *Admission {
	if in == nil {
		return nil
	}
	out := new(Admission)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterQueue) DeepCopyInto(out *ClusterQueue) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterQueue.
func (in *ClusterQueue) DeepCopy() *ClusterQueue {
	if in == nil {
		return nil
	}
	out := new(ClusterQueue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterQueue) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterQueueList) DeepCopyInto(out *ClusterQueueList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterQueue, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterQueueList.
func (in *ClusterQueueList) DeepCopy() *ClusterQueueList {
	if in == nil {
		return nil
	}
	out := new(ClusterQueueList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterQueueList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterQueuePreemption) DeepCopyInto(out *ClusterQueuePreemption) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterQueuePreemption.
func (in *ClusterQueuePreemption) DeepCopy() *ClusterQueuePreemption {
	if in == nil {
		return nil
	}
	out := new(ClusterQueuePreemption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterQueueSpec) DeepCopyInto(out *ClusterQueueSpec) {
	*out = *in
	if in.ResourceGroups != nil {
		in, out := &in.ResourceGroups, &out.ResourceGroups
		*out = make([]ResourceGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NamespaceSelector != nil {
		in, out := &in.NamespaceSelector, &out.NamespaceSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Preemption != nil {
		in, out := &in.Preemption, &out.Preemption
		*out = new(ClusterQueuePreemption)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterQueueSpec.
func (in *ClusterQueueSpec) DeepCopy() *ClusterQueueSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterQueueSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterQueueStatus) DeepCopyInto(out *ClusterQueueStatus) {
	*out = *in
	if in.FlavorsUsage != nil {
		in, out := &in.FlavorsUsage, &out.FlavorsUsage
		*out = make([]FlavorUsage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterQueueStatus.
func (in *ClusterQueueStatus) DeepCopy() *ClusterQueueStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterQueueStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlavorQuotas) DeepCopyInto(out *FlavorQuotas) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]ResourceQuota, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlavorQuotas.
func (in *FlavorQuotas) DeepCopy() *FlavorQuotas {
	if in == nil {
		return nil
	}
	out := new(FlavorQuotas)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlavorUsage) DeepCopyInto(out *FlavorUsage) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]ResourceUsage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlavorUsage.
func (in *FlavorUsage) DeepCopy() *FlavorUsage {
	if in == nil {
		return nil
	}
	out := new(FlavorUsage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalQueue) DeepCopyInto(out *LocalQueue) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalQueue.
func (in *LocalQueue) DeepCopy() *LocalQueue {
	if in == nil {
		return nil
	}
	out := new(LocalQueue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalQueue) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalQueueFlavorUsage) DeepCopyInto(out *LocalQueueFlavorUsage) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]LocalQueueResourceUsage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalQueueFlavorUsage.
func (in *LocalQueueFlavorUsage) DeepCopy() *LocalQueueFlavorUsage {
	if in == nil {
		return nil
	}
	out := new(LocalQueueFlavorUsage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalQueueList) DeepCopyInto(out *LocalQueueList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LocalQueue, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalQueueList.
func (in *LocalQueueList) DeepCopy() *LocalQueueList {
	if in == nil {
		return nil
	}
	out := new(LocalQueueList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalQueueList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalQueueResourceUsage) DeepCopyInto(out *LocalQueueResourceUsage) {
	*out = *in
	out.Total = in.Total.DeepCopy()
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalQueueResourceUsage.
func (in *LocalQueueResourceUsage) DeepCopy() *LocalQueueResourceUsage {
	if in == nil {
		return nil
	}
	out := new(LocalQueueResourceUsage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalQueueSpec) DeepCopyInto(out *LocalQueueSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalQueueSpec.
func (in *LocalQueueSpec) DeepCopy() *LocalQueueSpec {
	if in == nil {
		return nil
	}
	out := new(LocalQueueSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalQueueStatus) DeepCopyInto(out *LocalQueueStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FlavorUsage != nil {
		in, out := &in.FlavorUsage, &out.FlavorUsage
		*out = make([]LocalQueueFlavorUsage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalQueueStatus.
func (in *LocalQueueStatus) DeepCopy() *LocalQueueStatus {
	if in == nil {
		return nil
	}
	out := new(LocalQueueStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodSet) DeepCopyInto(out *PodSet) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	if in.MinCount != nil {
		in, out := &in.MinCount, &out.MinCount
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodSet.
func (in *PodSet) DeepCopy() *PodSet {
	if in == nil {
		return nil
	}
	out := new(PodSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodSetAssignment) DeepCopyInto(out *PodSetAssignment) {
	*out = *in
	if in.Flavors != nil {
		in, out := &in.Flavors, &out.Flavors
		*out = make(map[corev1.ResourceName]ResourceFlavorReference, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ResourceUsage != nil {
		in, out := &in.ResourceUsage, &out.ResourceUsage
		*out = make(corev1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.Count != nil {
		in, out := &in.Count, &out.Count
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodSetAssignment.
func (in *PodSetAssignment) DeepCopy() *PodSetAssignment {
	if in == nil {
		return nil
	}
	out := new(PodSetAssignment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReclaimablePod) DeepCopyInto(out *ReclaimablePod) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReclaimablePod.
func (in *ReclaimablePod) DeepCopy() *ReclaimablePod {
	if in == nil {
		return nil
	}
	out := new(ReclaimablePod)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceFlavor) DeepCopyInto(out *ResourceFlavor) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceFlavor.
func (in *ResourceFlavor) DeepCopy() *ResourceFlavor {
	if in == nil {
		return nil
	}
	out := new(ResourceFlavor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceFlavor) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceFlavorList) DeepCopyInto(out *ResourceFlavorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ResourceFlavor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceFlavorList.
func (in *ResourceFlavorList) DeepCopy() *ResourceFlavorList {
	if in == nil {
		return nil
	}
	out := new(ResourceFlavorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceFlavorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceFlavorSpec) DeepCopyInto(out *ResourceFlavorSpec) {
	*out = *in
	if in.NodeLabels != nil {
		in, out := &in.NodeLabels, &out.NodeLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NodeTaints != nil {
		in, out := &in.NodeTaints, &out.NodeTaints
		*out = make([]corev1.Taint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceFlavorSpec.
func (in *ResourceFlavorSpec) DeepCopy() *ResourceFlavorSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceFlavorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceGroup) DeepCopyInto(out *ResourceGroup) {
	*out = *in
	if in.CoveredResources != nil {
		in, out := &in.CoveredResources, &out.CoveredResources
		*out = make([]corev1.ResourceName, len(*in))
		copy(*out, *in)
	}
	if in.Flavors != nil {
		in, out := &in.Flavors, &out.Flavors
		*out = make([]FlavorQuotas, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceGroup.
func (in *ResourceGroup) DeepCopy() *ResourceGroup {
	if in == nil {
		return nil
	}
	out := new(ResourceGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceQuota) DeepCopyInto(out *ResourceQuota) {
	*out = *in
	out.NominalQuota = in.NominalQuota.DeepCopy()
	if in.BorrowingLimit != nil {
		in, out := &in.BorrowingLimit, &out.BorrowingLimit
		x := (*in).DeepCopy()
		*out = &x
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceQuota.
func (in *ResourceQuota) DeepCopy() *ResourceQuota {
	if in == nil {
		return nil
	}
	out := new(ResourceQuota)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceUsage) DeepCopyInto(out *ResourceUsage) {
	*out = *in
	out.Total = in.Total.DeepCopy()
	out.Borrowed = in.Borrowed.DeepCopy()
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceUsage.
func (in *ResourceUsage) DeepCopy() *ResourceUsage {
	if in == nil {
		return nil
	}
	out := new(ResourceUsage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Workload) DeepCopyInto(out *Workload) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Workload.
func (in *Workload) DeepCopy() *Workload {
	if in == nil {
		return nil
	}
	out := new(Workload)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Workload) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadList) DeepCopyInto(out *WorkloadList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Workload, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadList.
func (in *WorkloadList) DeepCopy() *WorkloadList {
	if in == nil {
		return nil
	}
	out := new(WorkloadList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkloadList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadSpec) DeepCopyInto(out *WorkloadSpec) {
	*out = *in
	if in.PodSets != nil {
		in, out := &in.PodSets, &out.PodSets
		*out = make([]PodSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadSpec.
func (in *WorkloadSpec) DeepCopy() *WorkloadSpec {
	if in == nil {
		return nil
	}
	out := new(WorkloadSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadStatus) DeepCopyInto(out *WorkloadStatus) {
	*out = *in
	if in.Admission != nil {
		in, out := &in.Admission, &out.Admission
		*out = new(Admission)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ReclaimablePods != nil {
		in, out := &in.ReclaimablePods, &out.ReclaimablePods
		*out = make([]ReclaimablePod, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadStatus.
func (in *WorkloadStatus) DeepCopy() *WorkloadStatus {
	if in == nil {
		return nil
	}
	out := new(WorkloadStatus)
	in.DeepCopyInto(out)
	return out
}
