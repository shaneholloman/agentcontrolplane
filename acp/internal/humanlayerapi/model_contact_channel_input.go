/*
FastAPI

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package humanlayerapi

import (
	"encoding/json"
)

// checks if the ContactChannelInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContactChannelInput{}

// ContactChannelInput struct for ContactChannelInput
type ContactChannelInput struct {
	Slack    NullableSlackContactChannelInput `json:"slack,omitempty"`
	Sms      NullableSMSContactChannel        `json:"sms,omitempty"`
	Whatsapp NullableWhatsAppContactChannel   `json:"whatsapp,omitempty"`
	Email    NullableEmailContactChannel      `json:"email,omitempty"`
}

// NewContactChannelInput instantiates a new ContactChannelInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactChannelInput() *ContactChannelInput {
	this := ContactChannelInput{}
	return &this
}

// NewContactChannelInputWithDefaults instantiates a new ContactChannelInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactChannelInputWithDefaults() *ContactChannelInput {
	this := ContactChannelInput{}
	return &this
}

// GetSlack returns the Slack field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContactChannelInput) GetSlack() SlackContactChannelInput {
	if o == nil || IsNil(o.Slack.Get()) {
		var ret SlackContactChannelInput
		return ret
	}
	return *o.Slack.Get()
}

// GetSlackOk returns a tuple with the Slack field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContactChannelInput) GetSlackOk() (*SlackContactChannelInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Slack.Get(), o.Slack.IsSet()
}

// HasSlack returns a boolean if a field has been set.
func (o *ContactChannelInput) HasSlack() bool {
	if o != nil && o.Slack.IsSet() {
		return true
	}

	return false
}

// SetSlack gets a reference to the given NullableSlackContactChannelInput and assigns it to the Slack field.
func (o *ContactChannelInput) SetSlack(v SlackContactChannelInput) {
	o.Slack.Set(&v)
}

// SetSlackNil sets the value for Slack to be an explicit nil
func (o *ContactChannelInput) SetSlackNil() {
	o.Slack.Set(nil)
}

// UnsetSlack ensures that no value is present for Slack, not even an explicit nil
func (o *ContactChannelInput) UnsetSlack() {
	o.Slack.Unset()
}

// GetSms returns the Sms field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContactChannelInput) GetSms() SMSContactChannel {
	if o == nil || IsNil(o.Sms.Get()) {
		var ret SMSContactChannel
		return ret
	}
	return *o.Sms.Get()
}

// GetSmsOk returns a tuple with the Sms field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContactChannelInput) GetSmsOk() (*SMSContactChannel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sms.Get(), o.Sms.IsSet()
}

// HasSms returns a boolean if a field has been set.
func (o *ContactChannelInput) HasSms() bool {
	if o != nil && o.Sms.IsSet() {
		return true
	}

	return false
}

// SetSms gets a reference to the given NullableSMSContactChannel and assigns it to the Sms field.
func (o *ContactChannelInput) SetSms(v SMSContactChannel) {
	o.Sms.Set(&v)
}

// SetSmsNil sets the value for Sms to be an explicit nil
func (o *ContactChannelInput) SetSmsNil() {
	o.Sms.Set(nil)
}

// UnsetSms ensures that no value is present for Sms, not even an explicit nil
func (o *ContactChannelInput) UnsetSms() {
	o.Sms.Unset()
}

// GetWhatsapp returns the Whatsapp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContactChannelInput) GetWhatsapp() WhatsAppContactChannel {
	if o == nil || IsNil(o.Whatsapp.Get()) {
		var ret WhatsAppContactChannel
		return ret
	}
	return *o.Whatsapp.Get()
}

// GetWhatsappOk returns a tuple with the Whatsapp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContactChannelInput) GetWhatsappOk() (*WhatsAppContactChannel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Whatsapp.Get(), o.Whatsapp.IsSet()
}

// HasWhatsapp returns a boolean if a field has been set.
func (o *ContactChannelInput) HasWhatsapp() bool {
	if o != nil && o.Whatsapp.IsSet() {
		return true
	}

	return false
}

// SetWhatsapp gets a reference to the given NullableWhatsAppContactChannel and assigns it to the Whatsapp field.
func (o *ContactChannelInput) SetWhatsapp(v WhatsAppContactChannel) {
	o.Whatsapp.Set(&v)
}

// SetWhatsappNil sets the value for Whatsapp to be an explicit nil
func (o *ContactChannelInput) SetWhatsappNil() {
	o.Whatsapp.Set(nil)
}

// UnsetWhatsapp ensures that no value is present for Whatsapp, not even an explicit nil
func (o *ContactChannelInput) UnsetWhatsapp() {
	o.Whatsapp.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContactChannelInput) GetEmail() EmailContactChannel {
	if o == nil || IsNil(o.Email.Get()) {
		var ret EmailContactChannel
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContactChannelInput) GetEmailOk() (*EmailContactChannel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *ContactChannelInput) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableEmailContactChannel and assigns it to the Email field.
func (o *ContactChannelInput) SetEmail(v EmailContactChannel) {
	o.Email.Set(&v)
}

// SetEmailNil sets the value for Email to be an explicit nil
func (o *ContactChannelInput) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *ContactChannelInput) UnsetEmail() {
	o.Email.Unset()
}

func (o ContactChannelInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContactChannelInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Slack.IsSet() {
		toSerialize["slack"] = o.Slack.Get()
	}
	if o.Sms.IsSet() {
		toSerialize["sms"] = o.Sms.Get()
	}
	if o.Whatsapp.IsSet() {
		toSerialize["whatsapp"] = o.Whatsapp.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	return toSerialize, nil
}

type NullableContactChannelInput struct {
	value *ContactChannelInput
	isSet bool
}

func (v NullableContactChannelInput) Get() *ContactChannelInput {
	return v.value
}

func (v *NullableContactChannelInput) Set(val *ContactChannelInput) {
	v.value = val
	v.isSet = true
}

func (v NullableContactChannelInput) IsSet() bool {
	return v.isSet
}

func (v *NullableContactChannelInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactChannelInput(val *ContactChannelInput) *NullableContactChannelInput {
	return &NullableContactChannelInput{value: val, isSet: true}
}

func (v NullableContactChannelInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactChannelInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
