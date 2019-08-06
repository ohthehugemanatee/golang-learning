package reflection

import (
	"reflect"
	"testing"
)

type City struct {
	Name  string
	Stats CityStats
}

type CityStats struct {
	Population int
	Nickname   string
}

func TestWalk(t *testing.T) {
	testCases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			Name: "Simple single string value",
			Input: struct {
				string
			}{
				"value1",
			},
			ExpectedCalls: []string{"value1"},
		},
		{
			Name: "Multiple string fields",
			Input: struct {
				Name string
				City string
			}{
				"value1",
				"value2",
			},
			ExpectedCalls: []string{"value1", "value2"},
		},
		{
			Name: "Mixed type fields",
			Input: struct {
				Name       string
				City       string
				Population int
			}{
				"value1",
				"value2",
				110,
			},
			ExpectedCalls: []string{"value1", "value2"},
		},
		{
			Name: "Nested fields",
			Input: City{
				"value1",
				CityStats{
					111,
					"value2",
				},
			},
			ExpectedCalls: []string{"value1", "value2"},
		},
		{
			Name: "Pointers",
			Input: &City{
				"value1",
				CityStats{
					111,
					"value2",
				},
			},
			ExpectedCalls: []string{"value1", "value2"},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			got := []string{}
			Walk(testCase.Input, func(str string) {
				got = append(got, str)
			})
			if !reflect.DeepEqual(testCase.ExpectedCalls, got) {
				t.Errorf("Wrong value in function call. Expected %v, got %v", testCase.ExpectedCalls, got)
			}
		})
	}
}
