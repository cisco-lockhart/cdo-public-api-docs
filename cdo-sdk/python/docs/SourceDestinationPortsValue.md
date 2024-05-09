# SourceDestinationPortsValue


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**source** | [**PortValue**](PortValue.md) |  | [optional] 
**destination** | [**PortValue**](PortValue.md) |  | [optional] 

## Example

```python
from cdo_sdk_python.models.source_destination_ports_value import SourceDestinationPortsValue

# TODO update the JSON string below
json = "{}"
# create an instance of SourceDestinationPortsValue from a JSON string
source_destination_ports_value_instance = SourceDestinationPortsValue.from_json(json)
# print the JSON string representation of the object
print(SourceDestinationPortsValue.to_json())

# convert the object into a dict
source_destination_ports_value_dict = source_destination_ports_value_instance.to_dict()
# create an instance of SourceDestinationPortsValue from a dict
source_destination_ports_value_form_dict = source_destination_ports_value.from_dict(source_destination_ports_value_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


