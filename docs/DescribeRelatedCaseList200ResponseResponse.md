# DescribeRelatedCaseList200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cases** | Pointer to [**[]RelatedCase**](RelatedCase.md) | 事项关联的测试用例 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeRelatedCaseList200ResponseResponse

`func NewDescribeRelatedCaseList200ResponseResponse() *DescribeRelatedCaseList200ResponseResponse`

NewDescribeRelatedCaseList200ResponseResponse instantiates a new DescribeRelatedCaseList200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeRelatedCaseList200ResponseResponseWithDefaults

`func NewDescribeRelatedCaseList200ResponseResponseWithDefaults() *DescribeRelatedCaseList200ResponseResponse`

NewDescribeRelatedCaseList200ResponseResponseWithDefaults instantiates a new DescribeRelatedCaseList200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCases

`func (o *DescribeRelatedCaseList200ResponseResponse) GetCases() []RelatedCase`

GetCases returns the Cases field if non-nil, zero value otherwise.

### GetCasesOk

`func (o *DescribeRelatedCaseList200ResponseResponse) GetCasesOk() (*[]RelatedCase, bool)`

GetCasesOk returns a tuple with the Cases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCases

`func (o *DescribeRelatedCaseList200ResponseResponse) SetCases(v []RelatedCase)`

SetCases sets Cases field to given value.

### HasCases

`func (o *DescribeRelatedCaseList200ResponseResponse) HasCases() bool`

HasCases returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeRelatedCaseList200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeRelatedCaseList200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeRelatedCaseList200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeRelatedCaseList200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


