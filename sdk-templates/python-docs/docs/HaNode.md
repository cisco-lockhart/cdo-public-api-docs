# HaNode


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**chassis_serial** | **str** | The serial number of the chassis on the device. | [optional] 
**name** | **str** | The name of the node on the device. | [optional] 
**role** | **str** | (cdFMC-managed FTDs only) Information on the current role of the node in the HA Pair. | [optional] 
**serial** | **str** | The serial number of the node on the device. This is typically used for licensing, and is not the same as the chassis&#39; serial number. | [optional] 
**software_version** | **str** | The version of the software running on the device. | [optional] 
**status** | **str** | The status of the HA node. | [optional] 
**uid_on_fmc** | **str** | The unique identifier, represented as a UUID, of the device on the FMC. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.ha_node import HaNode

# TODO update the JSON string below
json = "{}"
# create an instance of HaNode from a JSON string
ha_node_instance = HaNode.from_json(json)
# print the JSON string representation of the object
print(HaNode.to_json())

# convert the object into a dict
ha_node_dict = ha_node_instance.to_dict()
# create an instance of HaNode from a dict
ha_node_form_dict = ha_node.from_dict(ha_node_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


