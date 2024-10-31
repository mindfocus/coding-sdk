# DeleteSshKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SshKeyId** | **int32** | ssh id | 

## Methods

### NewDeleteSshKeyRequest

`func NewDeleteSshKeyRequest(sshKeyId int32, ) *DeleteSshKeyRequest`

NewDeleteSshKeyRequest instantiates a new DeleteSshKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteSshKeyRequestWithDefaults

`func NewDeleteSshKeyRequestWithDefaults() *DeleteSshKeyRequest`

NewDeleteSshKeyRequestWithDefaults instantiates a new DeleteSshKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSshKeyId

`func (o *DeleteSshKeyRequest) GetSshKeyId() int32`

GetSshKeyId returns the SshKeyId field if non-nil, zero value otherwise.

### GetSshKeyIdOk

`func (o *DeleteSshKeyRequest) GetSshKeyIdOk() (*int32, bool)`

GetSshKeyIdOk returns a tuple with the SshKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyId

`func (o *DeleteSshKeyRequest) SetSshKeyId(v int32)`

SetSshKeyId sets SshKeyId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


