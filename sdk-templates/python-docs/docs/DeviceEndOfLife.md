# DeviceEndOfLife


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**end_of_life_details** | [**DeviceEndOfLifeDetails**](DeviceEndOfLifeDetails.md) |  | [optional] 
**model_numbers** | **List[str]** | A list of hardware models subject to the End-of-Life evaluation. | [optional] 
**recommendations** | [**List[DeviceEndOfLifeRecommendation]**](DeviceEndOfLifeRecommendation.md) | End-Of-Life appropriate upgrade recommendations for the hardware models. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.device_end_of_life import DeviceEndOfLife

# TODO update the JSON string below
json = "{}"
# create an instance of DeviceEndOfLife from a JSON string
device_end_of_life_instance = DeviceEndOfLife.from_json(json)
# print the JSON string representation of the object
print(DeviceEndOfLife.to_json())

# convert the object into a dict
device_end_of_life_dict = device_end_of_life_instance.to_dict()
# create an instance of DeviceEndOfLife from a dict
device_end_of_life_form_dict = device_end_of_life.from_dict(device_end_of_life_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


