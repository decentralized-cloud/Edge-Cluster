// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: edge-cluster-pod-messages.proto

package __

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//*
// The different error types
type PodConditionType int32

const (
	// ContainersReady indicates whether all containers in the pod are ready.
	PodConditionType_ContainersReady PodConditionType = 0
	// PodInitialized means that all init containers in the pod have started successfully.
	PodConditionType_PodInitialized PodConditionType = 1
	// PodReady means the pod is able to service requests and should be added to the
	// load balancing pools of all matching services.
	PodConditionType_PodReady PodConditionType = 2
	// PodScheduled represents status of the scheduling process for this pod.
	PodConditionType_PodScheduled PodConditionType = 3
)

// Enum value maps for PodConditionType.
var (
	PodConditionType_name = map[int32]string{
		0: "ContainersReady",
		1: "PodInitialized",
		2: "PodReady",
		3: "PodScheduled",
	}
	PodConditionType_value = map[string]int32{
		"ContainersReady": 0,
		"PodInitialized":  1,
		"PodReady":        2,
		"PodScheduled":    3,
	}
)

func (x PodConditionType) Enum() *PodConditionType {
	p := new(PodConditionType)
	*p = x
	return p
}

func (x PodConditionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PodConditionType) Descriptor() protoreflect.EnumDescriptor {
	return file_edge_cluster_pod_messages_proto_enumTypes[0].Descriptor()
}

func (PodConditionType) Type() protoreflect.EnumType {
	return &file_edge_cluster_pod_messages_proto_enumTypes[0]
}

func (x PodConditionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PodConditionType.Descriptor instead.
func (PodConditionType) EnumDescriptor() ([]byte, []int) {
	return file_edge_cluster_pod_messages_proto_rawDescGZIP(), []int{0}
}

//*
// Declares the specification of the desired behavior of the existing edge cluster pod
type PodSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the node where the Pod is deployed into
	NodeName string `protobuf:"bytes,1,opt,name=nodeName,proto3" json:"nodeName,omitempty"`
}

func (x *PodSpec) Reset() {
	*x = PodSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_edge_cluster_pod_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PodSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PodSpec) ProtoMessage() {}

func (x *PodSpec) ProtoReflect() protoreflect.Message {
	mi := &file_edge_cluster_pod_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PodSpec.ProtoReflect.Descriptor instead.
func (*PodSpec) Descriptor() ([]byte, []int) {
	return file_edge_cluster_pod_messages_proto_rawDescGZIP(), []int{0}
}

func (x *PodSpec) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

//*
// Declares the most recently observed status of the existing edge cluster pod
type PodCondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Type is the type of the condition.
	Type PodConditionType `protobuf:"varint,1,opt,name=type,proto3,enum=edgecluster.PodConditionType" json:"type,omitempty"`
	// Status is the status of the condition.
	Status ConditionStatus `protobuf:"varint,2,opt,name=status,proto3,enum=edgecluster.ConditionStatus" json:"status,omitempty"`
	// Last time we got an update on a given condition.
	LastProbeTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=LastProbeTime,proto3" json:"LastProbeTime,omitempty"`
	// Last time the condition transitioned from one status to another.
	LastTransitionTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=LastTransitionTime,proto3" json:"LastTransitionTime,omitempty"`
	// Unique, one-word, CamelCase reason for the condition's last transition.
	Reason string `protobuf:"bytes,5,opt,name=reason,proto3" json:"reason,omitempty"`
	// Human-readable message indicating details about last transition.
	Message string `protobuf:"bytes,6,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PodCondition) Reset() {
	*x = PodCondition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_edge_cluster_pod_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PodCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PodCondition) ProtoMessage() {}

func (x *PodCondition) ProtoReflect() protoreflect.Message {
	mi := &file_edge_cluster_pod_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PodCondition.ProtoReflect.Descriptor instead.
func (*PodCondition) Descriptor() ([]byte, []int) {
	return file_edge_cluster_pod_messages_proto_rawDescGZIP(), []int{1}
}

func (x *PodCondition) GetType() PodConditionType {
	if x != nil {
		return x.Type
	}
	return PodConditionType_ContainersReady
}

func (x *PodCondition) GetStatus() ConditionStatus {
	if x != nil {
		return x.Status
	}
	return ConditionStatus_ConditionTrue
}

func (x *PodCondition) GetLastProbeTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastProbeTime
	}
	return nil
}

func (x *PodCondition) GetLastTransitionTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastTransitionTime
	}
	return nil
}

func (x *PodCondition) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *PodCondition) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

//*
// Declares the most recently observed status of the existing edge cluster pod
type PodStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// IP address allocated to the pod. Routable at least within the cluster.
	HostIP string `protobuf:"bytes,1,opt,name=hostIP,proto3" json:"hostIP,omitempty"`
	// IP address allocated to the pod. Routable at least within the cluster.
	PodIP string `protobuf:"bytes,2,opt,name=podIP,proto3" json:"podIP,omitempty"`
	// Current service state of pod.
	Conditions []*PodCondition `protobuf:"bytes,3,rep,name=conditions,proto3" json:"conditions,omitempty"`
}

func (x *PodStatus) Reset() {
	*x = PodStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_edge_cluster_pod_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PodStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PodStatus) ProtoMessage() {}

func (x *PodStatus) ProtoReflect() protoreflect.Message {
	mi := &file_edge_cluster_pod_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PodStatus.ProtoReflect.Descriptor instead.
func (*PodStatus) Descriptor() ([]byte, []int) {
	return file_edge_cluster_pod_messages_proto_rawDescGZIP(), []int{2}
}

func (x *PodStatus) GetHostIP() string {
	if x != nil {
		return x.HostIP
	}
	return ""
}

func (x *PodStatus) GetPodIP() string {
	if x != nil {
		return x.PodIP
	}
	return ""
}

func (x *PodStatus) GetConditions() []*PodCondition {
	if x != nil {
		return x.Conditions
	}
	return nil
}

//*
// Declares the details of an existing edge cluster pod
type EdgeClusterPod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The pod metadata
	Metadata *ObjectMeta `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// The most recently observed status of the pod
	Status *PodStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	// The specification of the desired behavior of the pod.
	Spec *PodSpec `protobuf:"bytes,3,opt,name=spec,proto3" json:"spec,omitempty"`
}

func (x *EdgeClusterPod) Reset() {
	*x = EdgeClusterPod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_edge_cluster_pod_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EdgeClusterPod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EdgeClusterPod) ProtoMessage() {}

func (x *EdgeClusterPod) ProtoReflect() protoreflect.Message {
	mi := &file_edge_cluster_pod_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EdgeClusterPod.ProtoReflect.Descriptor instead.
func (*EdgeClusterPod) Descriptor() ([]byte, []int) {
	return file_edge_cluster_pod_messages_proto_rawDescGZIP(), []int{3}
}

func (x *EdgeClusterPod) GetMetadata() *ObjectMeta {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *EdgeClusterPod) GetStatus() *PodStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *EdgeClusterPod) GetSpec() *PodSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

//*
// ListEdgeClusterNodePodsRequest to list an existing edge cluster pods details
type ListEdgeClusterPodsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique edge cluster identifier
	EdgeClusterID string `protobuf:"bytes,1,opt,name=edgeClusterID,proto3" json:"edgeClusterID,omitempty"`
	// Optional, if provided, will be used to filter pods under the given namespace
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Optional, if provided, will be used to filter pods deployed to the given node
	NodeName string `protobuf:"bytes,3,opt,name=nodeName,proto3" json:"nodeName,omitempty"`
}

func (x *ListEdgeClusterPodsRequest) Reset() {
	*x = ListEdgeClusterPodsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_edge_cluster_pod_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEdgeClusterPodsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEdgeClusterPodsRequest) ProtoMessage() {}

func (x *ListEdgeClusterPodsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_edge_cluster_pod_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEdgeClusterPodsRequest.ProtoReflect.Descriptor instead.
func (*ListEdgeClusterPodsRequest) Descriptor() ([]byte, []int) {
	return file_edge_cluster_pod_messages_proto_rawDescGZIP(), []int{4}
}

func (x *ListEdgeClusterPodsRequest) GetEdgeClusterID() string {
	if x != nil {
		return x.EdgeClusterID
	}
	return ""
}

func (x *ListEdgeClusterPodsRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ListEdgeClusterPodsRequest) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

//*
// Response contains the result of listing an existing edge cluster pods details
type ListEdgeClusterPodsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Indicate whether the operation has any error
	Error Error `protobuf:"varint,1,opt,name=error,proto3,enum=edgecluster.Error" json:"error,omitempty"`
	// Contains error message if the operation was unsuccessful
	ErrorMessage string `protobuf:"bytes,2,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	// The list of an existing edge cluster pods details
	Pods []*EdgeClusterPod `protobuf:"bytes,3,rep,name=pods,proto3" json:"pods,omitempty"`
}

func (x *ListEdgeClusterPodsResponse) Reset() {
	*x = ListEdgeClusterPodsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_edge_cluster_pod_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEdgeClusterPodsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEdgeClusterPodsResponse) ProtoMessage() {}

func (x *ListEdgeClusterPodsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_edge_cluster_pod_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEdgeClusterPodsResponse.ProtoReflect.Descriptor instead.
func (*ListEdgeClusterPodsResponse) Descriptor() ([]byte, []int) {
	return file_edge_cluster_pod_messages_proto_rawDescGZIP(), []int{5}
}

func (x *ListEdgeClusterPodsResponse) GetError() Error {
	if x != nil {
		return x.Error
	}
	return Error_NO_ERROR
}

func (x *ListEdgeClusterPodsResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *ListEdgeClusterPodsResponse) GetPods() []*EdgeClusterPod {
	if x != nil {
		return x.Pods
	}
	return nil
}

var File_edge_cluster_pod_messages_proto protoreflect.FileDescriptor

var file_edge_cluster_pod_messages_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x65, 0x64, 0x67, 0x65, 0x2d, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2d, 0x70,
	0x6f, 0x64, 0x2d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0b, 0x65, 0x64, 0x67, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x1a, 0x0d,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25,
	0x0a, 0x07, 0x50, 0x6f, 0x64, 0x53, 0x70, 0x65, 0x63, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x6f, 0x64,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xb7, 0x02, 0x0a, 0x0c, 0x50, 0x6f, 0x64, 0x43, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x2e, 0x50, 0x6f, 0x64, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x65, 0x64, 0x67, 0x65,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x40, 0x0a, 0x0d, 0x4c, 0x61, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0d, 0x4c, 0x61, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x4a, 0x0a, 0x12, 0x4c, 0x61, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x12, 0x4c, 0x61, 0x73, 0x74, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x74, 0x0a, 0x09, 0x50, 0x6f, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x68, 0x6f, 0x73, 0x74, 0x49, 0x50, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x6f,
	0x73, 0x74, 0x49, 0x50, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x64, 0x49, 0x50, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x6f, 0x64, 0x49, 0x50, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x65, 0x64, 0x67, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x6f, 0x64,
	0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x9f, 0x01, 0x0a, 0x0e, 0x45, 0x64, 0x67, 0x65, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x6f, 0x64, 0x12, 0x33, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x64, 0x67,
	0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d,
	0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2e, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x65, 0x64, 0x67, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x6f, 0x64, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x28, 0x0a,
	0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65, 0x64,
	0x67, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x50, 0x6f, 0x64, 0x53, 0x70, 0x65,
	0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0x7c, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74, 0x45,
	0x64, 0x67, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x6f, 0x64, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x65, 0x64, 0x67, 0x65, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x64,
	0x67, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x6f, 0x64,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x9c, 0x01, 0x0a, 0x1b, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x64,
	0x67, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x22, 0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x70, 0x6f, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e,
	0x45, 0x64, 0x67, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x6f, 0x64, 0x52, 0x04,
	0x70, 0x6f, 0x64, 0x73, 0x2a, 0x5b, 0x0a, 0x10, 0x50, 0x6f, 0x64, 0x43, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x52, 0x65, 0x61, 0x64, 0x79, 0x10, 0x00, 0x12, 0x12, 0x0a,
	0x0e, 0x50, 0x6f, 0x64, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x10,
	0x01, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x6f, 0x64, 0x52, 0x65, 0x61, 0x64, 0x79, 0x10, 0x02, 0x12,
	0x10, 0x0a, 0x0c, 0x50, 0x6f, 0x64, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x10,
	0x03, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_edge_cluster_pod_messages_proto_rawDescOnce sync.Once
	file_edge_cluster_pod_messages_proto_rawDescData = file_edge_cluster_pod_messages_proto_rawDesc
)

func file_edge_cluster_pod_messages_proto_rawDescGZIP() []byte {
	file_edge_cluster_pod_messages_proto_rawDescOnce.Do(func() {
		file_edge_cluster_pod_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_edge_cluster_pod_messages_proto_rawDescData)
	})
	return file_edge_cluster_pod_messages_proto_rawDescData
}

var file_edge_cluster_pod_messages_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_edge_cluster_pod_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_edge_cluster_pod_messages_proto_goTypes = []interface{}{
	(PodConditionType)(0),               // 0: edgecluster.PodConditionType
	(*PodSpec)(nil),                     // 1: edgecluster.PodSpec
	(*PodCondition)(nil),                // 2: edgecluster.PodCondition
	(*PodStatus)(nil),                   // 3: edgecluster.PodStatus
	(*EdgeClusterPod)(nil),              // 4: edgecluster.EdgeClusterPod
	(*ListEdgeClusterPodsRequest)(nil),  // 5: edgecluster.ListEdgeClusterPodsRequest
	(*ListEdgeClusterPodsResponse)(nil), // 6: edgecluster.ListEdgeClusterPodsResponse
	(ConditionStatus)(0),                // 7: edgecluster.ConditionStatus
	(*timestamppb.Timestamp)(nil),       // 8: google.protobuf.Timestamp
	(*ObjectMeta)(nil),                  // 9: edgecluster.ObjectMeta
	(Error)(0),                          // 10: edgecluster.Error
}
var file_edge_cluster_pod_messages_proto_depIdxs = []int32{
	0,  // 0: edgecluster.PodCondition.type:type_name -> edgecluster.PodConditionType
	7,  // 1: edgecluster.PodCondition.status:type_name -> edgecluster.ConditionStatus
	8,  // 2: edgecluster.PodCondition.LastProbeTime:type_name -> google.protobuf.Timestamp
	8,  // 3: edgecluster.PodCondition.LastTransitionTime:type_name -> google.protobuf.Timestamp
	2,  // 4: edgecluster.PodStatus.conditions:type_name -> edgecluster.PodCondition
	9,  // 5: edgecluster.EdgeClusterPod.metadata:type_name -> edgecluster.ObjectMeta
	3,  // 6: edgecluster.EdgeClusterPod.status:type_name -> edgecluster.PodStatus
	1,  // 7: edgecluster.EdgeClusterPod.spec:type_name -> edgecluster.PodSpec
	10, // 8: edgecluster.ListEdgeClusterPodsResponse.error:type_name -> edgecluster.Error
	4,  // 9: edgecluster.ListEdgeClusterPodsResponse.pods:type_name -> edgecluster.EdgeClusterPod
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_edge_cluster_pod_messages_proto_init() }
func file_edge_cluster_pod_messages_proto_init() {
	if File_edge_cluster_pod_messages_proto != nil {
		return
	}
	file_commons_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_edge_cluster_pod_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PodSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_edge_cluster_pod_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PodCondition); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_edge_cluster_pod_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PodStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_edge_cluster_pod_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EdgeClusterPod); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_edge_cluster_pod_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEdgeClusterPodsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_edge_cluster_pod_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEdgeClusterPodsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_edge_cluster_pod_messages_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_edge_cluster_pod_messages_proto_goTypes,
		DependencyIndexes: file_edge_cluster_pod_messages_proto_depIdxs,
		EnumInfos:         file_edge_cluster_pod_messages_proto_enumTypes,
		MessageInfos:      file_edge_cluster_pod_messages_proto_msgTypes,
	}.Build()
	File_edge_cluster_pod_messages_proto = out.File
	file_edge_cluster_pod_messages_proto_rawDesc = nil
	file_edge_cluster_pod_messages_proto_goTypes = nil
	file_edge_cluster_pod_messages_proto_depIdxs = nil
}
