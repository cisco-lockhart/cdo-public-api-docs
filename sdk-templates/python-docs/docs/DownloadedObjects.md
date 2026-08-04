# DownloadedObjects


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**groups** | **int** |  | [optional] 
**users** | **int** |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.downloaded_objects import DownloadedObjects

# TODO update the JSON string below
json = "{}"
# create an instance of DownloadedObjects from a JSON string
downloaded_objects_instance = DownloadedObjects.from_json(json)
# print the JSON string representation of the object
print(DownloadedObjects.to_json())

# convert the object into a dict
downloaded_objects_dict = downloaded_objects_instance.to_dict()
# create an instance of DownloadedObjects from a dict
downloaded_objects_form_dict = downloaded_objects.from_dict(downloaded_objects_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


