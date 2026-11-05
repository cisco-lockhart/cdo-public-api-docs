# FtdDeviceRestoreInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**device_uids** | **List[str]** | The unique identifiers, represented as UUIDs, of the devices in the deployment run to restore. Every device must be part of the deployment run and restorable. | 

## Example

```python
from scc_firewall_manager_sdk.models.ftd_device_restore_input import FtdDeviceRestoreInput

# TODO update the JSON string below
json = "{}"
# create an instance of FtdDeviceRestoreInput from a JSON string
ftd_device_restore_input_instance = FtdDeviceRestoreInput.from_json(json)
# print the JSON string representation of the object
print(FtdDeviceRestoreInput.to_json())

# convert the object into a dict
ftd_device_restore_input_dict = ftd_device_restore_input_instance.to_dict()
# create an instance of FtdDeviceRestoreInput from a dict
ftd_device_restore_input_form_dict = ftd_device_restore_input.from_dict(ftd_device_restore_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


