//go:build !windows

/*
   Copyright The containerd Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package dialer

import (
	"errors"
	"fmt"
	"net"
	"strings"
	"syscall"
	"time"
)

// DialAddress returns the address with unix:// prepended to the
// provided address
func DialAddress(address string) string {
	return fmt.Sprintf("unix://%s", address)
}

func isNoent(err error) bool {
	return errors.Is(err, syscall.ENOENT)
}

func dialer(address string, timeout time.Duration) (net.Conn, error) {
	address = strings.TrimPrefix(address, "unix://")
	return net.DialTimeout("unix", address, timeout)
}
