// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type NetdnsSpec struct {
	// +kubebuilder:validation:Optional
	Schedule *SchedulePlan `json:"schedule,omitempty"`

	// +kubebuilder:validation:Optional
	SourceAgentNodeSelector *metav1.LabelSelector `json:"sourceAgentNodeSelector,omitempty"`

	// +kubebuilder:validation:Optional
	Target *NetDnsTarget `json:"target,omitempty"`

	// +kubebuilder:validation:Optional
	Request *NetdnsRequest `json:"request,omitempty"`

	// +kubebuilder:validation:Optional
	SuccessCondition *NetSuccessCondition `json:"success,omitempty"`
}

type NetDnsTarget struct {
	// +kubebuilder:validation:Optional
	NetDnsTargetUser *NetDnsTargetUserSpec `json:"targetUser,omitempty"`
	// +kubebuilder:validation:Optional
	NetDnsTargetDns *NetDnsTargetDnsSpec `json:"targetDns,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=udp
	// +kubebuilder:validation:Type:=string
	// +kubebuilder:validation:Enum=udp;tcp;tcp-tls
	Protocol *string `json:"protocol,omitempty"`
}

type NetDnsTargetUserSpec struct {
	// +kubebuilder:validation:Optional
	Server *string `json:"server,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:default=53
	Port *int `json:"port,omitempty"`
}

type NetDnsTargetDnsSpec struct {
	// +kubebuilder:validation:Optional
	ServiceNamespacedName *string `json:"serviceNamespaceName,omitempty"`
	// +kubebuilder:default=true
	// +kubebuilder:validation:Optional
	TestIPv4 *bool `json:"testIPv4,omitempty"`
	// +kubebuilder:default=false
	// +kubebuilder:validation:Optional
	TestIPv6 *bool `json:"testIPv6,omitempty"`
}

type NetdnsRequest struct {

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=2
	// +kubebuilder:validation:Minimum=1
	DurationInSecond *uint64 `json:"durationInSecond,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=5
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=20
	QPS *uint64 `json:"qps,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=5
	// +kubebuilder:validation:Minimum=1
	PerRequestTimeoutInMS *uint64 `json:"perRequestTimeoutInMS,omitempty"`

	// +kubebuilder:default=kubernetes.default.svc.cluster.local
	// +kubebuilder:validation:Optional
	Domain string `json:"domain"`
}

// scope(Namespaced or Cluster)
// +kubebuilder:resource:categories={spiderdoctor},path="netdnss",singular="netdns",scope="Cluster"
// +kubebuilder:printcolumn:JSONPath=".status.finish",description="finish",name="finish",type=boolean
// +kubebuilder:printcolumn:JSONPath=".status.expectedRound",description="expectedRound",name="expectedRound",type=integer
// +kubebuilder:printcolumn:JSONPath=".status.doneRound",description="doneRound",name="doneRound",type=integer
// +kubebuilder:printcolumn:JSONPath=".status.lastRoundStatus",description="lastRoundStatus",name="lastRoundStatus",type=string
// +kubebuilder:printcolumn:JSONPath=".spec.schedule.schedule",description="schedule",name="schedule",type=string
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +genclient
// +genclient:nonNamespaced

type Netdns struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   NetdnsSpec `json:"spec,omitempty"`
	Status TaskStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

type NetdnsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Netdns `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Netdns{}, &NetdnsList{})
}
