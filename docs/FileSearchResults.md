# FileSearchResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** | Total number of matching files | [optional] 
**Results** | Pointer to [**[]File**](File.md) |  | [optional] 

## Methods

### NewFileSearchResults

`func NewFileSearchResults() *FileSearchResults`

NewFileSearchResults instantiates a new FileSearchResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileSearchResultsWithDefaults

`func NewFileSearchResultsWithDefaults() *FileSearchResults`

NewFileSearchResultsWithDefaults instantiates a new FileSearchResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *FileSearchResults) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *FileSearchResults) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *FileSearchResults) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *FileSearchResults) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetResults

`func (o *FileSearchResults) GetResults() []File`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *FileSearchResults) GetResultsOk() (*[]File, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *FileSearchResults) SetResults(v []File)`

SetResults sets Results field to given value.

### HasResults

`func (o *FileSearchResults) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


