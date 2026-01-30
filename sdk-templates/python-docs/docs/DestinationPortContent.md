# DestinationPortContent

The destination port.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**elements** | **List[str]** | The list of elements. | [optional] 
**name** | **str** | The name | 
**type** | **str** | The type of the policy object. | [optional] 
**uid** | **str** | The unique identifier. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.destination_port_content import DestinationPortContent

# TODO update the JSON string below
json = "{}"
# create an instance of DestinationPortContent from a JSON string
destination_port_content_instance = DestinationPortContent.from_json(json)
# print the JSON string representation of the object
print(DestinationPortContent.to_json())

# convert the object into a dict
destination_port_content_dict = destination_port_content_instance.to_dict()
# create an instance of DestinationPortContent from a dict
destination_port_content_form_dict = destination_port_content.from_dict(destination_port_content_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


