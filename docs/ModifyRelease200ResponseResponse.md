# ModifyRelease200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Release** | Pointer to [**OpenApiRelease**](OpenApiRelease.md) |  | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewModifyRelease200ResponseResponse

`func NewModifyRelease200ResponseResponse() *ModifyRelease200ResponseResponse`

NewModifyRelease200ResponseResponse instantiates a new ModifyRelease200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyRelease200ResponseResponseWithDefaults

`func NewModifyRelease200ResponseResponseWithDefaults() *ModifyRelease200ResponseResponse`

NewModifyRelease200ResponseResponseWithDefaults instantiates a new ModifyRelease200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRelease

`func (o *ModifyRelease200ResponseResponse) GetRelease() OpenApiRelease`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *ModifyRelease200ResponseResponse) GetReleaseOk() (*OpenApiRelease, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *ModifyRelease200ResponseResponse) SetRelease(v OpenApiRelease)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *ModifyRelease200ResponseResponse) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetRequestId

`func (o *ModifyRelease200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ModifyRelease200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ModifyRelease200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ModifyRelease200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


