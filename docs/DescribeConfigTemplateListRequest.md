# DescribeConfigTemplateListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CooperateMode** | Pointer to **string** | 配置方案协作类型，包括 SCRUM 和 CLASSIC | [optional] 
**Keyword** | Pointer to **string** | 关键字 | [optional] 
**PageNumber** | **int64** | 页码 | 
**PageSize** | **int64** | 分页大小 | 
**ProjectName** | **string** | 项目名称 | 
**TemplateType** | Pointer to **string** | 配置方案类型，全局配置方案取值 GLOBAL ，项目配置方案取值 PROJECT，不填默认为 GLOBAL | [optional] 

## Methods

### NewDescribeConfigTemplateListRequest

`func NewDescribeConfigTemplateListRequest(pageNumber int64, pageSize int64, projectName string, ) *DescribeConfigTemplateListRequest`

NewDescribeConfigTemplateListRequest instantiates a new DescribeConfigTemplateListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeConfigTemplateListRequestWithDefaults

`func NewDescribeConfigTemplateListRequestWithDefaults() *DescribeConfigTemplateListRequest`

NewDescribeConfigTemplateListRequestWithDefaults instantiates a new DescribeConfigTemplateListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCooperateMode

`func (o *DescribeConfigTemplateListRequest) GetCooperateMode() string`

GetCooperateMode returns the CooperateMode field if non-nil, zero value otherwise.

### GetCooperateModeOk

`func (o *DescribeConfigTemplateListRequest) GetCooperateModeOk() (*string, bool)`

GetCooperateModeOk returns a tuple with the CooperateMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCooperateMode

`func (o *DescribeConfigTemplateListRequest) SetCooperateMode(v string)`

SetCooperateMode sets CooperateMode field to given value.

### HasCooperateMode

`func (o *DescribeConfigTemplateListRequest) HasCooperateMode() bool`

HasCooperateMode returns a boolean if a field has been set.

### GetKeyword

`func (o *DescribeConfigTemplateListRequest) GetKeyword() string`

GetKeyword returns the Keyword field if non-nil, zero value otherwise.

### GetKeywordOk

`func (o *DescribeConfigTemplateListRequest) GetKeywordOk() (*string, bool)`

GetKeywordOk returns a tuple with the Keyword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyword

`func (o *DescribeConfigTemplateListRequest) SetKeyword(v string)`

SetKeyword sets Keyword field to given value.

### HasKeyword

`func (o *DescribeConfigTemplateListRequest) HasKeyword() bool`

HasKeyword returns a boolean if a field has been set.

### GetPageNumber

`func (o *DescribeConfigTemplateListRequest) GetPageNumber() int64`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *DescribeConfigTemplateListRequest) GetPageNumberOk() (*int64, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *DescribeConfigTemplateListRequest) SetPageNumber(v int64)`

SetPageNumber sets PageNumber field to given value.


### GetPageSize

`func (o *DescribeConfigTemplateListRequest) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DescribeConfigTemplateListRequest) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DescribeConfigTemplateListRequest) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.


### GetProjectName

`func (o *DescribeConfigTemplateListRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *DescribeConfigTemplateListRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *DescribeConfigTemplateListRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetTemplateType

`func (o *DescribeConfigTemplateListRequest) GetTemplateType() string`

GetTemplateType returns the TemplateType field if non-nil, zero value otherwise.

### GetTemplateTypeOk

`func (o *DescribeConfigTemplateListRequest) GetTemplateTypeOk() (*string, bool)`

GetTemplateTypeOk returns a tuple with the TemplateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateType

`func (o *DescribeConfigTemplateListRequest) SetTemplateType(v string)`

SetTemplateType sets TemplateType field to given value.

### HasTemplateType

`func (o *DescribeConfigTemplateListRequest) HasTemplateType() bool`

HasTemplateType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


