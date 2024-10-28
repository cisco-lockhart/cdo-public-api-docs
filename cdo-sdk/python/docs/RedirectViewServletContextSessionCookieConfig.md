# RedirectViewServletContextSessionCookieConfig


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**http_only** | **bool** |  | [optional] 
**name** | **str** |  | [optional] 
**path** | **str** |  | [optional] 
**comment** | **str** |  | [optional] 
**domain** | **str** |  | [optional] 
**max_age** | **int** |  | [optional] 
**secure** | **bool** |  | [optional] 

## Example

```python
from cdo_sdk_python.models.redirect_view_servlet_context_session_cookie_config import RedirectViewServletContextSessionCookieConfig

# TODO update the JSON string below
json = "{}"
# create an instance of RedirectViewServletContextSessionCookieConfig from a JSON string
redirect_view_servlet_context_session_cookie_config_instance = RedirectViewServletContextSessionCookieConfig.from_json(json)
# print the JSON string representation of the object
print(RedirectViewServletContextSessionCookieConfig.to_json())

# convert the object into a dict
redirect_view_servlet_context_session_cookie_config_dict = redirect_view_servlet_context_session_cookie_config_instance.to_dict()
# create an instance of RedirectViewServletContextSessionCookieConfig from a dict
redirect_view_servlet_context_session_cookie_config_form_dict = redirect_view_servlet_context_session_cookie_config.from_dict(redirect_view_servlet_context_session_cookie_config_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


