# Icmp6Value


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**icmp6_type** | **str** |  | [optional] 
**icmp6_code** | **str** |  | [optional] 

## Example

```python
from openapi_client.models.icmp6_value import Icmp6Value

# TODO update the JSON string below
json = "{}"
# create an instance of Icmp6Value from a JSON string
icmp6_value_instance = Icmp6Value.from_json(json)
# print the JSON string representation of the object
print(Icmp6Value.to_json())

# convert the object into a dict
icmp6_value_dict = icmp6_value_instance.to_dict()
# create an instance of Icmp6Value from a dict
icmp6_value_form_dict = icmp6_value.from_dict(icmp6_value_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


