# CreateMemberSshKey200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SshKeyInfo** | Pointer to [**SshKeyInfo**](SshKeyInfo.md) |  | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewCreateMemberSshKey200ResponseResponse

`func NewCreateMemberSshKey200ResponseResponse() *CreateMemberSshKey200ResponseResponse`

NewCreateMemberSshKey200ResponseResponse instantiates a new CreateMemberSshKey200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMemberSshKey200ResponseResponseWithDefaults

`func NewCreateMemberSshKey200ResponseResponseWithDefaults() *CreateMemberSshKey200ResponseResponse`

NewCreateMemberSshKey200ResponseResponseWithDefaults instantiates a new CreateMemberSshKey200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSshKeyInfo

`func (o *CreateMemberSshKey200ResponseResponse) GetSshKeyInfo() SshKeyInfo`

GetSshKeyInfo returns the SshKeyInfo field if non-nil, zero value otherwise.

### GetSshKeyInfoOk

`func (o *CreateMemberSshKey200ResponseResponse) GetSshKeyInfoOk() (*SshKeyInfo, bool)`

GetSshKeyInfoOk returns a tuple with the SshKeyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyInfo

`func (o *CreateMemberSshKey200ResponseResponse) SetSshKeyInfo(v SshKeyInfo)`

SetSshKeyInfo sets SshKeyInfo field to given value.

### HasSshKeyInfo

`func (o *CreateMemberSshKey200ResponseResponse) HasSshKeyInfo() bool`

HasSshKeyInfo returns a boolean if a field has been set.

### GetRequestId

`func (o *CreateMemberSshKey200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CreateMemberSshKey200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CreateMemberSshKey200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *CreateMemberSshKey200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


