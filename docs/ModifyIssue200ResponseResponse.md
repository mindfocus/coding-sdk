# ModifyIssue200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Issue** | Pointer to [**IssueDetail**](IssueDetail.md) |  | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewModifyIssue200ResponseResponse

`func NewModifyIssue200ResponseResponse() *ModifyIssue200ResponseResponse`

NewModifyIssue200ResponseResponse instantiates a new ModifyIssue200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyIssue200ResponseResponseWithDefaults

`func NewModifyIssue200ResponseResponseWithDefaults() *ModifyIssue200ResponseResponse`

NewModifyIssue200ResponseResponseWithDefaults instantiates a new ModifyIssue200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssue

`func (o *ModifyIssue200ResponseResponse) GetIssue() IssueDetail`

GetIssue returns the Issue field if non-nil, zero value otherwise.

### GetIssueOk

`func (o *ModifyIssue200ResponseResponse) GetIssueOk() (*IssueDetail, bool)`

GetIssueOk returns a tuple with the Issue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssue

`func (o *ModifyIssue200ResponseResponse) SetIssue(v IssueDetail)`

SetIssue sets Issue field to given value.

### HasIssue

`func (o *ModifyIssue200ResponseResponse) HasIssue() bool`

HasIssue returns a boolean if a field has been set.

### GetRequestId

`func (o *ModifyIssue200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ModifyIssue200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ModifyIssue200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ModifyIssue200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


