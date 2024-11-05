// Code generated by bpf2go; DO NOT EDIT.
//go:build arm64

package httptracer

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type bpfConnectionInfoT struct {
	S_addr [16]uint8
	D_addr [16]uint8
	S_port uint16
	D_port uint16
}

type bpfEgressKeyT struct {
	S_port uint16
	D_port uint16
}

type bpfTcHttpCtx struct {
	XtraBytes uint32
	State     uint8
}

type bpfTpInfoPidT struct {
	Tp struct {
		TraceId  [16]uint8
		SpanId   [8]uint8
		ParentId [8]uint8
		Ts       uint64
		Flags    uint8
		_        [7]byte
	}
	Pid   uint32
	Valid uint8
	_     [3]byte
}

// loadBpf returns the embedded CollectionSpec for bpf.
func loadBpf() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_BpfBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load bpf: %w", err)
	}

	return spec, err
}

// loadBpfObjects loads bpf and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*bpfObjects
//	*bpfPrograms
//	*bpfMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadBpfObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadBpf()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// bpfSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfSpecs struct {
	bpfProgramSpecs
	bpfMapSpecs
}

// bpfSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfProgramSpecs struct {
	TcHttpEgress  *ebpf.ProgramSpec `ebpf:"tc_http_egress"`
	TcHttpIngress *ebpf.ProgramSpec `ebpf:"tc_http_ingress"`
}

// bpfMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfMapSpecs struct {
	IncomingTraceMap *ebpf.MapSpec `ebpf:"incoming_trace_map"`
	OutgoingTraceMap *ebpf.MapSpec `ebpf:"outgoing_trace_map"`
	TcHttpCtxMap     *ebpf.MapSpec `ebpf:"tc_http_ctx_map"`
	TraceMap         *ebpf.MapSpec `ebpf:"trace_map"`
}

// bpfObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfObjects struct {
	bpfPrograms
	bpfMaps
}

func (o *bpfObjects) Close() error {
	return _BpfClose(
		&o.bpfPrograms,
		&o.bpfMaps,
	)
}

// bpfMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfMaps struct {
	IncomingTraceMap *ebpf.Map `ebpf:"incoming_trace_map"`
	OutgoingTraceMap *ebpf.Map `ebpf:"outgoing_trace_map"`
	TcHttpCtxMap     *ebpf.Map `ebpf:"tc_http_ctx_map"`
	TraceMap         *ebpf.Map `ebpf:"trace_map"`
}

func (m *bpfMaps) Close() error {
	return _BpfClose(
		m.IncomingTraceMap,
		m.OutgoingTraceMap,
		m.TcHttpCtxMap,
		m.TraceMap,
	)
}

// bpfPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfPrograms struct {
	TcHttpEgress  *ebpf.Program `ebpf:"tc_http_egress"`
	TcHttpIngress *ebpf.Program `ebpf:"tc_http_ingress"`
}

func (p *bpfPrograms) Close() error {
	return _BpfClose(
		p.TcHttpEgress,
		p.TcHttpIngress,
	)
}

func _BpfClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed bpf_arm64_bpfel.o
var _BpfBytes []byte