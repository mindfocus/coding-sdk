# ModifyGitFiles200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GitCommit** | Pointer to [**GitCommit**](GitCommit.md) |  | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewModifyGitFiles200ResponseResponse

`func NewModifyGitFiles200ResponseResponse() *ModifyGitFiles200ResponseResponse`

NewModifyGitFiles200ResponseResponse instantiates a new ModifyGitFiles200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyGitFiles200ResponseResponseWithDefaults

`func NewModifyGitFiles200ResponseResponseWithDefaults() *ModifyGitFiles200ResponseResponse`

NewModifyGitFiles200ResponseResponseWithDefaults instantiates a new ModifyGitFiles200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGitCommit

`func (o *ModifyGitFiles200ResponseResponse) GetGitCommit() GitCommit`

GetGitCommit returns the GitCommit field if non-nil, zero value otherwise.

### GetGitCommitOk

`func (o *ModifyGitFiles200ResponseResponse) GetGitCommitOk() (*GitCommit, bool)`

GetGitCommitOk returns a tuple with the GitCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitCommit

`func (o *ModifyGitFiles200ResponseResponse) SetGitCommit(v GitCommit)`

SetGitCommit sets GitCommit field to given value.

### HasGitCommit

`func (o *ModifyGitFiles200ResponseResponse) HasGitCommit() bool`

HasGitCommit returns a boolean if a field has been set.

### GetRequestId

`func (o *ModifyGitFiles200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ModifyGitFiles200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ModifyGitFiles200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ModifyGitFiles200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


