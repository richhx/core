// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: localblobstore

package localblobstore

import (
	"time"
	"v.io/v23/vdl"
	vdltime "v.io/v23/vdlroot/time"
	"v.io/x/ref/services/syncbase/server/interfaces"
)

var _ = __VDLInit() // Must be first; see __VDLInit comments for details.

//////////////////////////////////////////////////
// Type definitions

// A BlobMetadata describes information that syncbase stores for a blob it
// holds, independent of the actual content.  Compare with a Signpost, which
// may be stored for a blob that the current device does not hold (and
// indicates where it may be found).  (See
// v.io/x/ref/services/syncbase/server/interfaces/sync_types.vdl for the
// Signpost definition.)
type BlobMetadata struct {
	OwnerShares interfaces.BlobSharesBySyncgroup // >0 for any group => syncbase must keep blob.
	Referenced  time.Time                        // When structured-store reference to blob last seen.
	Accessed    time.Time                        // Last attempted access.
}

func (BlobMetadata) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/localblobstore.BlobMetadata"`
}) {
}

func (x BlobMetadata) VDLIsZero() bool {
	if len(x.OwnerShares) != 0 {
		return false
	}
	if !x.Referenced.IsZero() {
		return false
	}
	if !x.Accessed.IsZero() {
		return false
	}
	return true
}

func (x BlobMetadata) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(__VDLType_struct_1); err != nil {
		return err
	}
	if len(x.OwnerShares) != 0 {
		if err := enc.NextField(0); err != nil {
			return err
		}
		if err := x.OwnerShares.VDLWrite(enc); err != nil {
			return err
		}
	}
	if !x.Referenced.IsZero() {
		if err := enc.NextField(1); err != nil {
			return err
		}
		var wire vdltime.Time
		if err := vdltime.TimeFromNative(&wire, x.Referenced); err != nil {
			return err
		}
		if err := wire.VDLWrite(enc); err != nil {
			return err
		}
	}
	if !x.Accessed.IsZero() {
		if err := enc.NextField(2); err != nil {
			return err
		}
		var wire vdltime.Time
		if err := vdltime.TimeFromNative(&wire, x.Accessed); err != nil {
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

func (x *BlobMetadata) VDLRead(dec vdl.Decoder) error {
	*x = BlobMetadata{}
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
			if err := x.OwnerShares.VDLRead(dec); err != nil {
				return err
			}
		case 1:
			var wire vdltime.Time
			if err := wire.VDLRead(dec); err != nil {
				return err
			}
			if err := vdltime.TimeToNative(wire, &x.Referenced); err != nil {
				return err
			}
		case 2:
			var wire vdltime.Time
			if err := wire.VDLRead(dec); err != nil {
				return err
			}
			if err := vdltime.TimeToNative(wire, &x.Accessed); err != nil {
				return err
			}
		}
	}
}

// A PerSyncgroup is blob-related data stored per syncgroup.
// It includes information that helps syncgroup members decide whether
// a peer makes a better or worse owner of a blob.
type PerSyncgroup struct {
	Priority interfaces.SgPriority
}

func (PerSyncgroup) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/localblobstore.PerSyncgroup"`
}) {
}

func (x PerSyncgroup) VDLIsZero() bool {
	if !x.Priority.VDLIsZero() {
		return false
	}
	return true
}

func (x PerSyncgroup) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(__VDLType_struct_4); err != nil {
		return err
	}
	if !x.Priority.VDLIsZero() {
		if err := enc.NextField(0); err != nil {
			return err
		}
		if err := x.Priority.VDLWrite(enc); err != nil {
			return err
		}
	}
	if err := enc.NextField(-1); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *PerSyncgroup) VDLRead(dec vdl.Decoder) error {
	*x = PerSyncgroup{}
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
			if err := x.Priority.VDLRead(dec); err != nil {
				return err
			}
		}
	}
}

// Hold type definitions in package-level variables, for better performance.
var (
	__VDLType_struct_1 *vdl.Type
	__VDLType_map_2    *vdl.Type
	__VDLType_struct_3 *vdl.Type
	__VDLType_struct_4 *vdl.Type
	__VDLType_struct_5 *vdl.Type
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
	vdl.Register((*BlobMetadata)(nil))
	vdl.Register((*PerSyncgroup)(nil))

	// Initialize type definitions.
	__VDLType_struct_1 = vdl.TypeOf((*BlobMetadata)(nil)).Elem()
	__VDLType_map_2 = vdl.TypeOf((*interfaces.BlobSharesBySyncgroup)(nil))
	__VDLType_struct_3 = vdl.TypeOf((*vdltime.Time)(nil)).Elem()
	__VDLType_struct_4 = vdl.TypeOf((*PerSyncgroup)(nil)).Elem()
	__VDLType_struct_5 = vdl.TypeOf((*interfaces.SgPriority)(nil)).Elem()

	return struct{}{}
}
