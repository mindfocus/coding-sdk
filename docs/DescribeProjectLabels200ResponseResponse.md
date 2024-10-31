# DescribeProjectLabels200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectList** | Pointer to [**[]Project**](Project.md) | 项目列表信息 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeProjectLabels200ResponseResponse

`func NewDescribeProjectLabels200ResponseResponse() *DescribeProjectLabels200ResponseResponse`

NewDescribeProjectLabels200ResponseResponse instantiates a new DescribeProjectLabels200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeProjectLabels200ResponseResponseWithDefaults

`func NewDescribeProjectLabels200ResponseResponseWithDefaults() *DescribeProjectLabels200ResponseResponse`

NewDescribeProjectLabels200ResponseResponseWithDefaults instantiates a new DescribeProjectLabels200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectList

`func (o *DescribeProjectLabels200ResponseResponse) GetProjectList() []Project`

GetProjectList returns the ProjectList field if non-nil, zero value otherwise.

### GetProjectListOk

`func (o *DescribeProjectLabels200ResponseResponse) GetProjectListOk() (*[]Project, bool)`

GetProjectListOk returns a tuple with the ProjectList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectList

`func (o *DescribeProjectLabels200ResponseResponse) SetProjectList(v []Project)`

SetProjectList sets ProjectList field to given value.

### HasProjectList

`func (o *DescribeProjectLabels200ResponseResponse) HasProjectList() bool`

HasProjectList returns a boolean if a field has been set.

### SetProjectListNil

`func (o *DescribeProjectLabels200ResponseResponse) SetProjectListNil(b bool)`

 SetProjectListNil sets the value for ProjectList to be an explicit nil

### UnsetProjectList
`func (o *DescribeProjectLabels200ResponseResponse) UnsetProjectList()`

UnsetProjectList ensures that no value is present for ProjectList, not even an explicit nil
### GetRequestId

`func (o *DescribeProjectLabels200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeProjectLabels200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeProjectLabels200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeProjectLabels200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


