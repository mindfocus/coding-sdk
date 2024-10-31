# ModifyCodingCIAgentEnableRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolId** | **int64** | 节点池 ID | [default to 0]
**Enable** | **bool** | 是否启用：启用true、禁用false | [default to false]
**Id** | **int64** | 节点ID | [default to 0]

## Methods

### NewModifyCodingCIAgentEnableRequest

`func NewModifyCodingCIAgentEnableRequest(poolId int64, enable bool, id int64, ) *ModifyCodingCIAgentEnableRequest`

NewModifyCodingCIAgentEnableRequest instantiates a new ModifyCodingCIAgentEnableRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyCodingCIAgentEnableRequestWithDefaults

`func NewModifyCodingCIAgentEnableRequestWithDefaults() *ModifyCodingCIAgentEnableRequest`

NewModifyCodingCIAgentEnableRequestWithDefaults instantiates a new ModifyCodingCIAgentEnableRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoolId

`func (o *ModifyCodingCIAgentEnableRequest) GetPoolId() int64`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *ModifyCodingCIAgentEnableRequest) GetPoolIdOk() (*int64, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *ModifyCodingCIAgentEnableRequest) SetPoolId(v int64)`

SetPoolId sets PoolId field to given value.


### GetEnable

`func (o *ModifyCodingCIAgentEnableRequest) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *ModifyCodingCIAgentEnableRequest) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *ModifyCodingCIAgentEnableRequest) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetId

`func (o *ModifyCodingCIAgentEnableRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModifyCodingCIAgentEnableRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModifyCodingCIAgentEnableRequest) SetId(v int64)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


