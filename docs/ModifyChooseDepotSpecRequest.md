# ModifyChooseDepotSpecRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotPath** | **string** | 仓库路径 | 
**DepotSpecName** | **string** | 选择的仓库规范名字（不使用仓库规范时，填空字符串） | 

## Methods

### NewModifyChooseDepotSpecRequest

`func NewModifyChooseDepotSpecRequest(depotPath string, depotSpecName string, ) *ModifyChooseDepotSpecRequest`

NewModifyChooseDepotSpecRequest instantiates a new ModifyChooseDepotSpecRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyChooseDepotSpecRequestWithDefaults

`func NewModifyChooseDepotSpecRequestWithDefaults() *ModifyChooseDepotSpecRequest`

NewModifyChooseDepotSpecRequestWithDefaults instantiates a new ModifyChooseDepotSpecRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotPath

`func (o *ModifyChooseDepotSpecRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *ModifyChooseDepotSpecRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *ModifyChooseDepotSpecRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.


### GetDepotSpecName

`func (o *ModifyChooseDepotSpecRequest) GetDepotSpecName() string`

GetDepotSpecName returns the DepotSpecName field if non-nil, zero value otherwise.

### GetDepotSpecNameOk

`func (o *ModifyChooseDepotSpecRequest) GetDepotSpecNameOk() (*string, bool)`

GetDepotSpecNameOk returns a tuple with the DepotSpecName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotSpecName

`func (o *ModifyChooseDepotSpecRequest) SetDepotSpecName(v string)`

SetDepotSpecName sets DepotSpecName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


