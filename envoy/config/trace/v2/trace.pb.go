// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: trace.proto

/*
	Package v2 is a generated protocol buffer package.

	It is generated from these files:
		trace.proto

	It has these top-level messages:
		Tracing
		LightstepConfig
		ZipkinConfig
		DynamicOtConfig
		TraceServiceConfig
*/
package v2

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import envoy_api_v2_core1 "envoy/api/v2/core"
import google_protobuf "google/protobuf"
import _ "github.com/lyft/protoc-gen-validate/validate"

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

// The tracing configuration specifies global
// settings for the HTTP tracer used by Envoy. The configuration is defined by
// the :ref:`Bootstrap <envoy_api_msg_config.bootstrap.v2.Bootstrap>` :ref:`tracing
// <envoy_api_field_config.bootstrap.v2.Bootstrap.tracing>` field. Envoy may support other tracers
// in the future, but right now the HTTP tracer is the only one supported.
type Tracing struct {
	// Provides configuration for the HTTP tracer.
	Http *Tracing_Http `protobuf:"bytes,1,opt,name=http" json:"http,omitempty"`
}

func (m *Tracing) Reset()                    { *m = Tracing{} }
func (m *Tracing) String() string            { return proto.CompactTextString(m) }
func (*Tracing) ProtoMessage()               {}
func (*Tracing) Descriptor() ([]byte, []int) { return fileDescriptorTrace, []int{0} }

func (m *Tracing) GetHttp() *Tracing_Http {
	if m != nil {
		return m.Http
	}
	return nil
}

type Tracing_Http struct {
	// The name of the HTTP trace driver to instantiate. The name must match a
	// supported HTTP trace driver. *envoy.lightstep*, *envoy.zipkin*, and
	// *envoy.dynamic.ot* are built-in trace drivers.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Trace driver specific configuration which depends on the driver being
	// instantiated. See the :ref:`LightstepConfig
	// <envoy_api_msg_config.trace.v2.LightstepConfig>`, :ref:`ZipkinConfig
	// <envoy_api_msg_config.trace.v2.ZipkinConfig>`, and :ref:`DynamicOtConfig
	// <envoy_api_msg_config.trace.v2.DynamicOtConfig>` trace drivers for examples.
	Config *google_protobuf.Struct `protobuf:"bytes,2,opt,name=config" json:"config,omitempty"`
}

func (m *Tracing_Http) Reset()                    { *m = Tracing_Http{} }
func (m *Tracing_Http) String() string            { return proto.CompactTextString(m) }
func (*Tracing_Http) ProtoMessage()               {}
func (*Tracing_Http) Descriptor() ([]byte, []int) { return fileDescriptorTrace, []int{0, 0} }

func (m *Tracing_Http) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Tracing_Http) GetConfig() *google_protobuf.Struct {
	if m != nil {
		return m.Config
	}
	return nil
}

// Configuration for the LightStep tracer.
type LightstepConfig struct {
	// The cluster manager cluster that hosts the LightStep collectors.
	CollectorCluster string `protobuf:"bytes,1,opt,name=collector_cluster,json=collectorCluster,proto3" json:"collector_cluster,omitempty"`
	// File containing the access token to the `LightStep
	// <http://lightstep.com/>`_ API.
	AccessTokenFile string `protobuf:"bytes,2,opt,name=access_token_file,json=accessTokenFile,proto3" json:"access_token_file,omitempty"`
}

func (m *LightstepConfig) Reset()                    { *m = LightstepConfig{} }
func (m *LightstepConfig) String() string            { return proto.CompactTextString(m) }
func (*LightstepConfig) ProtoMessage()               {}
func (*LightstepConfig) Descriptor() ([]byte, []int) { return fileDescriptorTrace, []int{1} }

func (m *LightstepConfig) GetCollectorCluster() string {
	if m != nil {
		return m.CollectorCluster
	}
	return ""
}

func (m *LightstepConfig) GetAccessTokenFile() string {
	if m != nil {
		return m.AccessTokenFile
	}
	return ""
}

type ZipkinConfig struct {
	// The cluster manager cluster that hosts the Zipkin collectors. Note that the
	// Zipkin cluster must be defined in the :ref:`Bootstrap static cluster
	// resources <envoy_api_field_config.bootstrap.v2.Bootstrap.StaticResources.clusters>`.
	CollectorCluster string `protobuf:"bytes,1,opt,name=collector_cluster,json=collectorCluster,proto3" json:"collector_cluster,omitempty"`
	// The API endpoint of the Zipkin service where the spans will be sent. When
	// using a standard Zipkin installation, the API endpoint is typically
	// /api/v1/spans, which is the default value.
	CollectorEndpoint string `protobuf:"bytes,2,opt,name=collector_endpoint,json=collectorEndpoint,proto3" json:"collector_endpoint,omitempty"`
	// Determines whether a 128bit trace id will be used when creating a new
	// trace instance. The default value is false, which will result in a 64 bit trace id being used.
	TraceId_128Bit bool `protobuf:"varint,3,opt,name=trace_id_128bit,json=traceId128bit,proto3" json:"trace_id_128bit,omitempty"`
}

func (m *ZipkinConfig) Reset()                    { *m = ZipkinConfig{} }
func (m *ZipkinConfig) String() string            { return proto.CompactTextString(m) }
func (*ZipkinConfig) ProtoMessage()               {}
func (*ZipkinConfig) Descriptor() ([]byte, []int) { return fileDescriptorTrace, []int{2} }

func (m *ZipkinConfig) GetCollectorCluster() string {
	if m != nil {
		return m.CollectorCluster
	}
	return ""
}

func (m *ZipkinConfig) GetCollectorEndpoint() string {
	if m != nil {
		return m.CollectorEndpoint
	}
	return ""
}

func (m *ZipkinConfig) GetTraceId_128Bit() bool {
	if m != nil {
		return m.TraceId_128Bit
	}
	return false
}

// DynamicOtConfig is used to dynamically load a tracer from a shared library
// that implements the `OpenTracing dynamic loading API
// <https://github.com/opentracing/opentracing-cpp>`_.
type DynamicOtConfig struct {
	// Dynamic library implementing the `OpenTracing API
	// <https://github.com/opentracing/opentracing-cpp>`_.
	Library string `protobuf:"bytes,1,opt,name=library,proto3" json:"library,omitempty"`
	// The configuration to use when creating a tracer from the given dynamic
	// library.
	Config *google_protobuf.Struct `protobuf:"bytes,2,opt,name=config" json:"config,omitempty"`
}

func (m *DynamicOtConfig) Reset()                    { *m = DynamicOtConfig{} }
func (m *DynamicOtConfig) String() string            { return proto.CompactTextString(m) }
func (*DynamicOtConfig) ProtoMessage()               {}
func (*DynamicOtConfig) Descriptor() ([]byte, []int) { return fileDescriptorTrace, []int{3} }

func (m *DynamicOtConfig) GetLibrary() string {
	if m != nil {
		return m.Library
	}
	return ""
}

func (m *DynamicOtConfig) GetConfig() *google_protobuf.Struct {
	if m != nil {
		return m.Config
	}
	return nil
}

// Configuration structure.
type TraceServiceConfig struct {
	// The upstream gRPC cluster that hosts the metrics service.
	GrpcService *envoy_api_v2_core1.GrpcService `protobuf:"bytes,1,opt,name=grpc_service,json=grpcService" json:"grpc_service,omitempty"`
}

func (m *TraceServiceConfig) Reset()                    { *m = TraceServiceConfig{} }
func (m *TraceServiceConfig) String() string            { return proto.CompactTextString(m) }
func (*TraceServiceConfig) ProtoMessage()               {}
func (*TraceServiceConfig) Descriptor() ([]byte, []int) { return fileDescriptorTrace, []int{4} }

func (m *TraceServiceConfig) GetGrpcService() *envoy_api_v2_core1.GrpcService {
	if m != nil {
		return m.GrpcService
	}
	return nil
}

func init() {
	proto.RegisterType((*Tracing)(nil), "envoy.config.trace.v2.Tracing")
	proto.RegisterType((*Tracing_Http)(nil), "envoy.config.trace.v2.Tracing.Http")
	proto.RegisterType((*LightstepConfig)(nil), "envoy.config.trace.v2.LightstepConfig")
	proto.RegisterType((*ZipkinConfig)(nil), "envoy.config.trace.v2.ZipkinConfig")
	proto.RegisterType((*DynamicOtConfig)(nil), "envoy.config.trace.v2.DynamicOtConfig")
	proto.RegisterType((*TraceServiceConfig)(nil), "envoy.config.trace.v2.TraceServiceConfig")
}
func (m *Tracing) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Tracing) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Http != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTrace(dAtA, i, uint64(m.Http.Size()))
		n1, err := m.Http.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *Tracing_Http) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Tracing_Http) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTrace(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Config != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTrace(dAtA, i, uint64(m.Config.Size()))
		n2, err := m.Config.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *LightstepConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LightstepConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.CollectorCluster) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTrace(dAtA, i, uint64(len(m.CollectorCluster)))
		i += copy(dAtA[i:], m.CollectorCluster)
	}
	if len(m.AccessTokenFile) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTrace(dAtA, i, uint64(len(m.AccessTokenFile)))
		i += copy(dAtA[i:], m.AccessTokenFile)
	}
	return i, nil
}

func (m *ZipkinConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ZipkinConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.CollectorCluster) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTrace(dAtA, i, uint64(len(m.CollectorCluster)))
		i += copy(dAtA[i:], m.CollectorCluster)
	}
	if len(m.CollectorEndpoint) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTrace(dAtA, i, uint64(len(m.CollectorEndpoint)))
		i += copy(dAtA[i:], m.CollectorEndpoint)
	}
	if m.TraceId_128Bit {
		dAtA[i] = 0x18
		i++
		if m.TraceId_128Bit {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *DynamicOtConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DynamicOtConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Library) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTrace(dAtA, i, uint64(len(m.Library)))
		i += copy(dAtA[i:], m.Library)
	}
	if m.Config != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTrace(dAtA, i, uint64(m.Config.Size()))
		n3, err := m.Config.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}

func (m *TraceServiceConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TraceServiceConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.GrpcService != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTrace(dAtA, i, uint64(m.GrpcService.Size()))
		n4, err := m.GrpcService.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}

func encodeVarintTrace(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Tracing) Size() (n int) {
	var l int
	_ = l
	if m.Http != nil {
		l = m.Http.Size()
		n += 1 + l + sovTrace(uint64(l))
	}
	return n
}

func (m *Tracing_Http) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTrace(uint64(l))
	}
	if m.Config != nil {
		l = m.Config.Size()
		n += 1 + l + sovTrace(uint64(l))
	}
	return n
}

func (m *LightstepConfig) Size() (n int) {
	var l int
	_ = l
	l = len(m.CollectorCluster)
	if l > 0 {
		n += 1 + l + sovTrace(uint64(l))
	}
	l = len(m.AccessTokenFile)
	if l > 0 {
		n += 1 + l + sovTrace(uint64(l))
	}
	return n
}

func (m *ZipkinConfig) Size() (n int) {
	var l int
	_ = l
	l = len(m.CollectorCluster)
	if l > 0 {
		n += 1 + l + sovTrace(uint64(l))
	}
	l = len(m.CollectorEndpoint)
	if l > 0 {
		n += 1 + l + sovTrace(uint64(l))
	}
	if m.TraceId_128Bit {
		n += 2
	}
	return n
}

func (m *DynamicOtConfig) Size() (n int) {
	var l int
	_ = l
	l = len(m.Library)
	if l > 0 {
		n += 1 + l + sovTrace(uint64(l))
	}
	if m.Config != nil {
		l = m.Config.Size()
		n += 1 + l + sovTrace(uint64(l))
	}
	return n
}

func (m *TraceServiceConfig) Size() (n int) {
	var l int
	_ = l
	if m.GrpcService != nil {
		l = m.GrpcService.Size()
		n += 1 + l + sovTrace(uint64(l))
	}
	return n
}

func sovTrace(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTrace(x uint64) (n int) {
	return sovTrace(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Tracing) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTrace
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
			return fmt.Errorf("proto: Tracing: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Tracing: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Http", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrace
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
				return ErrInvalidLengthTrace
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Http == nil {
				m.Http = &Tracing_Http{}
			}
			if err := m.Http.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTrace(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTrace
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
func (m *Tracing_Http) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTrace
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
			return fmt.Errorf("proto: Http: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Http: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrace
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
				return ErrInvalidLengthTrace
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Config", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrace
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
				return ErrInvalidLengthTrace
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Config == nil {
				m.Config = &google_protobuf.Struct{}
			}
			if err := m.Config.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTrace(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTrace
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
func (m *LightstepConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTrace
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
			return fmt.Errorf("proto: LightstepConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LightstepConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollectorCluster", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrace
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
				return ErrInvalidLengthTrace
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CollectorCluster = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccessTokenFile", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrace
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
				return ErrInvalidLengthTrace
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccessTokenFile = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTrace(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTrace
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
func (m *ZipkinConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTrace
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
			return fmt.Errorf("proto: ZipkinConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ZipkinConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollectorCluster", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrace
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
				return ErrInvalidLengthTrace
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CollectorCluster = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollectorEndpoint", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrace
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
				return ErrInvalidLengthTrace
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CollectorEndpoint = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TraceId_128Bit", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrace
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
			m.TraceId_128Bit = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTrace(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTrace
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
func (m *DynamicOtConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTrace
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
			return fmt.Errorf("proto: DynamicOtConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DynamicOtConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Library", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrace
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
				return ErrInvalidLengthTrace
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Library = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Config", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrace
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
				return ErrInvalidLengthTrace
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Config == nil {
				m.Config = &google_protobuf.Struct{}
			}
			if err := m.Config.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTrace(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTrace
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
func (m *TraceServiceConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTrace
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
			return fmt.Errorf("proto: TraceServiceConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TraceServiceConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GrpcService", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTrace
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
				return ErrInvalidLengthTrace
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.GrpcService == nil {
				m.GrpcService = &envoy_api_v2_core1.GrpcService{}
			}
			if err := m.GrpcService.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTrace(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTrace
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
func skipTrace(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTrace
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
					return 0, ErrIntOverflowTrace
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
					return 0, ErrIntOverflowTrace
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
				return 0, ErrInvalidLengthTrace
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTrace
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
				next, err := skipTrace(dAtA[start:])
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
	ErrInvalidLengthTrace = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTrace   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("trace.proto", fileDescriptorTrace) }

var fileDescriptorTrace = []byte{
	// 453 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xcd, 0x6e, 0x13, 0x31,
	0x14, 0x85, 0xe5, 0x34, 0xb4, 0xf4, 0xa6, 0x28, 0xa9, 0x25, 0xd4, 0x28, 0x82, 0x28, 0x4a, 0x11,
	0xea, 0xca, 0x23, 0x06, 0x01, 0x5d, 0xb7, 0xfc, 0x4a, 0x20, 0xa4, 0x69, 0xc5, 0xa2, 0x9b, 0x91,
	0xe3, 0x38, 0x53, 0xab, 0x53, 0xdb, 0xf2, 0xdc, 0x8c, 0x94, 0x1d, 0x6b, 0x1e, 0x81, 0x27, 0xe0,
	0x19, 0x58, 0xb1, 0x64, 0xc9, 0x23, 0xa0, 0xec, 0x78, 0x0b, 0x34, 0xb6, 0x43, 0x8b, 0xc2, 0x06,
	0x75, 0x17, 0xdf, 0x73, 0xce, 0xbd, 0x9f, 0x32, 0x07, 0x3a, 0xe8, 0xb8, 0x90, 0xcc, 0x3a, 0x83,
	0x86, 0xde, 0x95, 0xba, 0x36, 0x0b, 0x26, 0x8c, 0x9e, 0xa9, 0x82, 0x05, 0xa5, 0x4e, 0x07, 0x0f,
	0xfc, 0x38, 0xe1, 0x56, 0x25, 0x75, 0x9a, 0x08, 0xe3, 0x64, 0x52, 0x38, 0x2b, 0xf2, 0x4a, 0xba,
	0x5a, 0xad, 0xc2, 0x83, 0x7b, 0x85, 0x31, 0x45, 0x29, 0x13, 0xff, 0x9a, 0xcc, 0x67, 0x49, 0x85,
	0x6e, 0x2e, 0x30, 0xaa, 0x7b, 0x35, 0x2f, 0xd5, 0x94, 0xa3, 0x4c, 0x56, 0x3f, 0x82, 0x30, 0xfe,
	0x4c, 0x60, 0xeb, 0xd4, 0x71, 0xa1, 0x74, 0x41, 0x9f, 0x41, 0xfb, 0x1c, 0xd1, 0xf6, 0xc9, 0x88,
	0x1c, 0x74, 0xd2, 0x7d, 0xf6, 0x4f, 0x1c, 0x16, 0xdd, 0xec, 0x35, 0xa2, 0xcd, 0x7c, 0x60, 0xf0,
	0x01, 0xda, 0xcd, 0x8b, 0xde, 0x87, 0xb6, 0xe6, 0x97, 0xd2, 0x2f, 0xd8, 0x3e, 0xda, 0xfe, 0xfa,
	0xeb, 0xdb, 0x46, 0xdb, 0xb5, 0x46, 0x24, 0xf3, 0x63, 0x9a, 0xc0, 0x66, 0x58, 0xd6, 0x6f, 0xf9,
	0x0b, 0x7b, 0x2c, 0x30, 0xb3, 0x15, 0x33, 0x3b, 0xf1, 0xcc, 0x59, 0xb4, 0x8d, 0x3f, 0x12, 0xe8,
	0xbe, 0x55, 0xc5, 0x39, 0x56, 0x28, 0xed, 0xb1, 0x9f, 0xd1, 0xa7, 0xb0, 0x2b, 0x4c, 0x59, 0x4a,
	0x81, 0xc6, 0xe5, 0xa2, 0x9c, 0x57, 0x28, 0xdd, 0xfa, 0xc1, 0xde, 0x1f, 0xcf, 0x71, 0xb0, 0xd0,
	0x27, 0xb0, 0xcb, 0x85, 0x90, 0x55, 0x95, 0xa3, 0xb9, 0x90, 0x3a, 0x9f, 0xa9, 0x52, 0x7a, 0x8e,
	0xbf, 0x72, 0xdd, 0xe0, 0x39, 0x6d, 0x2c, 0x2f, 0x55, 0x29, 0xc7, 0x5f, 0x08, 0xec, 0x9c, 0x29,
	0x7b, 0xa1, 0xf4, 0x0d, 0xef, 0x1f, 0x02, 0xbd, 0xca, 0x49, 0x3d, 0xb5, 0x46, 0x69, 0x5c, 0x07,
	0xb8, 0x5a, 0xfe, 0x22, 0x7a, 0xe8, 0x43, 0xe8, 0xfa, 0x3f, 0x3f, 0x57, 0xd3, 0xfc, 0x51, 0x7a,
	0x38, 0x51, 0xd8, 0xdf, 0x18, 0x91, 0x83, 0xdb, 0xd9, 0x1d, 0x3f, 0x7e, 0x33, 0x0d, 0xc3, 0x71,
	0x01, 0xdd, 0xe7, 0x0b, 0xcd, 0x2f, 0x95, 0x78, 0x8f, 0x11, 0x76, 0x1f, 0xb6, 0x4a, 0x35, 0x71,
	0xdc, 0x2d, 0xd6, 0x11, 0x57, 0xca, 0xff, 0x7f, 0x16, 0x01, 0xb4, 0x29, 0x81, 0x3c, 0x09, 0x05,
	0x8c, 0xb7, 0xde, 0xc1, 0xce, 0xf5, 0x5a, 0xc6, 0x16, 0x0d, 0x63, 0x8b, 0xb8, 0x55, 0x4d, 0x79,
	0x9a, 0xf6, 0xb2, 0x57, 0xce, 0x8a, 0x98, 0x3d, 0x82, 0x06, 0xe8, 0xd6, 0x27, 0xd2, 0xea, 0x91,
	0xac, 0x53, 0x5c, 0x13, 0x7a, 0xdf, 0x97, 0x43, 0xf2, 0x63, 0x39, 0x24, 0x3f, 0x97, 0x43, 0x72,
	0xd6, 0xaa, 0xd3, 0xc9, 0xa6, 0xe7, 0x79, 0xfc, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xc4, 0xf6, 0xad,
	0xfc, 0x34, 0x03, 0x00, 0x00,
}
