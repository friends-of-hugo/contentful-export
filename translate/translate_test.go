package translate

import (
	"reflect"
	"testing"

	"github.com/bhsi-cinch/contentful-hugo/mapper"
)

func TestConvertContent(t *testing.T) {
	tests := []struct {
		Map      map[string]interface{}
		fields   []mapper.TypeField
		expected Content
	}{
		{
			map[string]interface{}{
				"key": "value",
			},
			[]mapper.TypeField{
				mapper.TypeField{"key", "", "String", false, false, false, false},
			},
			Content{
				map[string]interface{}{
					"key": "value",
				},
				"",
				"",
			},
		},
		{
			map[string]interface{}{
				"key":         "value",
				"mainContent": "This is test main content\nand one more line",
				"slug":        "my-test-slug",
			},
			[]mapper.TypeField{
				mapper.TypeField{"key", "", "String", false, false, false, false},
				mapper.TypeField{"mainContent", "", "String", false, false, false, false},
				mapper.TypeField{"slug", "", "String", false, false, false, false},
			},
			Content{
				map[string]interface{}{
					"key":  "value",
					"slug": "my-test-slug",
				},
				"This is test main content\nand one more line",
				"my-test-slug",
			},
		},
	}

	tc := TranslationConfig{}

	for _, test := range tests {
		result := tc.convertContent(test.Map, test.fields)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("convertContent(%v, %v) incorrect, expected %v, got %v", test.Map, test.fields, test.expected, result)
		}

	}
}

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
		field    mapper.TypeField
		expected interface{}
	}{
		{
			"Unchanged",
			mapper.TypeField{"", "", "default", false, false, false, false},
			"Unchanged",
		},
		{
			[]interface{}{
				map[string]interface{}{"sys": map[string]interface{}{"id": "test-id-1", "linkType": "Entry"}},
				map[string]interface{}{"sys": map[string]interface{}{"id": "test-id-2", "linkType": "Entry"}},
				map[string]interface{}{"sys": map[string]interface{}{"id": "test-id-3", "linkType": "Entry"}},
			},
			mapper.TypeField{"", "", "Array", false, false, false, false},
			[]string{"test-id-1.md", "test-id-2.md", "test-id-3.md"},
		},
	}

	tc := TranslationConfig{}

	for _, test := range tests {
		result := tc.translateField(test.value, test.field)

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("translateField(%v, %v) incorrect, expected %v, got %v", test.value, test.field.Type, test.expected, result)
		}
	}

}
