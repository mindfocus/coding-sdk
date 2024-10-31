# DescribeWikiRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Iid** | **int64** | wiki编号 | 
**ProjectName** | **string** | 项目名称 | 
**VersionId** | **int64** | 版本号 | 

## Methods

### NewDescribeWikiRequest

`func NewDescribeWikiRequest(iid int64, projectName string, versionId int64, ) *DescribeWikiRequest`

NewDescribeWikiRequest instantiates a new DescribeWikiRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeWikiRequestWithDefaults

`func NewDescribeWikiRequestWithDefaults() *DescribeWikiRequest`

NewDescribeWikiRequestWithDefaults instantiates a new DescribeWikiRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIid

`func (o *DescribeWikiRequest) GetIid() int64`

GetIid returns the Iid field if non-nil, zero value otherwise.

### GetIidOk

`func (o *DescribeWikiRequest) GetIidOk() (*int64, bool)`

GetIidOk returns a tuple with the Iid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIid

`func (o *DescribeWikiRequest) SetIid(v int64)`

SetIid sets Iid field to given value.


### GetProjectName

`func (o *DescribeWikiRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *DescribeWikiRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *DescribeWikiRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetVersionId

`func (o *DescribeWikiRequest) GetVersionId() int64`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *DescribeWikiRequest) GetVersionIdOk() (*int64, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *DescribeWikiRequest) SetVersionId(v int64)`

SetVersionId sets VersionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


