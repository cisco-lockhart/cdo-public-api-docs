# DeviceEndOfLifeRecommendation


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**data_sheet_url** | **str** | The Link to the Product Datasheet. | [optional] 
**form_factor** | **str** | The physical or virtual design and deployment format of the device. | [optional] 
**interfaces** | **str** | The physical or virtual ports through which the device communicates with networks. | [optional] 
**product_label** | **str** | Cisco product labels specify the firewall series name, model number, and included security feature set (e.g., Threat Defense, VPN), providing clear identification of the device’s capabilities, performance level, and intended deployment use case. | [optional] 
**product_number** | **str** | A unique identifier assigned to a specific Cisco device model or configuration, used to specify its features, performance tier, and hardware specifications | [optional] 
**throughput** | **str** | The rate at which the firewall can process network traffic. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.device_end_of_life_recommendation import DeviceEndOfLifeRecommendation

# TODO update the JSON string below
json = "{}"
# create an instance of DeviceEndOfLifeRecommendation from a JSON string
device_end_of_life_recommendation_instance = DeviceEndOfLifeRecommendation.from_json(json)
# print the JSON string representation of the object
print(DeviceEndOfLifeRecommendation.to_json())

# convert the object into a dict
device_end_of_life_recommendation_dict = device_end_of_life_recommendation_instance.to_dict()
# create an instance of DeviceEndOfLifeRecommendation from a dict
device_end_of_life_recommendation_form_dict = device_end_of_life_recommendation.from_dict(device_end_of_life_recommendation_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


