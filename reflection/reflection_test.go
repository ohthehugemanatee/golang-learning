package reflection

import (
	"reflect"
	"testing"
)

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
	}
	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			got := []string{}
			Walk(testCase.Input, func(str string) {
				got = append(got, str)
			})
			if !reflect.DeepEqual(testCase.ExpectedCalls, got) {
				t.Errorf("Wrong value in function call. Expected %s, got %s", testCase.ExpectedCalls[0], got[0])
			}
		})
	}
}
