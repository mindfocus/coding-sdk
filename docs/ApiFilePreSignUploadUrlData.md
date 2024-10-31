# ApiFilePreSignUploadUrlData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthToken** | **string** | token信息 | [default to ""]
**Headers** | **string** | 请求头信息 | [default to ""]
**StorageKey** | **string** | 文件存储的key信息 | [default to ""]
**UploadLink** | **string** | 上传链接 | [default to ""]

## Methods

### NewApiFilePreSignUploadUrlData

`func NewApiFilePreSignUploadUrlData(authToken string, headers string, storageKey string, uploadLink string, ) *ApiFilePreSignUploadUrlData`

NewApiFilePreSignUploadUrlData instantiates a new ApiFilePreSignUploadUrlData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiFilePreSignUploadUrlDataWithDefaults

`func NewApiFilePreSignUploadUrlDataWithDefaults() *ApiFilePreSignUploadUrlData`

NewApiFilePreSignUploadUrlDataWithDefaults instantiates a new ApiFilePreSignUploadUrlData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthToken

`func (o *ApiFilePreSignUploadUrlData) GetAuthToken() string`

GetAuthToken returns the AuthToken field if non-nil, zero value otherwise.

### GetAuthTokenOk

`func (o *ApiFilePreSignUploadUrlData) GetAuthTokenOk() (*string, bool)`

GetAuthTokenOk returns a tuple with the AuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthToken

`func (o *ApiFilePreSignUploadUrlData) SetAuthToken(v string)`

SetAuthToken sets AuthToken field to given value.


### GetHeaders

`func (o *ApiFilePreSignUploadUrlData) GetHeaders() string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *ApiFilePreSignUploadUrlData) GetHeadersOk() (*string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *ApiFilePreSignUploadUrlData) SetHeaders(v string)`

SetHeaders sets Headers field to given value.


### GetStorageKey

`func (o *ApiFilePreSignUploadUrlData) GetStorageKey() string`

GetStorageKey returns the StorageKey field if non-nil, zero value otherwise.

### GetStorageKeyOk

`func (o *ApiFilePreSignUploadUrlData) GetStorageKeyOk() (*string, bool)`

GetStorageKeyOk returns a tuple with the StorageKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageKey

`func (o *ApiFilePreSignUploadUrlData) SetStorageKey(v string)`

SetStorageKey sets StorageKey field to given value.


### GetUploadLink

`func (o *ApiFilePreSignUploadUrlData) GetUploadLink() string`

GetUploadLink returns the UploadLink field if non-nil, zero value otherwise.

### GetUploadLinkOk

`func (o *ApiFilePreSignUploadUrlData) GetUploadLinkOk() (*string, bool)`

GetUploadLinkOk returns a tuple with the UploadLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadLink

`func (o *ApiFilePreSignUploadUrlData) SetUploadLink(v string)`

SetUploadLink sets UploadLink field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


