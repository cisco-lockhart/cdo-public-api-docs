# CopyFilesResult


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**completed_acknowledged** | **int** |  | [optional] 
**completed_awaiting_acknowledgement** | **int** |  | [optional] 
**requested** | **int** |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.copy_files_result import CopyFilesResult

# TODO update the JSON string below
json = "{}"
# create an instance of CopyFilesResult from a JSON string
copy_files_result_instance = CopyFilesResult.from_json(json)
# print the JSON string representation of the object
print(CopyFilesResult.to_json())

# convert the object into a dict
copy_files_result_dict = copy_files_result_instance.to_dict()
# create an instance of CopyFilesResult from a dict
copy_files_result_form_dict = copy_files_result.from_dict(copy_files_result_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


