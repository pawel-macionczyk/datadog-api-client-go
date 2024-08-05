// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DORAIncidentRequestData The JSON:API data.
type DORAIncidentRequestData struct {
	// Attributes to create a DORA incident event.
	Attributes DORAIncidentRequestAttributes `json:"attributes"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDORAIncidentRequestData instantiates a new DORAIncidentRequestData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDORAIncidentRequestData(attributes DORAIncidentRequestAttributes) *DORAIncidentRequestData {
	this := DORAIncidentRequestData{}
	this.Attributes = attributes
	return &this
}

// NewDORAIncidentRequestDataWithDefaults instantiates a new DORAIncidentRequestData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDORAIncidentRequestDataWithDefaults() *DORAIncidentRequestData {
	this := DORAIncidentRequestData{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *DORAIncidentRequestData) GetAttributes() DORAIncidentRequestAttributes {
	if o == nil {
		var ret DORAIncidentRequestAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *DORAIncidentRequestData) GetAttributesOk() (*DORAIncidentRequestAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *DORAIncidentRequestData) SetAttributes(v DORAIncidentRequestAttributes) {
	o.Attributes = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DORAIncidentRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DORAIncidentRequestData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *DORAIncidentRequestAttributes `json:"attributes"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
