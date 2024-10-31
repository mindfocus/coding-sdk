# SshKeyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int32** | 创建时间 | [optional] 
**ExpirationDate** | Pointer to **string** | 过期时间 | [optional] [default to ""]
**FingerPrint** | Pointer to **string** | 指纹信息 | [optional] [default to ""]
**HasExpired** | Pointer to **bool** | 是否过期 | [optional] [default to false]
**Id** | Pointer to **int32** | 公钥Id | [optional] 
**OwnerId** | Pointer to **int32** | 公钥所属者Id | [optional] 
**Title** | Pointer to **string** | 公钥标题 | [optional] [default to ""]

## Methods

### NewSshKeyInfo

`func NewSshKeyInfo() *SshKeyInfo`

NewSshKeyInfo instantiates a new SshKeyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshKeyInfoWithDefaults

`func NewSshKeyInfoWithDefaults() *SshKeyInfo`

NewSshKeyInfoWithDefaults instantiates a new SshKeyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *SshKeyInfo) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SshKeyInfo) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SshKeyInfo) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SshKeyInfo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExpirationDate

`func (o *SshKeyInfo) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *SshKeyInfo) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *SshKeyInfo) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *SshKeyInfo) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetFingerPrint

`func (o *SshKeyInfo) GetFingerPrint() string`

GetFingerPrint returns the FingerPrint field if non-nil, zero value otherwise.

### GetFingerPrintOk

`func (o *SshKeyInfo) GetFingerPrintOk() (*string, bool)`

GetFingerPrintOk returns a tuple with the FingerPrint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerPrint

`func (o *SshKeyInfo) SetFingerPrint(v string)`

SetFingerPrint sets FingerPrint field to given value.

### HasFingerPrint

`func (o *SshKeyInfo) HasFingerPrint() bool`

HasFingerPrint returns a boolean if a field has been set.

### GetHasExpired

`func (o *SshKeyInfo) GetHasExpired() bool`

GetHasExpired returns the HasExpired field if non-nil, zero value otherwise.

### GetHasExpiredOk

`func (o *SshKeyInfo) GetHasExpiredOk() (*bool, bool)`

GetHasExpiredOk returns a tuple with the HasExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasExpired

`func (o *SshKeyInfo) SetHasExpired(v bool)`

SetHasExpired sets HasExpired field to given value.

### HasHasExpired

`func (o *SshKeyInfo) HasHasExpired() bool`

HasHasExpired returns a boolean if a field has been set.

### GetId

`func (o *SshKeyInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SshKeyInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SshKeyInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SshKeyInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOwnerId

`func (o *SshKeyInfo) GetOwnerId() int32`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *SshKeyInfo) GetOwnerIdOk() (*int32, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *SshKeyInfo) SetOwnerId(v int32)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *SshKeyInfo) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetTitle

`func (o *SshKeyInfo) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SshKeyInfo) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SshKeyInfo) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SshKeyInfo) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


