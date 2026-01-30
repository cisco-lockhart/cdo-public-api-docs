# ServletRegistration


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**class_name** | **str** |  | [optional] 
**init_parameters** | **Dict[str, str]** |  | [optional] 
**mappings** | **List[str]** |  | [optional] 
**name** | **str** |  | [optional] 
**run_as_role** | **str** |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.servlet_registration import ServletRegistration

# TODO update the JSON string below
json = "{}"
# create an instance of ServletRegistration from a JSON string
servlet_registration_instance = ServletRegistration.from_json(json)
# print the JSON string representation of the object
print(ServletRegistration.to_json())

# convert the object into a dict
servlet_registration_dict = servlet_registration_instance.to_dict()
# create an instance of ServletRegistration from a dict
servlet_registration_from_dict = ServletRegistration.from_dict(servlet_registration_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


