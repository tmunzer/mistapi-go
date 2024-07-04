/*
Mist API

> Version: **2406.1.14** > > Date: **July 3, 2024**  ---  ### Additional Documentation * [Mist Automation Guide](https://www.juniper.net/documentation/us/en/software/mist/automation-integration/index.html) * [Mist Location SDK](https://www.juniper.net/documentation/us/en/software/mist/location_services/topics/concept/mist-how-get-mist-sdk.html) * [Mist Product Updates](https://www.mist.com/documentation/category/product-updates/)  ---  ### Helpful Resources * [API Sandbox and Exercises](https://api-class.mist.com/) * [Postman Collection, Runners and Webhook Samples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace) * [API Demo Apps](https://apps.mist-lab.fr/) * [Juniper Blog](https://blogs.juniper.net/)  --- 

API version: 2406.1.14
Contact: tmunzer@juniper.net
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mistapigo

import (
	"encoding/json"
	"fmt"
)

// checks if the Ticket type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Ticket{}

// Ticket Support Ticket
type Ticket struct {
	CaseNumber *string `json:"case_number,omitempty"`
	Comments []TicketComment `json:"comments,omitempty"`
	CreatedAt *int32 `json:"created_at,omitempty"`
	Id *string `json:"id,omitempty"`
	Requester *string `json:"requester,omitempty"`
	// email of the requester
	RequesterEmail *string `json:"requester_email,omitempty"`
	Status *TicketStatus `json:"status,omitempty"`
	Subject string `json:"subject"`
	// question (default) / bug / critical
	Type string `json:"type"`
	UpdatedAt *int32 `json:"updated_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Ticket Ticket

// NewTicket instantiates a new Ticket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTicket(subject string, type_ string) *Ticket {
	this := Ticket{}
	this.Subject = subject
	this.Type = type_
	return &this
}

// NewTicketWithDefaults instantiates a new Ticket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTicketWithDefaults() *Ticket {
	this := Ticket{}
	return &this
}

// GetCaseNumber returns the CaseNumber field value if set, zero value otherwise.
func (o *Ticket) GetCaseNumber() string {
	if o == nil || IsNil(o.CaseNumber) {
		var ret string
		return ret
	}
	return *o.CaseNumber
}

// GetCaseNumberOk returns a tuple with the CaseNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticket) GetCaseNumberOk() (*string, bool) {
	if o == nil || IsNil(o.CaseNumber) {
		return nil, false
	}
	return o.CaseNumber, true
}

// HasCaseNumber returns a boolean if a field has been set.
func (o *Ticket) HasCaseNumber() bool {
	if o != nil && !IsNil(o.CaseNumber) {
		return true
	}

	return false
}

// SetCaseNumber gets a reference to the given string and assigns it to the CaseNumber field.
func (o *Ticket) SetCaseNumber(v string) {
	o.CaseNumber = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *Ticket) GetComments() []TicketComment {
	if o == nil || IsNil(o.Comments) {
		var ret []TicketComment
		return ret
	}
	return o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticket) GetCommentsOk() ([]TicketComment, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *Ticket) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given []TicketComment and assigns it to the Comments field.
func (o *Ticket) SetComments(v []TicketComment) {
	o.Comments = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Ticket) GetCreatedAt() int32 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticket) GetCreatedAtOk() (*int32, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Ticket) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int32 and assigns it to the CreatedAt field.
func (o *Ticket) SetCreatedAt(v int32) {
	o.CreatedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Ticket) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticket) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Ticket) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Ticket) SetId(v string) {
	o.Id = &v
}

// GetRequester returns the Requester field value if set, zero value otherwise.
func (o *Ticket) GetRequester() string {
	if o == nil || IsNil(o.Requester) {
		var ret string
		return ret
	}
	return *o.Requester
}

// GetRequesterOk returns a tuple with the Requester field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticket) GetRequesterOk() (*string, bool) {
	if o == nil || IsNil(o.Requester) {
		return nil, false
	}
	return o.Requester, true
}

// HasRequester returns a boolean if a field has been set.
func (o *Ticket) HasRequester() bool {
	if o != nil && !IsNil(o.Requester) {
		return true
	}

	return false
}

// SetRequester gets a reference to the given string and assigns it to the Requester field.
func (o *Ticket) SetRequester(v string) {
	o.Requester = &v
}

// GetRequesterEmail returns the RequesterEmail field value if set, zero value otherwise.
func (o *Ticket) GetRequesterEmail() string {
	if o == nil || IsNil(o.RequesterEmail) {
		var ret string
		return ret
	}
	return *o.RequesterEmail
}

// GetRequesterEmailOk returns a tuple with the RequesterEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticket) GetRequesterEmailOk() (*string, bool) {
	if o == nil || IsNil(o.RequesterEmail) {
		return nil, false
	}
	return o.RequesterEmail, true
}

// HasRequesterEmail returns a boolean if a field has been set.
func (o *Ticket) HasRequesterEmail() bool {
	if o != nil && !IsNil(o.RequesterEmail) {
		return true
	}

	return false
}

// SetRequesterEmail gets a reference to the given string and assigns it to the RequesterEmail field.
func (o *Ticket) SetRequesterEmail(v string) {
	o.RequesterEmail = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Ticket) GetStatus() TicketStatus {
	if o == nil || IsNil(o.Status) {
		var ret TicketStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticket) GetStatusOk() (*TicketStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Ticket) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given TicketStatus and assigns it to the Status field.
func (o *Ticket) SetStatus(v TicketStatus) {
	o.Status = &v
}

// GetSubject returns the Subject field value
func (o *Ticket) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *Ticket) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *Ticket) SetSubject(v string) {
	o.Subject = v
}

// GetType returns the Type field value
func (o *Ticket) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Ticket) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Ticket) SetType(v string) {
	o.Type = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Ticket) GetUpdatedAt() int32 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret int32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticket) GetUpdatedAtOk() (*int32, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Ticket) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int32 and assigns it to the UpdatedAt field.
func (o *Ticket) SetUpdatedAt(v int32) {
	o.UpdatedAt = &v
}

func (o Ticket) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Ticket) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CaseNumber) {
		toSerialize["case_number"] = o.CaseNumber
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Requester) {
		toSerialize["requester"] = o.Requester
	}
	if !IsNil(o.RequesterEmail) {
		toSerialize["requester_email"] = o.RequesterEmail
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	toSerialize["subject"] = o.Subject
	toSerialize["type"] = o.Type
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Ticket) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"subject",
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTicket := _Ticket{}

	err = json.Unmarshal(data, &varTicket)

	if err != nil {
		return err
	}

	*o = Ticket(varTicket)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "case_number")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "id")
		delete(additionalProperties, "requester")
		delete(additionalProperties, "requester_email")
		delete(additionalProperties, "status")
		delete(additionalProperties, "subject")
		delete(additionalProperties, "type")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTicket struct {
	value *Ticket
	isSet bool
}

func (v NullableTicket) Get() *Ticket {
	return v.value
}

func (v *NullableTicket) Set(val *Ticket) {
	v.value = val
	v.isSet = true
}

func (v NullableTicket) IsSet() bool {
	return v.isSet
}

func (v *NullableTicket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTicket(val *Ticket) *NullableTicket {
	return &NullableTicket{value: val, isSet: true}
}

func (v NullableTicket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTicket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

