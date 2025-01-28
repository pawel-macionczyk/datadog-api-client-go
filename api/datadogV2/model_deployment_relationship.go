// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeploymentRelationship Information pointing to the app's publication status.
type DeploymentRelationship struct {
	// Data object containing the deployment ID.
	Data *DeploymentRelationshipData `json:"data,omitempty"`
	// Metadata object containing the publication creation information.
	Meta *DeploymentMetadata `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDeploymentRelationship instantiates a new DeploymentRelationship object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeploymentRelationship() *DeploymentRelationship {
	this := DeploymentRelationship{}
	return &this
}

// NewDeploymentRelationshipWithDefaults instantiates a new DeploymentRelationship object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDeploymentRelationshipWithDefaults() *DeploymentRelationship {
	this := DeploymentRelationship{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DeploymentRelationship) GetData() DeploymentRelationshipData {
	if o == nil || o.Data == nil {
		var ret DeploymentRelationshipData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRelationship) GetDataOk() (*DeploymentRelationshipData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DeploymentRelationship) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given DeploymentRelationshipData and assigns it to the Data field.
func (o *DeploymentRelationship) SetData(v DeploymentRelationshipData) {
	o.Data = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *DeploymentRelationship) GetMeta() DeploymentMetadata {
	if o == nil || o.Meta == nil {
		var ret DeploymentMetadata
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentRelationship) GetMetaOk() (*DeploymentMetadata, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *DeploymentRelationship) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given DeploymentMetadata and assigns it to the Meta field.
func (o *DeploymentRelationship) SetMeta(v DeploymentMetadata) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DeploymentRelationship) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DeploymentRelationship) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *DeploymentRelationshipData `json:"data,omitempty"`
		Meta *DeploymentMetadata         `json:"meta,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Data != nil && all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Data = all.Data
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = all.Meta

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
