# RedirectViewServletContext


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**session_timeout** | **int** |  | [optional] 
**class_loader** | [**ApplicationContextClassLoaderParentUnnamedModuleClassLoader**](ApplicationContextClassLoaderParentUnnamedModuleClassLoader.md) |  | [optional] 
**major_version** | **int** |  | [optional] 
**minor_version** | **int** |  | [optional] 
**attribute_names** | **object** |  | [optional] 
**context_path** | **str** |  | [optional] 
**init_parameter_names** | **object** |  | [optional] 
**session_tracking_modes** | **List[str]** |  | [optional] 
**servlet_names** | **object** |  | [optional] 
**server_info** | **str** |  | [optional] 
**servlet_context_name** | **str** |  | [optional] 
**servlet_registrations** | [**Dict[str, RedirectViewServletContextServletRegistrationsValue]**](RedirectViewServletContextServletRegistrationsValue.md) |  | [optional] 
**filter_registrations** | [**Dict[str, RedirectViewServletContextFilterRegistrationsValue]**](RedirectViewServletContextFilterRegistrationsValue.md) |  | [optional] 
**session_cookie_config** | [**RedirectViewServletContextSessionCookieConfig**](RedirectViewServletContextSessionCookieConfig.md) |  | [optional] 
**default_session_tracking_modes** | **List[str]** |  | [optional] 
**effective_session_tracking_modes** | **List[str]** |  | [optional] 
**jsp_config_descriptor** | [**RedirectViewServletContextJspConfigDescriptor**](RedirectViewServletContextJspConfigDescriptor.md) |  | [optional] 
**virtual_server_name** | **str** |  | [optional] 
**request_character_encoding** | **str** |  | [optional] 
**response_character_encoding** | **str** |  | [optional] 
**effective_major_version** | **int** |  | [optional] 
**effective_minor_version** | **int** |  | [optional] 
**servlets** | **object** |  | [optional] 

## Example

```python
from cdo_sdk_python.models.redirect_view_servlet_context import RedirectViewServletContext

# TODO update the JSON string below
json = "{}"
# create an instance of RedirectViewServletContext from a JSON string
redirect_view_servlet_context_instance = RedirectViewServletContext.from_json(json)
# print the JSON string representation of the object
print(RedirectViewServletContext.to_json())

# convert the object into a dict
redirect_view_servlet_context_dict = redirect_view_servlet_context_instance.to_dict()
# create an instance of RedirectViewServletContext from a dict
redirect_view_servlet_context_form_dict = redirect_view_servlet_context.from_dict(redirect_view_servlet_context_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


