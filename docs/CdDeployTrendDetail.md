# CdDeployTrendDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | Pointer to **string** | 应用名称 | [optional] [default to ""]
**CdDeployTrend** | Pointer to [**[]CdDeployTrend**](CdDeployTrend.md) | 应用发布趋势统计 | [optional] 
**SuccessCount** | Pointer to **int64** | 成功发布次数 | [optional] 
**TotalCount** | Pointer to **int64** | 发布总次数 | [optional] 

## Methods

### NewCdDeployTrendDetail

`func NewCdDeployTrendDetail() *CdDeployTrendDetail`

NewCdDeployTrendDetail instantiates a new CdDeployTrendDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdDeployTrendDetailWithDefaults

`func NewCdDeployTrendDetailWithDefaults() *CdDeployTrendDetail`

NewCdDeployTrendDetailWithDefaults instantiates a new CdDeployTrendDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *CdDeployTrendDetail) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *CdDeployTrendDetail) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *CdDeployTrendDetail) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *CdDeployTrendDetail) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetCdDeployTrend

`func (o *CdDeployTrendDetail) GetCdDeployTrend() []CdDeployTrend`

GetCdDeployTrend returns the CdDeployTrend field if non-nil, zero value otherwise.

### GetCdDeployTrendOk

`func (o *CdDeployTrendDetail) GetCdDeployTrendOk() (*[]CdDeployTrend, bool)`

GetCdDeployTrendOk returns a tuple with the CdDeployTrend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdDeployTrend

`func (o *CdDeployTrendDetail) SetCdDeployTrend(v []CdDeployTrend)`

SetCdDeployTrend sets CdDeployTrend field to given value.

### HasCdDeployTrend

`func (o *CdDeployTrendDetail) HasCdDeployTrend() bool`

HasCdDeployTrend returns a boolean if a field has been set.

### GetSuccessCount

`func (o *CdDeployTrendDetail) GetSuccessCount() int64`

GetSuccessCount returns the SuccessCount field if non-nil, zero value otherwise.

### GetSuccessCountOk

`func (o *CdDeployTrendDetail) GetSuccessCountOk() (*int64, bool)`

GetSuccessCountOk returns a tuple with the SuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessCount

`func (o *CdDeployTrendDetail) SetSuccessCount(v int64)`

SetSuccessCount sets SuccessCount field to given value.

### HasSuccessCount

`func (o *CdDeployTrendDetail) HasSuccessCount() bool`

HasSuccessCount returns a boolean if a field has been set.

### GetTotalCount

`func (o *CdDeployTrendDetail) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *CdDeployTrendDetail) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *CdDeployTrendDetail) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *CdDeployTrendDetail) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


