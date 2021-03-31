// Copyright 2017 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// +build cgo

package d2xx

import _ "periph.io/x/extra/hostextra/d2xx/linux_amd64"

/*
#cgo LDFLAGS: ${SRCDIR}/linux_amd64/libftd2xx.a
*/
import "C"
