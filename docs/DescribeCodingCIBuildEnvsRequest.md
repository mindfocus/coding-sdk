# DescribeCodingCIBuildEnvsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildId** | **int32** | 构建ID | 
**Type** | **string** | SYSTEM（系统内置环境变量） Param（触发构建输入环境变量） Env（构建计划填写环境变量） | 

## Methods

### NewDescribeCodingCIBuildEnvsRequest

`func NewDescribeCodingCIBuildEnvsRequest(buildId int32, type_ string, ) *DescribeCodingCIBuildEnvsRequest`

NewDescribeCodingCIBuildEnvsRequest instantiates a new DescribeCodingCIBuildEnvsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeCodingCIBuildEnvsRequestWithDefaults

`func NewDescribeCodingCIBuildEnvsRequestWithDefaults() *DescribeCodingCIBuildEnvsRequest`

NewDescribeCodingCIBuildEnvsRequestWithDefaults instantiates a new DescribeCodingCIBuildEnvsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildId

`func (o *DescribeCodingCIBuildEnvsRequest) GetBuildId() int32`

GetBuildId returns the BuildId field if non-nil, zero value otherwise.

### GetBuildIdOk

`func (o *DescribeCodingCIBuildEnvsRequest) GetBuildIdOk() (*int32, bool)`

GetBuildIdOk returns a tuple with the BuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildId

`func (o *DescribeCodingCIBuildEnvsRequest) SetBuildId(v int32)`

SetBuildId sets BuildId field to given value.


### GetType

`func (o *DescribeCodingCIBuildEnvsRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DescribeCodingCIBuildEnvsRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DescribeCodingCIBuildEnvsRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


