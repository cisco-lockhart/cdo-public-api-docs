# PortValue


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**literal** | **str** | The literal port or range of port values | [optional] 

## Example

```python
from cdo_sdk_python.models.port_value import PortValue

# TODO update the JSON string below
json = "{}"
# create an instance of PortValue from a JSON string
port_value_instance = PortValue.from_json(json)
# print the JSON string representation of the object
print(PortValue.to_json())

# convert the object into a dict
port_value_dict = port_value_instance.to_dict()
# create an instance of PortValue from a dict
port_value_form_dict = port_value.from_dict(port_value_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


