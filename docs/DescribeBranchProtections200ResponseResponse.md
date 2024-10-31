# DescribeBranchProtections200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BranchProtections** | Pointer to [**[]BranchProtection**](BranchProtection.md) | 保护分支规则列表 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeBranchProtections200ResponseResponse

`func NewDescribeBranchProtections200ResponseResponse() *DescribeBranchProtections200ResponseResponse`

NewDescribeBranchProtections200ResponseResponse instantiates a new DescribeBranchProtections200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeBranchProtections200ResponseResponseWithDefaults

`func NewDescribeBranchProtections200ResponseResponseWithDefaults() *DescribeBranchProtections200ResponseResponse`

NewDescribeBranchProtections200ResponseResponseWithDefaults instantiates a new DescribeBranchProtections200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranchProtections

`func (o *DescribeBranchProtections200ResponseResponse) GetBranchProtections() []BranchProtection`

GetBranchProtections returns the BranchProtections field if non-nil, zero value otherwise.

### GetBranchProtectionsOk

`func (o *DescribeBranchProtections200ResponseResponse) GetBranchProtectionsOk() (*[]BranchProtection, bool)`

GetBranchProtectionsOk returns a tuple with the BranchProtections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchProtections

`func (o *DescribeBranchProtections200ResponseResponse) SetBranchProtections(v []BranchProtection)`

SetBranchProtections sets BranchProtections field to given value.

### HasBranchProtections

`func (o *DescribeBranchProtections200ResponseResponse) HasBranchProtections() bool`

HasBranchProtections returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeBranchProtections200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeBranchProtections200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeBranchProtections200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeBranchProtections200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


