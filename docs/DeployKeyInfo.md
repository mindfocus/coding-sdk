# DeployKeyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowWrite** | Pointer to **NullableBool** | 是否授予写入权限 | [optional] [default to false]
**CreatedAt** | Pointer to **NullableInt64** | 创建时间 | [optional] 
**DepotId** | Pointer to **NullableInt64** | 仓库 Id | [optional] 
**ExpirationDate** | Pointer to **NullableString** | 过期时间 | [optional] [default to ""]
**FingerPrint** | Pointer to **NullableString** | key 指纹 | [optional] [default to ""]
**HasExpired** | Pointer to **NullableBool** | 是否过期 | [optional] [default to false]
**KeyId** | Pointer to **NullableInt64** | SSH Key Id | [optional] 
**OwnerName** | Pointer to **NullableString** | 所属者名字 | [optional] [default to ""]
**Title** | Pointer to **NullableString** | key 标题 | [optional] [default to ""]

## Methods

### NewDeployKeyInfo

`func NewDeployKeyInfo() *DeployKeyInfo`

NewDeployKeyInfo instantiates a new DeployKeyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployKeyInfoWithDefaults

`func NewDeployKeyInfoWithDefaults() *DeployKeyInfo`

NewDeployKeyInfoWithDefaults instantiates a new DeployKeyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowWrite

`func (o *DeployKeyInfo) GetAllowWrite() bool`

GetAllowWrite returns the AllowWrite field if non-nil, zero value otherwise.

### GetAllowWriteOk

`func (o *DeployKeyInfo) GetAllowWriteOk() (*bool, bool)`

GetAllowWriteOk returns a tuple with the AllowWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowWrite

`func (o *DeployKeyInfo) SetAllowWrite(v bool)`

SetAllowWrite sets AllowWrite field to given value.

### HasAllowWrite

`func (o *DeployKeyInfo) HasAllowWrite() bool`

HasAllowWrite returns a boolean if a field has been set.

### SetAllowWriteNil

`func (o *DeployKeyInfo) SetAllowWriteNil(b bool)`

 SetAllowWriteNil sets the value for AllowWrite to be an explicit nil

### UnsetAllowWrite
`func (o *DeployKeyInfo) UnsetAllowWrite()`

UnsetAllowWrite ensures that no value is present for AllowWrite, not even an explicit nil
### GetCreatedAt

`func (o *DeployKeyInfo) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DeployKeyInfo) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DeployKeyInfo) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DeployKeyInfo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *DeployKeyInfo) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *DeployKeyInfo) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDepotId

`func (o *DeployKeyInfo) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *DeployKeyInfo) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *DeployKeyInfo) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.

### HasDepotId

`func (o *DeployKeyInfo) HasDepotId() bool`

HasDepotId returns a boolean if a field has been set.

### SetDepotIdNil

`func (o *DeployKeyInfo) SetDepotIdNil(b bool)`

 SetDepotIdNil sets the value for DepotId to be an explicit nil

### UnsetDepotId
`func (o *DeployKeyInfo) UnsetDepotId()`

UnsetDepotId ensures that no value is present for DepotId, not even an explicit nil
### GetExpirationDate

`func (o *DeployKeyInfo) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *DeployKeyInfo) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *DeployKeyInfo) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *DeployKeyInfo) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### SetExpirationDateNil

`func (o *DeployKeyInfo) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *DeployKeyInfo) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
### GetFingerPrint

`func (o *DeployKeyInfo) GetFingerPrint() string`

GetFingerPrint returns the FingerPrint field if non-nil, zero value otherwise.

### GetFingerPrintOk

`func (o *DeployKeyInfo) GetFingerPrintOk() (*string, bool)`

GetFingerPrintOk returns a tuple with the FingerPrint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerPrint

`func (o *DeployKeyInfo) SetFingerPrint(v string)`

SetFingerPrint sets FingerPrint field to given value.

### HasFingerPrint

`func (o *DeployKeyInfo) HasFingerPrint() bool`

HasFingerPrint returns a boolean if a field has been set.

### SetFingerPrintNil

`func (o *DeployKeyInfo) SetFingerPrintNil(b bool)`

 SetFingerPrintNil sets the value for FingerPrint to be an explicit nil

### UnsetFingerPrint
`func (o *DeployKeyInfo) UnsetFingerPrint()`

UnsetFingerPrint ensures that no value is present for FingerPrint, not even an explicit nil
### GetHasExpired

`func (o *DeployKeyInfo) GetHasExpired() bool`

GetHasExpired returns the HasExpired field if non-nil, zero value otherwise.

### GetHasExpiredOk

`func (o *DeployKeyInfo) GetHasExpiredOk() (*bool, bool)`

GetHasExpiredOk returns a tuple with the HasExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasExpired

`func (o *DeployKeyInfo) SetHasExpired(v bool)`

SetHasExpired sets HasExpired field to given value.

### HasHasExpired

`func (o *DeployKeyInfo) HasHasExpired() bool`

HasHasExpired returns a boolean if a field has been set.

### SetHasExpiredNil

`func (o *DeployKeyInfo) SetHasExpiredNil(b bool)`

 SetHasExpiredNil sets the value for HasExpired to be an explicit nil

### UnsetHasExpired
`func (o *DeployKeyInfo) UnsetHasExpired()`

UnsetHasExpired ensures that no value is present for HasExpired, not even an explicit nil
### GetKeyId

`func (o *DeployKeyInfo) GetKeyId() int64`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *DeployKeyInfo) GetKeyIdOk() (*int64, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *DeployKeyInfo) SetKeyId(v int64)`

SetKeyId sets KeyId field to given value.

### HasKeyId

`func (o *DeployKeyInfo) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.

### SetKeyIdNil

`func (o *DeployKeyInfo) SetKeyIdNil(b bool)`

 SetKeyIdNil sets the value for KeyId to be an explicit nil

### UnsetKeyId
`func (o *DeployKeyInfo) UnsetKeyId()`

UnsetKeyId ensures that no value is present for KeyId, not even an explicit nil
### GetOwnerName

`func (o *DeployKeyInfo) GetOwnerName() string`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *DeployKeyInfo) GetOwnerNameOk() (*string, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *DeployKeyInfo) SetOwnerName(v string)`

SetOwnerName sets OwnerName field to given value.

### HasOwnerName

`func (o *DeployKeyInfo) HasOwnerName() bool`

HasOwnerName returns a boolean if a field has been set.

### SetOwnerNameNil

`func (o *DeployKeyInfo) SetOwnerNameNil(b bool)`

 SetOwnerNameNil sets the value for OwnerName to be an explicit nil

### UnsetOwnerName
`func (o *DeployKeyInfo) UnsetOwnerName()`

UnsetOwnerName ensures that no value is present for OwnerName, not even an explicit nil
### GetTitle

`func (o *DeployKeyInfo) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DeployKeyInfo) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DeployKeyInfo) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DeployKeyInfo) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *DeployKeyInfo) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *DeployKeyInfo) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


