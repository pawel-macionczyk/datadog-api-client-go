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

// SyntheticsSslCertificateSubject struct for SyntheticsSslCertificateSubject
type SyntheticsSslCertificateSubject struct {
	C       *string `json:"C,omitempty"`
	CN      *string `json:"CN,omitempty"`
	L       *string `json:"L,omitempty"`
	O       *string `json:"O,omitempty"`
	OU      *string `json:"OU,omitempty"`
	ST      *string `json:"ST,omitempty"`
	AltName *string `json:"altName,omitempty"`
}

// NewSyntheticsSslCertificateSubject instantiates a new SyntheticsSslCertificateSubject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticsSslCertificateSubject() *SyntheticsSslCertificateSubject {
	this := SyntheticsSslCertificateSubject{}
	return &this
}

// NewSyntheticsSslCertificateSubjectWithDefaults instantiates a new SyntheticsSslCertificateSubject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticsSslCertificateSubjectWithDefaults() *SyntheticsSslCertificateSubject {
	this := SyntheticsSslCertificateSubject{}
	return &this
}

// GetC returns the C field value if set, zero value otherwise.
func (o *SyntheticsSslCertificateSubject) GetC() string {
	if o == nil || o.C == nil {
		var ret string
		return ret
	}
	return *o.C
}

// GetCOk returns a tuple with the C field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsSslCertificateSubject) GetCOk() (string, bool) {
	if o == nil || o.C == nil {
		var ret string
		return ret, false
	}
	return *o.C, true
}

// HasC returns a boolean if a field has been set.
func (o *SyntheticsSslCertificateSubject) HasC() bool {
	if o != nil && o.C != nil {
		return true
	}

	return false
}

// SetC gets a reference to the given string and assigns it to the C field.
func (o *SyntheticsSslCertificateSubject) SetC(v string) {
	o.C = &v
}

// GetCN returns the CN field value if set, zero value otherwise.
func (o *SyntheticsSslCertificateSubject) GetCN() string {
	if o == nil || o.CN == nil {
		var ret string
		return ret
	}
	return *o.CN
}

// GetCNOk returns a tuple with the CN field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsSslCertificateSubject) GetCNOk() (string, bool) {
	if o == nil || o.CN == nil {
		var ret string
		return ret, false
	}
	return *o.CN, true
}

// HasCN returns a boolean if a field has been set.
func (o *SyntheticsSslCertificateSubject) HasCN() bool {
	if o != nil && o.CN != nil {
		return true
	}

	return false
}

// SetCN gets a reference to the given string and assigns it to the CN field.
func (o *SyntheticsSslCertificateSubject) SetCN(v string) {
	o.CN = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *SyntheticsSslCertificateSubject) GetL() string {
	if o == nil || o.L == nil {
		var ret string
		return ret
	}
	return *o.L
}

// GetLOk returns a tuple with the L field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsSslCertificateSubject) GetLOk() (string, bool) {
	if o == nil || o.L == nil {
		var ret string
		return ret, false
	}
	return *o.L, true
}

// HasL returns a boolean if a field has been set.
func (o *SyntheticsSslCertificateSubject) HasL() bool {
	if o != nil && o.L != nil {
		return true
	}

	return false
}

// SetL gets a reference to the given string and assigns it to the L field.
func (o *SyntheticsSslCertificateSubject) SetL(v string) {
	o.L = &v
}

// GetO returns the O field value if set, zero value otherwise.
func (o *SyntheticsSslCertificateSubject) GetO() string {
	if o == nil || o.O == nil {
		var ret string
		return ret
	}
	return *o.O
}

// GetOOk returns a tuple with the O field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsSslCertificateSubject) GetOOk() (string, bool) {
	if o == nil || o.O == nil {
		var ret string
		return ret, false
	}
	return *o.O, true
}

// HasO returns a boolean if a field has been set.
func (o *SyntheticsSslCertificateSubject) HasO() bool {
	if o != nil && o.O != nil {
		return true
	}

	return false
}

// SetO gets a reference to the given string and assigns it to the O field.
func (o *SyntheticsSslCertificateSubject) SetO(v string) {
	o.O = &v
}

// GetOU returns the OU field value if set, zero value otherwise.
func (o *SyntheticsSslCertificateSubject) GetOU() string {
	if o == nil || o.OU == nil {
		var ret string
		return ret
	}
	return *o.OU
}

// GetOUOk returns a tuple with the OU field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsSslCertificateSubject) GetOUOk() (string, bool) {
	if o == nil || o.OU == nil {
		var ret string
		return ret, false
	}
	return *o.OU, true
}

// HasOU returns a boolean if a field has been set.
func (o *SyntheticsSslCertificateSubject) HasOU() bool {
	if o != nil && o.OU != nil {
		return true
	}

	return false
}

// SetOU gets a reference to the given string and assigns it to the OU field.
func (o *SyntheticsSslCertificateSubject) SetOU(v string) {
	o.OU = &v
}

// GetST returns the ST field value if set, zero value otherwise.
func (o *SyntheticsSslCertificateSubject) GetST() string {
	if o == nil || o.ST == nil {
		var ret string
		return ret
	}
	return *o.ST
}

// GetSTOk returns a tuple with the ST field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsSslCertificateSubject) GetSTOk() (string, bool) {
	if o == nil || o.ST == nil {
		var ret string
		return ret, false
	}
	return *o.ST, true
}

// HasST returns a boolean if a field has been set.
func (o *SyntheticsSslCertificateSubject) HasST() bool {
	if o != nil && o.ST != nil {
		return true
	}

	return false
}

// SetST gets a reference to the given string and assigns it to the ST field.
func (o *SyntheticsSslCertificateSubject) SetST(v string) {
	o.ST = &v
}

// GetAltName returns the AltName field value if set, zero value otherwise.
func (o *SyntheticsSslCertificateSubject) GetAltName() string {
	if o == nil || o.AltName == nil {
		var ret string
		return ret
	}
	return *o.AltName
}

// GetAltNameOk returns a tuple with the AltName field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsSslCertificateSubject) GetAltNameOk() (string, bool) {
	if o == nil || o.AltName == nil {
		var ret string
		return ret, false
	}
	return *o.AltName, true
}

// HasAltName returns a boolean if a field has been set.
func (o *SyntheticsSslCertificateSubject) HasAltName() bool {
	if o != nil && o.AltName != nil {
		return true
	}

	return false
}

// SetAltName gets a reference to the given string and assigns it to the AltName field.
func (o *SyntheticsSslCertificateSubject) SetAltName(v string) {
	o.AltName = &v
}

type NullableSyntheticsSslCertificateSubject struct {
	Value        SyntheticsSslCertificateSubject
	ExplicitNull bool
}

func (v NullableSyntheticsSslCertificateSubject) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableSyntheticsSslCertificateSubject) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
