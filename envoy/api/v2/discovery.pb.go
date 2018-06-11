// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: discovery.proto

/*
	Package v2 is a generated protocol buffer package.

	It is generated from these files:
		discovery.proto

	It has these top-level messages:
		DiscoveryRequest
		DiscoveryResponse
*/
package v2

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import envoy_api_v2_core "envoy/api/v2/core"
import google_protobuf5 "google/protobuf"
import google_rpc "google.golang.org/genproto/googleapis/rpc/status"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// A DiscoveryRequest requests a set of versioned resources of the same type for
// a given Envoy node on some API.
type DiscoveryRequest struct {
	// The version_info provided in the request messages will be the version_info
	// received with the most recent successfully processed response or empty on
	// the first request. It is expected that no new request is sent after a
	// response is received until the Envoy instance is ready to ACK/NACK the new
	// configuration. ACK/NACK takes place by returning the new API config version
	// as applied or the previous API config version respectively. Each type_url
	// (see below) has an independent version associated with it.
	VersionInfo string `protobuf:"bytes,1,opt,name=version_info,json=versionInfo,proto3" json:"version_info,omitempty"`
	// The node making the request.
	Node *envoy_api_v2_core.Node `protobuf:"bytes,2,opt,name=node" json:"node,omitempty"`
	// List of resources to subscribe to, e.g. list of cluster names or a route
	// configuration name. If this is empty, all resources for the API are
	// returned. LDS/CDS expect empty resource_names, since this is global
	// discovery for the Envoy instance. The LDS and CDS responses will then imply
	// a number of resources that need to be fetched via EDS/RDS, which will be
	// explicitly enumerated in resource_names.
	ResourceNames []string `protobuf:"bytes,3,rep,name=resource_names,json=resourceNames" json:"resource_names,omitempty"`
	// Type of the resource that is being requested, e.g.
	// "type.googleapis.com/envoy.api.v2.ClusterLoadAssignment". This is implicit
	// in requests made via singleton xDS APIs such as CDS, LDS, etc. but is
	// required for ADS.
	TypeUrl string `protobuf:"bytes,4,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	// nonce corresponding to DiscoveryResponse being ACK/NACKed. See above
	// discussion on version_info and the DiscoveryResponse nonce comment. This
	// may be empty if no nonce is available, e.g. at startup or for non-stream
	// xDS implementations.
	ResponseNonce string `protobuf:"bytes,5,opt,name=response_nonce,json=responseNonce,proto3" json:"response_nonce,omitempty"`
	// This is populated when the previous :ref:`DiscoveryResponse <envoy_api_msg_DiscoveryResponse>`
	// failed to update configuration. The *message* field in *error_details* provides the Envoy
	// internal exception related to the failure. It is only intended for consumption during manual
	// debugging, the string provided is not guaranteed to be stable across Envoy versions.
	ErrorDetail *google_rpc.Status `protobuf:"bytes,6,opt,name=error_detail,json=errorDetail" json:"error_detail,omitempty"`
}

func (m *DiscoveryRequest) Reset()                    { *m = DiscoveryRequest{} }
func (m *DiscoveryRequest) String() string            { return proto.CompactTextString(m) }
func (*DiscoveryRequest) ProtoMessage()               {}
func (*DiscoveryRequest) Descriptor() ([]byte, []int) { return fileDescriptorDiscovery, []int{0} }

func (m *DiscoveryRequest) GetVersionInfo() string {
	if m != nil {
		return m.VersionInfo
	}
	return ""
}

func (m *DiscoveryRequest) GetNode() *envoy_api_v2_core.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *DiscoveryRequest) GetResourceNames() []string {
	if m != nil {
		return m.ResourceNames
	}
	return nil
}

func (m *DiscoveryRequest) GetTypeUrl() string {
	if m != nil {
		return m.TypeUrl
	}
	return ""
}

func (m *DiscoveryRequest) GetResponseNonce() string {
	if m != nil {
		return m.ResponseNonce
	}
	return ""
}

func (m *DiscoveryRequest) GetErrorDetail() *google_rpc.Status {
	if m != nil {
		return m.ErrorDetail
	}
	return nil
}

type DiscoveryResponse struct {
	// The version of the response data.
	VersionInfo string `protobuf:"bytes,1,opt,name=version_info,json=versionInfo,proto3" json:"version_info,omitempty"`
	// The response resources. These resources are typed and depend on the API being called.
	Resources []google_protobuf5.Any `protobuf:"bytes,2,rep,name=resources" json:"resources"`
	// [#not-implemented-hide:]
	// Canary is used to support two Envoy command line flags:
	//
	// * --terminate-on-canary-transition-failure. When set, Envoy is able to
	//   terminate if it detects that configuration is stuck at canary. Consider
	//   this example sequence of updates:
	//   - Management server applies a canary config successfully.
	//   - Management server rolls back to a production config.
	//   - Envoy rejects the new production config.
	//   Since there is no sensible way to continue receiving configuration
	//   updates, Envoy will then terminate and apply production config from a
	//   clean slate.
	// * --dry-run-canary. When set, a canary response will never be applied, only
	//   validated via a dry run.
	Canary bool `protobuf:"varint,3,opt,name=canary,proto3" json:"canary,omitempty"`
	// Type URL for resources. This must be consistent with the type_url in the
	// Any messages for resources if resources is non-empty. This effectively
	// identifies the xDS API when muxing over ADS.
	TypeUrl string `protobuf:"bytes,4,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	// For gRPC based subscriptions, the nonce provides a way to explicitly ack a
	// specific DiscoveryResponse in a following DiscoveryRequest. Additional
	// messages may have been sent by Envoy to the management server for the
	// previous version on the stream prior to this DiscoveryResponse, that were
	// unprocessed at response send time. The nonce allows the management server
	// to ignore any further DiscoveryRequests for the previous version until a
	// DiscoveryRequest bearing the nonce. The nonce is optional and is not
	// required for non-stream based xDS implementations.
	Nonce string `protobuf:"bytes,5,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (m *DiscoveryResponse) Reset()                    { *m = DiscoveryResponse{} }
func (m *DiscoveryResponse) String() string            { return proto.CompactTextString(m) }
func (*DiscoveryResponse) ProtoMessage()               {}
func (*DiscoveryResponse) Descriptor() ([]byte, []int) { return fileDescriptorDiscovery, []int{1} }

func (m *DiscoveryResponse) GetVersionInfo() string {
	if m != nil {
		return m.VersionInfo
	}
	return ""
}

func (m *DiscoveryResponse) GetResources() []google_protobuf5.Any {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *DiscoveryResponse) GetCanary() bool {
	if m != nil {
		return m.Canary
	}
	return false
}

func (m *DiscoveryResponse) GetTypeUrl() string {
	if m != nil {
		return m.TypeUrl
	}
	return ""
}

func (m *DiscoveryResponse) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

func init() {
	proto.RegisterType((*DiscoveryRequest)(nil), "envoy.api.v2.DiscoveryRequest")
	proto.RegisterType((*DiscoveryResponse)(nil), "envoy.api.v2.DiscoveryResponse")
}
func (this *DiscoveryRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DiscoveryRequest)
	if !ok {
		that2, ok := that.(DiscoveryRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.VersionInfo != that1.VersionInfo {
		return false
	}
	if !this.Node.Equal(that1.Node) {
		return false
	}
	if len(this.ResourceNames) != len(that1.ResourceNames) {
		return false
	}
	for i := range this.ResourceNames {
		if this.ResourceNames[i] != that1.ResourceNames[i] {
			return false
		}
	}
	if this.TypeUrl != that1.TypeUrl {
		return false
	}
	if this.ResponseNonce != that1.ResponseNonce {
		return false
	}
	if !this.ErrorDetail.Equal(that1.ErrorDetail) {
		return false
	}
	return true
}
func (this *DiscoveryResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DiscoveryResponse)
	if !ok {
		that2, ok := that.(DiscoveryResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.VersionInfo != that1.VersionInfo {
		return false
	}
	if len(this.Resources) != len(that1.Resources) {
		return false
	}
	for i := range this.Resources {
		if !this.Resources[i].Equal(&that1.Resources[i]) {
			return false
		}
	}
	if this.Canary != that1.Canary {
		return false
	}
	if this.TypeUrl != that1.TypeUrl {
		return false
	}
	if this.Nonce != that1.Nonce {
		return false
	}
	return true
}
func (m *DiscoveryRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DiscoveryRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.VersionInfo) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintDiscovery(dAtA, i, uint64(len(m.VersionInfo)))
		i += copy(dAtA[i:], m.VersionInfo)
	}
	if m.Node != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintDiscovery(dAtA, i, uint64(m.Node.Size()))
		n1, err := m.Node.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.ResourceNames) > 0 {
		for _, s := range m.ResourceNames {
			dAtA[i] = 0x1a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.TypeUrl) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintDiscovery(dAtA, i, uint64(len(m.TypeUrl)))
		i += copy(dAtA[i:], m.TypeUrl)
	}
	if len(m.ResponseNonce) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintDiscovery(dAtA, i, uint64(len(m.ResponseNonce)))
		i += copy(dAtA[i:], m.ResponseNonce)
	}
	if m.ErrorDetail != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintDiscovery(dAtA, i, uint64(m.ErrorDetail.Size()))
		n2, err := m.ErrorDetail.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *DiscoveryResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DiscoveryResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.VersionInfo) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintDiscovery(dAtA, i, uint64(len(m.VersionInfo)))
		i += copy(dAtA[i:], m.VersionInfo)
	}
	if len(m.Resources) > 0 {
		for _, msg := range m.Resources {
			dAtA[i] = 0x12
			i++
			i = encodeVarintDiscovery(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.Canary {
		dAtA[i] = 0x18
		i++
		if m.Canary {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.TypeUrl) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintDiscovery(dAtA, i, uint64(len(m.TypeUrl)))
		i += copy(dAtA[i:], m.TypeUrl)
	}
	if len(m.Nonce) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintDiscovery(dAtA, i, uint64(len(m.Nonce)))
		i += copy(dAtA[i:], m.Nonce)
	}
	return i, nil
}

func encodeVarintDiscovery(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *DiscoveryRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.VersionInfo)
	if l > 0 {
		n += 1 + l + sovDiscovery(uint64(l))
	}
	if m.Node != nil {
		l = m.Node.Size()
		n += 1 + l + sovDiscovery(uint64(l))
	}
	if len(m.ResourceNames) > 0 {
		for _, s := range m.ResourceNames {
			l = len(s)
			n += 1 + l + sovDiscovery(uint64(l))
		}
	}
	l = len(m.TypeUrl)
	if l > 0 {
		n += 1 + l + sovDiscovery(uint64(l))
	}
	l = len(m.ResponseNonce)
	if l > 0 {
		n += 1 + l + sovDiscovery(uint64(l))
	}
	if m.ErrorDetail != nil {
		l = m.ErrorDetail.Size()
		n += 1 + l + sovDiscovery(uint64(l))
	}
	return n
}

func (m *DiscoveryResponse) Size() (n int) {
	var l int
	_ = l
	l = len(m.VersionInfo)
	if l > 0 {
		n += 1 + l + sovDiscovery(uint64(l))
	}
	if len(m.Resources) > 0 {
		for _, e := range m.Resources {
			l = e.Size()
			n += 1 + l + sovDiscovery(uint64(l))
		}
	}
	if m.Canary {
		n += 2
	}
	l = len(m.TypeUrl)
	if l > 0 {
		n += 1 + l + sovDiscovery(uint64(l))
	}
	l = len(m.Nonce)
	if l > 0 {
		n += 1 + l + sovDiscovery(uint64(l))
	}
	return n
}

func sovDiscovery(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozDiscovery(x uint64) (n int) {
	return sovDiscovery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DiscoveryRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDiscovery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DiscoveryRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DiscoveryRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VersionInfo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiscovery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDiscovery
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VersionInfo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Node", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiscovery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDiscovery
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Node == nil {
				m.Node = &envoy_api_v2_core.Node{}
			}
			if err := m.Node.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResourceNames", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiscovery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDiscovery
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResourceNames = append(m.ResourceNames, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TypeUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiscovery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDiscovery
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TypeUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResponseNonce", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiscovery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDiscovery
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResponseNonce = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrorDetail", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiscovery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDiscovery
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ErrorDetail == nil {
				m.ErrorDetail = &google_rpc.Status{}
			}
			if err := m.ErrorDetail.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDiscovery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDiscovery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DiscoveryResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDiscovery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DiscoveryResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DiscoveryResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VersionInfo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiscovery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDiscovery
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VersionInfo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resources", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiscovery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDiscovery
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Resources = append(m.Resources, google_protobuf5.Any{})
			if err := m.Resources[len(m.Resources)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Canary", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiscovery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Canary = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TypeUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiscovery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDiscovery
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TypeUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiscovery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDiscovery
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nonce = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDiscovery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDiscovery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipDiscovery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDiscovery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDiscovery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDiscovery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthDiscovery
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDiscovery
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipDiscovery(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthDiscovery = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDiscovery   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("discovery.proto", fileDescriptorDiscovery) }

var fileDescriptorDiscovery = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0xcd, 0xaa, 0xd3, 0x40,
	0x18, 0x75, 0xda, 0xde, 0x7a, 0x3b, 0xed, 0xf5, 0x67, 0x28, 0xde, 0xb9, 0x45, 0x62, 0xbc, 0x20,
	0x04, 0x84, 0x19, 0x88, 0x08, 0x6e, 0xbd, 0xdc, 0x8d, 0x9b, 0x2e, 0x22, 0x6e, 0xdc, 0x84, 0x69,
	0xf2, 0x35, 0x04, 0xe2, 0x7c, 0x71, 0x26, 0x09, 0x64, 0xeb, 0xd3, 0xf8, 0x0a, 0xbe, 0x41, 0x97,
	0x3e, 0x81, 0x48, 0x9e, 0x44, 0x32, 0x49, 0xb0, 0x2b, 0x71, 0x37, 0xe7, 0x7c, 0x87, 0xef, 0x9c,
	0xef, 0x0c, 0x7d, 0x9c, 0xe6, 0x36, 0xc1, 0x06, 0x4c, 0x2b, 0x4a, 0x83, 0x15, 0xb2, 0x0d, 0xe8,
	0x06, 0x5b, 0xa1, 0xca, 0x5c, 0x34, 0xe1, 0xee, 0xb9, 0x43, 0x52, 0x95, 0xb9, 0x6c, 0x42, 0x99,
	0xa0, 0x01, 0x79, 0x50, 0x16, 0x06, 0xed, 0xee, 0x26, 0x43, 0xcc, 0x0a, 0x90, 0x0e, 0x1d, 0xea,
	0xa3, 0x54, 0x7a, 0x5c, 0xb3, 0xbb, 0x1e, 0x47, 0xa6, 0x4c, 0xa4, 0xad, 0x54, 0x55, 0xdb, 0x71,
	0xb0, 0xcd, 0x30, 0x43, 0xf7, 0x94, 0xfd, 0x6b, 0x60, 0x6f, 0xbf, 0xcd, 0xe8, 0x93, 0xfb, 0x29,
	0x49, 0x04, 0x5f, 0x6b, 0xb0, 0x15, 0x7b, 0x49, 0x37, 0x0d, 0x18, 0x9b, 0xa3, 0x8e, 0x73, 0x7d,
	0x44, 0x4e, 0x7c, 0x12, 0xac, 0xa2, 0xf5, 0xc8, 0x7d, 0xd0, 0x47, 0x64, 0xaf, 0xe9, 0x42, 0x63,
	0x0a, 0x7c, 0xe6, 0x93, 0x60, 0x1d, 0x5e, 0x8b, 0xf3, 0xf0, 0xa2, 0x8f, 0x2b, 0xf6, 0x98, 0x42,
	0xe4, 0x44, 0xec, 0x15, 0x7d, 0x64, 0xc0, 0x62, 0x6d, 0x12, 0x88, 0xb5, 0xfa, 0x02, 0x96, 0xcf,
	0xfd, 0x79, 0xb0, 0x8a, 0xae, 0x26, 0x76, 0xdf, 0x93, 0xec, 0x86, 0x5e, 0x56, 0x6d, 0x09, 0x71,
	0x6d, 0x0a, 0xbe, 0x70, 0x96, 0x0f, 0x7b, 0xfc, 0xc9, 0x14, 0xe3, 0x86, 0x12, 0xb5, 0x85, 0x58,
	0xa3, 0x4e, 0x80, 0x5f, 0x38, 0xc1, 0xd5, 0xc4, 0xee, 0x7b, 0x92, 0xbd, 0xa5, 0x1b, 0x30, 0x06,
	0x4d, 0x9c, 0x42, 0xa5, 0xf2, 0x82, 0x2f, 0x5d, 0x3a, 0x26, 0x86, 0x4e, 0x84, 0x29, 0x13, 0xf1,
	0xd1, 0x75, 0x12, 0xad, 0x9d, 0xee, 0xde, 0xc9, 0x6e, 0x7f, 0x10, 0xfa, 0xf4, 0xac, 0x84, 0x61,
	0xe3, 0xff, 0xb4, 0xf0, 0x8e, 0xae, 0xa6, 0x13, 0x2c, 0x9f, 0xf9, 0xf3, 0x60, 0x1d, 0x6e, 0x27,
	0xb3, 0xe9, 0x6f, 0xc4, 0x7b, 0xdd, 0xde, 0x2d, 0x4e, 0xbf, 0x5e, 0x3c, 0x88, 0xfe, 0x8a, 0xd9,
	0x33, 0xba, 0x4c, 0x94, 0x56, 0xa6, 0xe5, 0x73, 0x9f, 0x04, 0x97, 0xd1, 0x88, 0xfe, 0xd5, 0xc1,
	0x96, 0x5e, 0x9c, 0x9f, 0x3e, 0x80, 0xbb, 0xed, 0xf7, 0xce, 0x23, 0xa7, 0xce, 0x23, 0x3f, 0x3b,
	0x8f, 0xfc, 0xee, 0x3c, 0xf2, 0x79, 0xd6, 0x84, 0x87, 0xa5, 0x73, 0x7f, 0xf3, 0x27, 0x00, 0x00,
	0xff, 0xff, 0x76, 0x0a, 0x11, 0x5e, 0x66, 0x02, 0x00, 0x00,
}
