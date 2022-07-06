/*
TrueNAS RESTful API

Go SDK for interacting with TrueNAS APIs (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package truenas

import (
	"encoding/json"
)

// CreateISCSIInitiatorParams struct for CreateISCSIInitiatorParams
type CreateISCSIInitiatorParams struct {
	Initiators []string `json:"initiators,omitempty"`
	AuthNetwork []string `json:"auth_network,omitempty"`
	Comment *string `json:"comment,omitempty"`
}

// NewCreateISCSIInitiatorParams instantiates a new CreateISCSIInitiatorParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateISCSIInitiatorParams() *CreateISCSIInitiatorParams {
	this := CreateISCSIInitiatorParams{}
	return &this
}

// NewCreateISCSIInitiatorParamsWithDefaults instantiates a new CreateISCSIInitiatorParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateISCSIInitiatorParamsWithDefaults() *CreateISCSIInitiatorParams {
	this := CreateISCSIInitiatorParams{}
	return &this
}

// GetInitiators returns the Initiators field value if set, zero value otherwise.
func (o *CreateISCSIInitiatorParams) GetInitiators() []string {
	if o == nil || o.Initiators == nil {
		var ret []string
		return ret
	}
	return o.Initiators
}

// GetInitiatorsOk returns a tuple with the Initiators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateISCSIInitiatorParams) GetInitiatorsOk() ([]string, bool) {
	if o == nil || o.Initiators == nil {
		return nil, false
	}
	return o.Initiators, true
}

// HasInitiators returns a boolean if a field has been set.
func (o *CreateISCSIInitiatorParams) HasInitiators() bool {
	if o != nil && o.Initiators != nil {
		return true
	}

	return false
}

// SetInitiators gets a reference to the given []string and assigns it to the Initiators field.
func (o *CreateISCSIInitiatorParams) SetInitiators(v []string) {
	o.Initiators = v
}

// GetAuthNetwork returns the AuthNetwork field value if set, zero value otherwise.
func (o *CreateISCSIInitiatorParams) GetAuthNetwork() []string {
	if o == nil || o.AuthNetwork == nil {
		var ret []string
		return ret
	}
	return o.AuthNetwork
}

// GetAuthNetworkOk returns a tuple with the AuthNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateISCSIInitiatorParams) GetAuthNetworkOk() ([]string, bool) {
	if o == nil || o.AuthNetwork == nil {
		return nil, false
	}
	return o.AuthNetwork, true
}

// HasAuthNetwork returns a boolean if a field has been set.
func (o *CreateISCSIInitiatorParams) HasAuthNetwork() bool {
	if o != nil && o.AuthNetwork != nil {
		return true
	}

	return false
}

// SetAuthNetwork gets a reference to the given []string and assigns it to the AuthNetwork field.
func (o *CreateISCSIInitiatorParams) SetAuthNetwork(v []string) {
	o.AuthNetwork = v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *CreateISCSIInitiatorParams) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateISCSIInitiatorParams) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *CreateISCSIInitiatorParams) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *CreateISCSIInitiatorParams) SetComment(v string) {
	o.Comment = &v
}

func (o CreateISCSIInitiatorParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Initiators != nil {
		toSerialize["initiators"] = o.Initiators
	}
	if o.AuthNetwork != nil {
		toSerialize["auth_network"] = o.AuthNetwork
	}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	return json.Marshal(toSerialize)
}

type NullableCreateISCSIInitiatorParams struct {
	value *CreateISCSIInitiatorParams
	isSet bool
}

func (v NullableCreateISCSIInitiatorParams) Get() *CreateISCSIInitiatorParams {
	return v.value
}

func (v *NullableCreateISCSIInitiatorParams) Set(val *CreateISCSIInitiatorParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateISCSIInitiatorParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateISCSIInitiatorParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateISCSIInitiatorParams(val *CreateISCSIInitiatorParams) *NullableCreateISCSIInitiatorParams {
	return &NullableCreateISCSIInitiatorParams{value: val, isSet: true}
}

func (v NullableCreateISCSIInitiatorParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateISCSIInitiatorParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


