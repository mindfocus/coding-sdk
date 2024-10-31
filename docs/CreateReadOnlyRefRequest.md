# CreateReadOnlyRefRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotPath** | **string** | 仓库路径 | 
**RefName** | **string** | 分支名称 | 

## Methods

### NewCreateReadOnlyRefRequest

`func NewCreateReadOnlyRefRequest(depotPath string, refName string, ) *CreateReadOnlyRefRequest`

NewCreateReadOnlyRefRequest instantiates a new CreateReadOnlyRefRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateReadOnlyRefRequestWithDefaults

`func NewCreateReadOnlyRefRequestWithDefaults() *CreateReadOnlyRefRequest`

NewCreateReadOnlyRefRequestWithDefaults instantiates a new CreateReadOnlyRefRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotPath

`func (o *CreateReadOnlyRefRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *CreateReadOnlyRefRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *CreateReadOnlyRefRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.


### GetRefName

`func (o *CreateReadOnlyRefRequest) GetRefName() string`

GetRefName returns the RefName field if non-nil, zero value otherwise.

### GetRefNameOk

`func (o *CreateReadOnlyRefRequest) GetRefNameOk() (*string, bool)`

GetRefNameOk returns a tuple with the RefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefName

`func (o *CreateReadOnlyRefRequest) SetRefName(v string)`

SetRefName sets RefName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


