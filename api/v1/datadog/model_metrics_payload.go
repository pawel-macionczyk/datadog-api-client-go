/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"bytes"
	"encoding/json"
)

// MetricsPayload struct for MetricsPayload
type MetricsPayload struct {
	// A list of time series to submit to Datadog
	Series *[]Series `json:"series,omitempty"`
}

// NewMetricsPayload instantiates a new MetricsPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsPayload() *MetricsPayload {
	this := MetricsPayload{}
	return &this
}

// NewMetricsPayloadWithDefaults instantiates a new MetricsPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsPayloadWithDefaults() *MetricsPayload {
	this := MetricsPayload{}
	return &this
}

// GetSeries returns the Series field value if set, zero value otherwise.
func (o *MetricsPayload) GetSeries() []Series {
	if o == nil || o.Series == nil {
		var ret []Series
		return ret
	}
	return *o.Series
}

// GetSeriesOk returns a tuple with the Series field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MetricsPayload) GetSeriesOk() ([]Series, bool) {
	if o == nil || o.Series == nil {
		var ret []Series
		return ret, false
	}
	return *o.Series, true
}

// HasSeries returns a boolean if a field has been set.
func (o *MetricsPayload) HasSeries() bool {
	if o != nil && o.Series != nil {
		return true
	}

	return false
}

// SetSeries gets a reference to the given []Series and assigns it to the Series field.
func (o *MetricsPayload) SetSeries(v []Series) {
	o.Series = &v
}

type NullableMetricsPayload struct {
	Value        MetricsPayload
	ExplicitNull bool
}

func (v NullableMetricsPayload) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableMetricsPayload) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
