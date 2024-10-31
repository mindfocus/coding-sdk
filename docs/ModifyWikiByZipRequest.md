# ModifyWikiByZipRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthToken** | **string** | 验证文件的token | 
**FileName** | **string** | 文件名 | 
**Iid** | **int64** | wiki的Iid | 
**Key** | **string** | 上传文件的uuid名称 b5d0d8e0-3aca-11eb-8673-a9b6d94ca755.png | 
**ProjectName** | **string** | 项目名称 | 
**Time** | **int32** | 获取token的时间 | 

## Methods

### NewModifyWikiByZipRequest

`func NewModifyWikiByZipRequest(authToken string, fileName string, iid int64, key string, projectName string, time int32, ) *ModifyWikiByZipRequest`

NewModifyWikiByZipRequest instantiates a new ModifyWikiByZipRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyWikiByZipRequestWithDefaults

`func NewModifyWikiByZipRequestWithDefaults() *ModifyWikiByZipRequest`

NewModifyWikiByZipRequestWithDefaults instantiates a new ModifyWikiByZipRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthToken

`func (o *ModifyWikiByZipRequest) GetAuthToken() string`

GetAuthToken returns the AuthToken field if non-nil, zero value otherwise.

### GetAuthTokenOk

`func (o *ModifyWikiByZipRequest) GetAuthTokenOk() (*string, bool)`

GetAuthTokenOk returns a tuple with the AuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthToken

`func (o *ModifyWikiByZipRequest) SetAuthToken(v string)`

SetAuthToken sets AuthToken field to given value.


### GetFileName

`func (o *ModifyWikiByZipRequest) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *ModifyWikiByZipRequest) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *ModifyWikiByZipRequest) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetIid

`func (o *ModifyWikiByZipRequest) GetIid() int64`

GetIid returns the Iid field if non-nil, zero value otherwise.

### GetIidOk

`func (o *ModifyWikiByZipRequest) GetIidOk() (*int64, bool)`

GetIidOk returns a tuple with the Iid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIid

`func (o *ModifyWikiByZipRequest) SetIid(v int64)`

SetIid sets Iid field to given value.


### GetKey

`func (o *ModifyWikiByZipRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ModifyWikiByZipRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ModifyWikiByZipRequest) SetKey(v string)`

SetKey sets Key field to given value.


### GetProjectName

`func (o *ModifyWikiByZipRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *ModifyWikiByZipRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *ModifyWikiByZipRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetTime

`func (o *ModifyWikiByZipRequest) GetTime() int32`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ModifyWikiByZipRequest) GetTimeOk() (*int32, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ModifyWikiByZipRequest) SetTime(v int32)`

SetTime sets Time field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


