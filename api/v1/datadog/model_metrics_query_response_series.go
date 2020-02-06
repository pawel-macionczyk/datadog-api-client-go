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

// MetricsQueryResponseSeries struct for MetricsQueryResponseSeries
type MetricsQueryResponseSeries struct {
	// Aggregation type
	Aggr *string `json:"aggr,omitempty"`
	// Display name of the metric
	DisplayName *string `json:"display_name,omitempty"`
	// End of the time window, milliseconds since Unix epoch
	End *int64 `json:"end,omitempty"`
	// Metric expression
	Expression *string `json:"expression,omitempty"`
	// Number of seconds between data samples
	Interval *int64 `json:"interval,omitempty"`
	// Number of data samples
	Length *int64 `json:"length,omitempty"`
	// Metric name
	Metric *string `json:"metric,omitempty"`
	// List of points of the time series
	Pointlist *[][]float64 `json:"pointlist,omitempty"`
	// Metric scope, comma separated list of tags
	Scope *string `json:"scope,omitempty"`
	// Start of the time window, milliseconds since Unix epoch
	Start *int64 `json:"start,omitempty"`
	// Detailed information about the metric unit. First element describes the \"primary unit\" (e.g. `bytes` in `bytes per second`), second describes the \"per unit\" (e.g. `second` in `bytes per second`)
	Unit *[]MetricsQueryResponseUnit `json:"unit,omitempty"`
}

// NewMetricsQueryResponseSeries instantiates a new MetricsQueryResponseSeries object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsQueryResponseSeries() *MetricsQueryResponseSeries {
	this := MetricsQueryResponseSeries{}
	return &this
}

// NewMetricsQueryResponseSeriesWithDefaults instantiates a new MetricsQueryResponseSeries object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsQueryResponseSeriesWithDefaults() *MetricsQueryResponseSeries {
	this := MetricsQueryResponseSeries{}
	return &this
}

// GetAggr returns the Aggr field value if set, zero value otherwise.
func (o *MetricsQueryResponseSeries) GetAggr() string {
	if o == nil || o.Aggr == nil {
		var ret string
		return ret
	}
	return *o.Aggr
}

// GetAggrOk returns a tuple with the Aggr field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MetricsQueryResponseSeries) GetAggrOk() (string, bool) {
	if o == nil || o.Aggr == nil {
		var ret string
		return ret, false
	}
	return *o.Aggr, true
}

// HasAggr returns a boolean if a field has been set.
func (o *MetricsQueryResponseSeries) HasAggr() bool {
	if o != nil && o.Aggr != nil {
		return true
	}

	return false
}

// SetAggr gets a reference to the given string and assigns it to the Aggr field.
func (o *MetricsQueryResponseSeries) SetAggr(v string) {
	o.Aggr = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *MetricsQueryResponseSeries) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MetricsQueryResponseSeries) GetDisplayNameOk() (string, bool) {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret, false
	}
	return *o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MetricsQueryResponseSeries) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *MetricsQueryResponseSeries) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *MetricsQueryResponseSeries) GetEnd() int64 {
	if o == nil || o.End == nil {
		var ret int64
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MetricsQueryResponseSeries) GetEndOk() (int64, bool) {
	if o == nil || o.End == nil {
		var ret int64
		return ret, false
	}
	return *o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *MetricsQueryResponseSeries) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int64 and assigns it to the End field.
func (o *MetricsQueryResponseSeries) SetEnd(v int64) {
	o.End = &v
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *MetricsQueryResponseSeries) GetExpression() string {
	if o == nil || o.Expression == nil {
		var ret string
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MetricsQueryResponseSeries) GetExpressionOk() (string, bool) {
	if o == nil || o.Expression == nil {
		var ret string
		return ret, false
	}
	return *o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *MetricsQueryResponseSeries) HasExpression() bool {
	if o != nil && o.Expression != nil {
		return true
	}

	return false
}

// SetExpression gets a reference to the given string and assigns it to the Expression field.
func (o *MetricsQueryResponseSeries) SetExpression(v string) {
	o.Expression = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *MetricsQueryResponseSeries) GetInterval() int64 {
	if o == nil || o.Interval == nil {
		var ret int64
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MetricsQueryResponseSeries) GetIntervalOk() (int64, bool) {
	if o == nil || o.Interval == nil {
		var ret int64
		return ret, false
	}
	return *o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *MetricsQueryResponseSeries) HasInterval() bool {
	if o != nil && o.Interval != nil {
		return true
	}

	return false
}

// SetInterval gets a reference to the given int64 and assigns it to the Interval field.
func (o *MetricsQueryResponseSeries) SetInterval(v int64) {
	o.Interval = &v
}

// GetLength returns the Length field value if set, zero value otherwise.
func (o *MetricsQueryResponseSeries) GetLength() int64 {
	if o == nil || o.Length == nil {
		var ret int64
		return ret
	}
	return *o.Length
}

// GetLengthOk returns a tuple with the Length field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MetricsQueryResponseSeries) GetLengthOk() (int64, bool) {
	if o == nil || o.Length == nil {
		var ret int64
		return ret, false
	}
	return *o.Length, true
}

// HasLength returns a boolean if a field has been set.
func (o *MetricsQueryResponseSeries) HasLength() bool {
	if o != nil && o.Length != nil {
		return true
	}

	return false
}

// SetLength gets a reference to the given int64 and assigns it to the Length field.
func (o *MetricsQueryResponseSeries) SetLength(v int64) {
	o.Length = &v
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *MetricsQueryResponseSeries) GetMetric() string {
	if o == nil || o.Metric == nil {
		var ret string
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MetricsQueryResponseSeries) GetMetricOk() (string, bool) {
	if o == nil || o.Metric == nil {
		var ret string
		return ret, false
	}
	return *o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *MetricsQueryResponseSeries) HasMetric() bool {
	if o != nil && o.Metric != nil {
		return true
	}

	return false
}

// SetMetric gets a reference to the given string and assigns it to the Metric field.
func (o *MetricsQueryResponseSeries) SetMetric(v string) {
	o.Metric = &v
}

// GetPointlist returns the Pointlist field value if set, zero value otherwise.
func (o *MetricsQueryResponseSeries) GetPointlist() [][]float64 {
	if o == nil || o.Pointlist == nil {
		var ret [][]float64
		return ret
	}
	return *o.Pointlist
}

// GetPointlistOk returns a tuple with the Pointlist field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MetricsQueryResponseSeries) GetPointlistOk() ([][]float64, bool) {
	if o == nil || o.Pointlist == nil {
		var ret [][]float64
		return ret, false
	}
	return *o.Pointlist, true
}

// HasPointlist returns a boolean if a field has been set.
func (o *MetricsQueryResponseSeries) HasPointlist() bool {
	if o != nil && o.Pointlist != nil {
		return true
	}

	return false
}

// SetPointlist gets a reference to the given [][]float64 and assigns it to the Pointlist field.
func (o *MetricsQueryResponseSeries) SetPointlist(v [][]float64) {
	o.Pointlist = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *MetricsQueryResponseSeries) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MetricsQueryResponseSeries) GetScopeOk() (string, bool) {
	if o == nil || o.Scope == nil {
		var ret string
		return ret, false
	}
	return *o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *MetricsQueryResponseSeries) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *MetricsQueryResponseSeries) SetScope(v string) {
	o.Scope = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *MetricsQueryResponseSeries) GetStart() int64 {
	if o == nil || o.Start == nil {
		var ret int64
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MetricsQueryResponseSeries) GetStartOk() (int64, bool) {
	if o == nil || o.Start == nil {
		var ret int64
		return ret, false
	}
	return *o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *MetricsQueryResponseSeries) HasStart() bool {
	if o != nil && o.Start != nil {
		return true
	}

	return false
}

// SetStart gets a reference to the given int64 and assigns it to the Start field.
func (o *MetricsQueryResponseSeries) SetStart(v int64) {
	o.Start = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *MetricsQueryResponseSeries) GetUnit() []MetricsQueryResponseUnit {
	if o == nil || o.Unit == nil {
		var ret []MetricsQueryResponseUnit
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *MetricsQueryResponseSeries) GetUnitOk() ([]MetricsQueryResponseUnit, bool) {
	if o == nil || o.Unit == nil {
		var ret []MetricsQueryResponseUnit
		return ret, false
	}
	return *o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *MetricsQueryResponseSeries) HasUnit() bool {
	if o != nil && o.Unit != nil {
		return true
	}

	return false
}

// SetUnit gets a reference to the given []MetricsQueryResponseUnit and assigns it to the Unit field.
func (o *MetricsQueryResponseSeries) SetUnit(v []MetricsQueryResponseUnit) {
	o.Unit = &v
}

type NullableMetricsQueryResponseSeries struct {
	Value        MetricsQueryResponseSeries
	ExplicitNull bool
}

func (v NullableMetricsQueryResponseSeries) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableMetricsQueryResponseSeries) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
