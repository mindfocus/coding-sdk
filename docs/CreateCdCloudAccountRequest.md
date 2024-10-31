# CreateCdCloudAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | **string** | 云账号类型（可选值：KUBERNETES、TENCENT） | 
**Credential** | [**CloudAccountCredential**](CloudAccountCredential.md) |  | 
**Name** | **string** | 云账号名称 | 

## Methods

### NewCreateCdCloudAccountRequest

`func NewCreateCdCloudAccountRequest(cloudProvider string, credential CloudAccountCredential, name string, ) *CreateCdCloudAccountRequest`

NewCreateCdCloudAccountRequest instantiates a new CreateCdCloudAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCdCloudAccountRequestWithDefaults

`func NewCreateCdCloudAccountRequestWithDefaults() *CreateCdCloudAccountRequest`

NewCreateCdCloudAccountRequestWithDefaults instantiates a new CreateCdCloudAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *CreateCdCloudAccountRequest) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *CreateCdCloudAccountRequest) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *CreateCdCloudAccountRequest) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetCredential

`func (o *CreateCdCloudAccountRequest) GetCredential() CloudAccountCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *CreateCdCloudAccountRequest) GetCredentialOk() (*CloudAccountCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *CreateCdCloudAccountRequest) SetCredential(v CloudAccountCredential)`

SetCredential sets Credential field to given value.


### GetName

`func (o *CreateCdCloudAccountRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCdCloudAccountRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCdCloudAccountRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


