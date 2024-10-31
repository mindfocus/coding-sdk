# CreateWikiByZipRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthToken** | **string** | 验证文件的token | 
**FileName** | **string** | 文件名 | 
**Key** | **string** | 上传文件的uuid名称 b5d0d8e0-3aca-11eb-8673-a9b6d94ca755.png | 
**ParentIid** | **int64** | 父级iid | 
**ProjectName** | **string** | 项目名称 | 
**Time** | **int32** | 获取token的时间 | 

## Methods

### NewCreateWikiByZipRequest

`func NewCreateWikiByZipRequest(authToken string, fileName string, key string, parentIid int64, projectName string, time int32, ) *CreateWikiByZipRequest`

NewCreateWikiByZipRequest instantiates a new CreateWikiByZipRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWikiByZipRequestWithDefaults

`func NewCreateWikiByZipRequestWithDefaults() *CreateWikiByZipRequest`

NewCreateWikiByZipRequestWithDefaults instantiates a new CreateWikiByZipRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthToken

`func (o *CreateWikiByZipRequest) GetAuthToken() string`

GetAuthToken returns the AuthToken field if non-nil, zero value otherwise.

### GetAuthTokenOk

`func (o *CreateWikiByZipRequest) GetAuthTokenOk() (*string, bool)`

GetAuthTokenOk returns a tuple with the AuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthToken

`func (o *CreateWikiByZipRequest) SetAuthToken(v string)`

SetAuthToken sets AuthToken field to given value.


### GetFileName

`func (o *CreateWikiByZipRequest) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *CreateWikiByZipRequest) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *CreateWikiByZipRequest) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetKey

`func (o *CreateWikiByZipRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateWikiByZipRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateWikiByZipRequest) SetKey(v string)`

SetKey sets Key field to given value.


### GetParentIid

`func (o *CreateWikiByZipRequest) GetParentIid() int64`

GetParentIid returns the ParentIid field if non-nil, zero value otherwise.

### GetParentIidOk

`func (o *CreateWikiByZipRequest) GetParentIidOk() (*int64, bool)`

GetParentIidOk returns a tuple with the ParentIid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentIid

`func (o *CreateWikiByZipRequest) SetParentIid(v int64)`

SetParentIid sets ParentIid field to given value.


### GetProjectName

`func (o *CreateWikiByZipRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *CreateWikiByZipRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *CreateWikiByZipRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetTime

`func (o *CreateWikiByZipRequest) GetTime() int32`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *CreateWikiByZipRequest) GetTimeOk() (*int32, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *CreateWikiByZipRequest) SetTime(v int32)`

SetTime sets Time field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


