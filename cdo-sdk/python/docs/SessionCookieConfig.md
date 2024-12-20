# SessionCookieConfig


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**http_only** | **bool** |  | [optional] 
**name** | **str** |  | [optional] 
**path** | **str** |  | [optional] 
**attributes** | **Dict[str, str]** |  | [optional] 
**comment** | **str** |  | [optional] 
**max_age** | **int** |  | [optional] 
**secure** | **bool** |  | [optional] 
**domain** | **str** |  | [optional] 

## Example

```python
from cdo_sdk_python.models.session_cookie_config import SessionCookieConfig

# TODO update the JSON string below
json = "{}"
# create an instance of SessionCookieConfig from a JSON string
session_cookie_config_instance = SessionCookieConfig.from_json(json)
# print the JSON string representation of the object
print(SessionCookieConfig.to_json())

# convert the object into a dict
session_cookie_config_dict = session_cookie_config_instance.to_dict()
# create an instance of SessionCookieConfig from a dict
session_cookie_config_form_dict = session_cookie_config.from_dict(session_cookie_config_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


