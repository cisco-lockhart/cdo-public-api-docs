# ObjectContent

The content of an object. Which properties are populated is determined by the objectType of the enclosing object: network objects carry a literal, URL objects a url, service objects a protocol and serviceValue, and groups their literals and referencedObjectUids.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**dns_resolution** | **str** | For FQDN network objects sourced from FTD/cdFMC, the DNS resolution family that governs which DNS record types FTD resolves at policy-enforcement time. Populated only for FQDN objects; absent otherwise. | [optional] 
**literal** | **str** | The literal content of the network object | [optional] 
**literals** | [**List[SingleContent]**](SingleContent.md) | (groups only) The group&#39;s inline members. Each member has the same shape as the content of the matching single object: { \&quot;literal\&quot;: \&quot;10.10.10.0/24\&quot; } in a NETWORK_GROUP, { \&quot;url\&quot;: \&quot;https://www.cisco.com/\&quot; } in a URL_GROUP, and { \&quot;protocol\&quot;: \&quot;TCP\&quot;, \&quot;serviceValue\&quot;: { \&quot;literal\&quot;: \&quot;443\&quot; } } in a SERVICE_GROUP. | [optional] 
**protocol** | **str** | The service object protocol | [optional] 
**referenced_object_uids** | **List[str]** | (groups only) The UIDs of existing objects, or groups of the same kind, that are members of the group | [optional] 
**service_value** | [**ServiceObjectValueContent**](ServiceObjectValueContent.md) |  | [optional] 
**url** | **str** | The URL literal | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.object_content import ObjectContent

# TODO update the JSON string below
json = "{}"
# create an instance of ObjectContent from a JSON string
object_content_instance = ObjectContent.from_json(json)
# print the JSON string representation of the object
print(ObjectContent.to_json())

# convert the object into a dict
object_content_dict = object_content_instance.to_dict()
# create an instance of ObjectContent from a dict
object_content_form_dict = object_content.from_dict(object_content_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


