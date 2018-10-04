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

package v1alpha2

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientConnectionConfiguration) DeepCopyInto(out *ClientConnectionConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientConnectionConfiguration.
func (in *ClientConnectionConfiguration) DeepCopy() *ClientConnectionConfiguration {
	if in == nil {
		return nil
	}
	out := new(ClientConnectionConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterConfig) DeepCopyInto(out *ClusterConfig) {
	*out = *in
	if in.KubeAPIServer != nil {
		in, out := &in.KubeAPIServer, &out.KubeAPIServer
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.KubeDNS != nil {
		in, out := &in.KubeDNS, &out.KubeDNS
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.KubeControllerManager != nil {
		in, out := &in.KubeControllerManager, &out.KubeControllerManager
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.KubeScheduler != nil {
		in, out := &in.KubeScheduler, &out.KubeScheduler
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.KubeProxy != nil {
		in, out := &in.KubeProxy, &out.KubeProxy
		*out = new(KubeProxyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Kubelet != nil {
		in, out := &in.Kubelet, &out.Kubelet
		*out = new(KubeletConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.NetworkBackend != nil {
		in, out := &in.NetworkBackend, &out.NetworkBackend
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.KeepAlived != nil {
		in, out := &in.KeepAlived, &out.KeepAlived
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConfig.
func (in *ClusterConfig) DeepCopy() *ClusterConfig {
	if in == nil {
		return nil
	}
	out := new(ClusterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.EtcdCASecret != nil {
		in, out := &in.EtcdCASecret, &out.EtcdCASecret
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.APIServerCASecret != nil {
		in, out := &in.APIServerCASecret, &out.APIServerCASecret
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.FrontProxyCASecret != nil {
		in, out := &in.FrontProxyCASecret, &out.FrontProxyCASecret
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.ServiceAccountKeySecret != nil {
		in, out := &in.ServiceAccountKeySecret, &out.ServiceAccountKeySecret
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.BootstrapTokenSecret != nil {
		in, out := &in.BootstrapTokenSecret, &out.BootstrapTokenSecret
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.VIPConfiguration != nil {
		in, out := &in.VIPConfiguration, &out.VIPConfiguration
		*out = new(VIPConfiguration)
		**out = **in
	}
	if in.ClusterConfig != nil {
		in, out := &in.ClusterConfig, &out.ClusterConfig
		*out = new(ClusterConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterSpec) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.EtcdMembers != nil {
		in, out := &in.EtcdMembers, &out.EtcdMembers
		*out = make([]EtcdMember, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterStatus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

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
func (in *KubeProxyConfiguration) DeepCopyInto(out *KubeProxyConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.FeatureGates != nil {
		in, out := &in.FeatureGates, &out.FeatureGates
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.ClientConnection = in.ClientConnection
	in.IPTables.DeepCopyInto(&out.IPTables)
	out.IPVS = in.IPVS
	if in.OOMScoreAdj != nil {
		in, out := &in.OOMScoreAdj, &out.OOMScoreAdj
		*out = new(int32)
		**out = **in
	}
	out.UDPIdleTimeout = in.UDPIdleTimeout
	in.Conntrack.DeepCopyInto(&out.Conntrack)
	out.ConfigSyncPeriod = in.ConfigSyncPeriod
	if in.NodePortAddresses != nil {
		in, out := &in.NodePortAddresses, &out.NodePortAddresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeProxyConfiguration.
func (in *KubeProxyConfiguration) DeepCopy() *KubeProxyConfiguration {
	if in == nil {
		return nil
	}
	out := new(KubeProxyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeProxyConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeProxyConntrackConfiguration) DeepCopyInto(out *KubeProxyConntrackConfiguration) {
	*out = *in
	if in.Max != nil {
		in, out := &in.Max, &out.Max
		*out = new(int32)
		**out = **in
	}
	if in.MaxPerCore != nil {
		in, out := &in.MaxPerCore, &out.MaxPerCore
		*out = new(int32)
		**out = **in
	}
	if in.Min != nil {
		in, out := &in.Min, &out.Min
		*out = new(int32)
		**out = **in
	}
	if in.TCPEstablishedTimeout != nil {
		in, out := &in.TCPEstablishedTimeout, &out.TCPEstablishedTimeout
		*out = new(metav1.Duration)
		**out = **in
	}
	if in.TCPCloseWaitTimeout != nil {
		in, out := &in.TCPCloseWaitTimeout, &out.TCPCloseWaitTimeout
		*out = new(metav1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeProxyConntrackConfiguration.
func (in *KubeProxyConntrackConfiguration) DeepCopy() *KubeProxyConntrackConfiguration {
	if in == nil {
		return nil
	}
	out := new(KubeProxyConntrackConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeProxyIPTablesConfiguration) DeepCopyInto(out *KubeProxyIPTablesConfiguration) {
	*out = *in
	if in.MasqueradeBit != nil {
		in, out := &in.MasqueradeBit, &out.MasqueradeBit
		*out = new(int32)
		**out = **in
	}
	out.SyncPeriod = in.SyncPeriod
	out.MinSyncPeriod = in.MinSyncPeriod
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeProxyIPTablesConfiguration.
func (in *KubeProxyIPTablesConfiguration) DeepCopy() *KubeProxyIPTablesConfiguration {
	if in == nil {
		return nil
	}
	out := new(KubeProxyIPTablesConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeProxyIPVSConfiguration) DeepCopyInto(out *KubeProxyIPVSConfiguration) {
	*out = *in
	out.SyncPeriod = in.SyncPeriod
	out.MinSyncPeriod = in.MinSyncPeriod
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeProxyIPVSConfiguration.
func (in *KubeProxyIPVSConfiguration) DeepCopy() *KubeProxyIPVSConfiguration {
	if in == nil {
		return nil
	}
	out := new(KubeProxyIPVSConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeletAnonymousAuthentication) DeepCopyInto(out *KubeletAnonymousAuthentication) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeletAnonymousAuthentication.
func (in *KubeletAnonymousAuthentication) DeepCopy() *KubeletAnonymousAuthentication {
	if in == nil {
		return nil
	}
	out := new(KubeletAnonymousAuthentication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeletAuthentication) DeepCopyInto(out *KubeletAuthentication) {
	*out = *in
	out.X509 = in.X509
	in.Webhook.DeepCopyInto(&out.Webhook)
	in.Anonymous.DeepCopyInto(&out.Anonymous)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeletAuthentication.
func (in *KubeletAuthentication) DeepCopy() *KubeletAuthentication {
	if in == nil {
		return nil
	}
	out := new(KubeletAuthentication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeletAuthorization) DeepCopyInto(out *KubeletAuthorization) {
	*out = *in
	out.Webhook = in.Webhook
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeletAuthorization.
func (in *KubeletAuthorization) DeepCopy() *KubeletAuthorization {
	if in == nil {
		return nil
	}
	out := new(KubeletAuthorization)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeletConfiguration) DeepCopyInto(out *KubeletConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.SyncFrequency = in.SyncFrequency
	out.FileCheckFrequency = in.FileCheckFrequency
	out.HTTPCheckFrequency = in.HTTPCheckFrequency
	if in.StaticPodURLHeader != nil {
		in, out := &in.StaticPodURLHeader, &out.StaticPodURLHeader
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	if in.TLSCipherSuites != nil {
		in, out := &in.TLSCipherSuites, &out.TLSCipherSuites
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Authentication.DeepCopyInto(&out.Authentication)
	out.Authorization = in.Authorization
	if in.RegistryPullQPS != nil {
		in, out := &in.RegistryPullQPS, &out.RegistryPullQPS
		*out = new(int32)
		**out = **in
	}
	if in.EventRecordQPS != nil {
		in, out := &in.EventRecordQPS, &out.EventRecordQPS
		*out = new(int32)
		**out = **in
	}
	if in.EnableDebuggingHandlers != nil {
		in, out := &in.EnableDebuggingHandlers, &out.EnableDebuggingHandlers
		*out = new(bool)
		**out = **in
	}
	if in.HealthzPort != nil {
		in, out := &in.HealthzPort, &out.HealthzPort
		*out = new(int32)
		**out = **in
	}
	if in.OOMScoreAdj != nil {
		in, out := &in.OOMScoreAdj, &out.OOMScoreAdj
		*out = new(int32)
		**out = **in
	}
	if in.ClusterDNS != nil {
		in, out := &in.ClusterDNS, &out.ClusterDNS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.StreamingConnectionIdleTimeout = in.StreamingConnectionIdleTimeout
	out.NodeStatusUpdateFrequency = in.NodeStatusUpdateFrequency
	out.ImageMinimumGCAge = in.ImageMinimumGCAge
	if in.ImageGCHighThresholdPercent != nil {
		in, out := &in.ImageGCHighThresholdPercent, &out.ImageGCHighThresholdPercent
		*out = new(int32)
		**out = **in
	}
	if in.ImageGCLowThresholdPercent != nil {
		in, out := &in.ImageGCLowThresholdPercent, &out.ImageGCLowThresholdPercent
		*out = new(int32)
		**out = **in
	}
	out.VolumeStatsAggPeriod = in.VolumeStatsAggPeriod
	if in.CgroupsPerQOS != nil {
		in, out := &in.CgroupsPerQOS, &out.CgroupsPerQOS
		*out = new(bool)
		**out = **in
	}
	out.CPUManagerReconcilePeriod = in.CPUManagerReconcilePeriod
	out.RuntimeRequestTimeout = in.RuntimeRequestTimeout
	if in.PodPidsLimit != nil {
		in, out := &in.PodPidsLimit, &out.PodPidsLimit
		*out = new(int64)
		**out = **in
	}
	if in.CPUCFSQuota != nil {
		in, out := &in.CPUCFSQuota, &out.CPUCFSQuota
		*out = new(bool)
		**out = **in
	}
	if in.KubeAPIQPS != nil {
		in, out := &in.KubeAPIQPS, &out.KubeAPIQPS
		*out = new(int32)
		**out = **in
	}
	if in.SerializeImagePulls != nil {
		in, out := &in.SerializeImagePulls, &out.SerializeImagePulls
		*out = new(bool)
		**out = **in
	}
	if in.EvictionHard != nil {
		in, out := &in.EvictionHard, &out.EvictionHard
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.EvictionSoft != nil {
		in, out := &in.EvictionSoft, &out.EvictionSoft
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.EvictionSoftGracePeriod != nil {
		in, out := &in.EvictionSoftGracePeriod, &out.EvictionSoftGracePeriod
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.EvictionPressureTransitionPeriod = in.EvictionPressureTransitionPeriod
	if in.EvictionMinimumReclaim != nil {
		in, out := &in.EvictionMinimumReclaim, &out.EvictionMinimumReclaim
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.EnableControllerAttachDetach != nil {
		in, out := &in.EnableControllerAttachDetach, &out.EnableControllerAttachDetach
		*out = new(bool)
		**out = **in
	}
	if in.MakeIPTablesUtilChains != nil {
		in, out := &in.MakeIPTablesUtilChains, &out.MakeIPTablesUtilChains
		*out = new(bool)
		**out = **in
	}
	if in.IPTablesMasqueradeBit != nil {
		in, out := &in.IPTablesMasqueradeBit, &out.IPTablesMasqueradeBit
		*out = new(int32)
		**out = **in
	}
	if in.IPTablesDropBit != nil {
		in, out := &in.IPTablesDropBit, &out.IPTablesDropBit
		*out = new(int32)
		**out = **in
	}
	if in.FeatureGates != nil {
		in, out := &in.FeatureGates, &out.FeatureGates
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.FailSwapOn != nil {
		in, out := &in.FailSwapOn, &out.FailSwapOn
		*out = new(bool)
		**out = **in
	}
	if in.ContainerLogMaxFiles != nil {
		in, out := &in.ContainerLogMaxFiles, &out.ContainerLogMaxFiles
		*out = new(int32)
		**out = **in
	}
	if in.SystemReserved != nil {
		in, out := &in.SystemReserved, &out.SystemReserved
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.KubeReserved != nil {
		in, out := &in.KubeReserved, &out.KubeReserved
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.EnforceNodeAllocatable != nil {
		in, out := &in.EnforceNodeAllocatable, &out.EnforceNodeAllocatable
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeletConfiguration.
func (in *KubeletConfiguration) DeepCopy() *KubeletConfiguration {
	if in == nil {
		return nil
	}
	out := new(KubeletConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeletConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeletWebhookAuthentication) DeepCopyInto(out *KubeletWebhookAuthentication) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	out.CacheTTL = in.CacheTTL
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeletWebhookAuthentication.
func (in *KubeletWebhookAuthentication) DeepCopy() *KubeletWebhookAuthentication {
	if in == nil {
		return nil
	}
	out := new(KubeletWebhookAuthentication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeletWebhookAuthorization) DeepCopyInto(out *KubeletWebhookAuthorization) {
	*out = *in
	out.CacheAuthorizedTTL = in.CacheAuthorizedTTL
	out.CacheUnauthorizedTTL = in.CacheUnauthorizedTTL
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeletWebhookAuthorization.
func (in *KubeletWebhookAuthorization) DeepCopy() *KubeletWebhookAuthorization {
	if in == nil {
		return nil
	}
	out := new(KubeletWebhookAuthorization)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeletX509Authentication) DeepCopyInto(out *KubeletX509Authentication) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeletX509Authentication.
func (in *KubeletX509Authentication) DeepCopy() *KubeletX509Authentication {
	if in == nil {
		return nil
	}
	out := new(KubeletX509Authentication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineComponentVersions) DeepCopyInto(out *MachineComponentVersions) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineComponentVersions.
func (in *MachineComponentVersions) DeepCopy() *MachineComponentVersions {
	if in == nil {
		return nil
	}
	out := new(MachineComponentVersions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineSpec) DeepCopyInto(out *MachineSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]MachineRole, len(*in))
		copy(*out, *in)
	}
	if in.ComponentVersions != nil {
		in, out := &in.ComponentVersions, &out.ComponentVersions
		*out = new(MachineComponentVersions)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineSpec.
func (in *MachineSpec) DeepCopy() *MachineSpec {
	if in == nil {
		return nil
	}
	out := new(MachineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineSpec) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineStatus) DeepCopyInto(out *MachineStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.SSHConfig != nil {
		in, out := &in.SSHConfig, &out.SSHConfig
		*out = new(SSHConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.EtcdMember != nil {
		in, out := &in.EtcdMember, &out.EtcdMember
		*out = new(EtcdMember)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineStatus.
func (in *MachineStatus) DeepCopy() *MachineStatus {
	if in == nil {
		return nil
	}
	out := new(MachineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineStatus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProvisionedMachine) DeepCopyInto(out *ProvisionedMachine) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProvisionedMachine.
func (in *ProvisionedMachine) DeepCopy() *ProvisionedMachine {
	if in == nil {
		return nil
	}
	out := new(ProvisionedMachine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProvisionedMachine) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProvisionedMachineList) DeepCopyInto(out *ProvisionedMachineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ProvisionedMachine, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProvisionedMachineList.
func (in *ProvisionedMachineList) DeepCopy() *ProvisionedMachineList {
	if in == nil {
		return nil
	}
	out := new(ProvisionedMachineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProvisionedMachineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProvisionedMachineSpec) DeepCopyInto(out *ProvisionedMachineSpec) {
	*out = *in
	if in.SSHConfig != nil {
		in, out := &in.SSHConfig, &out.SSHConfig
		*out = new(SSHConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProvisionedMachineSpec.
func (in *ProvisionedMachineSpec) DeepCopy() *ProvisionedMachineSpec {
	if in == nil {
		return nil
	}
	out := new(ProvisionedMachineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProvisionedMachineStatus) DeepCopyInto(out *ProvisionedMachineStatus) {
	*out = *in
	if in.MachineRef != nil {
		in, out := &in.MachineRef, &out.MachineRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProvisionedMachineStatus.
func (in *ProvisionedMachineStatus) DeepCopy() *ProvisionedMachineStatus {
	if in == nil {
		return nil
	}
	out := new(ProvisionedMachineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSHConfig) DeepCopyInto(out *SSHConfig) {
	*out = *in
	if in.PublicKeys != nil {
		in, out := &in.PublicKeys, &out.PublicKeys
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.CredentialSecret = in.CredentialSecret
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
func (in *VIPConfiguration) DeepCopyInto(out *VIPConfiguration) {
	*out = *in
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
