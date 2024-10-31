# DescribeProgramProjects200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Program**](Program.md) | 项目集下项目列表数据 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeProgramProjects200ResponseResponse

`func NewDescribeProgramProjects200ResponseResponse() *DescribeProgramProjects200ResponseResponse`

NewDescribeProgramProjects200ResponseResponse instantiates a new DescribeProgramProjects200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeProgramProjects200ResponseResponseWithDefaults

`func NewDescribeProgramProjects200ResponseResponseWithDefaults() *DescribeProgramProjects200ResponseResponse`

NewDescribeProgramProjects200ResponseResponseWithDefaults instantiates a new DescribeProgramProjects200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DescribeProgramProjects200ResponseResponse) GetData() []Program`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DescribeProgramProjects200ResponseResponse) GetDataOk() (*[]Program, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DescribeProgramProjects200ResponseResponse) SetData(v []Program)`

SetData sets Data field to given value.

### HasData

`func (o *DescribeProgramProjects200ResponseResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeProgramProjects200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeProgramProjects200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeProgramProjects200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeProgramProjects200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

