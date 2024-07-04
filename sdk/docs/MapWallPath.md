# MapWallPath

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Coordinate** | Pointer to **string** |  | [optional] 
**Nodes** | Pointer to [**[]MapNode**](MapNode.md) |  | [optional] 

## Methods

### NewMapWallPath

`func NewMapWallPath() *MapWallPath`

NewMapWallPath instantiates a new MapWallPath object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMapWallPathWithDefaults

`func NewMapWallPathWithDefaults() *MapWallPath`

NewMapWallPathWithDefaults instantiates a new MapWallPath object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoordinate

`func (o *MapWallPath) GetCoordinate() string`

GetCoordinate returns the Coordinate field if non-nil, zero value otherwise.

### GetCoordinateOk

`func (o *MapWallPath) GetCoordinateOk() (*string, bool)`

GetCoordinateOk returns a tuple with the Coordinate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinate

`func (o *MapWallPath) SetCoordinate(v string)`

SetCoordinate sets Coordinate field to given value.

### HasCoordinate

`func (o *MapWallPath) HasCoordinate() bool`

HasCoordinate returns a boolean if a field has been set.

### GetNodes

`func (o *MapWallPath) GetNodes() []MapNode`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *MapWallPath) GetNodesOk() (*[]MapNode, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *MapWallPath) SetNodes(v []MapNode)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *MapWallPath) HasNodes() bool`

HasNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


