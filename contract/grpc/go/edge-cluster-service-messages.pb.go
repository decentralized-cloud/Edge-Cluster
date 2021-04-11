// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: edge-cluster-service-messages.proto

package edgecluster

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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
//  ServiceType string describes ingress methods for a service
type ServiceType int32

const (
	// ServiceTypeClusterIP means a service will only be accessible inside the
	// cluster, via the cluster IP.
	ServiceType_ServiceTypeClusterIP ServiceType = 0
	// ServiceTypeNodePort means a service will be exposed on one port of
	// every node, in addition to 'ClusterIP' type.
	ServiceType_ServiceTypeNodePort ServiceType = 1
	// ServiceTypeLoadBalancer means a service will be exposed via an
	// external load balancer (if the cloud provider supports it), in addition
	// to 'NodePort' type.
	ServiceType_ServiceTypeLoadBalancer ServiceType = 2
	// ServiceTypeExternalName means a service consists of only a reference to
	// an external name that kubedns or equivalent will return as a CNAME
	// record, with no exposing or proxying of any pods involved.
	ServiceType_ServiceTypeExternalName ServiceType = 3
)

// Enum value maps for ServiceType.
var (
	ServiceType_name = map[int32]string{
		0: "ServiceTypeClusterIP",
		1: "ServiceTypeNodePort",
		2: "ServiceTypeLoadBalancer",
		3: "ServiceTypeExternalName",
	}
	ServiceType_value = map[string]int32{
		"ServiceTypeClusterIP":    0,
		"ServiceTypeNodePort":     1,
		"ServiceTypeLoadBalancer": 2,
		"ServiceTypeExternalName": 3,
	}
)

func (x ServiceType) Enum() *ServiceType {
	p := new(ServiceType)
	*p = x
	return p
}

func (x ServiceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServiceType) Descriptor() protoreflect.EnumDescriptor {
	return file_edge_cluster_service_messages_proto_enumTypes[0].Descriptor()
}

func (ServiceType) Type() protoreflect.EnumType {
	return &file_edge_cluster_service_messages_proto_enumTypes[0]
}

func (x ServiceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServiceType.Descriptor instead.
func (ServiceType) EnumDescriptor() ([]byte, []int) {
	return file_edge_cluster_service_messages_proto_rawDescGZIP(), []int{0}
}

//*
// ServicePort contains information on service's port.
type ServicePort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of this port within the service
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The IP protocol for this port
	Protcol Protocol `protobuf:"varint,2,opt,name=protcol,proto3,enum=edgecluster.Protocol" json:"protcol,omitempty"`
	// The port that will be exposed by this service.
	Port int32 `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	// Number or name of the port to access on the pods targeted by the service.
	TargetPort string `protobuf:"bytes,4,opt,name=targetPort,proto3" json:"targetPort,omitempty"`
	// The port on each node on which this service is exposed when type is
	// NodePort or LoadBalancer
	NodePort int32 `protobuf:"varint,5,opt,name=nodePort,proto3" json:"nodePort,omitempty"`
}

func (x *ServicePort) Reset() {
	*x = ServicePort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_edge_cluster_service_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServicePort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServicePort) ProtoMessage() {}

func (x *ServicePort) ProtoReflect() protoreflect.Message {
	mi := &file_edge_cluster_service_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServicePort.ProtoReflect.Descriptor instead.
func (*ServicePort) Descriptor() ([]byte, []int) {
	return file_edge_cluster_service_messages_proto_rawDescGZIP(), []int{0}
}

func (x *ServicePort) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ServicePort) GetProtcol() Protocol {
	if x != nil {
		return x.Protcol
	}
	return Protocol_TCP
}

func (x *ServicePort) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *ServicePort) GetTargetPort() string {
	if x != nil {
		return x.TargetPort
	}
	return ""
}

func (x *ServicePort) GetNodePort() int32 {
	if x != nil {
		return x.NodePort
	}
	return 0
}

//*
// Declares the specification of the desired behavior of the existing edge cluster service
type ServiceSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of ports that are exposed by this service.
	Ports []*ServicePort `protobuf:"bytes,1,rep,name=ports,proto3" json:"ports,omitempty"`
	// clusterIPs is a list of IP addresses assigned to this service
	ClusterIPs []string `protobuf:"bytes,2,rep,name=clusterIPs,proto3" json:"clusterIPs,omitempty"`
	// type determines how the Service is exposed.
	Type ServiceType `protobuf:"varint,3,opt,name=type,proto3,enum=edgecluster.ServiceType" json:"type,omitempty"`
	// externalIPs is a list of IP addresses for which nodes in the cluster
	// will also accept traffic for this service.
	ExternalIPs []string `protobuf:"bytes,4,rep,name=externalIPs,proto3" json:"externalIPs,omitempty"`
	// externalName is the external reference that discovery mechanisms will
	// return as an alias for this service (e.g. a DNS CNAME record).
	ExternalName string `protobuf:"bytes,5,opt,name=externalName,proto3" json:"externalName,omitempty"`
}

func (x *ServiceSpec) Reset() {
	*x = ServiceSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_edge_cluster_service_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceSpec) ProtoMessage() {}

func (x *ServiceSpec) ProtoReflect() protoreflect.Message {
	mi := &file_edge_cluster_service_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceSpec.ProtoReflect.Descriptor instead.
func (*ServiceSpec) Descriptor() ([]byte, []int) {
	return file_edge_cluster_service_messages_proto_rawDescGZIP(), []int{1}
}

func (x *ServiceSpec) GetPorts() []*ServicePort {
	if x != nil {
		return x.Ports
	}
	return nil
}

func (x *ServiceSpec) GetClusterIPs() []string {
	if x != nil {
		return x.ClusterIPs
	}
	return nil
}

func (x *ServiceSpec) GetType() ServiceType {
	if x != nil {
		return x.Type
	}
	return ServiceType_ServiceTypeClusterIP
}

func (x *ServiceSpec) GetExternalIPs() []string {
	if x != nil {
		return x.ExternalIPs
	}
	return nil
}

func (x *ServiceSpec) GetExternalName() string {
	if x != nil {
		return x.ExternalName
	}
	return ""
}

//*
// Declares the most recently observed status of the existing edge cluster service
type ServiceStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// LoadBalancer contains the current status of the load-balancer,
	// if one is present.
	LoadBalancer *LoadBalancerStatus `protobuf:"bytes,1,opt,name=loadBalancer,proto3" json:"loadBalancer,omitempty"`
	// Current service state of service.
	Conditions []*ServiceCondition `protobuf:"bytes,2,rep,name=conditions,proto3" json:"conditions,omitempty"`
}

func (x *ServiceStatus) Reset() {
	*x = ServiceStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_edge_cluster_service_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceStatus) ProtoMessage() {}

func (x *ServiceStatus) ProtoReflect() protoreflect.Message {
	mi := &file_edge_cluster_service_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceStatus.ProtoReflect.Descriptor instead.
func (*ServiceStatus) Descriptor() ([]byte, []int) {
	return file_edge_cluster_service_messages_proto_rawDescGZIP(), []int{2}
}

func (x *ServiceStatus) GetLoadBalancer() *LoadBalancerStatus {
	if x != nil {
		return x.LoadBalancer
	}
	return nil
}

func (x *ServiceStatus) GetConditions() []*ServiceCondition {
	if x != nil {
		return x.Conditions
	}
	return nil
}

//*
// Declares the details of an existing edge cluster service
type EdgeClusterService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The service metadata
	Metadata *ObjectMeta `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// The most recently observed status of the service
	Status *ServiceStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	// The specification of the desired behavior of the service.
	Spec *ServiceSpec `protobuf:"bytes,3,opt,name=spec,proto3" json:"spec,omitempty"`
}

func (x *EdgeClusterService) Reset() {
	*x = EdgeClusterService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_edge_cluster_service_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EdgeClusterService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EdgeClusterService) ProtoMessage() {}

func (x *EdgeClusterService) ProtoReflect() protoreflect.Message {
	mi := &file_edge_cluster_service_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EdgeClusterService.ProtoReflect.Descriptor instead.
func (*EdgeClusterService) Descriptor() ([]byte, []int) {
	return file_edge_cluster_service_messages_proto_rawDescGZIP(), []int{3}
}

func (x *EdgeClusterService) GetMetadata() *ObjectMeta {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *EdgeClusterService) GetStatus() *ServiceStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *EdgeClusterService) GetSpec() *ServiceSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

//*
// ListEdgeClusterNodeServicesRequest to list an existing edge cluster services details
type ListEdgeClusterServicesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique edge cluster identifier
	EdgeClusterID string `protobuf:"bytes,1,opt,name=edgeClusterID,proto3" json:"edgeClusterID,omitempty"`
	// Optional, if provided, will be used to filter services under the given namespace
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Optional, if provided, will be used to filter services deployed to the given node
	NodeName string `protobuf:"bytes,3,opt,name=nodeName,proto3" json:"nodeName,omitempty"`
}

func (x *ListEdgeClusterServicesRequest) Reset() {
	*x = ListEdgeClusterServicesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_edge_cluster_service_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEdgeClusterServicesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEdgeClusterServicesRequest) ProtoMessage() {}

func (x *ListEdgeClusterServicesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_edge_cluster_service_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEdgeClusterServicesRequest.ProtoReflect.Descriptor instead.
func (*ListEdgeClusterServicesRequest) Descriptor() ([]byte, []int) {
	return file_edge_cluster_service_messages_proto_rawDescGZIP(), []int{4}
}

func (x *ListEdgeClusterServicesRequest) GetEdgeClusterID() string {
	if x != nil {
		return x.EdgeClusterID
	}
	return ""
}

func (x *ListEdgeClusterServicesRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ListEdgeClusterServicesRequest) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

//*
// Response contains the result of listing an existing edge cluster services details
type ListEdgeClusterServicesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Indicate whether the operation has any error
	Error Error `protobuf:"varint,1,opt,name=error,proto3,enum=edgecluster.Error" json:"error,omitempty"`
	// Contains error message if the operation was unsuccessful
	ErrorMessage string `protobuf:"bytes,2,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	// The list of an existing edge cluster services details
	Services []*EdgeClusterService `protobuf:"bytes,3,rep,name=services,proto3" json:"services,omitempty"`
}

func (x *ListEdgeClusterServicesResponse) Reset() {
	*x = ListEdgeClusterServicesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_edge_cluster_service_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEdgeClusterServicesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEdgeClusterServicesResponse) ProtoMessage() {}

func (x *ListEdgeClusterServicesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_edge_cluster_service_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEdgeClusterServicesResponse.ProtoReflect.Descriptor instead.
func (*ListEdgeClusterServicesResponse) Descriptor() ([]byte, []int) {
	return file_edge_cluster_service_messages_proto_rawDescGZIP(), []int{5}
}

func (x *ListEdgeClusterServicesResponse) GetError() Error {
	if x != nil {
		return x.Error
	}
	return Error_NO_ERROR
}

func (x *ListEdgeClusterServicesResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *ListEdgeClusterServicesResponse) GetServices() []*EdgeClusterService {
	if x != nil {
		return x.Services
	}
	return nil
}

var File_edge_cluster_service_messages_proto protoreflect.FileDescriptor

var file_edge_cluster_service_messages_proto_rawDesc = []byte{
	0x0a, 0x23, 0x65, 0x64, 0x67, 0x65, 0x2d, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2d, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x65, 0x64, 0x67, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x1a, 0x1a, 0x65, 0x64, 0x67, 0x65, 0x2d, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2,
	0x01, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x74, 0x63, 0x6f, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x74,
	0x63, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x50,
	0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x50,
	0x6f, 0x72, 0x74, 0x22, 0xd1, 0x01, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53,
	0x70, 0x65, 0x63, 0x12, 0x2e, 0x0a, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x05, 0x70, 0x6f,
	0x72, 0x74, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x50,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x49, 0x50, 0x73, 0x12, 0x2c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x18, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x50, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x49, 0x50, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x93, 0x01, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x43, 0x0a, 0x0c, 0x6c, 0x6f, 0x61,
	0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x4c, 0x6f,
	0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x0c, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x12, 0x3d,
	0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xab, 0x01,
	0x0a, 0x12, 0x45, 0x64, 0x67, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x64, 0x67, 0x65,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2c, 0x0a,
	0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x64,
	0x67, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0x80, 0x01, 0x0a, 0x1e,
	0x4c, 0x69, 0x73, 0x74, 0x45, 0x64, 0x67, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24,
	0x0a, 0x0d, 0x65, 0x64, 0x67, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x64, 0x67, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xac,
	0x01, 0x0a, 0x1f, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x64, 0x67, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x12, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x22, 0x0a, 0x0c,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x3b, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x45, 0x64, 0x67, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2a, 0x7a, 0x0a,
	0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x14,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x49, 0x50, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x10, 0x01, 0x12,
	0x1b, 0x0a, 0x17, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x4c, 0x6f,
	0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x10, 0x02, 0x12, 0x1b, 0x0a, 0x17,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x10, 0x03, 0x42, 0x0d, 0x5a, 0x0b, 0x65, 0x64, 0x67,
	0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_edge_cluster_service_messages_proto_rawDescOnce sync.Once
	file_edge_cluster_service_messages_proto_rawDescData = file_edge_cluster_service_messages_proto_rawDesc
)

func file_edge_cluster_service_messages_proto_rawDescGZIP() []byte {
	file_edge_cluster_service_messages_proto_rawDescOnce.Do(func() {
		file_edge_cluster_service_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_edge_cluster_service_messages_proto_rawDescData)
	})
	return file_edge_cluster_service_messages_proto_rawDescData
}

var file_edge_cluster_service_messages_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_edge_cluster_service_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_edge_cluster_service_messages_proto_goTypes = []interface{}{
	(ServiceType)(0),                        // 0: edgecluster.ServiceType
	(*ServicePort)(nil),                     // 1: edgecluster.ServicePort
	(*ServiceSpec)(nil),                     // 2: edgecluster.ServiceSpec
	(*ServiceStatus)(nil),                   // 3: edgecluster.ServiceStatus
	(*EdgeClusterService)(nil),              // 4: edgecluster.EdgeClusterService
	(*ListEdgeClusterServicesRequest)(nil),  // 5: edgecluster.ListEdgeClusterServicesRequest
	(*ListEdgeClusterServicesResponse)(nil), // 6: edgecluster.ListEdgeClusterServicesResponse
	(Protocol)(0),                           // 7: edgecluster.Protocol
	(*LoadBalancerStatus)(nil),              // 8: edgecluster.LoadBalancerStatus
	(*ServiceCondition)(nil),                // 9: edgecluster.ServiceCondition
	(*ObjectMeta)(nil),                      // 10: edgecluster.ObjectMeta
	(Error)(0),                              // 11: edgecluster.Error
}
var file_edge_cluster_service_messages_proto_depIdxs = []int32{
	7,  // 0: edgecluster.ServicePort.protcol:type_name -> edgecluster.Protocol
	1,  // 1: edgecluster.ServiceSpec.ports:type_name -> edgecluster.ServicePort
	0,  // 2: edgecluster.ServiceSpec.type:type_name -> edgecluster.ServiceType
	8,  // 3: edgecluster.ServiceStatus.loadBalancer:type_name -> edgecluster.LoadBalancerStatus
	9,  // 4: edgecluster.ServiceStatus.conditions:type_name -> edgecluster.ServiceCondition
	10, // 5: edgecluster.EdgeClusterService.metadata:type_name -> edgecluster.ObjectMeta
	3,  // 6: edgecluster.EdgeClusterService.status:type_name -> edgecluster.ServiceStatus
	2,  // 7: edgecluster.EdgeClusterService.spec:type_name -> edgecluster.ServiceSpec
	11, // 8: edgecluster.ListEdgeClusterServicesResponse.error:type_name -> edgecluster.Error
	4,  // 9: edgecluster.ListEdgeClusterServicesResponse.services:type_name -> edgecluster.EdgeClusterService
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_edge_cluster_service_messages_proto_init() }
func file_edge_cluster_service_messages_proto_init() {
	if File_edge_cluster_service_messages_proto != nil {
		return
	}
	file_edge_cluster_commons_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_edge_cluster_service_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServicePort); i {
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
		file_edge_cluster_service_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceSpec); i {
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
		file_edge_cluster_service_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceStatus); i {
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
		file_edge_cluster_service_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EdgeClusterService); i {
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
		file_edge_cluster_service_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEdgeClusterServicesRequest); i {
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
		file_edge_cluster_service_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEdgeClusterServicesResponse); i {
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
			RawDescriptor: file_edge_cluster_service_messages_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_edge_cluster_service_messages_proto_goTypes,
		DependencyIndexes: file_edge_cluster_service_messages_proto_depIdxs,
		EnumInfos:         file_edge_cluster_service_messages_proto_enumTypes,
		MessageInfos:      file_edge_cluster_service_messages_proto_msgTypes,
	}.Build()
	File_edge_cluster_service_messages_proto = out.File
	file_edge_cluster_service_messages_proto_rawDesc = nil
	file_edge_cluster_service_messages_proto_goTypes = nil
	file_edge_cluster_service_messages_proto_depIdxs = nil
}
