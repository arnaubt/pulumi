// Copyright 2025, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package httpstate

import (
	"os"
	"runtime"
	"testing"
)

func TestMain(m *testing.M) {
	if runtime.GOOS == "windows" {
		// Skip tests on Windows as they are not supported.
		// TODO[pulumi/pulumi#12345]: Remove this when Windows support is added.
		os.Stdout.WriteString("Skipping tests on Windows as they are not supported.\n")
		os.Exit(0)
	}
	// Run the tests.
	code := m.Run()

	// Exit with the code returned by the tests.
	os.Exit(code)
}
