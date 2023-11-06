// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/goccy/go-json"
)

// ToplistWidgetLegend Top list widget stacked legend behavior.
type ToplistWidgetLegend string

// List of ToplistWidgetLegend.
const (
	TOPLISTWIDGETLEGEND_AUTOMATIC ToplistWidgetLegend = "automatic"
	TOPLISTWIDGETLEGEND_INLINE    ToplistWidgetLegend = "inline"
	TOPLISTWIDGETLEGEND_NONE      ToplistWidgetLegend = "none"
)

var allowedToplistWidgetLegendEnumValues = []ToplistWidgetLegend{
	TOPLISTWIDGETLEGEND_AUTOMATIC,
	TOPLISTWIDGETLEGEND_INLINE,
	TOPLISTWIDGETLEGEND_NONE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ToplistWidgetLegend) GetAllowedValues() []ToplistWidgetLegend {
	return allowedToplistWidgetLegendEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ToplistWidgetLegend) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ToplistWidgetLegend(value)
	return nil
}

// NewToplistWidgetLegendFromValue returns a pointer to a valid ToplistWidgetLegend
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewToplistWidgetLegendFromValue(v string) (*ToplistWidgetLegend, error) {
	ev := ToplistWidgetLegend(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ToplistWidgetLegend: valid values are %v", v, allowedToplistWidgetLegendEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ToplistWidgetLegend) IsValid() bool {
	for _, existing := range allowedToplistWidgetLegendEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ToplistWidgetLegend value.
func (v ToplistWidgetLegend) Ptr() *ToplistWidgetLegend {
	return &v
}
