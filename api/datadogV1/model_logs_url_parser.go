// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsURLParser This processor extracts query parameters and other important parameters from a URL.
type LogsURLParser struct {
	// Whether or not the processor is enabled.
	IsEnabled *bool `json:"is_enabled,omitempty"`
	// Name of the processor.
	Name *string `json:"name,omitempty"`
	// Normalize the ending slashes or not.
	NormalizeEndingSlashes datadog.NullableBool `json:"normalize_ending_slashes,omitempty"`
	// Array of source attributes.
	Sources []string `json:"sources"`
	// Name of the parent attribute that contains all the extracted details from the `sources`.
	Target string `json:"target"`
	// Type of logs URL parser.
	Type LogsURLParserType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLogsURLParser instantiates a new LogsURLParser object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogsURLParser(sources []string, target string, typeVar LogsURLParserType) *LogsURLParser {
	this := LogsURLParser{}
	var isEnabled bool = false
	this.IsEnabled = &isEnabled
	var normalizeEndingSlashes bool = false
	this.NormalizeEndingSlashes = *datadog.NewNullableBool(&normalizeEndingSlashes)
	this.Sources = sources
	this.Target = target
	this.Type = typeVar
	return &this
}

// NewLogsURLParserWithDefaults instantiates a new LogsURLParser object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogsURLParserWithDefaults() *LogsURLParser {
	this := LogsURLParser{}
	var isEnabled bool = false
	this.IsEnabled = &isEnabled
	var normalizeEndingSlashes bool = false
	this.NormalizeEndingSlashes = *datadog.NewNullableBool(&normalizeEndingSlashes)
	var target string = "http.url_details"
	this.Target = target
	var typeVar LogsURLParserType = LOGSURLPARSERTYPE_URL_PARSER
	this.Type = typeVar
	return &this
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *LogsURLParser) GetIsEnabled() bool {
	if o == nil || o.IsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsURLParser) GetIsEnabledOk() (*bool, bool) {
	if o == nil || o.IsEnabled == nil {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *LogsURLParser) HasIsEnabled() bool {
	return o != nil && o.IsEnabled != nil
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *LogsURLParser) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LogsURLParser) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsURLParser) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LogsURLParser) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LogsURLParser) SetName(v string) {
	o.Name = &v
}

// GetNormalizeEndingSlashes returns the NormalizeEndingSlashes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogsURLParser) GetNormalizeEndingSlashes() bool {
	if o == nil || o.NormalizeEndingSlashes.Get() == nil {
		var ret bool
		return ret
	}
	return *o.NormalizeEndingSlashes.Get()
}

// GetNormalizeEndingSlashesOk returns a tuple with the NormalizeEndingSlashes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LogsURLParser) GetNormalizeEndingSlashesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NormalizeEndingSlashes.Get(), o.NormalizeEndingSlashes.IsSet()
}

// HasNormalizeEndingSlashes returns a boolean if a field has been set.
func (o *LogsURLParser) HasNormalizeEndingSlashes() bool {
	return o != nil && o.NormalizeEndingSlashes.IsSet()
}

// SetNormalizeEndingSlashes gets a reference to the given datadog.NullableBool and assigns it to the NormalizeEndingSlashes field.
func (o *LogsURLParser) SetNormalizeEndingSlashes(v bool) {
	o.NormalizeEndingSlashes.Set(&v)
}

// SetNormalizeEndingSlashesNil sets the value for NormalizeEndingSlashes to be an explicit nil.
func (o *LogsURLParser) SetNormalizeEndingSlashesNil() {
	o.NormalizeEndingSlashes.Set(nil)
}

// UnsetNormalizeEndingSlashes ensures that no value is present for NormalizeEndingSlashes, not even an explicit nil.
func (o *LogsURLParser) UnsetNormalizeEndingSlashes() {
	o.NormalizeEndingSlashes.Unset()
}

// GetSources returns the Sources field value.
func (o *LogsURLParser) GetSources() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value
// and a boolean to check if the value has been set.
func (o *LogsURLParser) GetSourcesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sources, true
}

// SetSources sets field value.
func (o *LogsURLParser) SetSources(v []string) {
	o.Sources = v
}

// GetTarget returns the Target field value.
func (o *LogsURLParser) GetTarget() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *LogsURLParser) GetTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value.
func (o *LogsURLParser) SetTarget(v string) {
	o.Target = v
}

// GetType returns the Type field value.
func (o *LogsURLParser) GetType() LogsURLParserType {
	if o == nil {
		var ret LogsURLParserType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LogsURLParser) GetTypeOk() (*LogsURLParserType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *LogsURLParser) SetType(v LogsURLParserType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogsURLParser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.IsEnabled != nil {
		toSerialize["is_enabled"] = o.IsEnabled
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.NormalizeEndingSlashes.IsSet() {
		toSerialize["normalize_ending_slashes"] = o.NormalizeEndingSlashes.Get()
	}
	toSerialize["sources"] = o.Sources
	toSerialize["target"] = o.Target
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LogsURLParser) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IsEnabled              *bool                `json:"is_enabled,omitempty"`
		Name                   *string              `json:"name,omitempty"`
		NormalizeEndingSlashes datadog.NullableBool `json:"normalize_ending_slashes,omitempty"`
		Sources                *[]string            `json:"sources"`
		Target                 *string              `json:"target"`
		Type                   *LogsURLParserType   `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Sources == nil {
		return fmt.Errorf("required field sources missing")
	}
	if all.Target == nil {
		return fmt.Errorf("required field target missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"is_enabled", "name", "normalize_ending_slashes", "sources", "target", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.IsEnabled = all.IsEnabled
	o.Name = all.Name
	o.NormalizeEndingSlashes = all.NormalizeEndingSlashes
	o.Sources = *all.Sources
	o.Target = *all.Target
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
