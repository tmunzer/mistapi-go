# ResponseMapImport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aps** | [**[]ResponseMapImportAp**](ResponseMapImportAp.md) |  | 
**Floorplans** | [**[]ResponseMapImportFloorplan**](ResponseMapImportFloorplan.md) |  | 
**ForSite** | Pointer to **bool** |  | [optional] [readonly] 
**SiteId** | **string** |  | [readonly] 
**Summary** | [**ResponseMapImportSummary**](ResponseMapImportSummary.md) |  | 

## Methods

### NewResponseMapImport

`func NewResponseMapImport(aps []ResponseMapImportAp, floorplans []ResponseMapImportFloorplan, siteId string, summary ResponseMapImportSummary, ) *ResponseMapImport`

NewResponseMapImport instantiates a new ResponseMapImport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseMapImportWithDefaults

`func NewResponseMapImportWithDefaults() *ResponseMapImport`

NewResponseMapImportWithDefaults instantiates a new ResponseMapImport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAps

`func (o *ResponseMapImport) GetAps() []ResponseMapImportAp`

GetAps returns the Aps field if non-nil, zero value otherwise.

### GetApsOk

`func (o *ResponseMapImport) GetApsOk() (*[]ResponseMapImportAp, bool)`

GetApsOk returns a tuple with the Aps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAps

`func (o *ResponseMapImport) SetAps(v []ResponseMapImportAp)`

SetAps sets Aps field to given value.


### GetFloorplans

`func (o *ResponseMapImport) GetFloorplans() []ResponseMapImportFloorplan`

GetFloorplans returns the Floorplans field if non-nil, zero value otherwise.

### GetFloorplansOk

`func (o *ResponseMapImport) GetFloorplansOk() (*[]ResponseMapImportFloorplan, bool)`

GetFloorplansOk returns a tuple with the Floorplans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloorplans

`func (o *ResponseMapImport) SetFloorplans(v []ResponseMapImportFloorplan)`

SetFloorplans sets Floorplans field to given value.


### GetForSite

`func (o *ResponseMapImport) GetForSite() bool`

GetForSite returns the ForSite field if non-nil, zero value otherwise.

### GetForSiteOk

`func (o *ResponseMapImport) GetForSiteOk() (*bool, bool)`

GetForSiteOk returns a tuple with the ForSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSite

`func (o *ResponseMapImport) SetForSite(v bool)`

SetForSite sets ForSite field to given value.

### HasForSite

`func (o *ResponseMapImport) HasForSite() bool`

HasForSite returns a boolean if a field has been set.

### GetSiteId

`func (o *ResponseMapImport) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *ResponseMapImport) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *ResponseMapImport) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.


### GetSummary

`func (o *ResponseMapImport) GetSummary() ResponseMapImportSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ResponseMapImport) GetSummaryOk() (*ResponseMapImportSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ResponseMapImport) SetSummary(v ResponseMapImportSummary)`

SetSummary sets Summary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


