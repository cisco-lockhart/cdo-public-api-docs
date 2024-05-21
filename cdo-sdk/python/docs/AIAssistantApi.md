# cdo_sdk_python.AIAssistantApi

All URIs are relative to *https://edge.us.cdo.cisco.com/api/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_conversation**](AIAssistantApi.md#get_conversation) | **GET** /v1/ai-assistant/conversations/{conversationUid} | Get Conversation
[**get_conversations**](AIAssistantApi.md#get_conversations) | **GET** /v1/ai-assistant/conversations | Get Conversations


# **get_conversation**
> AiAssistantConversation get_conversation(conversation_uid)

Get Conversation

Get an AI Assistant conversation by UID in the CDO tenant.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import cdo_sdk_python
from cdo_sdk_python.models.ai_assistant_conversation import AiAssistantConversation
from cdo_sdk_python.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.cdo.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = cdo_sdk_python.Configuration(
    host = "https://edge.us.cdo.cisco.com/api/rest"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = cdo_sdk_python.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with cdo_sdk_python.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = cdo_sdk_python.AIAssistantApi(api_client)
    conversation_uid = 'conversation_uid_example' # str | The unique identifier of the conversation in CDO.

    try:
        # Get Conversation
        api_response = api_instance.get_conversation(conversation_uid)
        print("The response of AIAssistantApi->get_conversation:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling AIAssistantApi->get_conversation: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **conversation_uid** | **str**| The unique identifier of the conversation in CDO. | 

### Return type

[**AiAssistantConversation**](AiAssistantConversation.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Conversation object |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_conversations**
> AiAssistantConversation get_conversations(limit=limit, offset=offset, sort=sort)

Get Conversations

Get a list of AI Assistant Conversations. Note: the total number of conversations is set to -1 as this information is currently unavailable.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import cdo_sdk_python
from cdo_sdk_python.models.ai_assistant_conversation import AiAssistantConversation
from cdo_sdk_python.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.cdo.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = cdo_sdk_python.Configuration(
    host = "https://edge.us.cdo.cisco.com/api/rest"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = cdo_sdk_python.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with cdo_sdk_python.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = cdo_sdk_python.AIAssistantApi(api_client)
    limit = '50' # str | The number of results to retrieve. (optional) (default to '50')
    offset = '0' # str | The offset of the results retrieved. The CDO API uses the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional) (default to '0')
    sort = ['name:DESC'] # List[str] | The fields to sort results by. (optional)

    try:
        # Get Conversations
        api_response = api_instance.get_conversations(limit=limit, offset=offset, sort=sort)
        print("The response of AIAssistantApi->get_conversations:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling AIAssistantApi->get_conversations: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **str**| The number of results to retrieve. | [optional] [default to &#39;50&#39;]
 **offset** | **str**| The offset of the results retrieved. The CDO API uses the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] [default to &#39;0&#39;]
 **sort** | [**List[str]**](str.md)| The fields to sort results by. | [optional] 

### Return type

[**AiAssistantConversation**](AiAssistantConversation.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of Conversations |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

