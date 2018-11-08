// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodSpec) DeepCopyInto(out *MongodSpec) {
	*out = *in
	if in.ResourcesSpec != nil {
		in, out := &in.ResourcesSpec, &out.ResourcesSpec
		*out = new(ResourcesSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.InMemory != nil {
		in, out := &in.InMemory, &out.InMemory
		*out = new(MongodSpecInMemory)
		**out = **in
	}
	if in.MMAPv1 != nil {
		in, out := &in.MMAPv1, &out.MMAPv1
		*out = new(MongodSpecMMAPv1)
		**out = **in
	}
	if in.WiredTiger != nil {
		in, out := &in.WiredTiger, &out.WiredTiger
		*out = new(MongodSpecWiredTiger)
		**out = **in
	}
	if in.OperationProfiling != nil {
		in, out := &in.OperationProfiling, &out.OperationProfiling
		*out = new(MongodSpecOperationProfiling)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodSpec.
func (in *MongodSpec) DeepCopy() *MongodSpec {
	if in == nil {
		return nil
	}
	out := new(MongodSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodSpecInMemory) DeepCopyInto(out *MongodSpecInMemory) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodSpecInMemory.
func (in *MongodSpecInMemory) DeepCopy() *MongodSpecInMemory {
	if in == nil {
		return nil
	}
	out := new(MongodSpecInMemory)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodSpecMMAPv1) DeepCopyInto(out *MongodSpecMMAPv1) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodSpecMMAPv1.
func (in *MongodSpecMMAPv1) DeepCopy() *MongodSpecMMAPv1 {
	if in == nil {
		return nil
	}
	out := new(MongodSpecMMAPv1)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodSpecOperationProfiling) DeepCopyInto(out *MongodSpecOperationProfiling) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodSpecOperationProfiling.
func (in *MongodSpecOperationProfiling) DeepCopy() *MongodSpecOperationProfiling {
	if in == nil {
		return nil
	}
	out := new(MongodSpecOperationProfiling)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodSpecWiredTiger) DeepCopyInto(out *MongodSpecWiredTiger) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodSpecWiredTiger.
func (in *MongodSpecWiredTiger) DeepCopy() *MongodSpecWiredTiger {
	if in == nil {
		return nil
	}
	out := new(MongodSpecWiredTiger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongosSpec) DeepCopyInto(out *MongosSpec) {
	*out = *in
	if in.ResourcesSpec != nil {
		in, out := &in.ResourcesSpec, &out.ResourcesSpec
		*out = new(ResourcesSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongosSpec.
func (in *MongosSpec) DeepCopy() *MongosSpec {
	if in == nil {
		return nil
	}
	out := new(MongosSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaServerMongoDB) DeepCopyInto(out *PerconaServerMongoDB) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaServerMongoDB.
func (in *PerconaServerMongoDB) DeepCopy() *PerconaServerMongoDB {
	if in == nil {
		return nil
	}
	out := new(PerconaServerMongoDB)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PerconaServerMongoDB) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaServerMongoDBList) DeepCopyInto(out *PerconaServerMongoDBList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PerconaServerMongoDB, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaServerMongoDBList.
func (in *PerconaServerMongoDBList) DeepCopy() *PerconaServerMongoDBList {
	if in == nil {
		return nil
	}
	out := new(PerconaServerMongoDBList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PerconaServerMongoDBList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaServerMongoDBSpec) DeepCopyInto(out *PerconaServerMongoDBSpec) {
	*out = *in
	if in.Mongod != nil {
		in, out := &in.Mongod, &out.Mongod
		*out = new(MongodSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Replsets != nil {
		in, out := &in.Replsets, &out.Replsets
		*out = make([]*ReplsetSpec, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ReplsetSpec)
				**out = **in
			}
		}
	}
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = new(SecretsSpec)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaServerMongoDBSpec.
func (in *PerconaServerMongoDBSpec) DeepCopy() *PerconaServerMongoDBSpec {
	if in == nil {
		return nil
	}
	out := new(PerconaServerMongoDBSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaServerMongoDBStatus) DeepCopyInto(out *PerconaServerMongoDBStatus) {
	*out = *in
	if in.Replsets != nil {
		in, out := &in.Replsets, &out.Replsets
		*out = make([]*ReplsetStatus, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ReplsetStatus)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaServerMongoDBStatus.
func (in *PerconaServerMongoDBStatus) DeepCopy() *PerconaServerMongoDBStatus {
	if in == nil {
		return nil
	}
	out := new(PerconaServerMongoDBStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplsetSpec) DeepCopyInto(out *ReplsetSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplsetSpec.
func (in *ReplsetSpec) DeepCopy() *ReplsetSpec {
	if in == nil {
		return nil
	}
	out := new(ReplsetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplsetStatus) DeepCopyInto(out *ReplsetStatus) {
	*out = *in
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplsetStatus.
func (in *ReplsetStatus) DeepCopy() *ReplsetStatus {
	if in == nil {
		return nil
	}
	out := new(ReplsetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSpecRequirements) DeepCopyInto(out *ResourceSpecRequirements) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSpecRequirements.
func (in *ResourceSpecRequirements) DeepCopy() *ResourceSpecRequirements {
	if in == nil {
		return nil
	}
	out := new(ResourceSpecRequirements)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcesSpec) DeepCopyInto(out *ResourcesSpec) {
	*out = *in
	if in.Limits != nil {
		in, out := &in.Limits, &out.Limits
		*out = new(ResourceSpecRequirements)
		**out = **in
	}
	if in.Requests != nil {
		in, out := &in.Requests, &out.Requests
		*out = new(ResourceSpecRequirements)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcesSpec.
func (in *ResourcesSpec) DeepCopy() *ResourcesSpec {
	if in == nil {
		return nil
	}
	out := new(ResourcesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretsSpec) DeepCopyInto(out *SecretsSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretsSpec.
func (in *SecretsSpec) DeepCopy() *SecretsSpec {
	if in == nil {
		return nil
	}
	out := new(SecretsSpec)
	in.DeepCopyInto(out)
	return out
}
