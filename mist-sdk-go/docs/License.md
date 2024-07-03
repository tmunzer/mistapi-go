# License

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amendments** | Pointer to [**[]LicenseAmendment**](LicenseAmendment.md) |  | [optional] [readonly] 
**Entitled** | Pointer to **map[string]int32** | Property key is license type (e.g. SUB-MAN) and Property value is the number of licenses entitled. | [optional] [readonly] 
**Licenses** | Pointer to [**[]LicenseSub**](LicenseSub.md) |  | [optional] 
**Summary** | Pointer to **map[string]int32** | Property key is license type (e.g. SUB-MAN) and Property value is the number of licenses consumed. | [optional] [readonly] 

## Methods

### NewLicense

`func NewLicense() *License`

NewLicense instantiates a new License object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseWithDefaults

`func NewLicenseWithDefaults() *License`

NewLicenseWithDefaults instantiates a new License object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmendments

`func (o *License) GetAmendments() []LicenseAmendment`

GetAmendments returns the Amendments field if non-nil, zero value otherwise.

### GetAmendmentsOk

`func (o *License) GetAmendmentsOk() (*[]LicenseAmendment, bool)`

GetAmendmentsOk returns a tuple with the Amendments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmendments

`func (o *License) SetAmendments(v []LicenseAmendment)`

SetAmendments sets Amendments field to given value.

### HasAmendments

`func (o *License) HasAmendments() bool`

HasAmendments returns a boolean if a field has been set.

### GetEntitled

`func (o *License) GetEntitled() map[string]int32`

GetEntitled returns the Entitled field if non-nil, zero value otherwise.

### GetEntitledOk

`func (o *License) GetEntitledOk() (*map[string]int32, bool)`

GetEntitledOk returns a tuple with the Entitled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitled

`func (o *License) SetEntitled(v map[string]int32)`

SetEntitled sets Entitled field to given value.

### HasEntitled

`func (o *License) HasEntitled() bool`

HasEntitled returns a boolean if a field has been set.

### GetLicenses

`func (o *License) GetLicenses() []LicenseSub`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *License) GetLicensesOk() (*[]LicenseSub, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *License) SetLicenses(v []LicenseSub)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *License) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.

### GetSummary

`func (o *License) GetSummary() map[string]int32`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *License) GetSummaryOk() (*map[string]int32, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *License) SetSummary(v map[string]int32)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *License) HasSummary() bool`

HasSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


