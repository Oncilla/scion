// Copyright 2018 ETH Zurich
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package state

import (
	"io/ioutil"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/scionproto/scion/go/lib/common"
)

var testMatrix Matrix

func loadFile(filename string, t *testing.T) common.RawBytes {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("Error loading config from '%s': %v", filename, err)
	}
	return b
}

func loadMatrix(filename string, t *testing.T) {
	m, err := MatrixFromFile(filename)
	if err != nil {
		t.Fatalf("Error loading config from '%s': %v", filename, err)
	}
	testMatrix = m
}

func Test_MatrixFromRaw(t *testing.T) {
	fn := "testdata/matrix.yml"
	Convey("Matrix is loaded correctly", t, func() {
		raw := loadFile(fn, t)
		m, err := MatrixFromRaw(raw)
		SoMsg("Err", err, ShouldBeNil)
		SoMsg("Entry 52-19", m[52][19], ShouldEqual, 5219)
		SoMsg("Entry 64-19", m[64][19], ShouldEqual, 6419)
		SoMsg("Entry 19-64", m[19][64], ShouldEqual, 1964)
		SoMsg("Entry 19-52", m[19][52], ShouldEqual, 1952)
		SoMsg("Entry 64-00", m[64][0], ShouldEqual, 6400)
		SoMsg("Entry 52-00", m[52][0], ShouldEqual, 5200)
		SoMsg("Entry 00-64", m[0][64], ShouldEqual, 64)
		SoMsg("Entry 00-52", m[0][52], ShouldEqual, 52)
		SoMsg("Entry 00-19", m[0][19], ShouldEqual, 19)
	})
}
