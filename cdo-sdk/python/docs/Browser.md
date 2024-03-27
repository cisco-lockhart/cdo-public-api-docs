# Browser

The web browser running on the client device.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | The name of the web browser running on the client device. | [optional] 
**version** | **str** | The version of the web browser running on the client device. | [optional] 

## Example

```python
from cdo_sdk_python.models.browser import Browser

# TODO update the JSON string below
json = "{}"
# create an instance of Browser from a JSON string
browser_instance = Browser.from_json(json)
# print the JSON string representation of the object
print(Browser.to_json())

# convert the object into a dict
browser_dict = browser_instance.to_dict()
# create an instance of Browser from a dict
browser_form_dict = browser.from_dict(browser_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


