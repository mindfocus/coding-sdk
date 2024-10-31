# Token

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthToken** | Pointer to **string** | 验证文件的Token（用于导入wiki zip import 使用） | [optional] [default to ""]
**Provider** | Pointer to **string** | cos存储对象 | [optional] [default to ""]
**SecretId** | Pointer to **string** | cos上传的Id | [optional] [default to ""]
**SecretKey** | Pointer to **string** | cos上传的key （用于导入wiki zip import 使用） | [optional] [default to ""]
**Time** | Pointer to **int32** | 获取token 的时间（用于导入wiki zip import使用） | [optional] 
**UpToken** | Pointer to **string** | 上传文件的Token | [optional] [default to ""]
**UploadLink** | Pointer to **string** | 上传地址 | [optional] [default to ""]

## Methods

### NewToken

`func NewToken() *Token`

NewToken instantiates a new Token object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenWithDefaults

`func NewTokenWithDefaults() *Token`

NewTokenWithDefaults instantiates a new Token object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthToken

`func (o *Token) GetAuthToken() string`

GetAuthToken returns the AuthToken field if non-nil, zero value otherwise.

### GetAuthTokenOk

`func (o *Token) GetAuthTokenOk() (*string, bool)`

GetAuthTokenOk returns a tuple with the AuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthToken

`func (o *Token) SetAuthToken(v string)`

SetAuthToken sets AuthToken field to given value.

### HasAuthToken

`func (o *Token) HasAuthToken() bool`

HasAuthToken returns a boolean if a field has been set.

### GetProvider

`func (o *Token) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *Token) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *Token) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *Token) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetSecretId

`func (o *Token) GetSecretId() string`

GetSecretId returns the SecretId field if non-nil, zero value otherwise.

### GetSecretIdOk

`func (o *Token) GetSecretIdOk() (*string, bool)`

GetSecretIdOk returns a tuple with the SecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretId

`func (o *Token) SetSecretId(v string)`

SetSecretId sets SecretId field to given value.

### HasSecretId

`func (o *Token) HasSecretId() bool`

HasSecretId returns a boolean if a field has been set.

### GetSecretKey

`func (o *Token) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *Token) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *Token) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *Token) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetTime

`func (o *Token) GetTime() int32`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *Token) GetTimeOk() (*int32, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *Token) SetTime(v int32)`

SetTime sets Time field to given value.

### HasTime

`func (o *Token) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetUpToken

`func (o *Token) GetUpToken() string`

GetUpToken returns the UpToken field if non-nil, zero value otherwise.

### GetUpTokenOk

`func (o *Token) GetUpTokenOk() (*string, bool)`

GetUpTokenOk returns a tuple with the UpToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpToken

`func (o *Token) SetUpToken(v string)`

SetUpToken sets UpToken field to given value.

### HasUpToken

`func (o *Token) HasUpToken() bool`

HasUpToken returns a boolean if a field has been set.

### GetUploadLink

`func (o *Token) GetUploadLink() string`

GetUploadLink returns the UploadLink field if non-nil, zero value otherwise.

### GetUploadLinkOk

`func (o *Token) GetUploadLinkOk() (*string, bool)`

GetUploadLinkOk returns a tuple with the UploadLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadLink

`func (o *Token) SetUploadLink(v string)`

SetUploadLink sets UploadLink field to given value.

### HasUploadLink

`func (o *Token) HasUploadLink() bool`

HasUploadLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


