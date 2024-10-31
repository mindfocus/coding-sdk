# ModifyCdCloudAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credential** | [**CloudAccountCredential**](CloudAccountCredential.md) |  | 
**Id** | **int64** | 云账号 ID | 
**Name** | **string** | 云账号名称 | 

## Methods

### NewModifyCdCloudAccountRequest

`func NewModifyCdCloudAccountRequest(credential CloudAccountCredential, id int64, name string, ) *ModifyCdCloudAccountRequest`

NewModifyCdCloudAccountRequest instantiates a new ModifyCdCloudAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyCdCloudAccountRequestWithDefaults

`func NewModifyCdCloudAccountRequestWithDefaults() *ModifyCdCloudAccountRequest`

NewModifyCdCloudAccountRequestWithDefaults instantiates a new ModifyCdCloudAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredential

`func (o *ModifyCdCloudAccountRequest) GetCredential() CloudAccountCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *ModifyCdCloudAccountRequest) GetCredentialOk() (*CloudAccountCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *ModifyCdCloudAccountRequest) SetCredential(v CloudAccountCredential)`

SetCredential sets Credential field to given value.


### GetId

`func (o *ModifyCdCloudAccountRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModifyCdCloudAccountRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModifyCdCloudAccountRequest) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *ModifyCdCloudAccountRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModifyCdCloudAccountRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModifyCdCloudAccountRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


