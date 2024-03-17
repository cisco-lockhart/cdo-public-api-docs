# RaVpnSessionPage


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**count** | **int** | The total number of results available. | [optional] 
**limit** | **int** | The number of results retrieved. | [optional] 
**offset** | **int** | The offset of the results retrieved. The CDO Public API uses the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] 
**items** | [**List[RaVpnSession]**](RaVpnSession.md) | The list of items retrieved. | [optional] 

## Example

```python
from cdo-python-sdk.models.ra_vpn_session_page import RaVpnSessionPage

# TODO update the JSON string below
json = "{}"
# create an instance of RaVpnSessionPage from a JSON string
ra_vpn_session_page_instance = RaVpnSessionPage.from_json(json)
# print the JSON string representation of the object
print(RaVpnSessionPage.to_json())

# convert the object into a dict
ra_vpn_session_page_dict = ra_vpn_session_page_instance.to_dict()
# create an instance of RaVpnSessionPage from a dict
ra_vpn_session_page_form_dict = ra_vpn_session_page.from_dict(ra_vpn_session_page_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


