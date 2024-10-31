# CdDeployTimeDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | Pointer to **string** | 应用名称 | [optional] [default to ""]
**CdDeployTime** | Pointer to [**CdDeployTime**](CdDeployTime.md) |  | [optional] 

## Methods

### NewCdDeployTimeDetail

`func NewCdDeployTimeDetail() *CdDeployTimeDetail`

NewCdDeployTimeDetail instantiates a new CdDeployTimeDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdDeployTimeDetailWithDefaults

`func NewCdDeployTimeDetailWithDefaults() *CdDeployTimeDetail`

NewCdDeployTimeDetailWithDefaults instantiates a new CdDeployTimeDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *CdDeployTimeDetail) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *CdDeployTimeDetail) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *CdDeployTimeDetail) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *CdDeployTimeDetail) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetCdDeployTime

`func (o *CdDeployTimeDetail) GetCdDeployTime() CdDeployTime`

GetCdDeployTime returns the CdDeployTime field if non-nil, zero value otherwise.

### GetCdDeployTimeOk

`func (o *CdDeployTimeDetail) GetCdDeployTimeOk() (*CdDeployTime, bool)`

GetCdDeployTimeOk returns a tuple with the CdDeployTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdDeployTime

`func (o *CdDeployTimeDetail) SetCdDeployTime(v CdDeployTime)`

SetCdDeployTime sets CdDeployTime field to given value.

### HasCdDeployTime

`func (o *CdDeployTimeDetail) HasCdDeployTime() bool`

HasCdDeployTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


