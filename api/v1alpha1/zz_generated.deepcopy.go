// +build !ignore_autogenerated

/*


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

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BOSHConfig) DeepCopyInto(out *BOSHConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Dependencies.DeepCopyInto(&out.Dependencies)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BOSHConfig.
func (in *BOSHConfig) DeepCopy() *BOSHConfig {
	if in == nil {
		return nil
	}
	out := new(BOSHConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BOSHConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BOSHConfigList) DeepCopyInto(out *BOSHConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BOSHConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BOSHConfigList.
func (in *BOSHConfigList) DeepCopy() *BOSHConfigList {
	if in == nil {
		return nil
	}
	out := new(BOSHConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BOSHConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BOSHConfigSpec) DeepCopyInto(out *BOSHConfigSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BOSHConfigSpec.
func (in *BOSHConfigSpec) DeepCopy() *BOSHConfigSpec {
	if in == nil {
		return nil
	}
	out := new(BOSHConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BOSHConfigStatus) DeepCopyInto(out *BOSHConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BOSHConfigStatus.
func (in *BOSHConfigStatus) DeepCopy() *BOSHConfigStatus {
	if in == nil {
		return nil
	}
	out := new(BOSHConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BOSHDeployment) DeepCopyInto(out *BOSHDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Dependencies.DeepCopyInto(&out.Dependencies)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BOSHDeployment.
func (in *BOSHDeployment) DeepCopy() *BOSHDeployment {
	if in == nil {
		return nil
	}
	out := new(BOSHDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BOSHDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BOSHDeploymentList) DeepCopyInto(out *BOSHDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BOSHDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BOSHDeploymentList.
func (in *BOSHDeploymentList) DeepCopy() *BOSHDeploymentList {
	if in == nil {
		return nil
	}
	out := new(BOSHDeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BOSHDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BOSHDeploymentSpec) DeepCopyInto(out *BOSHDeploymentSpec) {
	*out = *in
	if in.Ops != nil {
		in, out := &in.Ops, &out.Ops
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Vars != nil {
		in, out := &in.Vars, &out.Vars
		*out = make([]VariableSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BOSHDeploymentSpec.
func (in *BOSHDeploymentSpec) DeepCopy() *BOSHDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(BOSHDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BOSHDeploymentStatus) DeepCopyInto(out *BOSHDeploymentStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BOSHDeploymentStatus.
func (in *BOSHDeploymentStatus) DeepCopy() *BOSHDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(BOSHDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BOSHStemcell) DeepCopyInto(out *BOSHStemcell) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Dependencies.DeepCopyInto(&out.Dependencies)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BOSHStemcell.
func (in *BOSHStemcell) DeepCopy() *BOSHStemcell {
	if in == nil {
		return nil
	}
	out := new(BOSHStemcell)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BOSHStemcell) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BOSHStemcellList) DeepCopyInto(out *BOSHStemcellList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BOSHStemcell, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BOSHStemcellList.
func (in *BOSHStemcellList) DeepCopy() *BOSHStemcellList {
	if in == nil {
		return nil
	}
	out := new(BOSHStemcellList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BOSHStemcellList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BOSHStemcellSpec) DeepCopyInto(out *BOSHStemcellSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BOSHStemcellSpec.
func (in *BOSHStemcellSpec) DeepCopy() *BOSHStemcellSpec {
	if in == nil {
		return nil
	}
	out := new(BOSHStemcellSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BOSHStemcellStatus) DeepCopyInto(out *BOSHStemcellStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BOSHStemcellStatus.
func (in *BOSHStemcellStatus) DeepCopy() *BOSHStemcellStatus {
	if in == nil {
		return nil
	}
	out := new(BOSHStemcellStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigMapVariableSource) DeepCopyInto(out *ConfigMapVariableSource) {
	*out = *in
	if in.MapKeys != nil {
		in, out := &in.MapKeys, &out.MapKeys
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigMapVariableSource.
func (in *ConfigMapVariableSource) DeepCopy() *ConfigMapVariableSource {
	if in == nil {
		return nil
	}
	out := new(ConfigMapVariableSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DependencySpec) DeepCopyInto(out *DependencySpec) {
	*out = *in
	if in.Stemcell != nil {
		in, out := &in.Stemcell, &out.Stemcell
		*out = new(string)
		**out = **in
	}
	if in.Deployment != nil {
		in, out := &in.Deployment, &out.Deployment
		*out = new(string)
		**out = **in
	}
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DependencySpec.
func (in *DependencySpec) DeepCopy() *DependencySpec {
	if in == nil {
		return nil
	}
	out := new(DependencySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DependencySpecs) DeepCopyInto(out *DependencySpecs) {
	*out = *in
	if in.Dependencies != nil {
		in, out := &in.Dependencies, &out.Dependencies
		*out = make([]DependencySpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DependencySpecs.
func (in *DependencySpecs) DeepCopy() *DependencySpecs {
	if in == nil {
		return nil
	}
	out := new(DependencySpecs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretVariableSource) DeepCopyInto(out *SecretVariableSource) {
	*out = *in
	if in.MapKeys != nil {
		in, out := &in.MapKeys, &out.MapKeys
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretVariableSource.
func (in *SecretVariableSource) DeepCopy() *SecretVariableSource {
	if in == nil {
		return nil
	}
	out := new(SecretVariableSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VariableSource) DeepCopyInto(out *VariableSource) {
	*out = *in
	if in.ConfigMap != nil {
		in, out := &in.ConfigMap, &out.ConfigMap
		*out = new(ConfigMapVariableSource)
		(*in).DeepCopyInto(*out)
	}
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(SecretVariableSource)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VariableSource.
func (in *VariableSource) DeepCopy() *VariableSource {
	if in == nil {
		return nil
	}
	out := new(VariableSource)
	in.DeepCopyInto(out)
	return out
}
