# MapNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Edges** | Pointer to **map[string]string** |  | [optional] 
**Name** | **string** |  | 
**Position** | Pointer to [**MapNodePosition**](MapNodePosition.md) |  | [optional] 

## Methods

### NewMapNode

`func NewMapNode(name string, ) *MapNode`

NewMapNode instantiates a new MapNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMapNodeWithDefaults

`func NewMapNodeWithDefaults() *MapNode`

NewMapNodeWithDefaults instantiates a new MapNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEdges

`func (o *MapNode) GetEdges() map[string]string`

GetEdges returns the Edges field if non-nil, zero value otherwise.

### GetEdgesOk

`func (o *MapNode) GetEdgesOk() (*map[string]string, bool)`

GetEdgesOk returns a tuple with the Edges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdges

`func (o *MapNode) SetEdges(v map[string]string)`

SetEdges sets Edges field to given value.

### HasEdges

`func (o *MapNode) HasEdges() bool`

HasEdges returns a boolean if a field has been set.

### GetName

`func (o *MapNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MapNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MapNode) SetName(v string)`

SetName sets Name field to given value.


### GetPosition

`func (o *MapNode) GetPosition() MapNodePosition`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *MapNode) GetPositionOk() (*MapNodePosition, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *MapNode) SetPosition(v MapNodePosition)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *MapNode) HasPosition() bool`

HasPosition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


