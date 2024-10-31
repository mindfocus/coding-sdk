# DescribeCodingCIBuild200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Build** | Pointer to [**CodingCIBuild**](CodingCIBuild.md) |  | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeCodingCIBuild200ResponseResponse

`func NewDescribeCodingCIBuild200ResponseResponse() *DescribeCodingCIBuild200ResponseResponse`

NewDescribeCodingCIBuild200ResponseResponse instantiates a new DescribeCodingCIBuild200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeCodingCIBuild200ResponseResponseWithDefaults

`func NewDescribeCodingCIBuild200ResponseResponseWithDefaults() *DescribeCodingCIBuild200ResponseResponse`

NewDescribeCodingCIBuild200ResponseResponseWithDefaults instantiates a new DescribeCodingCIBuild200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuild

`func (o *DescribeCodingCIBuild200ResponseResponse) GetBuild() CodingCIBuild`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *DescribeCodingCIBuild200ResponseResponse) GetBuildOk() (*CodingCIBuild, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *DescribeCodingCIBuild200ResponseResponse) SetBuild(v CodingCIBuild)`

SetBuild sets Build field to given value.

### HasBuild

`func (o *DescribeCodingCIBuild200ResponseResponse) HasBuild() bool`

HasBuild returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeCodingCIBuild200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeCodingCIBuild200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeCodingCIBuild200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeCodingCIBuild200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


