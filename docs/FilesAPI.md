# \FilesAPI

All URIs are relative to *https://api.openbuckets.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchFiles**](FilesAPI.md#SearchFiles) | **Get** /api/v2/files | Search Files



## SearchFiles

> FileSearchResults SearchFiles(ctx).Keywords(keywords).Order(order).Direction(direction).FieldToSearch(fieldToSearch).FullPath(fullPath).Extensions(extensions).LastModifiedFrom(lastModifiedFrom).LastModifiedTo(lastModifiedTo).SizeFrom(sizeFrom).SizeTo(sizeTo).Start(start).Limit(limit).ExcludeBuckets(excludeBuckets).Buckets(buckets).StopExtensions(stopExtensions).Execute()

Search Files



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "https://github.com/openbuckets/sdk-go"
)

func main() {
    keywords := "org images -aws" // string | multiple keywords.\"-\" denotes stop keywords (optional)
    order := "size" // string | the sorting field for the search results (e.g., \"size\", \"lastModified\") (optional)
    direction := "desc" // string | the sorting direction for the search results (e.g., \"desc\" for descending) (optional)
    fieldToSearch := "desc" // string | taken into consideration if you provide any of the allowed values, \"cloudProvider\",\"fileExtension\",\"fileName\",\"fileUrl\",\"fullPath\" (optional)
    fullPath := float32(1) // float32 | include the full path in the search results (1 for true, 0 for false) (optional)
    extensions := "pdf,.env" // string | comma-separated list of file extensions to include (e.g., \"pdf,env\") (optional)
    lastModifiedFrom := "1682965800" // string | UNIX timestamp for the starting date of the last modification range (optional)
    lastModifiedTo := "1693420200" // string | UNIX timestamp for the ending date of the last modification rang (optional)
    sizeFrom := "15155035" // string | minimum file size in bytes (optional)
    sizeTo := "4538824351471" // string | maximum file size in bytes (optional)
    start := float32(0) // float32 | starting index for pagination (optional)
    limit := float32(20) // float32 | number of search results to return per page, based on your role.  If you send a value more than the allowed limit, we set it to the allowed limit. (optional)
    excludeBuckets := "45,54" // string | comma-separated list of bucket IDs to exclude from the search (optional)
    buckets := "buckets_example" // string | filter search results to specific bucket IDs (optional)
    stopExtensions := ".csv,.env" // string | comma-separated list of file extensions to exclude with or without \".\" (e.g., sql, .sql) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesAPI.SearchFiles(context.Background()).Keywords(keywords).Order(order).Direction(direction).FieldToSearch(fieldToSearch).FullPath(fullPath).Extensions(extensions).LastModifiedFrom(lastModifiedFrom).LastModifiedTo(lastModifiedTo).SizeFrom(sizeFrom).SizeTo(sizeTo).Start(start).Limit(limit).ExcludeBuckets(excludeBuckets).Buckets(buckets).StopExtensions(stopExtensions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.SearchFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchFiles`: FileSearchResults
    fmt.Fprintf(os.Stdout, "Response from `FilesAPI.SearchFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keywords** | **string** | multiple keywords.\&quot;-\&quot; denotes stop keywords | 
 **order** | **string** | the sorting field for the search results (e.g., \&quot;size\&quot;, \&quot;lastModified\&quot;) | 
 **direction** | **string** | the sorting direction for the search results (e.g., \&quot;desc\&quot; for descending) | 
 **fieldToSearch** | **string** | taken into consideration if you provide any of the allowed values, \&quot;cloudProvider\&quot;,\&quot;fileExtension\&quot;,\&quot;fileName\&quot;,\&quot;fileUrl\&quot;,\&quot;fullPath\&quot; | 
 **fullPath** | **float32** | include the full path in the search results (1 for true, 0 for false) | 
 **extensions** | **string** | comma-separated list of file extensions to include (e.g., \&quot;pdf,env\&quot;) | 
 **lastModifiedFrom** | **string** | UNIX timestamp for the starting date of the last modification range | 
 **lastModifiedTo** | **string** | UNIX timestamp for the ending date of the last modification rang | 
 **sizeFrom** | **string** | minimum file size in bytes | 
 **sizeTo** | **string** | maximum file size in bytes | 
 **start** | **float32** | starting index for pagination | 
 **limit** | **float32** | number of search results to return per page, based on your role.  If you send a value more than the allowed limit, we set it to the allowed limit. | 
 **excludeBuckets** | **string** | comma-separated list of bucket IDs to exclude from the search | 
 **buckets** | **string** | filter search results to specific bucket IDs | 
 **stopExtensions** | **string** | comma-separated list of file extensions to exclude with or without \&quot;.\&quot; (e.g., sql, .sql) | 

### Return type

[**FileSearchResults**](FileSearchResults.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

