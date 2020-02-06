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

// HistoryServiceLevelObjectiveOverall struct for HistoryServiceLevelObjectiveOverall
type HistoryServiceLevelObjectiveOverall struct {
	// For `monitor` based SLOs, this includes the aggregated history uptime time series.
	History *[][]float64 `json:"history,omitempty"`
	// For `monitor` based SLOs this represents the overall group.
	Name *string `json:"name,omitempty"`
	// A mapping of threshold `timeframe` to number of accurate decimals, regardless of the from && to timestamp.
	Precision *map[string]float64 `json:"precision,omitempty"`
	// For `monitor` based SLOs when `true` this indicates that a replay is in progress to give an accurate uptime calculation.
	Preview *bool `json:"preview,omitempty"`
	// The amount of decimal places the uptime value is accurate to for the given from and to timestamp.
	SpanPrecision *float64 `json:"span_precision,omitempty"`
	// The uptime value of the SLO history window.
	Uptime *float64 `json:"uptime,omitempty"`
}

// NewHistoryServiceLevelObjectiveOverall instantiates a new HistoryServiceLevelObjectiveOverall object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoryServiceLevelObjectiveOverall() *HistoryServiceLevelObjectiveOverall {
	this := HistoryServiceLevelObjectiveOverall{}
	return &this
}

// NewHistoryServiceLevelObjectiveOverallWithDefaults instantiates a new HistoryServiceLevelObjectiveOverall object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoryServiceLevelObjectiveOverallWithDefaults() *HistoryServiceLevelObjectiveOverall {
	this := HistoryServiceLevelObjectiveOverall{}
	return &this
}

// GetHistory returns the History field value if set, zero value otherwise.
func (o *HistoryServiceLevelObjectiveOverall) GetHistory() [][]float64 {
	if o == nil || o.History == nil {
		var ret [][]float64
		return ret
	}
	return *o.History
}

// GetHistoryOk returns a tuple with the History field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *HistoryServiceLevelObjectiveOverall) GetHistoryOk() ([][]float64, bool) {
	if o == nil || o.History == nil {
		var ret [][]float64
		return ret, false
	}
	return *o.History, true
}

// HasHistory returns a boolean if a field has been set.
func (o *HistoryServiceLevelObjectiveOverall) HasHistory() bool {
	if o != nil && o.History != nil {
		return true
	}

	return false
}

// SetHistory gets a reference to the given [][]float64 and assigns it to the History field.
func (o *HistoryServiceLevelObjectiveOverall) SetHistory(v [][]float64) {
	o.History = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HistoryServiceLevelObjectiveOverall) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *HistoryServiceLevelObjectiveOverall) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HistoryServiceLevelObjectiveOverall) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HistoryServiceLevelObjectiveOverall) SetName(v string) {
	o.Name = &v
}

// GetPrecision returns the Precision field value if set, zero value otherwise.
func (o *HistoryServiceLevelObjectiveOverall) GetPrecision() map[string]float64 {
	if o == nil || o.Precision == nil {
		var ret map[string]float64
		return ret
	}
	return *o.Precision
}

// GetPrecisionOk returns a tuple with the Precision field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *HistoryServiceLevelObjectiveOverall) GetPrecisionOk() (map[string]float64, bool) {
	if o == nil || o.Precision == nil {
		var ret map[string]float64
		return ret, false
	}
	return *o.Precision, true
}

// HasPrecision returns a boolean if a field has been set.
func (o *HistoryServiceLevelObjectiveOverall) HasPrecision() bool {
	if o != nil && o.Precision != nil {
		return true
	}

	return false
}

// SetPrecision gets a reference to the given map[string]float64 and assigns it to the Precision field.
func (o *HistoryServiceLevelObjectiveOverall) SetPrecision(v map[string]float64) {
	o.Precision = &v
}

// GetPreview returns the Preview field value if set, zero value otherwise.
func (o *HistoryServiceLevelObjectiveOverall) GetPreview() bool {
	if o == nil || o.Preview == nil {
		var ret bool
		return ret
	}
	return *o.Preview
}

// GetPreviewOk returns a tuple with the Preview field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *HistoryServiceLevelObjectiveOverall) GetPreviewOk() (bool, bool) {
	if o == nil || o.Preview == nil {
		var ret bool
		return ret, false
	}
	return *o.Preview, true
}

// HasPreview returns a boolean if a field has been set.
func (o *HistoryServiceLevelObjectiveOverall) HasPreview() bool {
	if o != nil && o.Preview != nil {
		return true
	}

	return false
}

// SetPreview gets a reference to the given bool and assigns it to the Preview field.
func (o *HistoryServiceLevelObjectiveOverall) SetPreview(v bool) {
	o.Preview = &v
}

// GetSpanPrecision returns the SpanPrecision field value if set, zero value otherwise.
func (o *HistoryServiceLevelObjectiveOverall) GetSpanPrecision() float64 {
	if o == nil || o.SpanPrecision == nil {
		var ret float64
		return ret
	}
	return *o.SpanPrecision
}

// GetSpanPrecisionOk returns a tuple with the SpanPrecision field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *HistoryServiceLevelObjectiveOverall) GetSpanPrecisionOk() (float64, bool) {
	if o == nil || o.SpanPrecision == nil {
		var ret float64
		return ret, false
	}
	return *o.SpanPrecision, true
}

// HasSpanPrecision returns a boolean if a field has been set.
func (o *HistoryServiceLevelObjectiveOverall) HasSpanPrecision() bool {
	if o != nil && o.SpanPrecision != nil {
		return true
	}

	return false
}

// SetSpanPrecision gets a reference to the given float64 and assigns it to the SpanPrecision field.
func (o *HistoryServiceLevelObjectiveOverall) SetSpanPrecision(v float64) {
	o.SpanPrecision = &v
}

// GetUptime returns the Uptime field value if set, zero value otherwise.
func (o *HistoryServiceLevelObjectiveOverall) GetUptime() float64 {
	if o == nil || o.Uptime == nil {
		var ret float64
		return ret
	}
	return *o.Uptime
}

// GetUptimeOk returns a tuple with the Uptime field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *HistoryServiceLevelObjectiveOverall) GetUptimeOk() (float64, bool) {
	if o == nil || o.Uptime == nil {
		var ret float64
		return ret, false
	}
	return *o.Uptime, true
}

// HasUptime returns a boolean if a field has been set.
func (o *HistoryServiceLevelObjectiveOverall) HasUptime() bool {
	if o != nil && o.Uptime != nil {
		return true
	}

	return false
}

// SetUptime gets a reference to the given float64 and assigns it to the Uptime field.
func (o *HistoryServiceLevelObjectiveOverall) SetUptime(v float64) {
	o.Uptime = &v
}

type NullableHistoryServiceLevelObjectiveOverall struct {
	Value        HistoryServiceLevelObjectiveOverall
	ExplicitNull bool
}

func (v NullableHistoryServiceLevelObjectiveOverall) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableHistoryServiceLevelObjectiveOverall) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
