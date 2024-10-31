# DeleteMemberSshKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MemberUserId** | **int64** | 成员 Id | 
**SshKeyId** | **int32** | ssh id | 

## Methods

### NewDeleteMemberSshKeyRequest

`func NewDeleteMemberSshKeyRequest(memberUserId int64, sshKeyId int32, ) *DeleteMemberSshKeyRequest`

NewDeleteMemberSshKeyRequest instantiates a new DeleteMemberSshKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteMemberSshKeyRequestWithDefaults

`func NewDeleteMemberSshKeyRequestWithDefaults() *DeleteMemberSshKeyRequest`

NewDeleteMemberSshKeyRequestWithDefaults instantiates a new DeleteMemberSshKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemberUserId

`func (o *DeleteMemberSshKeyRequest) GetMemberUserId() int64`

GetMemberUserId returns the MemberUserId field if non-nil, zero value otherwise.

### GetMemberUserIdOk

`func (o *DeleteMemberSshKeyRequest) GetMemberUserIdOk() (*int64, bool)`

GetMemberUserIdOk returns a tuple with the MemberUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberUserId

`func (o *DeleteMemberSshKeyRequest) SetMemberUserId(v int64)`

SetMemberUserId sets MemberUserId field to given value.


### GetSshKeyId

`func (o *DeleteMemberSshKeyRequest) GetSshKeyId() int32`

GetSshKeyId returns the SshKeyId field if non-nil, zero value otherwise.

### GetSshKeyIdOk

`func (o *DeleteMemberSshKeyRequest) GetSshKeyIdOk() (*int32, bool)`

GetSshKeyIdOk returns a tuple with the SshKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyId

`func (o *DeleteMemberSshKeyRequest) SetSshKeyId(v int32)`

SetSshKeyId sets SshKeyId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


