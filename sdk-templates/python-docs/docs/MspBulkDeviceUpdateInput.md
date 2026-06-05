# MspBulkDeviceUpdateInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**device_uids** | **List[str]** | A set of device UIDs to bulk update. | [optional] 
**maintenance_window_duration_minutes** | **int** | Duration of the maintenance window in minutes. | [optional] 
**maintenance_window_start_cron** | **str** | CRON expression defining the scheduled start time of the maintenance window in format: Minute, Hour, Day of Month, Month, Day of Week. | [optional] 
**opted_in_to_asa_health_metrics** | **bool** | Specify whether to enable or disable health metrics collection (SDC-managed ASA devices only). | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.msp_bulk_device_update_input import MspBulkDeviceUpdateInput

# TODO update the JSON string below
json = "{}"
# create an instance of MspBulkDeviceUpdateInput from a JSON string
msp_bulk_device_update_input_instance = MspBulkDeviceUpdateInput.from_json(json)
# print the JSON string representation of the object
print(MspBulkDeviceUpdateInput.to_json())

# convert the object into a dict
msp_bulk_device_update_input_dict = msp_bulk_device_update_input_instance.to_dict()
# create an instance of MspBulkDeviceUpdateInput from a dict
msp_bulk_device_update_input_form_dict = msp_bulk_device_update_input.from_dict(msp_bulk_device_update_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


