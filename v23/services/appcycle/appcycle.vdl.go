// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: appcycle

// Package appcycle defines interfaces for managing application processes.
package appcycle

import (
	"io"
	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/rpc"
	"v.io/v23/vdl"
)

var _ = __VDLInit() // Must be first; see __VDLInit comments for details.

//////////////////////////////////////////////////
// Type definitions

// Task is streamed by Stop to provide the client with a sense of the progress
// of the shutdown.
// The meaning of Progress and Goal are up to the developer (the server provides
// the framework with values for these).  The recommended meanings are:
// - Progress: how far along the shutdown sequence the server is.  This should
//   be a monotonically increasing number.
// - Goal: when Progress reaches this value, the shutdown is expected to
//   complete.  This should not change during a stream, but could change if
//   e.g. new shutdown tasks are triggered that were not forseen at the outset
//   of the shutdown.
type Task struct {
	Progress int32
	Goal     int32
}

func (Task) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/services/appcycle.Task"`
}) {
}

func (x Task) VDLIsZero() bool {
	return x == Task{}
}

func (x Task) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(__VDLType_struct_1); err != nil {
		return err
	}
	if x.Progress != 0 {
		if err := enc.NextFieldValueInt(0, vdl.Int32Type, int64(x.Progress)); err != nil {
			return err
		}
	}
	if x.Goal != 0 {
		if err := enc.NextFieldValueInt(1, vdl.Int32Type, int64(x.Goal)); err != nil {
			return err
		}
	}
	if err := enc.NextField(-1); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *Task) VDLRead(dec vdl.Decoder) error {
	*x = Task{}
	if err := dec.StartValue(__VDLType_struct_1); err != nil {
		return err
	}
	decType := dec.Type()
	for {
		index, err := dec.NextField()
		switch {
		case err != nil:
			return err
		case index == -1:
			return dec.FinishValue()
		}
		if decType != __VDLType_struct_1 {
			index = __VDLType_struct_1.FieldIndexByName(decType.Field(index).Name)
			if index == -1 {
				if err := dec.SkipValue(); err != nil {
					return err
				}
				continue
			}
		}
		switch index {
		case 0:
			switch value, err := dec.ReadValueInt(32); {
			case err != nil:
				return err
			default:
				x.Progress = int32(value)
			}
		case 1:
			switch value, err := dec.ReadValueInt(32); {
			case err != nil:
				return err
			default:
				x.Goal = int32(value)
			}
		}
	}
}

//////////////////////////////////////////////////
// Interface definitions

// AppCycleClientMethods is the client interface
// containing AppCycle methods.
//
// AppCycle interfaces with the process running a vanadium runtime.
type AppCycleClientMethods interface {
	// Stop initiates shutdown of the server.  It streams back periodic
	// updates to give the client an idea of how the shutdown is
	// progressing.
	Stop(*context.T, ...rpc.CallOpt) (AppCycleStopClientCall, error)
	// ForceStop tells the server to shut down right away.  It can be issued
	// while a Stop is outstanding if for example the client does not want
	// to wait any longer.
	ForceStop(*context.T, ...rpc.CallOpt) error
}

// AppCycleClientStub adds universal methods to AppCycleClientMethods.
type AppCycleClientStub interface {
	AppCycleClientMethods
	rpc.UniversalServiceMethods
}

// AppCycleClient returns a client stub for AppCycle.
func AppCycleClient(name string) AppCycleClientStub {
	return implAppCycleClientStub{name}
}

type implAppCycleClientStub struct {
	name string
}

func (c implAppCycleClientStub) Stop(ctx *context.T, opts ...rpc.CallOpt) (ocall AppCycleStopClientCall, err error) {
	var call rpc.ClientCall
	if call, err = v23.GetClient(ctx).StartCall(ctx, c.name, "Stop", nil, opts...); err != nil {
		return
	}
	ocall = &implAppCycleStopClientCall{ClientCall: call}
	return
}

func (c implAppCycleClientStub) ForceStop(ctx *context.T, opts ...rpc.CallOpt) (err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "ForceStop", nil, nil, opts...)
	return
}

// AppCycleStopClientStream is the client stream for AppCycle.Stop.
type AppCycleStopClientStream interface {
	// RecvStream returns the receiver side of the AppCycle.Stop client stream.
	RecvStream() interface {
		// Advance stages an item so that it may be retrieved via Value.  Returns
		// true iff there is an item to retrieve.  Advance must be called before
		// Value is called.  May block if an item is not available.
		Advance() bool
		// Value returns the item that was staged by Advance.  May panic if Advance
		// returned false or was not called.  Never blocks.
		Value() Task
		// Err returns any error encountered by Advance.  Never blocks.
		Err() error
	}
}

// AppCycleStopClientCall represents the call returned from AppCycle.Stop.
type AppCycleStopClientCall interface {
	AppCycleStopClientStream
	// Finish blocks until the server is done, and returns the positional return
	// values for call.
	//
	// Finish returns immediately if the call has been canceled; depending on the
	// timing the output could either be an error signaling cancelation, or the
	// valid positional return values from the server.
	//
	// Calling Finish is mandatory for releasing stream resources, unless the call
	// has been canceled or any of the other methods return an error.  Finish should
	// be called at most once.
	Finish() error
}

type implAppCycleStopClientCall struct {
	rpc.ClientCall
	valRecv Task
	errRecv error
}

func (c *implAppCycleStopClientCall) RecvStream() interface {
	Advance() bool
	Value() Task
	Err() error
} {
	return implAppCycleStopClientCallRecv{c}
}

type implAppCycleStopClientCallRecv struct {
	c *implAppCycleStopClientCall
}

func (c implAppCycleStopClientCallRecv) Advance() bool {
	c.c.valRecv = Task{}
	c.c.errRecv = c.c.Recv(&c.c.valRecv)
	return c.c.errRecv == nil
}
func (c implAppCycleStopClientCallRecv) Value() Task {
	return c.c.valRecv
}
func (c implAppCycleStopClientCallRecv) Err() error {
	if c.c.errRecv == io.EOF {
		return nil
	}
	return c.c.errRecv
}
func (c *implAppCycleStopClientCall) Finish() (err error) {
	err = c.ClientCall.Finish()
	return
}

// AppCycleServerMethods is the interface a server writer
// implements for AppCycle.
//
// AppCycle interfaces with the process running a vanadium runtime.
type AppCycleServerMethods interface {
	// Stop initiates shutdown of the server.  It streams back periodic
	// updates to give the client an idea of how the shutdown is
	// progressing.
	Stop(*context.T, AppCycleStopServerCall) error
	// ForceStop tells the server to shut down right away.  It can be issued
	// while a Stop is outstanding if for example the client does not want
	// to wait any longer.
	ForceStop(*context.T, rpc.ServerCall) error
}

// AppCycleServerStubMethods is the server interface containing
// AppCycle methods, as expected by rpc.Server.
// The only difference between this interface and AppCycleServerMethods
// is the streaming methods.
type AppCycleServerStubMethods interface {
	// Stop initiates shutdown of the server.  It streams back periodic
	// updates to give the client an idea of how the shutdown is
	// progressing.
	Stop(*context.T, *AppCycleStopServerCallStub) error
	// ForceStop tells the server to shut down right away.  It can be issued
	// while a Stop is outstanding if for example the client does not want
	// to wait any longer.
	ForceStop(*context.T, rpc.ServerCall) error
}

// AppCycleServerStub adds universal methods to AppCycleServerStubMethods.
type AppCycleServerStub interface {
	AppCycleServerStubMethods
	// Describe the AppCycle interfaces.
	Describe__() []rpc.InterfaceDesc
}

// AppCycleServer returns a server stub for AppCycle.
// It converts an implementation of AppCycleServerMethods into
// an object that may be used by rpc.Server.
func AppCycleServer(impl AppCycleServerMethods) AppCycleServerStub {
	stub := implAppCycleServerStub{
		impl: impl,
	}
	// Initialize GlobState; always check the stub itself first, to handle the
	// case where the user has the Glob method defined in their VDL source.
	if gs := rpc.NewGlobState(stub); gs != nil {
		stub.gs = gs
	} else if gs := rpc.NewGlobState(impl); gs != nil {
		stub.gs = gs
	}
	return stub
}

type implAppCycleServerStub struct {
	impl AppCycleServerMethods
	gs   *rpc.GlobState
}

func (s implAppCycleServerStub) Stop(ctx *context.T, call *AppCycleStopServerCallStub) error {
	return s.impl.Stop(ctx, call)
}

func (s implAppCycleServerStub) ForceStop(ctx *context.T, call rpc.ServerCall) error {
	return s.impl.ForceStop(ctx, call)
}

func (s implAppCycleServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implAppCycleServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{AppCycleDesc}
}

// AppCycleDesc describes the AppCycle interface.
var AppCycleDesc rpc.InterfaceDesc = descAppCycle

// descAppCycle hides the desc to keep godoc clean.
var descAppCycle = rpc.InterfaceDesc{
	Name:    "AppCycle",
	PkgPath: "v.io/v23/services/appcycle",
	Doc:     "// AppCycle interfaces with the process running a vanadium runtime.",
	Methods: []rpc.MethodDesc{
		{
			Name: "Stop",
			Doc:  "// Stop initiates shutdown of the server.  It streams back periodic\n// updates to give the client an idea of how the shutdown is\n// progressing.",
		},
		{
			Name: "ForceStop",
			Doc:  "// ForceStop tells the server to shut down right away.  It can be issued\n// while a Stop is outstanding if for example the client does not want\n// to wait any longer.",
		},
	},
}

// AppCycleStopServerStream is the server stream for AppCycle.Stop.
type AppCycleStopServerStream interface {
	// SendStream returns the send side of the AppCycle.Stop server stream.
	SendStream() interface {
		// Send places the item onto the output stream.  Returns errors encountered
		// while sending.  Blocks if there is no buffer space; will unblock when
		// buffer space is available.
		Send(item Task) error
	}
}

// AppCycleStopServerCall represents the context passed to AppCycle.Stop.
type AppCycleStopServerCall interface {
	rpc.ServerCall
	AppCycleStopServerStream
}

// AppCycleStopServerCallStub is a wrapper that converts rpc.StreamServerCall into
// a typesafe stub that implements AppCycleStopServerCall.
type AppCycleStopServerCallStub struct {
	rpc.StreamServerCall
}

// Init initializes AppCycleStopServerCallStub from rpc.StreamServerCall.
func (s *AppCycleStopServerCallStub) Init(call rpc.StreamServerCall) {
	s.StreamServerCall = call
}

// SendStream returns the send side of the AppCycle.Stop server stream.
func (s *AppCycleStopServerCallStub) SendStream() interface {
	Send(item Task) error
} {
	return implAppCycleStopServerCallSend{s}
}

type implAppCycleStopServerCallSend struct {
	s *AppCycleStopServerCallStub
}

func (s implAppCycleStopServerCallSend) Send(item Task) error {
	return s.s.Send(item)
}

// Hold type definitions in package-level variables, for better performance.
var (
	__VDLType_struct_1 *vdl.Type
)

var __VDLInitCalled bool

// __VDLInit performs vdl initialization.  It is safe to call multiple times.
// If you have an init ordering issue, just insert the following line verbatim
// into your source files in this package, right after the "package foo" clause:
//
//    var _ = __VDLInit()
//
// The purpose of this function is to ensure that vdl initialization occurs in
// the right order, and very early in the init sequence.  In particular, vdl
// registration and package variable initialization needs to occur before
// functions like vdl.TypeOf will work properly.
//
// This function returns a dummy value, so that it can be used to initialize the
// first var in the file, to take advantage of Go's defined init order.
func __VDLInit() struct{} {
	if __VDLInitCalled {
		return struct{}{}
	}
	__VDLInitCalled = true

	// Register types.
	vdl.Register((*Task)(nil))

	// Initialize type definitions.
	__VDLType_struct_1 = vdl.TypeOf((*Task)(nil)).Elem()

	return struct{}{}
}
