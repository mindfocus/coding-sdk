# DescribeMemberSshKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MemberUserId** | **int64** | 成员 Id | 
**SshKeyId** | Pointer to **int64** | SSH 公钥 Id | [optional] 

## Methods

### NewDescribeMemberSshKeyRequest

`func NewDescribeMemberSshKeyRequest(memberUserId int64, ) *DescribeMemberSshKeyRequest`

NewDescribeMemberSshKeyRequest instantiates a new DescribeMemberSshKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeMemberSshKeyRequestWithDefaults

`func NewDescribeMemberSshKeyRequestWithDefaults() *DescribeMemberSshKeyRequest`

NewDescribeMemberSshKeyRequestWithDefaults instantiates a new DescribeMemberSshKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemberUserId

`func (o *DescribeMemberSshKeyRequest) GetMemberUserId() int64`

GetMemberUserId returns the MemberUserId field if non-nil, zero value otherwise.

### GetMemberUserIdOk

`func (o *DescribeMemberSshKeyRequest) GetMemberUserIdOk() (*int64, bool)`

GetMemberUserIdOk returns a tuple with the MemberUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberUserId

`func (o *DescribeMemberSshKeyRequest) SetMemberUserId(v int64)`

SetMemberUserId sets MemberUserId field to given value.


### GetSshKeyId

`func (o *DescribeMemberSshKeyRequest) GetSshKeyId() int64`

GetSshKeyId returns the SshKeyId field if non-nil, zero value otherwise.

### GetSshKeyIdOk

`func (o *DescribeMemberSshKeyRequest) GetSshKeyIdOk() (*int64, bool)`

GetSshKeyIdOk returns a tuple with the SshKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyId

`func (o *DescribeMemberSshKeyRequest) SetSshKeyId(v int64)`

SetSshKeyId sets SshKeyId field to given value.

### HasSshKeyId

`func (o *DescribeMemberSshKeyRequest) HasSshKeyId() bool`

HasSshKeyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


