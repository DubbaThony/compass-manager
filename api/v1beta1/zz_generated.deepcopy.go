//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompassManagerMapping) DeepCopyInto(out *CompassManagerMapping) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompassManagerMapping.
func (in *CompassManagerMapping) DeepCopy() *CompassManagerMapping {
	if in == nil {
		return nil
	}
	out := new(CompassManagerMapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CompassManagerMapping) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompassManagerMappingList) DeepCopyInto(out *CompassManagerMappingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CompassManagerMapping, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompassManagerMappingList.
func (in *CompassManagerMappingList) DeepCopy() *CompassManagerMappingList {
	if in == nil {
		return nil
	}
	out := new(CompassManagerMappingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CompassManagerMappingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompassManagerMappingSpec) DeepCopyInto(out *CompassManagerMappingSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompassManagerMappingSpec.
func (in *CompassManagerMappingSpec) DeepCopy() *CompassManagerMappingSpec {
	if in == nil {
		return nil
	}
	out := new(CompassManagerMappingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompassManagerMappingStatus) DeepCopyInto(out *CompassManagerMappingStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompassManagerMappingStatus.
func (in *CompassManagerMappingStatus) DeepCopy() *CompassManagerMappingStatus {
	if in == nil {
		return nil
	}
	out := new(CompassManagerMappingStatus)
	in.DeepCopyInto(out)
	return out
}
