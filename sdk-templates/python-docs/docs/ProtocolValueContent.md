# ProtocolValueContent


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**protocol_value** | **str** | The literal for a custom protocol value | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.protocol_value_content import ProtocolValueContent

# TODO update the JSON string below
json = "{}"
# create an instance of ProtocolValueContent from a JSON string
protocol_value_content_instance = ProtocolValueContent.from_json(json)
# print the JSON string representation of the object
print(ProtocolValueContent.to_json())

# convert the object into a dict
protocol_value_content_dict = protocol_value_content_instance.to_dict()
# create an instance of ProtocolValueContent from a dict
protocol_value_content_form_dict = protocol_value_content.from_dict(protocol_value_content_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


