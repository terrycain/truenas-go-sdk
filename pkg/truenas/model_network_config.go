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

// checks if the NetworkConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkConfig{}

// NetworkConfig struct for NetworkConfig
type NetworkConfig struct {
	Id                   *int32                            `json:"id,omitempty"`
	Hostname             *string                           `json:"hostname,omitempty"`
	Domain               *string                           `json:"domain,omitempty"`
	Ipv4gateway          *string                           `json:"ipv4gateway,omitempty"`
	Ipv6gateway          *string                           `json:"ipv6gateway,omitempty"`
	Nameserver1          *string                           `json:"nameserver1,omitempty"`
	Nameserver2          *string                           `json:"nameserver2,omitempty"`
	Nameserver3          *string                           `json:"nameserver3,omitempty"`
	Httpproxy            *string                           `json:"httpproxy,omitempty"`
	NetwaitEnabled       *bool                             `json:"netwait_enabled,omitempty"`
	NetwaitIp            []string                          `json:"netwait_ip,omitempty"`
	Hosts                *string                           `json:"hosts,omitempty"`
	Domains              []string                          `json:"domains,omitempty"`
	ServiceAnnouncement  *NetworkConfigServiceAnnouncement `json:"service_announcement,omitempty"`
	HostnameLocal        *string                           `json:"hostname_local,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NetworkConfig NetworkConfig

// NewNetworkConfig instantiates a new NetworkConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkConfig() *NetworkConfig {
	this := NetworkConfig{}
	return &this
}

// NewNetworkConfigWithDefaults instantiates a new NetworkConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkConfigWithDefaults() *NetworkConfig {
	this := NetworkConfig{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NetworkConfig) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfig) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NetworkConfig) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *NetworkConfig) SetId(v int32) {
	o.Id = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *NetworkConfig) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfig) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *NetworkConfig) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *NetworkConfig) SetHostname(v string) {
	o.Hostname = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *NetworkConfig) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfig) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *NetworkConfig) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *NetworkConfig) SetDomain(v string) {
	o.Domain = &v
}

// GetIpv4gateway returns the Ipv4gateway field value if set, zero value otherwise.
func (o *NetworkConfig) GetIpv4gateway() string {
	if o == nil || IsNil(o.Ipv4gateway) {
		var ret string
		return ret
	}
	return *o.Ipv4gateway
}

// GetIpv4gatewayOk returns a tuple with the Ipv4gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfig) GetIpv4gatewayOk() (*string, bool) {
	if o == nil || IsNil(o.Ipv4gateway) {
		return nil, false
	}
	return o.Ipv4gateway, true
}

// HasIpv4gateway returns a boolean if a field has been set.
func (o *NetworkConfig) HasIpv4gateway() bool {
	if o != nil && !IsNil(o.Ipv4gateway) {
		return true
	}

	return false
}

// SetIpv4gateway gets a reference to the given string and assigns it to the Ipv4gateway field.
func (o *NetworkConfig) SetIpv4gateway(v string) {
	o.Ipv4gateway = &v
}

// GetIpv6gateway returns the Ipv6gateway field value if set, zero value otherwise.
func (o *NetworkConfig) GetIpv6gateway() string {
	if o == nil || IsNil(o.Ipv6gateway) {
		var ret string
		return ret
	}
	return *o.Ipv6gateway
}

// GetIpv6gatewayOk returns a tuple with the Ipv6gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfig) GetIpv6gatewayOk() (*string, bool) {
	if o == nil || IsNil(o.Ipv6gateway) {
		return nil, false
	}
	return o.Ipv6gateway, true
}

// HasIpv6gateway returns a boolean if a field has been set.
func (o *NetworkConfig) HasIpv6gateway() bool {
	if o != nil && !IsNil(o.Ipv6gateway) {
		return true
	}

	return false
}

// SetIpv6gateway gets a reference to the given string and assigns it to the Ipv6gateway field.
func (o *NetworkConfig) SetIpv6gateway(v string) {
	o.Ipv6gateway = &v
}

// GetNameserver1 returns the Nameserver1 field value if set, zero value otherwise.
func (o *NetworkConfig) GetNameserver1() string {
	if o == nil || IsNil(o.Nameserver1) {
		var ret string
		return ret
	}
	return *o.Nameserver1
}

// GetNameserver1Ok returns a tuple with the Nameserver1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfig) GetNameserver1Ok() (*string, bool) {
	if o == nil || IsNil(o.Nameserver1) {
		return nil, false
	}
	return o.Nameserver1, true
}

// HasNameserver1 returns a boolean if a field has been set.
func (o *NetworkConfig) HasNameserver1() bool {
	if o != nil && !IsNil(o.Nameserver1) {
		return true
	}

	return false
}

// SetNameserver1 gets a reference to the given string and assigns it to the Nameserver1 field.
func (o *NetworkConfig) SetNameserver1(v string) {
	o.Nameserver1 = &v
}

// GetNameserver2 returns the Nameserver2 field value if set, zero value otherwise.
func (o *NetworkConfig) GetNameserver2() string {
	if o == nil || IsNil(o.Nameserver2) {
		var ret string
		return ret
	}
	return *o.Nameserver2
}

// GetNameserver2Ok returns a tuple with the Nameserver2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfig) GetNameserver2Ok() (*string, bool) {
	if o == nil || IsNil(o.Nameserver2) {
		return nil, false
	}
	return o.Nameserver2, true
}

// HasNameserver2 returns a boolean if a field has been set.
func (o *NetworkConfig) HasNameserver2() bool {
	if o != nil && !IsNil(o.Nameserver2) {
		return true
	}

	return false
}

// SetNameserver2 gets a reference to the given string and assigns it to the Nameserver2 field.
func (o *NetworkConfig) SetNameserver2(v string) {
	o.Nameserver2 = &v
}

// GetNameserver3 returns the Nameserver3 field value if set, zero value otherwise.
func (o *NetworkConfig) GetNameserver3() string {
	if o == nil || IsNil(o.Nameserver3) {
		var ret string
		return ret
	}
	return *o.Nameserver3
}

// GetNameserver3Ok returns a tuple with the Nameserver3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfig) GetNameserver3Ok() (*string, bool) {
	if o == nil || IsNil(o.Nameserver3) {
		return nil, false
	}
	return o.Nameserver3, true
}

// HasNameserver3 returns a boolean if a field has been set.
func (o *NetworkConfig) HasNameserver3() bool {
	if o != nil && !IsNil(o.Nameserver3) {
		return true
	}

	return false
}

// SetNameserver3 gets a reference to the given string and assigns it to the Nameserver3 field.
func (o *NetworkConfig) SetNameserver3(v string) {
	o.Nameserver3 = &v
}

// GetHttpproxy returns the Httpproxy field value if set, zero value otherwise.
func (o *NetworkConfig) GetHttpproxy() string {
	if o == nil || IsNil(o.Httpproxy) {
		var ret string
		return ret
	}
	return *o.Httpproxy
}

// GetHttpproxyOk returns a tuple with the Httpproxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfig) GetHttpproxyOk() (*string, bool) {
	if o == nil || IsNil(o.Httpproxy) {
		return nil, false
	}
	return o.Httpproxy, true
}

// HasHttpproxy returns a boolean if a field has been set.
func (o *NetworkConfig) HasHttpproxy() bool {
	if o != nil && !IsNil(o.Httpproxy) {
		return true
	}

	return false
}

// SetHttpproxy gets a reference to the given string and assigns it to the Httpproxy field.
func (o *NetworkConfig) SetHttpproxy(v string) {
	o.Httpproxy = &v
}

// GetNetwaitEnabled returns the NetwaitEnabled field value if set, zero value otherwise.
func (o *NetworkConfig) GetNetwaitEnabled() bool {
	if o == nil || IsNil(o.NetwaitEnabled) {
		var ret bool
		return ret
	}
	return *o.NetwaitEnabled
}

// GetNetwaitEnabledOk returns a tuple with the NetwaitEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfig) GetNetwaitEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.NetwaitEnabled) {
		return nil, false
	}
	return o.NetwaitEnabled, true
}

// HasNetwaitEnabled returns a boolean if a field has been set.
func (o *NetworkConfig) HasNetwaitEnabled() bool {
	if o != nil && !IsNil(o.NetwaitEnabled) {
		return true
	}

	return false
}

// SetNetwaitEnabled gets a reference to the given bool and assigns it to the NetwaitEnabled field.
func (o *NetworkConfig) SetNetwaitEnabled(v bool) {
	o.NetwaitEnabled = &v
}

// GetNetwaitIp returns the NetwaitIp field value if set, zero value otherwise.
func (o *NetworkConfig) GetNetwaitIp() []string {
	if o == nil || IsNil(o.NetwaitIp) {
		var ret []string
		return ret
	}
	return o.NetwaitIp
}

// GetNetwaitIpOk returns a tuple with the NetwaitIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfig) GetNetwaitIpOk() ([]string, bool) {
	if o == nil || IsNil(o.NetwaitIp) {
		return nil, false
	}
	return o.NetwaitIp, true
}

// HasNetwaitIp returns a boolean if a field has been set.
func (o *NetworkConfig) HasNetwaitIp() bool {
	if o != nil && !IsNil(o.NetwaitIp) {
		return true
	}

	return false
}

// SetNetwaitIp gets a reference to the given []string and assigns it to the NetwaitIp field.
func (o *NetworkConfig) SetNetwaitIp(v []string) {
	o.NetwaitIp = v
}

// GetHosts returns the Hosts field value if set, zero value otherwise.
func (o *NetworkConfig) GetHosts() string {
	if o == nil || IsNil(o.Hosts) {
		var ret string
		return ret
	}
	return *o.Hosts
}

// GetHostsOk returns a tuple with the Hosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfig) GetHostsOk() (*string, bool) {
	if o == nil || IsNil(o.Hosts) {
		return nil, false
	}
	return o.Hosts, true
}

// HasHosts returns a boolean if a field has been set.
func (o *NetworkConfig) HasHosts() bool {
	if o != nil && !IsNil(o.Hosts) {
		return true
	}

	return false
}

// SetHosts gets a reference to the given string and assigns it to the Hosts field.
func (o *NetworkConfig) SetHosts(v string) {
	o.Hosts = &v
}

// GetDomains returns the Domains field value if set, zero value otherwise.
func (o *NetworkConfig) GetDomains() []string {
	if o == nil || IsNil(o.Domains) {
		var ret []string
		return ret
	}
	return o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfig) GetDomainsOk() ([]string, bool) {
	if o == nil || IsNil(o.Domains) {
		return nil, false
	}
	return o.Domains, true
}

// HasDomains returns a boolean if a field has been set.
func (o *NetworkConfig) HasDomains() bool {
	if o != nil && !IsNil(o.Domains) {
		return true
	}

	return false
}

// SetDomains gets a reference to the given []string and assigns it to the Domains field.
func (o *NetworkConfig) SetDomains(v []string) {
	o.Domains = v
}

// GetServiceAnnouncement returns the ServiceAnnouncement field value if set, zero value otherwise.
func (o *NetworkConfig) GetServiceAnnouncement() NetworkConfigServiceAnnouncement {
	if o == nil || IsNil(o.ServiceAnnouncement) {
		var ret NetworkConfigServiceAnnouncement
		return ret
	}
	return *o.ServiceAnnouncement
}

// GetServiceAnnouncementOk returns a tuple with the ServiceAnnouncement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfig) GetServiceAnnouncementOk() (*NetworkConfigServiceAnnouncement, bool) {
	if o == nil || IsNil(o.ServiceAnnouncement) {
		return nil, false
	}
	return o.ServiceAnnouncement, true
}

// HasServiceAnnouncement returns a boolean if a field has been set.
func (o *NetworkConfig) HasServiceAnnouncement() bool {
	if o != nil && !IsNil(o.ServiceAnnouncement) {
		return true
	}

	return false
}

// SetServiceAnnouncement gets a reference to the given NetworkConfigServiceAnnouncement and assigns it to the ServiceAnnouncement field.
func (o *NetworkConfig) SetServiceAnnouncement(v NetworkConfigServiceAnnouncement) {
	o.ServiceAnnouncement = &v
}

// GetHostnameLocal returns the HostnameLocal field value if set, zero value otherwise.
func (o *NetworkConfig) GetHostnameLocal() string {
	if o == nil || IsNil(o.HostnameLocal) {
		var ret string
		return ret
	}
	return *o.HostnameLocal
}

// GetHostnameLocalOk returns a tuple with the HostnameLocal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfig) GetHostnameLocalOk() (*string, bool) {
	if o == nil || IsNil(o.HostnameLocal) {
		return nil, false
	}
	return o.HostnameLocal, true
}

// HasHostnameLocal returns a boolean if a field has been set.
func (o *NetworkConfig) HasHostnameLocal() bool {
	if o != nil && !IsNil(o.HostnameLocal) {
		return true
	}

	return false
}

// SetHostnameLocal gets a reference to the given string and assigns it to the HostnameLocal field.
func (o *NetworkConfig) SetHostnameLocal(v string) {
	o.HostnameLocal = &v
}

func (o NetworkConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.Ipv4gateway) {
		toSerialize["ipv4gateway"] = o.Ipv4gateway
	}
	if !IsNil(o.Ipv6gateway) {
		toSerialize["ipv6gateway"] = o.Ipv6gateway
	}
	if !IsNil(o.Nameserver1) {
		toSerialize["nameserver1"] = o.Nameserver1
	}
	if !IsNil(o.Nameserver2) {
		toSerialize["nameserver2"] = o.Nameserver2
	}
	if !IsNil(o.Nameserver3) {
		toSerialize["nameserver3"] = o.Nameserver3
	}
	if !IsNil(o.Httpproxy) {
		toSerialize["httpproxy"] = o.Httpproxy
	}
	if !IsNil(o.NetwaitEnabled) {
		toSerialize["netwait_enabled"] = o.NetwaitEnabled
	}
	if !IsNil(o.NetwaitIp) {
		toSerialize["netwait_ip"] = o.NetwaitIp
	}
	if !IsNil(o.Hosts) {
		toSerialize["hosts"] = o.Hosts
	}
	if !IsNil(o.Domains) {
		toSerialize["domains"] = o.Domains
	}
	if !IsNil(o.ServiceAnnouncement) {
		toSerialize["service_announcement"] = o.ServiceAnnouncement
	}
	if !IsNil(o.HostnameLocal) {
		toSerialize["hostname_local"] = o.HostnameLocal
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NetworkConfig) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkConfig := _NetworkConfig{}

	err = json.Unmarshal(bytes, &varNetworkConfig)

	if err != nil {
		return err
	}

	*o = NetworkConfig(varNetworkConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "hostname")
		delete(additionalProperties, "domain")
		delete(additionalProperties, "ipv4gateway")
		delete(additionalProperties, "ipv6gateway")
		delete(additionalProperties, "nameserver1")
		delete(additionalProperties, "nameserver2")
		delete(additionalProperties, "nameserver3")
		delete(additionalProperties, "httpproxy")
		delete(additionalProperties, "netwait_enabled")
		delete(additionalProperties, "netwait_ip")
		delete(additionalProperties, "hosts")
		delete(additionalProperties, "domains")
		delete(additionalProperties, "service_announcement")
		delete(additionalProperties, "hostname_local")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNetworkConfig struct {
	value *NetworkConfig
	isSet bool
}

func (v NullableNetworkConfig) Get() *NetworkConfig {
	return v.value
}

func (v *NullableNetworkConfig) Set(val *NetworkConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkConfig(val *NetworkConfig) *NullableNetworkConfig {
	return &NullableNetworkConfig{value: val, isSet: true}
}

func (v NullableNetworkConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}