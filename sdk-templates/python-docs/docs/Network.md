# Network


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **str** | The unique identifier, represented as a UUID, of the Meraki Dashboard Network. | [optional] 
**name** | **str** | The name of the Meraki Dashboard Network. | [optional] 
**timezone** | **str** | The timezone of the Meraki Dashboard Network. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.network import Network

# TODO update the JSON string below
json = "{}"
# create an instance of Network from a JSON string
network_instance = Network.from_json(json)
# print the JSON string representation of the object
print(Network.to_json())

# convert the object into a dict
network_dict = network_instance.to_dict()
# create an instance of Network from a dict
network_form_dict = network.from_dict(network_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


