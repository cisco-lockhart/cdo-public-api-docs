# SourceNetworkContent

The source network.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**elements** | **List[str]** | The list of elements. | [optional] 
**name** | **str** | The name | 
**type** | **str** | The type of the policy object. | [optional] 
**uid** | **str** | The unique identifier. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.source_network_content import SourceNetworkContent

# TODO update the JSON string below
json = "{}"
# create an instance of SourceNetworkContent from a JSON string
source_network_content_instance = SourceNetworkContent.from_json(json)
# print the JSON string representation of the object
print(SourceNetworkContent.to_json())

# convert the object into a dict
source_network_content_dict = source_network_content_instance.to_dict()
# create an instance of SourceNetworkContent from a dict
source_network_content_form_dict = source_network_content.from_dict(source_network_content_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


