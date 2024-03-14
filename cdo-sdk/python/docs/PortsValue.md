# PortsValue


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**op** | **str** | The operator applied to the list of ports | [optional] 
**ports** | **List[str]** | The list of ports | [optional] 

## Example

```python
from openapi_client.models.ports_value import PortsValue

# TODO update the JSON string below
json = "{}"
# create an instance of PortsValue from a JSON string
ports_value_instance = PortsValue.from_json(json)
# print the JSON string representation of the object
print(PortsValue.to_json())

# convert the object into a dict
ports_value_dict = ports_value_instance.to_dict()
# create an instance of PortsValue from a dict
ports_value_form_dict = ports_value.from_dict(ports_value_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


