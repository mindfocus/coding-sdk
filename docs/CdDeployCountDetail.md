# CdDeployCountDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | Pointer to **string** | 应用名称 | [optional] [default to ""]
**CdDeployCount** | Pointer to [**CdDeployCount**](CdDeployCount.md) |  | [optional] 

## Methods

### NewCdDeployCountDetail

`func NewCdDeployCountDetail() *CdDeployCountDetail`

NewCdDeployCountDetail instantiates a new CdDeployCountDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdDeployCountDetailWithDefaults

`func NewCdDeployCountDetailWithDefaults() *CdDeployCountDetail`

NewCdDeployCountDetailWithDefaults instantiates a new CdDeployCountDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *CdDeployCountDetail) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *CdDeployCountDetail) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *CdDeployCountDetail) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *CdDeployCountDetail) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetCdDeployCount

`func (o *CdDeployCountDetail) GetCdDeployCount() CdDeployCount`

GetCdDeployCount returns the CdDeployCount field if non-nil, zero value otherwise.

### GetCdDeployCountOk

`func (o *CdDeployCountDetail) GetCdDeployCountOk() (*CdDeployCount, bool)`

GetCdDeployCountOk returns a tuple with the CdDeployCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdDeployCount

`func (o *CdDeployCountDetail) SetCdDeployCount(v CdDeployCount)`

SetCdDeployCount sets CdDeployCount field to given value.

### HasCdDeployCount

`func (o *CdDeployCountDetail) HasCdDeployCount() bool`

HasCdDeployCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


