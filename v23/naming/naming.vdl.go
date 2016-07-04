// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: naming

package naming

import (
	"fmt"
	"v.io/v23/vdl"
	vdltime "v.io/v23/vdlroot/time"
	"v.io/v23/verror"
)

var _ = __VDLInit() // Must be first; see __VDLInit comments for details.

//////////////////////////////////////////////////
// Type definitions

// MountFlag is a bit mask of options to the mount call.
type MountFlag uint32

func (MountFlag) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/naming.MountFlag"`
}) {
}

func (x MountFlag) VDLIsZero() bool {
	return x == 0
}

func (x MountFlag) VDLWrite(enc vdl.Encoder) error {
	if err := enc.WriteValueUint(__VDLType_uint32_1, uint64(x)); err != nil {
		return err
	}
	return nil
}

func (x *MountFlag) VDLRead(dec vdl.Decoder) error {
	switch value, err := dec.ReadValueUint(32); {
	case err != nil:
		return err
	default:
		*x = MountFlag(value)
	}
	return nil
}

// MountedServer represents a server mounted on a specific name.
type MountedServer struct {
	// Server is the OA that's mounted.
	Server string
	// Deadline before the mount entry expires.
	Deadline vdltime.Deadline
}

func (MountedServer) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/naming.MountedServer"`
}) {
}

func (x MountedServer) VDLIsZero() bool {
	if x.Server != "" {
		return false
	}
	if !x.Deadline.Time.IsZero() {
		return false
	}
	return true
}

func (x MountedServer) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(__VDLType_struct_2); err != nil {
		return err
	}
	if x.Server != "" {
		if err := enc.NextFieldValueString(0, vdl.StringType, x.Server); err != nil {
			return err
		}
	}
	if !x.Deadline.Time.IsZero() {
		if err := enc.NextField(1); err != nil {
			return err
		}
		var wire vdltime.WireDeadline
		if err := vdltime.WireDeadlineFromNative(&wire, x.Deadline); err != nil {
			return err
		}
		if err := wire.VDLWrite(enc); err != nil {
			return err
		}
	}
	if err := enc.NextField(-1); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *MountedServer) VDLRead(dec vdl.Decoder) error {
	*x = MountedServer{}
	if err := dec.StartValue(__VDLType_struct_2); err != nil {
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
		if decType != __VDLType_struct_2 {
			index = __VDLType_struct_2.FieldIndexByName(decType.Field(index).Name)
			if index == -1 {
				if err := dec.SkipValue(); err != nil {
					return err
				}
				continue
			}
		}
		switch index {
		case 0:
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				x.Server = value
			}
		case 1:
			var wire vdltime.WireDeadline
			if err := wire.VDLRead(dec); err != nil {
				return err
			}
			if err := vdltime.WireDeadlineToNative(wire, &x.Deadline); err != nil {
				return err
			}
		}
	}
}

// MountEntry represents a given name mounted in the mounttable.
type MountEntry struct {
	// Name is the mounted name.
	Name string
	// Servers (if present) specifies the mounted names.
	Servers []MountedServer
	// ServesMountTable is true if the servers represent mount tables.
	ServesMountTable bool
	// IsLeaf is true if this entry represents a leaf object.
	IsLeaf bool
}

func (MountEntry) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/naming.MountEntry"`
}) {
}

func (x MountEntry) VDLIsZero() bool {
	if x.Name != "" {
		return false
	}
	if len(x.Servers) != 0 {
		return false
	}
	if x.ServesMountTable {
		return false
	}
	if x.IsLeaf {
		return false
	}
	return true
}

func (x MountEntry) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(__VDLType_struct_4); err != nil {
		return err
	}
	if x.Name != "" {
		if err := enc.NextFieldValueString(0, vdl.StringType, x.Name); err != nil {
			return err
		}
	}
	if len(x.Servers) != 0 {
		if err := enc.NextField(1); err != nil {
			return err
		}
		if err := __VDLWriteAnon_list_1(enc, x.Servers); err != nil {
			return err
		}
	}
	if x.ServesMountTable {
		if err := enc.NextFieldValueBool(2, vdl.BoolType, x.ServesMountTable); err != nil {
			return err
		}
	}
	if x.IsLeaf {
		if err := enc.NextFieldValueBool(3, vdl.BoolType, x.IsLeaf); err != nil {
			return err
		}
	}
	if err := enc.NextField(-1); err != nil {
		return err
	}
	return enc.FinishValue()
}

func __VDLWriteAnon_list_1(enc vdl.Encoder, x []MountedServer) error {
	if err := enc.StartValue(__VDLType_list_5); err != nil {
		return err
	}
	if err := enc.SetLenHint(len(x)); err != nil {
		return err
	}
	for _, elem := range x {
		if err := enc.NextEntry(false); err != nil {
			return err
		}
		if err := elem.VDLWrite(enc); err != nil {
			return err
		}
	}
	if err := enc.NextEntry(true); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *MountEntry) VDLRead(dec vdl.Decoder) error {
	*x = MountEntry{}
	if err := dec.StartValue(__VDLType_struct_4); err != nil {
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
		if decType != __VDLType_struct_4 {
			index = __VDLType_struct_4.FieldIndexByName(decType.Field(index).Name)
			if index == -1 {
				if err := dec.SkipValue(); err != nil {
					return err
				}
				continue
			}
		}
		switch index {
		case 0:
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				x.Name = value
			}
		case 1:
			if err := __VDLReadAnon_list_1(dec, &x.Servers); err != nil {
				return err
			}
		case 2:
			switch value, err := dec.ReadValueBool(); {
			case err != nil:
				return err
			default:
				x.ServesMountTable = value
			}
		case 3:
			switch value, err := dec.ReadValueBool(); {
			case err != nil:
				return err
			default:
				x.IsLeaf = value
			}
		}
	}
}

func __VDLReadAnon_list_1(dec vdl.Decoder, x *[]MountedServer) error {
	if err := dec.StartValue(__VDLType_list_5); err != nil {
		return err
	}
	if len := dec.LenHint(); len > 0 {
		*x = make([]MountedServer, 0, len)
	} else {
		*x = nil
	}
	for {
		switch done, err := dec.NextEntry(); {
		case err != nil:
			return err
		case done:
			return dec.FinishValue()
		default:
			var elem MountedServer
			if err := elem.VDLRead(dec); err != nil {
				return err
			}
			*x = append(*x, elem)
		}
	}
}

// GlobError is returned by namespace.Glob to indicate a subtree of the namespace
// that could not be traversed.
type GlobError struct {
	// Root of the subtree.
	Name string
	// The error that occurred fulfilling the request.
	Error error
}

func (GlobError) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/naming.GlobError"`
}) {
}

func (x GlobError) VDLIsZero() bool {
	return x == GlobError{}
}

func (x GlobError) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(__VDLType_struct_6); err != nil {
		return err
	}
	if x.Name != "" {
		if err := enc.NextFieldValueString(0, vdl.StringType, x.Name); err != nil {
			return err
		}
	}
	if x.Error != nil {
		if err := enc.NextField(1); err != nil {
			return err
		}
		if err := verror.VDLWrite(enc, x.Error); err != nil {
			return err
		}
	}
	if err := enc.NextField(-1); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *GlobError) VDLRead(dec vdl.Decoder) error {
	*x = GlobError{}
	if err := dec.StartValue(__VDLType_struct_6); err != nil {
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
		if decType != __VDLType_struct_6 {
			index = __VDLType_struct_6.FieldIndexByName(decType.Field(index).Name)
			if index == -1 {
				if err := dec.SkipValue(); err != nil {
					return err
				}
				continue
			}
		}
		switch index {
		case 0:
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				x.Name = value
			}
		case 1:
			if err := verror.VDLRead(dec, &x.Error); err != nil {
				return err
			}
		}
	}
}

type (
	// GlobReply represents any single field of the GlobReply union type.
	//
	// GlobReply is the data type returned by Glob__.
	GlobReply interface {
		// Index returns the field index.
		Index() int
		// Interface returns the field value as an interface.
		Interface() interface{}
		// Name returns the field name.
		Name() string
		// __VDLReflect describes the GlobReply union type.
		__VDLReflect(__GlobReplyReflect)
		VDLIsZero() bool
		VDLWrite(vdl.Encoder) error
	}
	// GlobReplyEntry represents field Entry of the GlobReply union type.
	GlobReplyEntry struct{ Value MountEntry }
	// GlobReplyError represents field Error of the GlobReply union type.
	GlobReplyError struct{ Value GlobError }
	// __GlobReplyReflect describes the GlobReply union type.
	__GlobReplyReflect struct {
		Name  string `vdl:"v.io/v23/naming.GlobReply"`
		Type  GlobReply
		Union struct {
			Entry GlobReplyEntry
			Error GlobReplyError
		}
	}
)

func (x GlobReplyEntry) Index() int                      { return 0 }
func (x GlobReplyEntry) Interface() interface{}          { return x.Value }
func (x GlobReplyEntry) Name() string                    { return "Entry" }
func (x GlobReplyEntry) __VDLReflect(__GlobReplyReflect) {}

func (x GlobReplyError) Index() int                      { return 1 }
func (x GlobReplyError) Interface() interface{}          { return x.Value }
func (x GlobReplyError) Name() string                    { return "Error" }
func (x GlobReplyError) __VDLReflect(__GlobReplyReflect) {}

func (x GlobReplyEntry) VDLIsZero() bool {
	return x.Value.VDLIsZero()
}

func (x GlobReplyError) VDLIsZero() bool {
	return false
}

func (x GlobReplyEntry) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(__VDLType_union_7); err != nil {
		return err
	}
	if err := enc.NextField(0); err != nil {
		return err
	}
	if err := x.Value.VDLWrite(enc); err != nil {
		return err
	}
	if err := enc.NextField(-1); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x GlobReplyError) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(__VDLType_union_7); err != nil {
		return err
	}
	if err := enc.NextField(1); err != nil {
		return err
	}
	if err := x.Value.VDLWrite(enc); err != nil {
		return err
	}
	if err := enc.NextField(-1); err != nil {
		return err
	}
	return enc.FinishValue()
}

func VDLReadGlobReply(dec vdl.Decoder, x *GlobReply) error {
	if err := dec.StartValue(__VDLType_union_7); err != nil {
		return err
	}
	decType := dec.Type()
	index, err := dec.NextField()
	switch {
	case err != nil:
		return err
	case index == -1:
		return fmt.Errorf("missing field in union %T, from %v", x, decType)
	}
	if decType != __VDLType_union_7 {
		name := decType.Field(index).Name
		index = __VDLType_union_7.FieldIndexByName(name)
		if index == -1 {
			return fmt.Errorf("field %q not in union %T, from %v", name, x, decType)
		}
	}
	switch index {
	case 0:
		var field GlobReplyEntry
		if err := field.Value.VDLRead(dec); err != nil {
			return err
		}
		*x = field
	case 1:
		var field GlobReplyError
		if err := field.Value.VDLRead(dec); err != nil {
			return err
		}
		*x = field
	}
	switch index, err := dec.NextField(); {
	case err != nil:
		return err
	case index != -1:
		return fmt.Errorf("extra field %d in union %T, from %v", index, x, dec.Type())
	}
	return dec.FinishValue()
}

type (
	// GlobChildrenReply represents any single field of the GlobChildrenReply union type.
	//
	// GlobChildrenReply is the data type returned by GlobChildren__.
	GlobChildrenReply interface {
		// Index returns the field index.
		Index() int
		// Interface returns the field value as an interface.
		Interface() interface{}
		// Name returns the field name.
		Name() string
		// __VDLReflect describes the GlobChildrenReply union type.
		__VDLReflect(__GlobChildrenReplyReflect)
		VDLIsZero() bool
		VDLWrite(vdl.Encoder) error
	}
	// GlobChildrenReplyName represents field Name of the GlobChildrenReply union type.
	GlobChildrenReplyName struct{ Value string }
	// GlobChildrenReplyError represents field Error of the GlobChildrenReply union type.
	GlobChildrenReplyError struct{ Value GlobError }
	// __GlobChildrenReplyReflect describes the GlobChildrenReply union type.
	__GlobChildrenReplyReflect struct {
		Name  string `vdl:"v.io/v23/naming.GlobChildrenReply"`
		Type  GlobChildrenReply
		Union struct {
			Name  GlobChildrenReplyName
			Error GlobChildrenReplyError
		}
	}
)

func (x GlobChildrenReplyName) Index() int                              { return 0 }
func (x GlobChildrenReplyName) Interface() interface{}                  { return x.Value }
func (x GlobChildrenReplyName) Name() string                            { return "Name" }
func (x GlobChildrenReplyName) __VDLReflect(__GlobChildrenReplyReflect) {}

func (x GlobChildrenReplyError) Index() int                              { return 1 }
func (x GlobChildrenReplyError) Interface() interface{}                  { return x.Value }
func (x GlobChildrenReplyError) Name() string                            { return "Error" }
func (x GlobChildrenReplyError) __VDLReflect(__GlobChildrenReplyReflect) {}

func (x GlobChildrenReplyName) VDLIsZero() bool {
	return x.Value == ""
}

func (x GlobChildrenReplyError) VDLIsZero() bool {
	return false
}

func (x GlobChildrenReplyName) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(__VDLType_union_8); err != nil {
		return err
	}
	if err := enc.NextFieldValueString(0, vdl.StringType, x.Value); err != nil {
		return err
	}
	if err := enc.NextField(-1); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x GlobChildrenReplyError) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(__VDLType_union_8); err != nil {
		return err
	}
	if err := enc.NextField(1); err != nil {
		return err
	}
	if err := x.Value.VDLWrite(enc); err != nil {
		return err
	}
	if err := enc.NextField(-1); err != nil {
		return err
	}
	return enc.FinishValue()
}

func VDLReadGlobChildrenReply(dec vdl.Decoder, x *GlobChildrenReply) error {
	if err := dec.StartValue(__VDLType_union_8); err != nil {
		return err
	}
	decType := dec.Type()
	index, err := dec.NextField()
	switch {
	case err != nil:
		return err
	case index == -1:
		return fmt.Errorf("missing field in union %T, from %v", x, decType)
	}
	if decType != __VDLType_union_8 {
		name := decType.Field(index).Name
		index = __VDLType_union_8.FieldIndexByName(name)
		if index == -1 {
			return fmt.Errorf("field %q not in union %T, from %v", name, x, decType)
		}
	}
	switch index {
	case 0:
		var field GlobChildrenReplyName
		switch value, err := dec.ReadValueString(); {
		case err != nil:
			return err
		default:
			field.Value = value
		}
		*x = field
	case 1:
		var field GlobChildrenReplyError
		if err := field.Value.VDLRead(dec); err != nil {
			return err
		}
		*x = field
	}
	switch index, err := dec.NextField(); {
	case err != nil:
		return err
	case index != -1:
		return fmt.Errorf("extra field %d in union %T, from %v", index, x, dec.Type())
	}
	return dec.FinishValue()
}

//////////////////////////////////////////////////
// Const definitions

const Replace = MountFlag(1) // Replace means the mount should replace what is currently at the mount point
const MT = MountFlag(2)      // MT means that the target server is a mount table.
const Leaf = MountFlag(4)    // Leaf means that the target server is a leaf.

// Hold type definitions in package-level variables, for better performance.
var (
	__VDLType_uint32_1 *vdl.Type
	__VDLType_struct_2 *vdl.Type
	__VDLType_struct_3 *vdl.Type
	__VDLType_struct_4 *vdl.Type
	__VDLType_list_5   *vdl.Type
	__VDLType_struct_6 *vdl.Type
	__VDLType_union_7  *vdl.Type
	__VDLType_union_8  *vdl.Type
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
	vdl.Register((*MountFlag)(nil))
	vdl.Register((*MountedServer)(nil))
	vdl.Register((*MountEntry)(nil))
	vdl.Register((*GlobError)(nil))
	vdl.Register((*GlobReply)(nil))
	vdl.Register((*GlobChildrenReply)(nil))

	// Initialize type definitions.
	__VDLType_uint32_1 = vdl.TypeOf((*MountFlag)(nil))
	__VDLType_struct_2 = vdl.TypeOf((*MountedServer)(nil)).Elem()
	__VDLType_struct_3 = vdl.TypeOf((*vdltime.WireDeadline)(nil)).Elem()
	__VDLType_struct_4 = vdl.TypeOf((*MountEntry)(nil)).Elem()
	__VDLType_list_5 = vdl.TypeOf((*[]MountedServer)(nil))
	__VDLType_struct_6 = vdl.TypeOf((*GlobError)(nil)).Elem()
	__VDLType_union_7 = vdl.TypeOf((*GlobReply)(nil))
	__VDLType_union_8 = vdl.TypeOf((*GlobChildrenReply)(nil))

	return struct{}{}
}
