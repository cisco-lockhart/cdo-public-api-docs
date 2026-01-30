# ProtocolContent

The protocol. Defaults to IP.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**elements** | **List[str]** | The list of elements. | [optional] 
**name** | **str** | The name | 
**type** | **str** | The type of the policy object. | [optional] 
**uid** | **str** | The unique identifier. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.protocol_content import ProtocolContent

# TODO update the JSON string below
json = "{}"
# create an instance of ProtocolContent from a JSON string
protocol_content_instance = ProtocolContent.from_json(json)
# print the JSON string representation of the object
print(ProtocolContent.to_json())

# convert the object into a dict
protocol_content_dict = protocol_content_instance.to_dict()
# create an instance of ProtocolContent from a dict
protocol_content_form_dict = protocol_content.from_dict(protocol_content_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


