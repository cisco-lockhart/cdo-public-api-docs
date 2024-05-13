# ServiceObjectValueContent

The value of the service object

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**type** | **str** |  | [optional] 
**icmp4_type** | **str** |  | [optional] 
**icmp4_code** | **str** |  | [optional] 
**icmp6_type** | **str** |  | [optional] 
**icmp6_code** | **str** |  | [optional] 
**source** | [**PortsValue**](PortsValue.md) |  | [optional] 
**destination** | [**PortsValue**](PortsValue.md) |  | [optional] 
**literal** | **str** | The literal port or range of port values | [optional] 
**protocol_value** | **str** | The literal for a custom protocol value | [optional] 

## Example

```python
from cdo_sdk_python.models.service_object_value_content import ServiceObjectValueContent

# TODO update the JSON string below
json = "{}"
# create an instance of ServiceObjectValueContent from a JSON string
service_object_value_content_instance = ServiceObjectValueContent.from_json(json)
# print the JSON string representation of the object
print(ServiceObjectValueContent.to_json())

# convert the object into a dict
service_object_value_content_dict = service_object_value_content_instance.to_dict()
# create an instance of ServiceObjectValueContent from a dict
service_object_value_content_form_dict = service_object_value_content.from_dict(service_object_value_content_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


