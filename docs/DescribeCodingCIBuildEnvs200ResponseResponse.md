# DescribeCodingCIBuildEnvs200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvList** | Pointer to [**[]CIJobEnv**](CIJobEnv.md) | 环境变量列表 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeCodingCIBuildEnvs200ResponseResponse

`func NewDescribeCodingCIBuildEnvs200ResponseResponse() *DescribeCodingCIBuildEnvs200ResponseResponse`

NewDescribeCodingCIBuildEnvs200ResponseResponse instantiates a new DescribeCodingCIBuildEnvs200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeCodingCIBuildEnvs200ResponseResponseWithDefaults

`func NewDescribeCodingCIBuildEnvs200ResponseResponseWithDefaults() *DescribeCodingCIBuildEnvs200ResponseResponse`

NewDescribeCodingCIBuildEnvs200ResponseResponseWithDefaults instantiates a new DescribeCodingCIBuildEnvs200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvList

`func (o *DescribeCodingCIBuildEnvs200ResponseResponse) GetEnvList() []CIJobEnv`

GetEnvList returns the EnvList field if non-nil, zero value otherwise.

### GetEnvListOk

`func (o *DescribeCodingCIBuildEnvs200ResponseResponse) GetEnvListOk() (*[]CIJobEnv, bool)`

GetEnvListOk returns a tuple with the EnvList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvList

`func (o *DescribeCodingCIBuildEnvs200ResponseResponse) SetEnvList(v []CIJobEnv)`

SetEnvList sets EnvList field to given value.

### HasEnvList

`func (o *DescribeCodingCIBuildEnvs200ResponseResponse) HasEnvList() bool`

HasEnvList returns a boolean if a field has been set.

### SetEnvListNil

`func (o *DescribeCodingCIBuildEnvs200ResponseResponse) SetEnvListNil(b bool)`

 SetEnvListNil sets the value for EnvList to be an explicit nil

### UnsetEnvList
`func (o *DescribeCodingCIBuildEnvs200ResponseResponse) UnsetEnvList()`

UnsetEnvList ensures that no value is present for EnvList, not even an explicit nil
### GetRequestId

`func (o *DescribeCodingCIBuildEnvs200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeCodingCIBuildEnvs200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeCodingCIBuildEnvs200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeCodingCIBuildEnvs200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


