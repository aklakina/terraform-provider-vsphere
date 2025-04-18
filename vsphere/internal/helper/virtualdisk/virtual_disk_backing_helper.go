// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualdisk

import (
	"github.com/vmware/govmomi/vim25/types"
	"reflect"
)

func VirtualDiskToSchemaPropsMap(backing types.BaseVirtualDeviceBackingInfo) map[string]interface{} {
	m := make(map[string]interface{})

	if backing == nil {
		return m
	}

	// Get reflect value of backing
	rv := reflect.Indirect(reflect.ValueOf(backing))

	// If it's a pointer, get the value it points to
	if rv.Kind() != reflect.Struct {
		return m
	}

	// Extract fields
	for i := 0; i < rv.NumField(); i++ {
		f := rv.Type().Field(i)
		v := rv.Field(i)

		// Only include exported fields.
		if f.IsExported() {
			m[f.Name] = v.Interface()
		}
	}

	return m
}
