// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListVulnerableAssetsResponse The expected response schema when listing vulnerable assets.
type ListVulnerableAssetsResponse struct {
	// List of vulnerable assets.
	Data []Asset `json:"data"`
	// The JSON:API links related to pagination.
	Links *Links `json:"links,omitempty"`
	// The metadata related to this request.
	Meta *Metadata `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewListVulnerableAssetsResponse instantiates a new ListVulnerableAssetsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewListVulnerableAssetsResponse(data []Asset) *ListVulnerableAssetsResponse {
	this := ListVulnerableAssetsResponse{}
	this.Data = data
	return &this
}

// NewListVulnerableAssetsResponseWithDefaults instantiates a new ListVulnerableAssetsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewListVulnerableAssetsResponseWithDefaults() *ListVulnerableAssetsResponse {
	this := ListVulnerableAssetsResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *ListVulnerableAssetsResponse) GetData() []Asset {
	if o == nil {
		var ret []Asset
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListVulnerableAssetsResponse) GetDataOk() (*[]Asset, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *ListVulnerableAssetsResponse) SetData(v []Asset) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ListVulnerableAssetsResponse) GetLinks() Links {
	if o == nil || o.Links == nil {
		var ret Links
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListVulnerableAssetsResponse) GetLinksOk() (*Links, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ListVulnerableAssetsResponse) HasLinks() bool {
	return o != nil && o.Links != nil
}

// SetLinks gets a reference to the given Links and assigns it to the Links field.
func (o *ListVulnerableAssetsResponse) SetLinks(v Links) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ListVulnerableAssetsResponse) GetMeta() Metadata {
	if o == nil || o.Meta == nil {
		var ret Metadata
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListVulnerableAssetsResponse) GetMetaOk() (*Metadata, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ListVulnerableAssetsResponse) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given Metadata and assigns it to the Meta field.
func (o *ListVulnerableAssetsResponse) SetMeta(v Metadata) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ListVulnerableAssetsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
	if o.Links != nil {
		toSerialize["links"] = o.Links
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
func (o *ListVulnerableAssetsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data  *[]Asset  `json:"data"`
		Links *Links    `json:"links,omitempty"`
		Meta  *Metadata `json:"meta,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "links", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Data = *all.Data
	if all.Links != nil && all.Links.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Links = all.Links
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
