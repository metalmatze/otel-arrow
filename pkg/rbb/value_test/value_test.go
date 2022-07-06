// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package value_test

import (
	"otel-arrow-adapter/pkg/rbb/value"
	"testing"
)

func TestNormalize(t *testing.T) {
	t.Parallel()

	// Test normalize on a scalar value
	s1 := value.Bool{
		Value: true,
	}
	s1.Normalize()
	if s1.Value != true {
		t.Errorf("Expected true, got %v", s1.Value)
	}

	// Test normalize on a struct value
	s2 := value.Struct{
		Fields: []value.Field{
			{
				Name:  "c",
				Value: &value.Bool{Value: false},
			},
			{
				Name:  "a",
				Value: &value.I64{Value: 123},
			},
			{
				Name:  "b",
				Value: &value.String{Value: "text"},
			},
		},
	}
	s2.Normalize()
	if s2.Fields[0].Name != "a" {
		t.Errorf("Expected a, got %v", s2.Fields[0].Name)
	}
	if s2.Fields[1].Name != "b" {
		t.Errorf("Expected b, got %v", s2.Fields[1].Name)
	}
	if s2.Fields[2].Name != "c" {
		t.Errorf("Expected c, got %v", s2.Fields[2].Name)
	}

	// Test normalize on a list value
	s3 := value.List{
		Values: []value.Value{
			&value.Bool{Value: true},
			&value.I64{Value: 1},
			&value.Struct{
				Fields: []value.Field{
					{
						Name:  "c",
						Value: &value.Bool{Value: false},
					},
					{
						Name:  "a",
						Value: &value.I64{Value: 123},
					},
					{
						Name:  "b",
						Value: &value.String{Value: "text"},
					},
				},
			},
		},
	}
	s3.Normalize()
	if s3.Values[2].(*value.Struct).Fields[0].Name != "a" {
		t.Errorf("Expected a, got %v", s2.Fields[0].Name)
	}
	if s3.Values[2].(*value.Struct).Fields[1].Name != "b" {
		t.Errorf("Expected b, got %v", s2.Fields[1].Name)
	}
	if s3.Values[2].(*value.Struct).Fields[2].Name != "c" {
		t.Errorf("Expected c, got %v", s2.Fields[2].Name)
	}
}