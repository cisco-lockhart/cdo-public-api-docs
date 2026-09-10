# BackupResult


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**num_devices** | **int** |  | [optional] 
**num_failed** | **int** |  | [optional] 
**num_running** | **int** |  | [optional] 
**num_successful** | **int** |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.backup_result import BackupResult

# TODO update the JSON string below
json = "{}"
# create an instance of BackupResult from a JSON string
backup_result_instance = BackupResult.from_json(json)
# print the JSON string representation of the object
print(BackupResult.to_json())

# convert the object into a dict
backup_result_dict = backup_result_instance.to_dict()
# create an instance of BackupResult from a dict
backup_result_form_dict = backup_result.from_dict(backup_result_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


