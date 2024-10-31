# DescribeCodingCIBuildMetrics200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]DescribeCodingCIBuildMetrics**](DescribeCodingCIBuildMetrics.md) | BuildMetrics 数组结构 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeCodingCIBuildMetrics200ResponseResponse

`func NewDescribeCodingCIBuildMetrics200ResponseResponse() *DescribeCodingCIBuildMetrics200ResponseResponse`

NewDescribeCodingCIBuildMetrics200ResponseResponse instantiates a new DescribeCodingCIBuildMetrics200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeCodingCIBuildMetrics200ResponseResponseWithDefaults

`func NewDescribeCodingCIBuildMetrics200ResponseResponseWithDefaults() *DescribeCodingCIBuildMetrics200ResponseResponse`

NewDescribeCodingCIBuildMetrics200ResponseResponseWithDefaults instantiates a new DescribeCodingCIBuildMetrics200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DescribeCodingCIBuildMetrics200ResponseResponse) GetData() []DescribeCodingCIBuildMetrics`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DescribeCodingCIBuildMetrics200ResponseResponse) GetDataOk() (*[]DescribeCodingCIBuildMetrics, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DescribeCodingCIBuildMetrics200ResponseResponse) SetData(v []DescribeCodingCIBuildMetrics)`

SetData sets Data field to given value.

### HasData

`func (o *DescribeCodingCIBuildMetrics200ResponseResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *DescribeCodingCIBuildMetrics200ResponseResponse) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *DescribeCodingCIBuildMetrics200ResponseResponse) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetRequestId

`func (o *DescribeCodingCIBuildMetrics200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeCodingCIBuildMetrics200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeCodingCIBuildMetrics200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeCodingCIBuildMetrics200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


