# MergeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotId** | Pointer to **int64** | 仓库id | [optional] 
**MergeRequestId** | Pointer to **int64** | 合并请求id | [optional] 
**MergeRequestInfo** | Pointer to [**MergeRequestInfo**](MergeRequestInfo.md) |  | [optional] 
**MergeRequestUrl** | Pointer to **NullableString** | 合并请求URl | [optional] [default to ""]
**ProjectId** | Pointer to **int64** | 项目id | [optional] 

## Methods

### NewMergeInfo

`func NewMergeInfo() *MergeInfo`

NewMergeInfo instantiates a new MergeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMergeInfoWithDefaults

`func NewMergeInfoWithDefaults() *MergeInfo`

NewMergeInfoWithDefaults instantiates a new MergeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotId

`func (o *MergeInfo) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *MergeInfo) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *MergeInfo) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.

### HasDepotId

`func (o *MergeInfo) HasDepotId() bool`

HasDepotId returns a boolean if a field has been set.

### GetMergeRequestId

`func (o *MergeInfo) GetMergeRequestId() int64`

GetMergeRequestId returns the MergeRequestId field if non-nil, zero value otherwise.

### GetMergeRequestIdOk

`func (o *MergeInfo) GetMergeRequestIdOk() (*int64, bool)`

GetMergeRequestIdOk returns a tuple with the MergeRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeRequestId

`func (o *MergeInfo) SetMergeRequestId(v int64)`

SetMergeRequestId sets MergeRequestId field to given value.

### HasMergeRequestId

`func (o *MergeInfo) HasMergeRequestId() bool`

HasMergeRequestId returns a boolean if a field has been set.

### GetMergeRequestInfo

`func (o *MergeInfo) GetMergeRequestInfo() MergeRequestInfo`

GetMergeRequestInfo returns the MergeRequestInfo field if non-nil, zero value otherwise.

### GetMergeRequestInfoOk

`func (o *MergeInfo) GetMergeRequestInfoOk() (*MergeRequestInfo, bool)`

GetMergeRequestInfoOk returns a tuple with the MergeRequestInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeRequestInfo

`func (o *MergeInfo) SetMergeRequestInfo(v MergeRequestInfo)`

SetMergeRequestInfo sets MergeRequestInfo field to given value.

### HasMergeRequestInfo

`func (o *MergeInfo) HasMergeRequestInfo() bool`

HasMergeRequestInfo returns a boolean if a field has been set.

### GetMergeRequestUrl

`func (o *MergeInfo) GetMergeRequestUrl() string`

GetMergeRequestUrl returns the MergeRequestUrl field if non-nil, zero value otherwise.

### GetMergeRequestUrlOk

`func (o *MergeInfo) GetMergeRequestUrlOk() (*string, bool)`

GetMergeRequestUrlOk returns a tuple with the MergeRequestUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeRequestUrl

`func (o *MergeInfo) SetMergeRequestUrl(v string)`

SetMergeRequestUrl sets MergeRequestUrl field to given value.

### HasMergeRequestUrl

`func (o *MergeInfo) HasMergeRequestUrl() bool`

HasMergeRequestUrl returns a boolean if a field has been set.

### SetMergeRequestUrlNil

`func (o *MergeInfo) SetMergeRequestUrlNil(b bool)`

 SetMergeRequestUrlNil sets the value for MergeRequestUrl to be an explicit nil

### UnsetMergeRequestUrl
`func (o *MergeInfo) UnsetMergeRequestUrl()`

UnsetMergeRequestUrl ensures that no value is present for MergeRequestUrl, not even an explicit nil
### GetProjectId

`func (o *MergeInfo) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *MergeInfo) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *MergeInfo) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *MergeInfo) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


