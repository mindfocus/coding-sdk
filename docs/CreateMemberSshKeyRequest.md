# CreateMemberSshKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** | 公钥内容 | 
**ExpirationDate** | Pointer to **string** | 过期时间，不填为永不过期 9999-12-31 | [optional] 
**MemberUserId** | **int64** | 成员 Id | 
**Title** | **string** | 导入的 ssh key 标识名 | 

## Methods

### NewCreateMemberSshKeyRequest

`func NewCreateMemberSshKeyRequest(content string, memberUserId int64, title string, ) *CreateMemberSshKeyRequest`

NewCreateMemberSshKeyRequest instantiates a new CreateMemberSshKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMemberSshKeyRequestWithDefaults

`func NewCreateMemberSshKeyRequestWithDefaults() *CreateMemberSshKeyRequest`

NewCreateMemberSshKeyRequestWithDefaults instantiates a new CreateMemberSshKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *CreateMemberSshKeyRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CreateMemberSshKeyRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CreateMemberSshKeyRequest) SetContent(v string)`

SetContent sets Content field to given value.


### GetExpirationDate

`func (o *CreateMemberSshKeyRequest) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *CreateMemberSshKeyRequest) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *CreateMemberSshKeyRequest) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *CreateMemberSshKeyRequest) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetMemberUserId

`func (o *CreateMemberSshKeyRequest) GetMemberUserId() int64`

GetMemberUserId returns the MemberUserId field if non-nil, zero value otherwise.

### GetMemberUserIdOk

`func (o *CreateMemberSshKeyRequest) GetMemberUserIdOk() (*int64, bool)`

GetMemberUserIdOk returns a tuple with the MemberUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberUserId

`func (o *CreateMemberSshKeyRequest) SetMemberUserId(v int64)`

SetMemberUserId sets MemberUserId field to given value.


### GetTitle

`func (o *CreateMemberSshKeyRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateMemberSshKeyRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateMemberSshKeyRequest) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


