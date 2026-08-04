# DirectoryDownloadEntry


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**completed_at** | **datetime** |  | [optional] 
**downloaded_objects** | [**DownloadedObjects**](DownloadedObjects.md) |  | [optional] 
**started_at** | **datetime** |  | [optional] 
**status** | **str** |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.directory_download_entry import DirectoryDownloadEntry

# TODO update the JSON string below
json = "{}"
# create an instance of DirectoryDownloadEntry from a JSON string
directory_download_entry_instance = DirectoryDownloadEntry.from_json(json)
# print the JSON string representation of the object
print(DirectoryDownloadEntry.to_json())

# convert the object into a dict
directory_download_entry_dict = directory_download_entry_instance.to_dict()
# create an instance of DirectoryDownloadEntry from a dict
directory_download_entry_form_dict = directory_download_entry.from_dict(directory_download_entry_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


