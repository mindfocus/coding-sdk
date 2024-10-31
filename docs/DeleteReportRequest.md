# DeleteReportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectName** | **string** | 项目名称 | 
**ReportId** | **int32** | 测试报告 ID | 

## Methods

### NewDeleteReportRequest

`func NewDeleteReportRequest(projectName string, reportId int32, ) *DeleteReportRequest`

NewDeleteReportRequest instantiates a new DeleteReportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteReportRequestWithDefaults

`func NewDeleteReportRequestWithDefaults() *DeleteReportRequest`

NewDeleteReportRequestWithDefaults instantiates a new DeleteReportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectName

`func (o *DeleteReportRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *DeleteReportRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *DeleteReportRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetReportId

`func (o *DeleteReportRequest) GetReportId() int32`

GetReportId returns the ReportId field if non-nil, zero value otherwise.

### GetReportIdOk

`func (o *DeleteReportRequest) GetReportIdOk() (*int32, bool)`

GetReportIdOk returns a tuple with the ReportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportId

`func (o *DeleteReportRequest) SetReportId(v int32)`

SetReportId sets ReportId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


