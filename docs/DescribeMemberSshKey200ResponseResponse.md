# DescribeMemberSshKey200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SshKeyInfos** | Pointer to [**[]SshKeyInfo**](SshKeyInfo.md) | SSH公钥详情列表 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeMemberSshKey200ResponseResponse

`func NewDescribeMemberSshKey200ResponseResponse() *DescribeMemberSshKey200ResponseResponse`

NewDescribeMemberSshKey200ResponseResponse instantiates a new DescribeMemberSshKey200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeMemberSshKey200ResponseResponseWithDefaults

`func NewDescribeMemberSshKey200ResponseResponseWithDefaults() *DescribeMemberSshKey200ResponseResponse`

NewDescribeMemberSshKey200ResponseResponseWithDefaults instantiates a new DescribeMemberSshKey200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSshKeyInfos

`func (o *DescribeMemberSshKey200ResponseResponse) GetSshKeyInfos() []SshKeyInfo`

GetSshKeyInfos returns the SshKeyInfos field if non-nil, zero value otherwise.

### GetSshKeyInfosOk

`func (o *DescribeMemberSshKey200ResponseResponse) GetSshKeyInfosOk() (*[]SshKeyInfo, bool)`

GetSshKeyInfosOk returns a tuple with the SshKeyInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyInfos

`func (o *DescribeMemberSshKey200ResponseResponse) SetSshKeyInfos(v []SshKeyInfo)`

SetSshKeyInfos sets SshKeyInfos field to given value.

### HasSshKeyInfos

`func (o *DescribeMemberSshKey200ResponseResponse) HasSshKeyInfos() bool`

HasSshKeyInfos returns a boolean if a field has been set.

### SetSshKeyInfosNil

`func (o *DescribeMemberSshKey200ResponseResponse) SetSshKeyInfosNil(b bool)`

 SetSshKeyInfosNil sets the value for SshKeyInfos to be an explicit nil

### UnsetSshKeyInfos
`func (o *DescribeMemberSshKey200ResponseResponse) UnsetSshKeyInfos()`

UnsetSshKeyInfos ensures that no value is present for SshKeyInfos, not even an explicit nil
### GetRequestId

`func (o *DescribeMemberSshKey200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeMemberSshKey200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeMemberSshKey200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeMemberSshKey200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


