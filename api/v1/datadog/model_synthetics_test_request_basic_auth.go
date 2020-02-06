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

// SyntheticsTestRequestBasicAuth struct for SyntheticsTestRequestBasicAuth
type SyntheticsTestRequestBasicAuth struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// NewSyntheticsTestRequestBasicAuth instantiates a new SyntheticsTestRequestBasicAuth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticsTestRequestBasicAuth(password string, username string) *SyntheticsTestRequestBasicAuth {
	this := SyntheticsTestRequestBasicAuth{}
	this.Password = password
	this.Username = username
	return &this
}

// NewSyntheticsTestRequestBasicAuthWithDefaults instantiates a new SyntheticsTestRequestBasicAuth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticsTestRequestBasicAuthWithDefaults() *SyntheticsTestRequestBasicAuth {
	this := SyntheticsTestRequestBasicAuth{}
	return &this
}

// GetPassword returns the Password field value
func (o *SyntheticsTestRequestBasicAuth) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// SetPassword sets field value
func (o *SyntheticsTestRequestBasicAuth) SetPassword(v string) {
	o.Password = v
}

// GetUsername returns the Username field value
func (o *SyntheticsTestRequestBasicAuth) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// SetUsername sets field value
func (o *SyntheticsTestRequestBasicAuth) SetUsername(v string) {
	o.Username = v
}

type NullableSyntheticsTestRequestBasicAuth struct {
	Value        SyntheticsTestRequestBasicAuth
	ExplicitNull bool
}

func (v NullableSyntheticsTestRequestBasicAuth) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableSyntheticsTestRequestBasicAuth) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
