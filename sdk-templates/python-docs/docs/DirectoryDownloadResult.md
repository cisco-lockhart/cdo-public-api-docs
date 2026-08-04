# DirectoryDownloadResult


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**downloads** | [**List[DirectoryDownloadEntry]**](DirectoryDownloadEntry.md) |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.directory_download_result import DirectoryDownloadResult

# TODO update the JSON string below
json = "{}"
# create an instance of DirectoryDownloadResult from a JSON string
directory_download_result_instance = DirectoryDownloadResult.from_json(json)
# print the JSON string representation of the object
print(DirectoryDownloadResult.to_json())

# convert the object into a dict
directory_download_result_dict = directory_download_result_instance.to_dict()
# create an instance of DirectoryDownloadResult from a dict
directory_download_result_form_dict = directory_download_result.from_dict(directory_download_result_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


