/*
TrueNAS RESTful API

Go SDK for interacting with TrueNAS APIs (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package truenas

import (
	"encoding/json"
	"fmt"
)

// checks if the ISCSIInitiator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ISCSIInitiator{}

// ISCSIInitiator struct for ISCSIInitiator
type ISCSIInitiator struct {
	Id          int32                    `json:"id"`
	Initiators  []map[string]interface{} `json:"initiators"`
	AuthNetwork []string                 `json:"auth_network"`
	Comment     string                   `json:"comment"`
}

type _ISCSIInitiator ISCSIInitiator

// NewISCSIInitiator instantiates a new ISCSIInitiator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewISCSIInitiator(id int32, initiators []map[string]interface{}, authNetwork []string, comment string) *ISCSIInitiator {
	this := ISCSIInitiator{}
	this.Id = id
	this.Initiators = initiators
	this.AuthNetwork = authNetwork
	this.Comment = comment
	return &this
}

// NewISCSIInitiatorWithDefaults instantiates a new ISCSIInitiator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewISCSIInitiatorWithDefaults() *ISCSIInitiator {
	this := ISCSIInitiator{}
	return &this
}

// GetId returns the Id field value
func (o *ISCSIInitiator) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ISCSIInitiator) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ISCSIInitiator) SetId(v int32) {
	o.Id = v
}

// GetInitiators returns the Initiators field value
func (o *ISCSIInitiator) GetInitiators() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Initiators
}

// GetInitiatorsOk returns a tuple with the Initiators field value
// and a boolean to check if the value has been set.
func (o *ISCSIInitiator) GetInitiatorsOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Initiators, true
}

// SetInitiators sets field value
func (o *ISCSIInitiator) SetInitiators(v []map[string]interface{}) {
	o.Initiators = v
}

// GetAuthNetwork returns the AuthNetwork field value
func (o *ISCSIInitiator) GetAuthNetwork() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AuthNetwork
}

// GetAuthNetworkOk returns a tuple with the AuthNetwork field value
// and a boolean to check if the value has been set.
func (o *ISCSIInitiator) GetAuthNetworkOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthNetwork, true
}

// SetAuthNetwork sets field value
func (o *ISCSIInitiator) SetAuthNetwork(v []string) {
	o.AuthNetwork = v
}

// GetComment returns the Comment field value
func (o *ISCSIInitiator) GetComment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Comment
}

// GetCommentOk returns a tuple with the Comment field value
// and a boolean to check if the value has been set.
func (o *ISCSIInitiator) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Comment, true
}

// SetComment sets field value
func (o *ISCSIInitiator) SetComment(v string) {
	o.Comment = v
}

func (o ISCSIInitiator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ISCSIInitiator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["initiators"] = o.Initiators
	toSerialize["auth_network"] = o.AuthNetwork
	toSerialize["comment"] = o.Comment
	return toSerialize, nil
}

func (o *ISCSIInitiator) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"initiators",
		"auth_network",
		"comment",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varISCSIInitiator := _ISCSIInitiator{}

	err = json.Unmarshal(bytes, &varISCSIInitiator)

	if err != nil {
		return err
	}

	*o = ISCSIInitiator(varISCSIInitiator)

	return err
}

type NullableISCSIInitiator struct {
	value *ISCSIInitiator
	isSet bool
}

func (v NullableISCSIInitiator) Get() *ISCSIInitiator {
	return v.value
}

func (v *NullableISCSIInitiator) Set(val *ISCSIInitiator) {
	v.value = val
	v.isSet = true
}

func (v NullableISCSIInitiator) IsSet() bool {
	return v.isSet
}

func (v *NullableISCSIInitiator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableISCSIInitiator(val *ISCSIInitiator) *NullableISCSIInitiator {
	return &NullableISCSIInitiator{value: val, isSet: true}
}

func (v NullableISCSIInitiator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableISCSIInitiator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}