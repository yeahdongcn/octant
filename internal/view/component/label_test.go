package component_test

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/heptio/developer-dash/internal/view/component"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Labels_Marshal(t *testing.T) {
	cases := []struct {
		name         string
		input        *component.Labels
		expectedPath string
	}{
		{
			name: "in general",
			input: component.NewLabels(map[string]string{
				"foo": "bar",
			}),
			expectedPath: "labels.json",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := json.Marshal(tc.input)
			require.NoError(t, err)

			expected, err := ioutil.ReadFile(filepath.Join("testdata", tc.expectedPath))
			require.NoError(t, err)

			assert.JSONEq(t, string(expected), string(got))
		})
	}
}

func Test_Labels_IsEmpty(t *testing.T) {
	cases := []struct {
		name    string
		input   *component.Labels
		isEmpty bool
	}{
		{
			name:    "empty (nil)",
			input:   component.NewLabels(nil),
			isEmpty: true,
		},
		{
			name:    "empty",
			input:   component.NewLabels(map[string]string{}),
			isEmpty: true,
		},
		{
			name: "not empty",
			input: component.NewLabels(map[string]string{
				"foo": "bar",
			}),
			isEmpty: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.isEmpty, tc.input.IsEmpty())
		})
	}
}

func Test_Labels_GetMetadata(t *testing.T) {
	input := component.NewLabels(map[string]string{
		"foo": "bar",
	})

	assert.Equal(t, "labels", input.GetMetadata().Type)
}

func Test_Labels_filter(t *testing.T) {
	input := component.NewLabels(map[string]string{
		"foo":                      "bar",
		"controller-uid":           "uid",
		"controller-revision-hash": "hash",
		"pod-template-generation":  "generation",
	})

	got, err := input.MarshalJSON()
	require.NoError(t, err)

	expected, err := ioutil.ReadFile(filepath.Join("testdata", "labels_filtered.json"))
	require.NoError(t, err)

	assert.JSONEq(t, string(expected), string(got))
}
