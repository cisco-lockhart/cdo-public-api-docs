# RedirectViewServletContextJspConfigDescriptor


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**taglibs** | [**List[RedirectViewServletContextJspConfigDescriptorTaglibsInner]**](RedirectViewServletContextJspConfigDescriptorTaglibsInner.md) |  | [optional] 
**jsp_property_groups** | [**List[RedirectViewServletContextJspConfigDescriptorJspPropertyGroupsInner]**](RedirectViewServletContextJspConfigDescriptorJspPropertyGroupsInner.md) |  | [optional] 

## Example

```python
from cdo_sdk_python.models.redirect_view_servlet_context_jsp_config_descriptor import RedirectViewServletContextJspConfigDescriptor

# TODO update the JSON string below
json = "{}"
# create an instance of RedirectViewServletContextJspConfigDescriptor from a JSON string
redirect_view_servlet_context_jsp_config_descriptor_instance = RedirectViewServletContextJspConfigDescriptor.from_json(json)
# print the JSON string representation of the object
print(RedirectViewServletContextJspConfigDescriptor.to_json())

# convert the object into a dict
redirect_view_servlet_context_jsp_config_descriptor_dict = redirect_view_servlet_context_jsp_config_descriptor_instance.to_dict()
# create an instance of RedirectViewServletContextJspConfigDescriptor from a dict
redirect_view_servlet_context_jsp_config_descriptor_form_dict = redirect_view_servlet_context_jsp_config_descriptor.from_dict(redirect_view_servlet_context_jsp_config_descriptor_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


