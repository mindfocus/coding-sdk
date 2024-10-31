# GrepLineData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrepLines** | Pointer to [**[]GrepLineInfo**](GrepLineInfo.md) | 代码片段详细信息 | [optional] 
**Page** | Pointer to [**PageInfo**](PageInfo.md) |  | [optional] 

## Methods

### NewGrepLineData

`func NewGrepLineData() *GrepLineData`

NewGrepLineData instantiates a new GrepLineData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrepLineDataWithDefaults

`func NewGrepLineDataWithDefaults() *GrepLineData`

NewGrepLineDataWithDefaults instantiates a new GrepLineData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrepLines

`func (o *GrepLineData) GetGrepLines() []GrepLineInfo`

GetGrepLines returns the GrepLines field if non-nil, zero value otherwise.

### GetGrepLinesOk

`func (o *GrepLineData) GetGrepLinesOk() (*[]GrepLineInfo, bool)`

GetGrepLinesOk returns a tuple with the GrepLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrepLines

`func (o *GrepLineData) SetGrepLines(v []GrepLineInfo)`

SetGrepLines sets GrepLines field to given value.

### HasGrepLines

`func (o *GrepLineData) HasGrepLines() bool`

HasGrepLines returns a boolean if a field has been set.

### SetGrepLinesNil

`func (o *GrepLineData) SetGrepLinesNil(b bool)`

 SetGrepLinesNil sets the value for GrepLines to be an explicit nil

### UnsetGrepLines
`func (o *GrepLineData) UnsetGrepLines()`

UnsetGrepLines ensures that no value is present for GrepLines, not even an explicit nil
### GetPage

`func (o *GrepLineData) GetPage() PageInfo`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GrepLineData) GetPageOk() (*PageInfo, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GrepLineData) SetPage(v PageInfo)`

SetPage sets Page field to given value.

### HasPage

`func (o *GrepLineData) HasPage() bool`

HasPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


