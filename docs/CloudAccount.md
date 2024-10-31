# CloudAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | Pointer to **string** | 云账号类型 | [optional] [default to ""]
**Credential** | Pointer to [**CloudAccountCredential**](CloudAccountCredential.md) |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** | 云账号状态错误信息 | [optional] [default to ""]
**Id** | Pointer to **int64** | 云账号 ID | [optional] 
**Name** | Pointer to **string** | 云账号名称 | [optional] [default to ""]
**Status** | Pointer to **string** | 云账号状态 | [optional] [default to ""]

## Methods

### NewCloudAccount

`func NewCloudAccount() *CloudAccount`

NewCloudAccount instantiates a new CloudAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAccountWithDefaults

`func NewCloudAccountWithDefaults() *CloudAccount`

NewCloudAccountWithDefaults instantiates a new CloudAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *CloudAccount) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *CloudAccount) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *CloudAccount) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *CloudAccount) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetCredential

`func (o *CloudAccount) GetCredential() CloudAccountCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *CloudAccount) GetCredentialOk() (*CloudAccountCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *CloudAccount) SetCredential(v CloudAccountCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *CloudAccount) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetErrorMessage

`func (o *CloudAccount) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *CloudAccount) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *CloudAccount) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *CloudAccount) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *CloudAccount) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *CloudAccount) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetId

`func (o *CloudAccount) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudAccount) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudAccount) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CloudAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CloudAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudAccount) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudAccount) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *CloudAccount) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudAccount) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudAccount) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudAccount) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


