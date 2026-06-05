# DeviceEndOfLifeDetails


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**end_of_life_announcement_date** | **datetime** | The date (UTC; represented using the RFC-3339 standard), the hardware End-Of_life was announced | [optional] 
**end_of_sale_date** | **datetime** | The date (UTC; represented using the RFC-3339 standard), the hardware is eligible for sale. | [optional] 
**end_of_software_maintenance_releases_date** | **datetime** | The date (UTC; represented using the RFC-3339 standard), the hardware is eligible for software maintenance. | [optional] 
**last_date_of_support** | **datetime** | The date (UTC; represented using the RFC-3339 standard), the hardware is eligible for support. | [optional] 
**url** | **str** | The hardware End-Of-Life notice URL with more information on the End-Of-Life details. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.device_end_of_life_details import DeviceEndOfLifeDetails

# TODO update the JSON string below
json = "{}"
# create an instance of DeviceEndOfLifeDetails from a JSON string
device_end_of_life_details_instance = DeviceEndOfLifeDetails.from_json(json)
# print the JSON string representation of the object
print(DeviceEndOfLifeDetails.to_json())

# convert the object into a dict
device_end_of_life_details_dict = device_end_of_life_details_instance.to_dict()
# create an instance of DeviceEndOfLifeDetails from a dict
device_end_of_life_details_form_dict = device_end_of_life_details.from_dict(device_end_of_life_details_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


