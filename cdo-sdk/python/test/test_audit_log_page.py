# coding: utf-8

"""
    CDO API

    Use the documentation to explore the endpoints CDO has to offer

    The version of the OpenAPI document: 1.1.0
    Contact: cdo.tac@cisco.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest

from cdo_sdk_python.models.audit_log_page import AuditLogPage

class TestAuditLogPage(unittest.TestCase):
    """AuditLogPage unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> AuditLogPage:
        """Test AuditLogPage
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `AuditLogPage`
        """
        model = AuditLogPage()
        if include_optional:
            return AuditLogPage(
                count = 100,
                limit = 50,
                offset = 0,
                items = [
                    cdo_sdk_python.models.audit_log.AuditLog(
                        uid = '7131daad-e813-4b8f-8f42-be1e241e8cdb', 
                        event_type = 'USER_LOGGED_IN', 
                        username = 'test@cisco.com', 
                        event_description = 'test@cisco.com logged in', 
                        event_time = '2024-06-26T20:44:06Z', 
                        roles = [ROLE_READ_ONLY, ROLE_DEPLOY_ONLY], )
                    ]
            )
        else:
            return AuditLogPage(
        )
        """

    def testAuditLogPage(self):
        """Test AuditLogPage"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()