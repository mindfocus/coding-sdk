# ApiIssueIssueAttachmentPreSign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthToken** | **string** | 认证凭据 | [default to ""]
**CosSecurityToken** | **string** | cos的临时token | [default to ""]
**StorageKey** | **string** | 文件名 | [default to ""]
**Time** | **int64** | 时间 | [default to 0]
**UploadLink** | **string** | 上传链接 | [default to ""]

## Methods

### NewApiIssueIssueAttachmentPreSign

`func NewApiIssueIssueAttachmentPreSign(authToken string, cosSecurityToken string, storageKey string, time int64, uploadLink string, ) *ApiIssueIssueAttachmentPreSign`

NewApiIssueIssueAttachmentPreSign instantiates a new ApiIssueIssueAttachmentPreSign object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiIssueIssueAttachmentPreSignWithDefaults

`func NewApiIssueIssueAttachmentPreSignWithDefaults() *ApiIssueIssueAttachmentPreSign`

NewApiIssueIssueAttachmentPreSignWithDefaults instantiates a new ApiIssueIssueAttachmentPreSign object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthToken

`func (o *ApiIssueIssueAttachmentPreSign) GetAuthToken() string`

GetAuthToken returns the AuthToken field if non-nil, zero value otherwise.

### GetAuthTokenOk

`func (o *ApiIssueIssueAttachmentPreSign) GetAuthTokenOk() (*string, bool)`

GetAuthTokenOk returns a tuple with the AuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthToken

`func (o *ApiIssueIssueAttachmentPreSign) SetAuthToken(v string)`

SetAuthToken sets AuthToken field to given value.


### GetCosSecurityToken

`func (o *ApiIssueIssueAttachmentPreSign) GetCosSecurityToken() string`

GetCosSecurityToken returns the CosSecurityToken field if non-nil, zero value otherwise.

### GetCosSecurityTokenOk

`func (o *ApiIssueIssueAttachmentPreSign) GetCosSecurityTokenOk() (*string, bool)`

GetCosSecurityTokenOk returns a tuple with the CosSecurityToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosSecurityToken

`func (o *ApiIssueIssueAttachmentPreSign) SetCosSecurityToken(v string)`

SetCosSecurityToken sets CosSecurityToken field to given value.


### GetStorageKey

`func (o *ApiIssueIssueAttachmentPreSign) GetStorageKey() string`

GetStorageKey returns the StorageKey field if non-nil, zero value otherwise.

### GetStorageKeyOk

`func (o *ApiIssueIssueAttachmentPreSign) GetStorageKeyOk() (*string, bool)`

GetStorageKeyOk returns a tuple with the StorageKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageKey

`func (o *ApiIssueIssueAttachmentPreSign) SetStorageKey(v string)`

SetStorageKey sets StorageKey field to given value.


### GetTime

`func (o *ApiIssueIssueAttachmentPreSign) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ApiIssueIssueAttachmentPreSign) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ApiIssueIssueAttachmentPreSign) SetTime(v int64)`

SetTime sets Time field to given value.


### GetUploadLink

`func (o *ApiIssueIssueAttachmentPreSign) GetUploadLink() string`

GetUploadLink returns the UploadLink field if non-nil, zero value otherwise.

### GetUploadLinkOk

`func (o *ApiIssueIssueAttachmentPreSign) GetUploadLinkOk() (*string, bool)`

GetUploadLinkOk returns a tuple with the UploadLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadLink

`func (o *ApiIssueIssueAttachmentPreSign) SetUploadLink(v string)`

SetUploadLink sets UploadLink field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


