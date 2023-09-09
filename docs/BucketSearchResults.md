# BucketSearchResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** | Total number of matching buckets | [optional] 
**Results** | Pointer to [**[]Bucket**](Bucket.md) |  | [optional] 

## Methods

### NewBucketSearchResults

`func NewBucketSearchResults() *BucketSearchResults`

NewBucketSearchResults instantiates a new BucketSearchResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketSearchResultsWithDefaults

`func NewBucketSearchResultsWithDefaults() *BucketSearchResults`

NewBucketSearchResultsWithDefaults instantiates a new BucketSearchResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *BucketSearchResults) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *BucketSearchResults) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *BucketSearchResults) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *BucketSearchResults) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetResults

`func (o *BucketSearchResults) GetResults() []Bucket`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BucketSearchResults) GetResultsOk() (*[]Bucket, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BucketSearchResults) SetResults(v []Bucket)`

SetResults sets Results field to given value.

### HasResults

`func (o *BucketSearchResults) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


