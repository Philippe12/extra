// Copyright 2016 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package winthermal

import (
	"github.com/StackExchange/wmi"
	"periph.io/x/conn/v3/driver/driverreg"
)

func initWindows() error {
	// https://msdn.microsoft.com/en-us/library/aa394493.aspx
	var obj interface{}
	if err := wmi.Query("SELECT * FROM Win32_TemperatureProbe", &obj); err != nil {
		return err
	}
	return nil
}

func init() {
	driverreg.MustRegister(&driver{})
}
