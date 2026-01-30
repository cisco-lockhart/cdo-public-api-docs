# Icmp4Value


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**icmp4_code** | **str** |  | [optional] 
**icmp4_type** | **str** |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.icmp4_value import Icmp4Value

# TODO update the JSON string below
json = "{}"
# create an instance of Icmp4Value from a JSON string
icmp4_value_instance = Icmp4Value.from_json(json)
# print the JSON string representation of the object
print(Icmp4Value.to_json())

# convert the object into a dict
icmp4_value_dict = icmp4_value_instance.to_dict()
# create an instance of Icmp4Value from a dict
icmp4_value_form_dict = icmp4_value.from_dict(icmp4_value_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


