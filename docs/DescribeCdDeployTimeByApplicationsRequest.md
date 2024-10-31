# DescribeCdDeployTimeByApplicationsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | **[]string** | 应用名列表 | 
**EndAt** | **string** | 结束时间（格式：yyyy-MM-dd HH:mm:ss） | 
**StartAt** | **string** | 开始时间（格式：yyyy-MM-dd HH:mm:ss） | 

## Methods

### NewDescribeCdDeployTimeByApplicationsRequest

`func NewDescribeCdDeployTimeByApplicationsRequest(application []string, endAt string, startAt string, ) *DescribeCdDeployTimeByApplicationsRequest`

NewDescribeCdDeployTimeByApplicationsRequest instantiates a new DescribeCdDeployTimeByApplicationsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeCdDeployTimeByApplicationsRequestWithDefaults

`func NewDescribeCdDeployTimeByApplicationsRequestWithDefaults() *DescribeCdDeployTimeByApplicationsRequest`

NewDescribeCdDeployTimeByApplicationsRequestWithDefaults instantiates a new DescribeCdDeployTimeByApplicationsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *DescribeCdDeployTimeByApplicationsRequest) GetApplication() []string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *DescribeCdDeployTimeByApplicationsRequest) GetApplicationOk() (*[]string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *DescribeCdDeployTimeByApplicationsRequest) SetApplication(v []string)`

SetApplication sets Application field to given value.


### GetEndAt

`func (o *DescribeCdDeployTimeByApplicationsRequest) GetEndAt() string`

GetEndAt returns the EndAt field if non-nil, zero value otherwise.

### GetEndAtOk

`func (o *DescribeCdDeployTimeByApplicationsRequest) GetEndAtOk() (*string, bool)`

GetEndAtOk returns a tuple with the EndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAt

`func (o *DescribeCdDeployTimeByApplicationsRequest) SetEndAt(v string)`

SetEndAt sets EndAt field to given value.


### GetStartAt

`func (o *DescribeCdDeployTimeByApplicationsRequest) GetStartAt() string`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *DescribeCdDeployTimeByApplicationsRequest) GetStartAtOk() (*string, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *DescribeCdDeployTimeByApplicationsRequest) SetStartAt(v string)`

SetStartAt sets StartAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


