# CreateFileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthToken** | **string** | 获取预签名URL接口返回的 AuthToken | [default to ""]
**StorageKey** | **string** | 获取预签名URL接口返回的 StorageKey | [default to ""]

## Methods

### NewCreateFileRequest

`func NewCreateFileRequest(authToken string, storageKey string, ) *CreateFileRequest`

NewCreateFileRequest instantiates a new CreateFileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFileRequestWithDefaults

`func NewCreateFileRequestWithDefaults() *CreateFileRequest`

NewCreateFileRequestWithDefaults instantiates a new CreateFileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthToken

`func (o *CreateFileRequest) GetAuthToken() string`

GetAuthToken returns the AuthToken field if non-nil, zero value otherwise.

### GetAuthTokenOk

`func (o *CreateFileRequest) GetAuthTokenOk() (*string, bool)`

GetAuthTokenOk returns a tuple with the AuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthToken

`func (o *CreateFileRequest) SetAuthToken(v string)`

SetAuthToken sets AuthToken field to given value.


### GetStorageKey

`func (o *CreateFileRequest) GetStorageKey() string`

GetStorageKey returns the StorageKey field if non-nil, zero value otherwise.

### GetStorageKeyOk

`func (o *CreateFileRequest) GetStorageKeyOk() (*string, bool)`

GetStorageKeyOk returns a tuple with the StorageKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageKey

`func (o *CreateFileRequest) SetStorageKey(v string)`

SetStorageKey sets StorageKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


