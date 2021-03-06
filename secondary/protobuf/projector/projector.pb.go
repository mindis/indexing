// Code generated by protoc-gen-go.
// source: projector.proto
// DO NOT EDIT!

package protobuf

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

// Requested by Coordinator/indexer to learn vbuckets
// hosted by kvnodes.
type VbmapRequest struct {
	Pool             *string  `protobuf:"bytes,1,req,name=pool" json:"pool,omitempty"`
	Bucket           *string  `protobuf:"bytes,2,req,name=bucket" json:"bucket,omitempty"`
	Kvaddrs          []string `protobuf:"bytes,3,rep,name=kvaddrs" json:"kvaddrs,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *VbmapRequest) Reset()         { *m = VbmapRequest{} }
func (m *VbmapRequest) String() string { return proto.CompactTextString(m) }
func (*VbmapRequest) ProtoMessage()    {}

func (m *VbmapRequest) GetPool() string {
	if m != nil && m.Pool != nil {
		return *m.Pool
	}
	return ""
}

func (m *VbmapRequest) GetBucket() string {
	if m != nil && m.Bucket != nil {
		return *m.Bucket
	}
	return ""
}

func (m *VbmapRequest) GetKvaddrs() []string {
	if m != nil {
		return m.Kvaddrs
	}
	return nil
}

type VbmapResponse struct {
	Kvaddrs          []string    `protobuf:"bytes,1,rep,name=kvaddrs" json:"kvaddrs,omitempty"`
	Kvvbnos          []*Vbuckets `protobuf:"bytes,2,rep,name=kvvbnos" json:"kvvbnos,omitempty"`
	Err              *Error      `protobuf:"bytes,3,opt,name=err" json:"err,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *VbmapResponse) Reset()         { *m = VbmapResponse{} }
func (m *VbmapResponse) String() string { return proto.CompactTextString(m) }
func (*VbmapResponse) ProtoMessage()    {}

func (m *VbmapResponse) GetKvaddrs() []string {
	if m != nil {
		return m.Kvaddrs
	}
	return nil
}

func (m *VbmapResponse) GetKvvbnos() []*Vbuckets {
	if m != nil {
		return m.Kvvbnos
	}
	return nil
}

func (m *VbmapResponse) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

// Requested by Coordinator during system-start,
// re-connect, rollback
type FailoverLogRequest struct {
	Pool             *string  `protobuf:"bytes,1,req,name=pool" json:"pool,omitempty"`
	Bucket           *string  `protobuf:"bytes,2,req,name=bucket" json:"bucket,omitempty"`
	Vbnos            []uint32 `protobuf:"varint,3,rep,name=vbnos" json:"vbnos,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *FailoverLogRequest) Reset()         { *m = FailoverLogRequest{} }
func (m *FailoverLogRequest) String() string { return proto.CompactTextString(m) }
func (*FailoverLogRequest) ProtoMessage()    {}

func (m *FailoverLogRequest) GetPool() string {
	if m != nil && m.Pool != nil {
		return *m.Pool
	}
	return ""
}

func (m *FailoverLogRequest) GetBucket() string {
	if m != nil && m.Bucket != nil {
		return *m.Bucket
	}
	return ""
}

func (m *FailoverLogRequest) GetVbnos() []uint32 {
	if m != nil {
		return m.Vbnos
	}
	return nil
}

type FailoverLogResponse struct {
	Logs             []*FailoverLog `protobuf:"bytes,1,rep,name=logs" json:"logs,omitempty"`
	Err              *Error         `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *FailoverLogResponse) Reset()         { *m = FailoverLogResponse{} }
func (m *FailoverLogResponse) String() string { return proto.CompactTextString(m) }
func (*FailoverLogResponse) ProtoMessage()    {}

func (m *FailoverLogResponse) GetLogs() []*FailoverLog {
	if m != nil {
		return m.Logs
	}
	return nil
}

func (m *FailoverLogResponse) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

// Requested by Coordinator or indexer to start a
// new mutation topic. Respond back with TopicResponse.
type MutationTopicRequest struct {
	Topic         *string     `protobuf:"bytes,1,req,name=topic" json:"topic,omitempty"`
	EndpointType  *string     `protobuf:"bytes,2,req,name=endpointType" json:"endpointType,omitempty"`
	ReqTimestamps []*TsVbuuid `protobuf:"bytes,3,rep,name=reqTimestamps" json:"reqTimestamps,omitempty"`
	// initial list of instances applicable for this topic
	Instances        []*Instance `protobuf:"bytes,4,rep,name=instances" json:"instances,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *MutationTopicRequest) Reset()         { *m = MutationTopicRequest{} }
func (m *MutationTopicRequest) String() string { return proto.CompactTextString(m) }
func (*MutationTopicRequest) ProtoMessage()    {}

func (m *MutationTopicRequest) GetTopic() string {
	if m != nil && m.Topic != nil {
		return *m.Topic
	}
	return ""
}

func (m *MutationTopicRequest) GetEndpointType() string {
	if m != nil && m.EndpointType != nil {
		return *m.EndpointType
	}
	return ""
}

func (m *MutationTopicRequest) GetReqTimestamps() []*TsVbuuid {
	if m != nil {
		return m.ReqTimestamps
	}
	return nil
}

func (m *MutationTopicRequest) GetInstances() []*Instance {
	if m != nil {
		return m.Instances
	}
	return nil
}

// Response back for
// MutationTopicRequest, RestartVbucketsRequest, AddBucketsRequest
type TopicResponse struct {
	Topic              *string     `protobuf:"bytes,1,opt,name=topic" json:"topic,omitempty"`
	InstanceIds        []uint64    `protobuf:"varint,2,rep,name=instanceIds" json:"instanceIds,omitempty"`
	ActiveTimestamps   []*TsVbuuid `protobuf:"bytes,3,rep,name=activeTimestamps" json:"activeTimestamps,omitempty"`
	RollbackTimestamps []*TsVbuuid `protobuf:"bytes,4,rep,name=rollbackTimestamps" json:"rollbackTimestamps,omitempty"`
	Err                *Error      `protobuf:"bytes,5,opt,name=err" json:"err,omitempty"`
	XXX_unrecognized   []byte      `json:"-"`
}

func (m *TopicResponse) Reset()         { *m = TopicResponse{} }
func (m *TopicResponse) String() string { return proto.CompactTextString(m) }
func (*TopicResponse) ProtoMessage()    {}

func (m *TopicResponse) GetTopic() string {
	if m != nil && m.Topic != nil {
		return *m.Topic
	}
	return ""
}

func (m *TopicResponse) GetInstanceIds() []uint64 {
	if m != nil {
		return m.InstanceIds
	}
	return nil
}

func (m *TopicResponse) GetActiveTimestamps() []*TsVbuuid {
	if m != nil {
		return m.ActiveTimestamps
	}
	return nil
}

func (m *TopicResponse) GetRollbackTimestamps() []*TsVbuuid {
	if m != nil {
		return m.RollbackTimestamps
	}
	return nil
}

func (m *TopicResponse) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

// RestartVbucketsRequest will restart a subset
// of vbuckets for each specified buckets.
// Respond back with TopicResponse
type RestartVbucketsRequest struct {
	Topic             *string     `protobuf:"bytes,1,req,name=topic" json:"topic,omitempty"`
	RestartTimestamps []*TsVbuuid `protobuf:"bytes,2,rep,name=restartTimestamps" json:"restartTimestamps,omitempty"`
	XXX_unrecognized  []byte      `json:"-"`
}

func (m *RestartVbucketsRequest) Reset()         { *m = RestartVbucketsRequest{} }
func (m *RestartVbucketsRequest) String() string { return proto.CompactTextString(m) }
func (*RestartVbucketsRequest) ProtoMessage()    {}

func (m *RestartVbucketsRequest) GetTopic() string {
	if m != nil && m.Topic != nil {
		return *m.Topic
	}
	return ""
}

func (m *RestartVbucketsRequest) GetRestartTimestamps() []*TsVbuuid {
	if m != nil {
		return m.RestartTimestamps
	}
	return nil
}

// ShutdownVbucketsRequest will shutdown a subset of vbuckets
// for each specified buckets. Respond back with TopicResponse
type ShutdownVbucketsRequest struct {
	Topic              *string     `protobuf:"bytes,1,req,name=topic" json:"topic,omitempty"`
	ShutdownTimestamps []*TsVbuuid `protobuf:"bytes,2,rep,name=shutdownTimestamps" json:"shutdownTimestamps,omitempty"`
	XXX_unrecognized   []byte      `json:"-"`
}

func (m *ShutdownVbucketsRequest) Reset()         { *m = ShutdownVbucketsRequest{} }
func (m *ShutdownVbucketsRequest) String() string { return proto.CompactTextString(m) }
func (*ShutdownVbucketsRequest) ProtoMessage()    {}

func (m *ShutdownVbucketsRequest) GetTopic() string {
	if m != nil && m.Topic != nil {
		return *m.Topic
	}
	return ""
}

func (m *ShutdownVbucketsRequest) GetShutdownTimestamps() []*TsVbuuid {
	if m != nil {
		return m.ShutdownTimestamps
	}
	return nil
}

// AddBucketsRequest will start vbucket-streams
// specified buckets and add them to the topic that
// has already started.
// Respond back with TopicResponse
type AddBucketsRequest struct {
	Topic         *string     `protobuf:"bytes,1,req,name=topic" json:"topic,omitempty"`
	ReqTimestamps []*TsVbuuid `protobuf:"bytes,2,rep,name=reqTimestamps" json:"reqTimestamps,omitempty"`
	// list of instances applicable for buckets.
	Instances        []*Instance `protobuf:"bytes,3,rep,name=instances" json:"instances,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *AddBucketsRequest) Reset()         { *m = AddBucketsRequest{} }
func (m *AddBucketsRequest) String() string { return proto.CompactTextString(m) }
func (*AddBucketsRequest) ProtoMessage()    {}

func (m *AddBucketsRequest) GetTopic() string {
	if m != nil && m.Topic != nil {
		return *m.Topic
	}
	return ""
}

func (m *AddBucketsRequest) GetReqTimestamps() []*TsVbuuid {
	if m != nil {
		return m.ReqTimestamps
	}
	return nil
}

func (m *AddBucketsRequest) GetInstances() []*Instance {
	if m != nil {
		return m.Instances
	}
	return nil
}

// DelBucketsRequest will shutdown vbucket-streams
// for specified buckets and remove the buckets from topic.
// Respond back with TopicResponse
type DelBucketsRequest struct {
	Topic            *string  `protobuf:"bytes,1,req,name=topic" json:"topic,omitempty"`
	Buckets          []string `protobuf:"bytes,2,rep,name=buckets" json:"buckets,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *DelBucketsRequest) Reset()         { *m = DelBucketsRequest{} }
func (m *DelBucketsRequest) String() string { return proto.CompactTextString(m) }
func (*DelBucketsRequest) ProtoMessage()    {}

func (m *DelBucketsRequest) GetTopic() string {
	if m != nil && m.Topic != nil {
		return *m.Topic
	}
	return ""
}

func (m *DelBucketsRequest) GetBuckets() []string {
	if m != nil {
		return m.Buckets
	}
	return nil
}

// AddInstancesRequest to add index-instances to a topic.
// Respond back with TopicResponse
type AddInstancesRequest struct {
	Topic            *string     `protobuf:"bytes,1,req,name=topic" json:"topic,omitempty"`
	Instances        []*Instance `protobuf:"bytes,2,rep,name=instances" json:"instances,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *AddInstancesRequest) Reset()         { *m = AddInstancesRequest{} }
func (m *AddInstancesRequest) String() string { return proto.CompactTextString(m) }
func (*AddInstancesRequest) ProtoMessage()    {}

func (m *AddInstancesRequest) GetTopic() string {
	if m != nil && m.Topic != nil {
		return *m.Topic
	}
	return ""
}

func (m *AddInstancesRequest) GetInstances() []*Instance {
	if m != nil {
		return m.Instances
	}
	return nil
}

// DelInstancesRequest to add index-instances to a topic.
// Respond back with TopicResponse
type DelInstancesRequest struct {
	Topic            *string  `protobuf:"bytes,1,req,name=topic" json:"topic,omitempty"`
	InstanceIds      []uint64 `protobuf:"varint,2,rep,name=instanceIds" json:"instanceIds,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *DelInstancesRequest) Reset()         { *m = DelInstancesRequest{} }
func (m *DelInstancesRequest) String() string { return proto.CompactTextString(m) }
func (*DelInstancesRequest) ProtoMessage()    {}

func (m *DelInstancesRequest) GetTopic() string {
	if m != nil && m.Topic != nil {
		return *m.Topic
	}
	return ""
}

func (m *DelInstancesRequest) GetInstanceIds() []uint64 {
	if m != nil {
		return m.InstanceIds
	}
	return nil
}

// Requested by indexer / coordinator to inform router to re-connect with
// downstream endpoint. Error message will be sent as response.
type RepairEndpointsRequest struct {
	Topic            *string  `protobuf:"bytes,1,req,name=topic" json:"topic,omitempty"`
	Endpoints        []string `protobuf:"bytes,2,rep,name=endpoints" json:"endpoints,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *RepairEndpointsRequest) Reset()         { *m = RepairEndpointsRequest{} }
func (m *RepairEndpointsRequest) String() string { return proto.CompactTextString(m) }
func (*RepairEndpointsRequest) ProtoMessage()    {}

func (m *RepairEndpointsRequest) GetTopic() string {
	if m != nil && m.Topic != nil {
		return *m.Topic
	}
	return ""
}

func (m *RepairEndpointsRequest) GetEndpoints() []string {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

// Requested by coordinator to should down a mutation topic and all KV
// connections active for that topic. Error message will be sent as response.
type ShutdownTopicRequest struct {
	Topic            *string `protobuf:"bytes,1,req,name=topic" json:"topic,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ShutdownTopicRequest) Reset()         { *m = ShutdownTopicRequest{} }
func (m *ShutdownTopicRequest) String() string { return proto.CompactTextString(m) }
func (*ShutdownTopicRequest) ProtoMessage()    {}

func (m *ShutdownTopicRequest) GetTopic() string {
	if m != nil && m.Topic != nil {
		return *m.Topic
	}
	return ""
}

// Generic instance, can be an index instance, xdcr, search etc ...
type Instance struct {
	IndexInstance    *IndexInst `protobuf:"bytes,1,opt,name=indexInstance" json:"indexInstance,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *Instance) Reset()         { *m = Instance{} }
func (m *Instance) String() string { return proto.CompactTextString(m) }
func (*Instance) ProtoMessage()    {}

func (m *Instance) GetIndexInstance() *IndexInst {
	if m != nil {
		return m.IndexInstance
	}
	return nil
}

// List of instances
type Instances struct {
	Instances        []*Instance `protobuf:"bytes,1,rep,name=instances" json:"instances,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *Instances) Reset()         { *m = Instances{} }
func (m *Instances) String() string { return proto.CompactTextString(m) }
func (*Instances) ProtoMessage()    {}

func (m *Instances) GetInstances() []*Instance {
	if m != nil {
		return m.Instances
	}
	return nil
}

func init() {
}
