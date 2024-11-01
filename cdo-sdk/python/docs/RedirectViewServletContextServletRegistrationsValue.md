# RedirectViewServletContextServletRegistrationsValue


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**mappings** | **List[str]** |  | [optional] 
**run_as_role** | **str** |  | [optional] 
**name** | **str** |  | [optional] 
**class_name** | **str** |  | [optional] 
**init_parameters** | **Dict[str, str]** |  | [optional] 

## Example

```python
from cdo_sdk_python.models.redirect_view_servlet_context_servlet_registrations_value import RedirectViewServletContextServletRegistrationsValue

# TODO update the JSON string below
json = "{}"
# create an instance of RedirectViewServletContextServletRegistrationsValue from a JSON string
redirect_view_servlet_context_servlet_registrations_value_instance = RedirectViewServletContextServletRegistrationsValue.from_json(json)
# print the JSON string representation of the object
print(RedirectViewServletContextServletRegistrationsValue.to_json())

# convert the object into a dict
redirect_view_servlet_context_servlet_registrations_value_dict = redirect_view_servlet_context_servlet_registrations_value_instance.to_dict()
# create an instance of RedirectViewServletContextServletRegistrationsValue from a dict
redirect_view_servlet_context_servlet_registrations_value_form_dict = redirect_view_servlet_context_servlet_registrations_value.from_dict(redirect_view_servlet_context_servlet_registrations_value_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


