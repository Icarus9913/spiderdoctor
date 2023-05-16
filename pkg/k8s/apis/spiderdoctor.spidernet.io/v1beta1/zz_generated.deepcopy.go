//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HttpTarget) DeepCopyInto(out *HttpTarget) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HttpTarget.
func (in *HttpTarget) DeepCopy() *HttpTarget {
	if in == nil {
		return nil
	}
	out := new(HttpTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetDnsTarget) DeepCopyInto(out *NetDnsTarget) {
	*out = *in
	if in.NetDnsTargetUser != nil {
		in, out := &in.NetDnsTargetUser, &out.NetDnsTargetUser
		*out = new(NetDnsTargetUserSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.NetDnsTargetDns != nil {
		in, out := &in.NetDnsTargetDns, &out.NetDnsTargetDns
		*out = new(NetDnsTargetDnsSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetDnsTarget.
func (in *NetDnsTarget) DeepCopy() *NetDnsTarget {
	if in == nil {
		return nil
	}
	out := new(NetDnsTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetDnsTargetDnsSpec) DeepCopyInto(out *NetDnsTargetDnsSpec) {
	*out = *in
	if in.ServiceNamespacedName != nil {
		in, out := &in.ServiceNamespacedName, &out.ServiceNamespacedName
		*out = new(string)
		**out = **in
	}
	if in.TestIPv4 != nil {
		in, out := &in.TestIPv4, &out.TestIPv4
		*out = new(bool)
		**out = **in
	}
	if in.TestIPv6 != nil {
		in, out := &in.TestIPv6, &out.TestIPv6
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetDnsTargetDnsSpec.
func (in *NetDnsTargetDnsSpec) DeepCopy() *NetDnsTargetDnsSpec {
	if in == nil {
		return nil
	}
	out := new(NetDnsTargetDnsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetDnsTargetUserSpec) DeepCopyInto(out *NetDnsTargetUserSpec) {
	*out = *in
	if in.Server != nil {
		in, out := &in.Server, &out.Server
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetDnsTargetUserSpec.
func (in *NetDnsTargetUserSpec) DeepCopy() *NetDnsTargetUserSpec {
	if in == nil {
		return nil
	}
	out := new(NetDnsTargetUserSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetSuccessCondition) DeepCopyInto(out *NetSuccessCondition) {
	*out = *in
	if in.SuccessRate != nil {
		in, out := &in.SuccessRate, &out.SuccessRate
		*out = new(float64)
		**out = **in
	}
	if in.MeanAccessDelayInMs != nil {
		in, out := &in.MeanAccessDelayInMs, &out.MeanAccessDelayInMs
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetSuccessCondition.
func (in *NetSuccessCondition) DeepCopy() *NetSuccessCondition {
	if in == nil {
		return nil
	}
	out := new(NetSuccessCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Netdns) DeepCopyInto(out *Netdns) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Netdns.
func (in *Netdns) DeepCopy() *Netdns {
	if in == nil {
		return nil
	}
	out := new(Netdns)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Netdns) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetdnsList) DeepCopyInto(out *NetdnsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Netdns, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetdnsList.
func (in *NetdnsList) DeepCopy() *NetdnsList {
	if in == nil {
		return nil
	}
	out := new(NetdnsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetdnsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetdnsRequest) DeepCopyInto(out *NetdnsRequest) {
	*out = *in
	if in.DurationInSecond != nil {
		in, out := &in.DurationInSecond, &out.DurationInSecond
		*out = new(uint64)
		**out = **in
	}
	if in.QPS != nil {
		in, out := &in.QPS, &out.QPS
		*out = new(uint64)
		**out = **in
	}
	if in.PerRequestTimeoutInMS != nil {
		in, out := &in.PerRequestTimeoutInMS, &out.PerRequestTimeoutInMS
		*out = new(uint64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetdnsRequest.
func (in *NetdnsRequest) DeepCopy() *NetdnsRequest {
	if in == nil {
		return nil
	}
	out := new(NetdnsRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetdnsSpec) DeepCopyInto(out *NetdnsSpec) {
	*out = *in
	if in.Schedule != nil {
		in, out := &in.Schedule, &out.Schedule
		*out = new(SchedulePlan)
		(*in).DeepCopyInto(*out)
	}
	if in.SourceAgentNodeSelector != nil {
		in, out := &in.SourceAgentNodeSelector, &out.SourceAgentNodeSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = new(NetDnsTarget)
		(*in).DeepCopyInto(*out)
	}
	if in.Request != nil {
		in, out := &in.Request, &out.Request
		*out = new(NetdnsRequest)
		(*in).DeepCopyInto(*out)
	}
	if in.SuccessCondition != nil {
		in, out := &in.SuccessCondition, &out.SuccessCondition
		*out = new(NetSuccessCondition)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetdnsSpec.
func (in *NetdnsSpec) DeepCopy() *NetdnsSpec {
	if in == nil {
		return nil
	}
	out := new(NetdnsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Nethttp) DeepCopyInto(out *Nethttp) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Nethttp.
func (in *Nethttp) DeepCopy() *Nethttp {
	if in == nil {
		return nil
	}
	out := new(Nethttp)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Nethttp) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NethttpList) DeepCopyInto(out *NethttpList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Nethttp, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NethttpList.
func (in *NethttpList) DeepCopy() *NethttpList {
	if in == nil {
		return nil
	}
	out := new(NethttpList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NethttpList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NethttpRequest) DeepCopyInto(out *NethttpRequest) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NethttpRequest.
func (in *NethttpRequest) DeepCopy() *NethttpRequest {
	if in == nil {
		return nil
	}
	out := new(NethttpRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NethttpSpec) DeepCopyInto(out *NethttpSpec) {
	*out = *in
	if in.Schedule != nil {
		in, out := &in.Schedule, &out.Schedule
		*out = new(SchedulePlan)
		(*in).DeepCopyInto(*out)
	}
	if in.SourceAgentNodeSelector != nil {
		in, out := &in.SourceAgentNodeSelector, &out.SourceAgentNodeSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = new(NethttpTarget)
		(*in).DeepCopyInto(*out)
	}
	if in.Request != nil {
		in, out := &in.Request, &out.Request
		*out = new(NethttpRequest)
		**out = **in
	}
	if in.SuccessCondition != nil {
		in, out := &in.SuccessCondition, &out.SuccessCondition
		*out = new(NetSuccessCondition)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NethttpSpec.
func (in *NethttpSpec) DeepCopy() *NethttpSpec {
	if in == nil {
		return nil
	}
	out := new(NethttpSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NethttpTarget) DeepCopyInto(out *NethttpTarget) {
	*out = *in
	if in.TargetUser != nil {
		in, out := &in.TargetUser, &out.TargetUser
		*out = new(HttpTarget)
		**out = **in
	}
	if in.TargetPod != nil {
		in, out := &in.TargetPod, &out.TargetPod
		*out = new(TargetPodSepc)
		(*in).DeepCopyInto(*out)
	}
	if in.TargetAgent != nil {
		in, out := &in.TargetAgent, &out.TargetAgent
		*out = new(TargetAgentSepc)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NethttpTarget.
func (in *NethttpTarget) DeepCopy() *NethttpTarget {
	if in == nil {
		return nil
	}
	out := new(NethttpTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PluginReport) DeepCopyInto(out *PluginReport) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PluginReport.
func (in *PluginReport) DeepCopy() *PluginReport {
	if in == nil {
		return nil
	}
	out := new(PluginReport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PluginReport) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PluginReportList) DeepCopyInto(out *PluginReportList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PluginReport, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PluginReportList.
func (in *PluginReportList) DeepCopy() *PluginReportList {
	if in == nil {
		return nil
	}
	out := new(PluginReportList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PluginReportList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PluginReportSpec) DeepCopyInto(out *PluginReportSpec) {
	*out = *in
	in.StartTimeStamp.DeepCopyInto(&out.StartTimeStamp)
	in.EndTimeStamp.DeepCopyInto(&out.EndTimeStamp)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PluginReportSpec.
func (in *PluginReportSpec) DeepCopy() *PluginReportSpec {
	if in == nil {
		return nil
	}
	out := new(PluginReportSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulePlan) DeepCopyInto(out *SchedulePlan) {
	*out = *in
	if in.Schedule != nil {
		in, out := &in.Schedule, &out.Schedule
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulePlan.
func (in *SchedulePlan) DeepCopy() *SchedulePlan {
	if in == nil {
		return nil
	}
	out := new(SchedulePlan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusHistoryRecord) DeepCopyInto(out *StatusHistoryRecord) {
	*out = *in
	in.StartTimeStamp.DeepCopyInto(&out.StartTimeStamp)
	if in.EndTimeStamp != nil {
		in, out := &in.EndTimeStamp, &out.EndTimeStamp
		*out = (*in).DeepCopy()
	}
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(string)
		**out = **in
	}
	in.DeadLineTimeStamp.DeepCopyInto(&out.DeadLineTimeStamp)
	if in.ExpectedActorNumber != nil {
		in, out := &in.ExpectedActorNumber, &out.ExpectedActorNumber
		*out = new(int)
		**out = **in
	}
	if in.FailedAgentNodeList != nil {
		in, out := &in.FailedAgentNodeList, &out.FailedAgentNodeList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SucceedAgentNodeList != nil {
		in, out := &in.SucceedAgentNodeList, &out.SucceedAgentNodeList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NotReportAgentNodeList != nil {
		in, out := &in.NotReportAgentNodeList, &out.NotReportAgentNodeList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusHistoryRecord.
func (in *StatusHistoryRecord) DeepCopy() *StatusHistoryRecord {
	if in == nil {
		return nil
	}
	out := new(StatusHistoryRecord)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetAgentSepc) DeepCopyInto(out *TargetAgentSepc) {
	*out = *in
	if in.TestIPv4 != nil {
		in, out := &in.TestIPv4, &out.TestIPv4
		*out = new(bool)
		**out = **in
	}
	if in.TestIPv6 != nil {
		in, out := &in.TestIPv6, &out.TestIPv6
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetAgentSepc.
func (in *TargetAgentSepc) DeepCopy() *TargetAgentSepc {
	if in == nil {
		return nil
	}
	out := new(TargetAgentSepc)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetPodSepc) DeepCopyInto(out *TargetPodSepc) {
	*out = *in
	in.PodLabelSelector.DeepCopyInto(&out.PodLabelSelector)
	if in.TestIPv4 != nil {
		in, out := &in.TestIPv4, &out.TestIPv4
		*out = new(bool)
		**out = **in
	}
	if in.TestIPv6 != nil {
		in, out := &in.TestIPv6, &out.TestIPv6
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetPodSepc.
func (in *TargetPodSepc) DeepCopy() *TargetPodSepc {
	if in == nil {
		return nil
	}
	out := new(TargetPodSepc)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskStatus) DeepCopyInto(out *TaskStatus) {
	*out = *in
	if in.ExpectedRound != nil {
		in, out := &in.ExpectedRound, &out.ExpectedRound
		*out = new(int64)
		**out = **in
	}
	if in.DoneRound != nil {
		in, out := &in.DoneRound, &out.DoneRound
		*out = new(int64)
		**out = **in
	}
	if in.LastRoundStatus != nil {
		in, out := &in.LastRoundStatus, &out.LastRoundStatus
		*out = new(string)
		**out = **in
	}
	if in.History != nil {
		in, out := &in.History, &out.History
		*out = make([]StatusHistoryRecord, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskStatus.
func (in *TaskStatus) DeepCopy() *TaskStatus {
	if in == nil {
		return nil
	}
	out := new(TaskStatus)
	in.DeepCopyInto(out)
	return out
}
