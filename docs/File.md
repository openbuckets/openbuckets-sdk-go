# File

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for the file | [optional] 
**Bucket** | Pointer to **string** | Name of the bucket | [optional] 
**BucketId** | Pointer to **int32** | Unique identifier for the bucket | [optional] 
**Filename** | Pointer to **string** | Name of the file | [optional] 
**FullPath** | Pointer to **string** | Full path to the file | [optional] 
**Url** | Pointer to **string** | URL to access the file | [optional] 
**Size** | Pointer to **int32** | Size of the file in bytes | [optional] 
**Type** | Pointer to **string** | Type of cloud storage | [optional] 
**LastModified** | Pointer to **int32** | Timestamp when file was last modified | [optional] 
**Container** | Pointer to **string** | Storage container name | [optional] 

## Methods

### NewFile

`func NewFile() *File`

NewFile instantiates a new File object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileWithDefaults

`func NewFileWithDefaults() *File`

NewFileWithDefaults instantiates a new File object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *File) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *File) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *File) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *File) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBucket

`func (o *File) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *File) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *File) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *File) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetBucketId

`func (o *File) GetBucketId() int32`

GetBucketId returns the BucketId field if non-nil, zero value otherwise.

### GetBucketIdOk

`func (o *File) GetBucketIdOk() (*int32, bool)`

GetBucketIdOk returns a tuple with the BucketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketId

`func (o *File) SetBucketId(v int32)`

SetBucketId sets BucketId field to given value.

### HasBucketId

`func (o *File) HasBucketId() bool`

HasBucketId returns a boolean if a field has been set.

### GetFilename

`func (o *File) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *File) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *File) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *File) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetFullPath

`func (o *File) GetFullPath() string`

GetFullPath returns the FullPath field if non-nil, zero value otherwise.

### GetFullPathOk

`func (o *File) GetFullPathOk() (*string, bool)`

GetFullPathOk returns a tuple with the FullPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullPath

`func (o *File) SetFullPath(v string)`

SetFullPath sets FullPath field to given value.

### HasFullPath

`func (o *File) HasFullPath() bool`

HasFullPath returns a boolean if a field has been set.

### GetUrl

`func (o *File) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *File) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *File) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *File) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetSize

`func (o *File) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *File) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *File) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *File) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetType

`func (o *File) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *File) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *File) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *File) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLastModified

`func (o *File) GetLastModified() int32`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *File) GetLastModifiedOk() (*int32, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *File) SetLastModified(v int32)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *File) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetContainer

`func (o *File) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *File) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *File) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *File) HasContainer() bool`

HasContainer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


