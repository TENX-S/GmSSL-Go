/*
 *  Copyright 2014-2023 The GmSSL Project. All Rights Reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the License); you may
 *  not use this file except in compliance with the License.
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 */
/* +build cgo */

package gmssl

/*
#cgo CFLAGS: -I./include
#cgo darwin LDFLAGS: -L./lib/darwin -lgmssl -framework Security
#cgo linux,amd64 LDFLAGS: -L./lib/linux-amd64 -lgmssl
#cgo linux,arm64 LDFLAGS: -L./lib/linux-arm64 -lgmssl
#cgo windows,amd64 LDFLAGS: -L./lib/windows-amd64 -lgmssl -lws2_32
*/
import "C"
