// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: edge-cluster-node-messages.proto

package edgecluster

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
// The valid conditions of node
type NodeConditionType int32

const (
	// NodeReady means kubelet is healthy and ready to accept pods.
	NodeConditionType_Ready NodeConditionType = 0
	// NodeMemoryPressure means the kubelet is under pressure due to insufficient available memory.
	NodeConditionType_MemoryPressure NodeConditionType = 1
	// NodeDiskPressure means the kubelet is under pressure due to insufficient available disk.
	NodeConditionType_DiskPressure NodeConditionType = 3
	// NodePIDPressure means the kubelet is under pressure due to insufficient available PID.
	NodeConditionType_PIDPressure NodeConditionType = 4
	// NodeNetworkUnavailable means that network for the node is not correctly configured.
	NodeConditionType_NetworkUnavailable NodeConditionType = 5
)

// Enum value maps for NodeConditionType.
var (
	NodeConditionType_name = map[int32]string{
		0: "Ready",
		1: "MemoryPressure",
		3: "DiskPressure",
		4: "PIDPressure",
		5: "NetworkUnavailable",
	}
	NodeConditionType_value = map[string]int32{
		"Ready":              0,
		"MemoryPressure":     1,
		"DiskPressure":       3,
		"PIDPressure":        4,
		"NetworkUnavailable": 5,
	}
)

func (x NodeConditionType) Enum() *NodeConditionType {
	p := new(NodeConditionType)
	*p = x
	return p
}

func (x NodeConditionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NodeConditionType) Descriptor() protoreflect.EnumDescriptor {
	return file_edge_cluster_node_messages_proto_enumTypes[0].Descriptor()
}

func (NodeConditionType) Type() protoreflect.EnumType {
	return &file_edge_cluster_node_messages_proto_enumTypes[0]
}

func (x NodeConditionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NodeConditionType.Descriptor instead.
func (NodeConditionType) EnumDescriptor() ([]byte, []int) {
	return file_edge_cluster_node_messages_proto_rawDescGZIP(), []int{0}
}

//*
// The valid address type of edge cluster node
type NodeAddressType int32

const (
	NodeAddressType_Hostname    NodeAddressType = 0
	NodeAddressType_ExternalIP  NodeAddressType = 1
	NodeAddressType_InternalIP  NodeAddressType = 2
	NodeAddressType_ExternalDNS NodeAddressType = 3
	NodeAddressType_InternalDNS NodeAddressType = 4
)

// Enum value maps for NodeAddressType.
var (
	NodeAddressType_name = map[int32]string{
		0: "Hostname",
		1: "ExternalIP",
		2: "InternalIP",
		3: "ExternalDNS",
		4: "InternalDNS",
	}
	NodeAddressType_value = map[string]int32{
		"Hostname":    0,
		"ExternalIP":  1,
		"InternalIP":  2,
		"ExternalDNS": 3,
		"InternalDNS": 4,
	}
)

func (x NodeAddressType) Enum() *NodeAddressType {
	p := new(NodeAddressType)
	*p = x
	return p
}

func (x NodeAddressType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NodeAddressType) Descriptor() protoreflect.EnumDescriptor {
	return file_edge_cluster_node_messages_proto_enumTypes[1].Descriptor()
}

func (NodeAddressType) Type() protoreflect.EnumType {
	return &file_edge_cluster_node_messages_proto_enumTypes[1]
}

func (x NodeAddressType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NodeAddressType.Descriptor instead.
func (NodeAddressType) EnumDescriptor() ([]byte, []int) {
	return file_edge_cluster_node_messages_proto_rawDescGZIP(), []int{1}
}

//*
// NodeCondition contains condition information for a node.
type NodeCondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Type is the type of the condition
	Type NodeConditionType `protobuf:"varint,1,opt,name=type,proto3,enum=edgecluster.NodeConditionType" json:"type,omitempty"`
	// Status is the status of the condition
	Status ConditionStatus `protobuf:"varint,2,opt,name=status,proto3,enum=edgecluster.ConditionStatus" json:"status,omitempty"`
	// Last time we got an update on a given condition.
	LastHeartbeatTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=LastHeartbeatTime,proto3" json:"LastHeartbeatTime,omitempty"`
	// Last time the condition transitioned from one status to another.
	LastTransitionTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=LastTransitionTime,proto3" json:"LastTransitionTime,omitempty"`
	// Unique, one-word, CamelCase reason for the condition last transition.
	Reason string `protobuf:"bytes,5,opt,name=Reason,proto3" json:"Reason,omitempty"`
	// Human-readable message indicating details about last transition.
	Message string `protobuf:"bytes,6,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *NodeCondition) Reset() {
	*x = NodeCondition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_edge_cluster_node_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeCondition) ProtoMessage() {}

func (x *NodeCondition) ProtoReflect() protoreflect.Message {
	mi := &file_edge_cluster_node_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeCondition.ProtoReflect.Descriptor instead.
func (*NodeCondition) Descriptor() ([]byte, []int) {
	return file_edge_cluster_node_messages_proto_rawDescGZIP(), []int{0}
}

func (x *NodeCondition) GetType() NodeConditionType {
	if x != nil {
		return x.Type
	}
	return NodeConditionType_Ready
}

func (x *NodeCondition) GetStatus() ConditionStatus {
	if x != nil {
		return x.Status
	}
	return ConditionStatus_ConditionTrue
}

func (x *NodeCondition) GetLastHeartbeatTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastHeartbeatTime
	}
	return nil
}

func (x *NodeCondition) GetLastTransitionTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastTransitionTime
	}
	return nil
}

func (x *NodeCondition) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *NodeCondition) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

//*
// EdgeClusterNodeAddress contains information for the edge cluster node's address.
type EdgeClusterNodeAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Edge cluster node address type, one of Hostname, ExternalIP or InternalIP.
	NodeAddressType NodeAddressType `protobuf:"varint,1,opt,name=nodeAddressType,proto3,enum=edgecluster.NodeAddressType" json:"nodeAddressType,omitempty"`
	// The node address.
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *EdgeClusterNodeAddress) Reset() {
	*x = EdgeClusterNodeAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_edge_cluster_node_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EdgeClusterNodeAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EdgeClusterNodeAddress) ProtoMessage() {}

func (x *EdgeClusterNodeAddress) ProtoReflect() protoreflect.Message {
	mi := &file_edge_cluster_node_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EdgeClusterNodeAddress.ProtoReflect.Descriptor instead.
func (*EdgeClusterNodeAddress) Descriptor() ([]byte, []int) {
	return file_edge_cluster_node_messages_proto_rawDescGZIP(), []int{1}
}

func (x *EdgeClusterNodeAddress) GetNodeAddressType() NodeAddressType {
	if x != nil {
		return x.NodeAddressType
	}
	return NodeAddressType_Hostname
}

func (x *EdgeClusterNodeAddress) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

//*
// NodeSystemInfo contains a set of ids/uuids to uniquely identify the node.
type NodeSystemInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// MachineID reported by the node. For unique machine identification
	// in the cluster this field is preferred.
	MachineID string `protobuf:"bytes,1,opt,name=machineID,proto3" json:"machineID,omitempty"`
	// SystemUUID reported by the node. For unique machine identification
	// MachineID is preferred. This field is specific to Red Hat hosts
	SystemUUID string `protobuf:"bytes,2,opt,name=systemUUID,proto3" json:"systemUUID,omitempty"`
	// Boot ID reported by the node.
	BootID string `protobuf:"bytes,3,opt,name=bootID,proto3" json:"bootID,omitempty"`
	// Kernel Version reported by the node from 'uname -r' (e.g. 3.16.0-0.bpo.4-amd64).
	KernelVersion string `protobuf:"bytes,4,opt,name=kernelVersion,proto3" json:"kernelVersion,omitempty"`
	// OS Image reported by the node from /etc/os-release (e.g. Debian GNU/Linux 7 (wheezy)).
	OsImage string `protobuf:"bytes,5,opt,name=osImage,proto3" json:"osImage,omitempty"`
	// ContainerRuntime Version reported by the node through runtime remote API (e.g. docker://1.5.0).
	ContainerRuntimeVersion string `protobuf:"bytes,6,opt,name=containerRuntimeVersion,proto3" json:"containerRuntimeVersion,omitempty"`
	// Kubelet Version reported by the node.
	KubeletVersion string `protobuf:"bytes,7,opt,name=kubeletVersion,proto3" json:"kubeletVersion,omitempty"`
	// KubeProxy Version reported by the node.
	KubeProxyVersion string `protobuf:"bytes,8,opt,name=kubeProxyVersion,proto3" json:"kubeProxyVersion,omitempty"`
	// The Operating System reported by the node
	OperatingSystem string `protobuf:"bytes,9,opt,name=operatingSystem,proto3" json:"operatingSystem,omitempty"`
	// The Architecture reported by the node
	Architecture string `protobuf:"bytes,10,opt,name=architecture,proto3" json:"architecture,omitempty"`
}

func (x *NodeSystemInfo) Reset() {
	*x = NodeSystemInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_edge_cluster_node_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeSystemInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeSystemInfo) ProtoMessage() {}

func (x *NodeSystemInfo) ProtoReflect() protoreflect.Message {
	mi := &file_edge_cluster_node_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeSystemInfo.ProtoReflect.Descriptor instead.
func (*NodeSystemInfo) Descriptor() ([]byte, []int) {
	return file_edge_cluster_node_messages_proto_rawDescGZIP(), []int{2}
}

func (x *NodeSystemInfo) GetMachineID() string {
	if x != nil {
		return x.MachineID
	}
	return ""
}

func (x *NodeSystemInfo) GetSystemUUID() string {
	if x != nil {
		return x.SystemUUID
	}
	return ""
}

func (x *NodeSystemInfo) GetBootID() string {
	if x != nil {
		return x.BootID
	}
	return ""
}

func (x *NodeSystemInfo) GetKernelVersion() string {
	if x != nil {
		return x.KernelVersion
	}
	return ""
}

func (x *NodeSystemInfo) GetOsImage() string {
	if x != nil {
		return x.OsImage
	}
	return ""
}

func (x *NodeSystemInfo) GetContainerRuntimeVersion() string {
	if x != nil {
		return x.ContainerRuntimeVersion
	}
	return ""
}

func (x *NodeSystemInfo) GetKubeletVersion() string {
	if x != nil {
		return x.KubeletVersion
	}
	return ""
}

func (x *NodeSystemInfo) GetKubeProxyVersion() string {
	if x != nil {
		return x.KubeProxyVersion
	}
	return ""
}

func (x *NodeSystemInfo) GetOperatingSystem() string {
	if x != nil {
		return x.OperatingSystem
	}
	return ""
}

func (x *NodeSystemInfo) GetArchitecture() string {
	if x != nil {
		return x.Architecture
	}
	return ""
}

//*
// NodeStatus is information about the current status of a node.
type NodeStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Conditions is an array of current observed node conditions.
	Conditions []*NodeCondition `protobuf:"bytes,1,rep,name=conditions,proto3" json:"conditions,omitempty"`
	// Addresses is the list of addresses reachable to the node.
	Addresses []*EdgeClusterNodeAddress `protobuf:"bytes,2,rep,name=addresses,proto3" json:"addresses,omitempty"`
	// NodeInfo is the set of ids/uuids to uniquely identify the node.
	NodeInfo *NodeSystemInfo `protobuf:"bytes,3,opt,name=nodeInfo,proto3" json:"nodeInfo,omitempty"`
}

func (x *NodeStatus) Reset() {
	*x = NodeStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_edge_cluster_node_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeStatus) ProtoMessage() {}

func (x *NodeStatus) ProtoReflect() protoreflect.Message {
	mi := &file_edge_cluster_node_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeStatus.ProtoReflect.Descriptor instead.
func (*NodeStatus) Descriptor() ([]byte, []int) {
	return file_edge_cluster_node_messages_proto_rawDescGZIP(), []int{3}
}

func (x *NodeStatus) GetConditions() []*NodeCondition {
	if x != nil {
		return x.Conditions
	}
	return nil
}

func (x *NodeStatus) GetAddresses() []*EdgeClusterNodeAddress {
	if x != nil {
		return x.Addresses
	}
	return nil
}

func (x *NodeStatus) GetNodeInfo() *NodeSystemInfo {
	if x != nil {
		return x.NodeInfo
	}
	return nil
}

//*
// Declares the details of an existing edge cluster node
type EdgeClusterNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The node metadata
	Metadata *ObjectMeta `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// The most recently observed status of the node
	Status *NodeStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *EdgeClusterNode) Reset() {
	*x = EdgeClusterNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_edge_cluster_node_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EdgeClusterNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EdgeClusterNode) ProtoMessage() {}

func (x *EdgeClusterNode) ProtoReflect() protoreflect.Message {
	mi := &file_edge_cluster_node_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EdgeClusterNode.ProtoReflect.Descriptor instead.
func (*EdgeClusterNode) Descriptor() ([]byte, []int) {
	return file_edge_cluster_node_messages_proto_rawDescGZIP(), []int{4}
}

func (x *EdgeClusterNode) GetMetadata() *ObjectMeta {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *EdgeClusterNode) GetStatus() *NodeStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

//*
// ListEdgeClusterNodesRequest to list an existing edge cluster nodes details
type ListEdgeClusterNodesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique edge cluster identifier
	EdgeClusterID string `protobuf:"bytes,1,opt,name=edgeClusterID,proto3" json:"edgeClusterID,omitempty"`
}

func (x *ListEdgeClusterNodesRequest) Reset() {
	*x = ListEdgeClusterNodesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_edge_cluster_node_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEdgeClusterNodesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEdgeClusterNodesRequest) ProtoMessage() {}

func (x *ListEdgeClusterNodesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_edge_cluster_node_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEdgeClusterNodesRequest.ProtoReflect.Descriptor instead.
func (*ListEdgeClusterNodesRequest) Descriptor() ([]byte, []int) {
	return file_edge_cluster_node_messages_proto_rawDescGZIP(), []int{5}
}

func (x *ListEdgeClusterNodesRequest) GetEdgeClusterID() string {
	if x != nil {
		return x.EdgeClusterID
	}
	return ""
}

//*
// Response contains the result of listing an existing edge cluster nodes details
type ListEdgeClusterNodesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Indicate whether the operation has any error
	Error Error `protobuf:"varint,1,opt,name=error,proto3,enum=edgecluster.Error" json:"error,omitempty"`
	// Contains error message if the operation was unsuccessful
	ErrorMessage string `protobuf:"bytes,2,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	// The list of an existing edge cluster nodes details
	Nodes []*EdgeClusterNode `protobuf:"bytes,3,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *ListEdgeClusterNodesResponse) Reset() {
	*x = ListEdgeClusterNodesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_edge_cluster_node_messages_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEdgeClusterNodesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEdgeClusterNodesResponse) ProtoMessage() {}

func (x *ListEdgeClusterNodesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_edge_cluster_node_messages_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEdgeClusterNodesResponse.ProtoReflect.Descriptor instead.
func (*ListEdgeClusterNodesResponse) Descriptor() ([]byte, []int) {
	return file_edge_cluster_node_messages_proto_rawDescGZIP(), []int{6}
}

func (x *ListEdgeClusterNodesResponse) GetError() Error {
	if x != nil {
		return x.Error
	}
	return Error_NO_ERROR
}

func (x *ListEdgeClusterNodesResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *ListEdgeClusterNodesResponse) GetNodes() []*EdgeClusterNode {
	if x != nil {
		return x.Nodes
	}
	return nil
}

var File_edge_cluster_node_messages_proto protoreflect.FileDescriptor

var file_edge_cluster_node_messages_proto_rawDesc = []byte{
	0x0a, 0x20, 0x65, 0x64, 0x67, 0x65, 0x2d, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2d, 0x6e,
	0x6f, 0x64, 0x65, 0x2d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0b, 0x65, 0x64, 0x67, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x1a,
	0x1a, 0x65, 0x64, 0x67, 0x65, 0x2d, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2d, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc1, 0x02, 0x0a,
	0x0d, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x65,
	0x64, 0x67, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x43,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x48, 0x0a, 0x11, 0x4c, 0x61, 0x73, 0x74,
	0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x11, 0x4c, 0x61, 0x73, 0x74, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x4a, 0x0a, 0x12, 0x4c, 0x61, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x12, 0x4c, 0x61, 0x73, 0x74,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x7a, 0x0a, 0x16, 0x45, 0x64, 0x67, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e,
	0x6f, 0x64, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x46, 0x0a, 0x0f, 0x6e, 0x6f,
	0x64, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0f, 0x6e, 0x6f, 0x64, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x82, 0x03, 0x0a,
	0x0e, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x49, 0x44, 0x12, 0x1e, 0x0a,
	0x0a, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x55, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x55, 0x49, 0x44, 0x12, 0x16, 0x0a,
	0x06, 0x62, 0x6f, 0x6f, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62,
	0x6f, 0x6f, 0x74, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6b, 0x65,
	0x72, 0x6e, 0x65, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6f,
	0x73, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x73,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x38, 0x0a, 0x17, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x26, 0x0a, 0x0e, 0x6b, 0x75, 0x62, 0x65, 0x6c, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6b, 0x75, 0x62, 0x65, 0x6c, 0x65, 0x74,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x10, 0x6b, 0x75, 0x62, 0x65, 0x50,
	0x72, 0x6f, 0x78, 0x79, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x6b, 0x75, 0x62, 0x65, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x22, 0x0a,
	0x0c, 0x61, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x22, 0xc4, 0x01, 0x0a, 0x0a, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x3a, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x41, 0x0a, 0x09,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x45, 0x64,
	0x67, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12,
	0x37, 0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e,
	0x4e, 0x6f, 0x64, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08,
	0x6e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x77, 0x0a, 0x0f, 0x45, 0x64, 0x67, 0x65,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x65, 0x64, 0x67, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x2f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x4e,
	0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x43, 0x0a, 0x1b, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x64, 0x67, 0x65, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x24, 0x0a, 0x0d, 0x65, 0x64, 0x67, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x64, 0x67, 0x65, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x49, 0x44, 0x22, 0xa0, 0x01, 0x0a, 0x1c, 0x4c, 0x69, 0x73, 0x74, 0x45,
	0x64, 0x67, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x2e, 0x45, 0x64, 0x67, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f,
	0x64, 0x65, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2a, 0x6d, 0x0a, 0x11, 0x4e, 0x6f, 0x64,
	0x65, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09,
	0x0a, 0x05, 0x52, 0x65, 0x61, 0x64, 0x79, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x50, 0x72, 0x65, 0x73, 0x73, 0x75, 0x72, 0x65, 0x10, 0x01, 0x12, 0x10, 0x0a,
	0x0c, 0x44, 0x69, 0x73, 0x6b, 0x50, 0x72, 0x65, 0x73, 0x73, 0x75, 0x72, 0x65, 0x10, 0x03, 0x12,
	0x0f, 0x0a, 0x0b, 0x50, 0x49, 0x44, 0x50, 0x72, 0x65, 0x73, 0x73, 0x75, 0x72, 0x65, 0x10, 0x04,
	0x12, 0x16, 0x0a, 0x12, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x55, 0x6e, 0x61, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x10, 0x05, 0x2a, 0x61, 0x0a, 0x0f, 0x4e, 0x6f, 0x64, 0x65,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x48,
	0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x45, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x50, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x50, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x4e, 0x53, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x4e, 0x53, 0x10, 0x04, 0x42, 0x0d, 0x5a, 0x0b, 0x65,
	0x64, 0x67, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_edge_cluster_node_messages_proto_rawDescOnce sync.Once
	file_edge_cluster_node_messages_proto_rawDescData = file_edge_cluster_node_messages_proto_rawDesc
)

func file_edge_cluster_node_messages_proto_rawDescGZIP() []byte {
	file_edge_cluster_node_messages_proto_rawDescOnce.Do(func() {
		file_edge_cluster_node_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_edge_cluster_node_messages_proto_rawDescData)
	})
	return file_edge_cluster_node_messages_proto_rawDescData
}

var file_edge_cluster_node_messages_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_edge_cluster_node_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_edge_cluster_node_messages_proto_goTypes = []interface{}{
	(NodeConditionType)(0),               // 0: edgecluster.NodeConditionType
	(NodeAddressType)(0),                 // 1: edgecluster.NodeAddressType
	(*NodeCondition)(nil),                // 2: edgecluster.NodeCondition
	(*EdgeClusterNodeAddress)(nil),       // 3: edgecluster.EdgeClusterNodeAddress
	(*NodeSystemInfo)(nil),               // 4: edgecluster.NodeSystemInfo
	(*NodeStatus)(nil),                   // 5: edgecluster.NodeStatus
	(*EdgeClusterNode)(nil),              // 6: edgecluster.EdgeClusterNode
	(*ListEdgeClusterNodesRequest)(nil),  // 7: edgecluster.ListEdgeClusterNodesRequest
	(*ListEdgeClusterNodesResponse)(nil), // 8: edgecluster.ListEdgeClusterNodesResponse
	(ConditionStatus)(0),                 // 9: edgecluster.ConditionStatus
	(*timestamppb.Timestamp)(nil),        // 10: google.protobuf.Timestamp
	(*ObjectMeta)(nil),                   // 11: edgecluster.ObjectMeta
	(Error)(0),                           // 12: edgecluster.Error
}
var file_edge_cluster_node_messages_proto_depIdxs = []int32{
	0,  // 0: edgecluster.NodeCondition.type:type_name -> edgecluster.NodeConditionType
	9,  // 1: edgecluster.NodeCondition.status:type_name -> edgecluster.ConditionStatus
	10, // 2: edgecluster.NodeCondition.LastHeartbeatTime:type_name -> google.protobuf.Timestamp
	10, // 3: edgecluster.NodeCondition.LastTransitionTime:type_name -> google.protobuf.Timestamp
	1,  // 4: edgecluster.EdgeClusterNodeAddress.nodeAddressType:type_name -> edgecluster.NodeAddressType
	2,  // 5: edgecluster.NodeStatus.conditions:type_name -> edgecluster.NodeCondition
	3,  // 6: edgecluster.NodeStatus.addresses:type_name -> edgecluster.EdgeClusterNodeAddress
	4,  // 7: edgecluster.NodeStatus.nodeInfo:type_name -> edgecluster.NodeSystemInfo
	11, // 8: edgecluster.EdgeClusterNode.metadata:type_name -> edgecluster.ObjectMeta
	5,  // 9: edgecluster.EdgeClusterNode.status:type_name -> edgecluster.NodeStatus
	12, // 10: edgecluster.ListEdgeClusterNodesResponse.error:type_name -> edgecluster.Error
	6,  // 11: edgecluster.ListEdgeClusterNodesResponse.nodes:type_name -> edgecluster.EdgeClusterNode
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_edge_cluster_node_messages_proto_init() }
func file_edge_cluster_node_messages_proto_init() {
	if File_edge_cluster_node_messages_proto != nil {
		return
	}
	file_edge_cluster_commons_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_edge_cluster_node_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeCondition); i {
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
		file_edge_cluster_node_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EdgeClusterNodeAddress); i {
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
		file_edge_cluster_node_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeSystemInfo); i {
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
		file_edge_cluster_node_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeStatus); i {
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
		file_edge_cluster_node_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EdgeClusterNode); i {
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
		file_edge_cluster_node_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEdgeClusterNodesRequest); i {
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
		file_edge_cluster_node_messages_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEdgeClusterNodesResponse); i {
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
			RawDescriptor: file_edge_cluster_node_messages_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_edge_cluster_node_messages_proto_goTypes,
		DependencyIndexes: file_edge_cluster_node_messages_proto_depIdxs,
		EnumInfos:         file_edge_cluster_node_messages_proto_enumTypes,
		MessageInfos:      file_edge_cluster_node_messages_proto_msgTypes,
	}.Build()
	File_edge_cluster_node_messages_proto = out.File
	file_edge_cluster_node_messages_proto_rawDesc = nil
	file_edge_cluster_node_messages_proto_goTypes = nil
	file_edge_cluster_node_messages_proto_depIdxs = nil
}
