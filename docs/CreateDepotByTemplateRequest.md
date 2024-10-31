# CreateDepotByTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotName** | **string** | 仓库名 | 
**Description** | Pointer to **string** | 仓库描述 | [optional] 
**ProjectId** | **int64** | 项目Id | 
**Template** | **string** | 仓库模板。目前支持：Spring, Ruby on Rails，Ruby Sinatra，Node.js，Android，Python，Hexo，Jekyll。对应模板参数分别为：SPRING,ROR,SINATRA,NODEJS,ANDROID,FLASK,CLOUD_API_HEXO,CLOUD_API_JEKYLL。如果设置了自定义模版，可以传入自定义模版的仓库Id | 
**UserId** | **int64** | 用户Id | 

## Methods

### NewCreateDepotByTemplateRequest

`func NewCreateDepotByTemplateRequest(depotName string, projectId int64, template string, userId int64, ) *CreateDepotByTemplateRequest`

NewCreateDepotByTemplateRequest instantiates a new CreateDepotByTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDepotByTemplateRequestWithDefaults

`func NewCreateDepotByTemplateRequestWithDefaults() *CreateDepotByTemplateRequest`

NewCreateDepotByTemplateRequestWithDefaults instantiates a new CreateDepotByTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotName

`func (o *CreateDepotByTemplateRequest) GetDepotName() string`

GetDepotName returns the DepotName field if non-nil, zero value otherwise.

### GetDepotNameOk

`func (o *CreateDepotByTemplateRequest) GetDepotNameOk() (*string, bool)`

GetDepotNameOk returns a tuple with the DepotName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotName

`func (o *CreateDepotByTemplateRequest) SetDepotName(v string)`

SetDepotName sets DepotName field to given value.


### GetDescription

`func (o *CreateDepotByTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateDepotByTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateDepotByTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateDepotByTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProjectId

`func (o *CreateDepotByTemplateRequest) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreateDepotByTemplateRequest) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreateDepotByTemplateRequest) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.


### GetTemplate

`func (o *CreateDepotByTemplateRequest) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *CreateDepotByTemplateRequest) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *CreateDepotByTemplateRequest) SetTemplate(v string)`

SetTemplate sets Template field to given value.


### GetUserId

`func (o *CreateDepotByTemplateRequest) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CreateDepotByTemplateRequest) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CreateDepotByTemplateRequest) SetUserId(v int64)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


