package dump

import (
	"reflect"
	"testing"
)

func TestRemoveItem(t *testing.T) {
	tests := []struct {
		initial       map[string]interface{}
		toDelete      string
		expectedValue string
		expectedMap   map[string]interface{}
	}{
		{
			map[string]interface{}{
				"one": "value-1",
				"two": "value-2",
			},
			"one",
			"value-1",
			map[string]interface{}{
				"two": "value-2",
			},
		},
		{
			map[string]interface{}{
				"two":   "value-2",
				"three": "value-3",
			},
			"one",
			"",
			map[string]interface{}{
				"two":   "value-2",
				"three": "value-3",
			},
		},
	}

	for _, test := range tests {
		resultValue := removeItem(test.initial, test.toDelete)

		if !reflect.DeepEqual(resultValue, test.expectedValue) {
			t.Errorf("removeItem(%v, %v) return value incorrect, expected %v, got %v", test.initial, test.toDelete, test.expectedValue, resultValue)
		}
		if !reflect.DeepEqual(test.initial, test.expectedMap) {
			t.Errorf("removeItem(%v, %v) resulting map incorrect, expected %v, got %v", test.initial, test.toDelete, test.expectedMap, test.initial)
		}

	}
}

func TestTranslateField(t *testing.T) {
	tests := []struct {
		value    interface{}
		field    TypeField
		expected interface{}
	}{
		{
			"Unchanged",
			TypeField{"", "", "default", false, false, false, false},
			"Unchanged",
		},
		{
			[]interface{}{
				map[string]interface{}{"sys": map[string]interface{}{"id": "test-id-1"}},
				map[string]interface{}{"sys": map[string]interface{}{"id": "test-id-2"}},
				map[string]interface{}{"sys": map[string]interface{}{"id": "test-id-3"}},
			},
			TypeField{"", "", "Array", false, false, false, false},
			[]string{"test-id-1.md", "test-id-2.md", "test-id-3.md"},
		},
	}

	for _, test := range tests {
		result := translateField(test.value, test.field)

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("translateField(%v, %v) incorrect, expected %v, got %v", test.value, test.field.Type, test.expected, result)
		}
	}

}
