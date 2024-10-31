# DescribeTestCaseListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keyword** | Pointer to **string** | 关键字搜索 | [optional] 
**Priority** | Pointer to **int32** | 优先级，默认2（中），可选值：0（紧急）,1（高）,2（中）,3（低） | [optional] 
**ProjectName** | **string** | 项目名称 | 
**SectionId** | Pointer to **int32** | 分组 ID | [optional] 
**TemplateType** | Pointer to **string** | 用例类型，可选值：STEPS(步骤用例)，TEXT(文本用例) | [optional] 

## Methods

### NewDescribeTestCaseListRequest

`func NewDescribeTestCaseListRequest(projectName string, ) *DescribeTestCaseListRequest`

NewDescribeTestCaseListRequest instantiates a new DescribeTestCaseListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeTestCaseListRequestWithDefaults

`func NewDescribeTestCaseListRequestWithDefaults() *DescribeTestCaseListRequest`

NewDescribeTestCaseListRequestWithDefaults instantiates a new DescribeTestCaseListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyword

`func (o *DescribeTestCaseListRequest) GetKeyword() string`

GetKeyword returns the Keyword field if non-nil, zero value otherwise.

### GetKeywordOk

`func (o *DescribeTestCaseListRequest) GetKeywordOk() (*string, bool)`

GetKeywordOk returns a tuple with the Keyword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyword

`func (o *DescribeTestCaseListRequest) SetKeyword(v string)`

SetKeyword sets Keyword field to given value.

### HasKeyword

`func (o *DescribeTestCaseListRequest) HasKeyword() bool`

HasKeyword returns a boolean if a field has been set.

### GetPriority

`func (o *DescribeTestCaseListRequest) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DescribeTestCaseListRequest) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DescribeTestCaseListRequest) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *DescribeTestCaseListRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProjectName

`func (o *DescribeTestCaseListRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *DescribeTestCaseListRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *DescribeTestCaseListRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetSectionId

`func (o *DescribeTestCaseListRequest) GetSectionId() int32`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *DescribeTestCaseListRequest) GetSectionIdOk() (*int32, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *DescribeTestCaseListRequest) SetSectionId(v int32)`

SetSectionId sets SectionId field to given value.

### HasSectionId

`func (o *DescribeTestCaseListRequest) HasSectionId() bool`

HasSectionId returns a boolean if a field has been set.

### GetTemplateType

`func (o *DescribeTestCaseListRequest) GetTemplateType() string`

GetTemplateType returns the TemplateType field if non-nil, zero value otherwise.

### GetTemplateTypeOk

`func (o *DescribeTestCaseListRequest) GetTemplateTypeOk() (*string, bool)`

GetTemplateTypeOk returns a tuple with the TemplateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateType

`func (o *DescribeTestCaseListRequest) SetTemplateType(v string)`

SetTemplateType sets TemplateType field to given value.

### HasTemplateType

`func (o *DescribeTestCaseListRequest) HasTemplateType() bool`

HasTemplateType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


