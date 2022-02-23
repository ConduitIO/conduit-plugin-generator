// Copyright © 2022 Meroxa, Inc.
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

package generator

import (
	"testing"
	"time"

	"github.com/matryer/is"
)

func TestParseFull(t *testing.T) {
	is := is.New(t)
	underTest, err := Parse(map[string]string{
		"recordCount": "-1",
		"readTime":    "5s",
		"fields":      "id:int,name:string,joined:time,admin:bool",
	})
	is.NoErr(err)
	is.Equal(int64(-1), underTest.RecordCount)
	is.Equal(5*time.Second, underTest.ReadTime)
	is.Equal(map[string]string{"id": "int", "name": "string", "joined": "time", "admin": "bool"}, underTest.Fields)
}

func TestParseFields_RequiredNotPresent(t *testing.T) {
	is := is.New(t)
	_, err := Parse(map[string]string{
		"recordCount": "100",
		"readTime":    "5s",
	})
	is.True(err != nil)
	is.Equal("no fields specified", err.Error())
}

func TestParseFields_OptionalNotPresent(t *testing.T) {
	is := is.New(t)
	_, err := Parse(map[string]string{
		"fields": "a:int",
	})
	is.NoErr(err)
}

func TestParseFields_MalformedFields_NoType(t *testing.T) {
	is := is.New(t)
	_, err := Parse(map[string]string{
		"fields": "abc:",
	})
	is.True(err != nil)
	is.Equal(`invalid field spec "abc:"`, err.Error())
}

func TestParseFields_MalformedFields_NameOnly(t *testing.T) {
	is := is.New(t)
	_, err := Parse(map[string]string{
		"fields": "abc",
	})
	is.True(err != nil)
	is.Equal(`invalid field spec "abc"`, err.Error())
}
