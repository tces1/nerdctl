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

package container

import (
	"testing"

	"github.com/containerd/nerdctl/v2/pkg/testutil"
	"github.com/containerd/nerdctl/v2/pkg/testutil/nerdtest"
	"github.com/containerd/nerdctl/v2/pkg/testutil/test"
)

func TestRunStargz(t *testing.T) {
	testCase := nerdtest.Setup()

	testCase.Require = test.Require(
		nerdtest.Stargz,
		test.Amd64,
		test.Not(nerdtest.Docker),
	)

	testCase.Command = test.Command("--snapshotter=stargz", "run", "--rm", testutil.FedoraESGZImage, "ls", "/.stargz-snapshotter")

	testCase.Expected = test.Expects(0, nil, nil)

	testCase.Run(t)
}
