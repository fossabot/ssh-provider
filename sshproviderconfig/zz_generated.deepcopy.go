// +build !ignore_autogenerated

/*
Copyright 2018 Platform 9 Systems, Inc.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package sshproviderconfig

import (
	net "net"

	runtime "k8s.io/apimachinery/pkg/runtime"
	v1alpha1 "k8s.io/kubernetes/pkg/kubelet/apis/kubeletconfig/v1alpha1"
	kubeproxyconfig_v1alpha1 "k8s.io/kubernetes/pkg/proxy/apis/kubeproxyconfig/v1alpha1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EtcdMember) DeepCopyInto(out *EtcdMember) {
	*out = *in
	if in.PeerURLs != nil {
		in, out := &in.PeerURLs, &out.PeerURLs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ClientURLs != nil {
		in, out := &in.ClientURLs, &out.ClientURLs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EtcdMember.
func (in *EtcdMember) DeepCopy() *EtcdMember {
	if in == nil {
		return nil
	}
	out := new(EtcdMember)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSHClusterProviderConfig) DeepCopyInto(out *SSHClusterProviderConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.VIPConfiguration != nil {
		in, out := &in.VIPConfiguration, &out.VIPConfiguration
		if *in == nil {
			*out = nil
		} else {
			*out = new(VIPConfiguration)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSHClusterProviderConfig.
func (in *SSHClusterProviderConfig) DeepCopy() *SSHClusterProviderConfig {
	if in == nil {
		return nil
	}
	out := new(SSHClusterProviderConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SSHClusterProviderConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSHClusterProviderStatus) DeepCopyInto(out *SSHClusterProviderStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSHClusterProviderStatus.
func (in *SSHClusterProviderStatus) DeepCopy() *SSHClusterProviderStatus {
	if in == nil {
		return nil
	}
	out := new(SSHClusterProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SSHClusterProviderStatus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSHConfig) DeepCopyInto(out *SSHConfig) {
	*out = *in
	if in.PublicKeys != nil {
		in, out := &in.PublicKeys, &out.PublicKeys
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSHConfig.
func (in *SSHConfig) DeepCopy() *SSHConfig {
	if in == nil {
		return nil
	}
	out := new(SSHConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSHMachineProviderConfig) DeepCopyInto(out *SSHMachineProviderConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.KubeletConfiguration != nil {
		in, out := &in.KubeletConfiguration, &out.KubeletConfiguration
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1alpha1.KubeletConfiguration)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.KubeProxyConfiguration != nil {
		in, out := &in.KubeProxyConfiguration, &out.KubeProxyConfiguration
		if *in == nil {
			*out = nil
		} else {
			*out = new(kubeproxyconfig_v1alpha1.KubeProxyConfiguration)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSHMachineProviderConfig.
func (in *SSHMachineProviderConfig) DeepCopy() *SSHMachineProviderConfig {
	if in == nil {
		return nil
	}
	out := new(SSHMachineProviderConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SSHMachineProviderConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSHMachineProviderStatus) DeepCopyInto(out *SSHMachineProviderStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.SSHConfig != nil {
		in, out := &in.SSHConfig, &out.SSHConfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(SSHConfig)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.EtcdMember != nil {
		in, out := &in.EtcdMember, &out.EtcdMember
		if *in == nil {
			*out = nil
		} else {
			*out = new(EtcdMember)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSHMachineProviderStatus.
func (in *SSHMachineProviderStatus) DeepCopy() *SSHMachineProviderStatus {
	if in == nil {
		return nil
	}
	out := new(SSHMachineProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SSHMachineProviderStatus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VIPConfiguration) DeepCopyInto(out *VIPConfiguration) {
	*out = *in
	if in.IP != nil {
		in, out := &in.IP, &out.IP
		*out = make(net.IP, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VIPConfiguration.
func (in *VIPConfiguration) DeepCopy() *VIPConfiguration {
	if in == nil {
		return nil
	}
	out := new(VIPConfiguration)
	in.DeepCopyInto(out)
	return out
}
