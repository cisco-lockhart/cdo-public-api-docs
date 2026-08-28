# MslaDevicePage


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**count** | **int** | The total number of results available. | [optional] 
**items** | [**List[MslaDeviceItemDto]**](MslaDeviceItemDto.md) | The list of items retrieved. | [optional] 
**last_polled_at** | **datetime** | Timestamp, in RFC 3339 format, of the most recent data refresh. Null if no data has been collected yet. | [optional] 
**limit** | **int** | The number of results retrieved. | [optional] 
**offset** | **int** | The offset of the results retrieved. The Security Cloud Control API uses the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.msla_device_page import MslaDevicePage

# TODO update the JSON string below
json = "{}"
# create an instance of MslaDevicePage from a JSON string
msla_device_page_instance = MslaDevicePage.from_json(json)
# print the JSON string representation of the object
print(MslaDevicePage.to_json())

# convert the object into a dict
msla_device_page_dict = msla_device_page_instance.to_dict()
# create an instance of MslaDevicePage from a dict
msla_device_page_form_dict = msla_device_page.from_dict(msla_device_page_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


