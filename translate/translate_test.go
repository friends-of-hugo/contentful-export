package translate

import (
	"reflect"
	"testing"

	"github.com/bhsi-cinch/contentful-hugo/mapper"
)

func TestTranslateFromMarkdown(t *testing.T) {
	yamlMarkdown := `
---
overriddentestbool: false
overriddentestint: 13
overriddenteststring: "please, overide me"
overriddentestslicenil: []
overriddentestslicebool: [false, false, false]
overriddentestsliceint: [13, 13, 13]		
overriddentestslicestring: ["please", "overide", "me"]
testbool: true
testint: 42
teststring: "test"
testslicenil: []
testslicebool: [true, false, true]
testsliceint: [1, 2, 3]		
testslicestring: ["one", "two", "thee"]
---
`
	tomlMarkdown := `
+++
overriddentestbool = false
overriddentestint = 13
overriddenteststring = "please, overide me"
overriddentestslicenil = []
overriddentestslicebool = [false, false, false]
overriddentestsliceint = [13, 13, 13]		
overriddentestslicestring = ["please", "overide", "me"]
testbool = true
testint = 42
teststring = "test"
testslicenil = []
testslicebool = [true, false, true]
testsliceint = [1, 2, 3]		
testslicestring = ["one", "two", "thee"]
+++
`
	extraMarkdown := "\n#header\n_italics_\n__bold__"
	archetypeYaml := yamlMarkdown + extraMarkdown
	archetypeToml := tomlMarkdown + extraMarkdown

	tests := []struct {
		encoding   string
		givenInput string
	}{
		{
			encoding:   "yaml",
			givenInput: archetypeYaml,
		},
		{
			encoding:   "toml",
			givenInput: archetypeToml,
		},
	}

	for _, test := range tests {
		tc := TranslationContext{TransConfig: TransConfig{Encoding: test.encoding}}
		result, err := tc.TranslateFromMarkdown(test.givenInput)
		if err != nil || result == nil {
			t.Errorf("TranslateFromMarkdown() failed...\n\nInput:\n%s\n\nActual Output:\n%v\n\nInner Error:\n%v", test.givenInput, result, err)
		}
	}
}

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
				mapper.TypeField{ID: "key", Name: "", Type: "String", Localized: false, Required: false, Disabled: false, Omitted: false},
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
				mapper.TypeField{ID: "key", Name: "", Type: "String", Localized: false, Required: false, Disabled: false, Omitted: false},
				mapper.TypeField{ID: "mainContent", Name: "", Type: "String", Localized: false, Required: false, Disabled: false, Omitted: false},
				mapper.TypeField{ID: "slug", Name: "", Type: "String", Localized: false, Required: false, Disabled: false, Omitted: false},
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

	tc := TranslationContext{}

	for _, test := range tests {
		result := tc.ConvertToContent(tc.MapContentValuesToTypeNames(test.Map, test.fields))
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("_.ConvertToContent( _.MapContentValuesToTypeNames(%v, %v) ) incorrect, expected %v, got %v", test.Map, test.fields, test.expected, result)
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
			mapper.TypeField{ID: "", Name: "", Type: "default", Localized: false, Required: false, Disabled: false, Omitted: false},
			"Unchanged",
		},
		{
			[]interface{}{
				map[string]interface{}{"sys": map[string]interface{}{"id": "test-id-1", "linkType": "Entry"}},
				map[string]interface{}{"sys": map[string]interface{}{"id": "test-id-2", "linkType": "Entry"}},
				map[string]interface{}{"sys": map[string]interface{}{"id": "test-id-3", "linkType": "Entry"}},
			},
			mapper.TypeField{ID: "", Name: "", Type: "Array", Localized: false, Required: false, Disabled: false, Omitted: false},
			[]string{"test-id-1.md", "test-id-2.md", "test-id-3.md"},
		},
	}

	tc := TranslationContext{}

	for _, test := range tests {
		result := tc.translateField(test.value, test.field)

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("translateField(%v, %v) incorrect, expected %v, got %v", test.value, test.field.Type, test.expected, result)
		}
	}

}
