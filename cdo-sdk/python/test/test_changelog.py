# coding: utf-8

"""
    Cisco Defense Orchestrator API

    Use the interactive documentation to explore the endpoints CDO has to offer

    The version of the OpenAPI document: 0.0.1
    Contact: cdo.tac@cisco.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest

from cdo_sdk_python.models.changelog import Changelog

class TestChangelog(unittest.TestCase):
    """Changelog unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> Changelog:
        """Test Changelog
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `Changelog`
        """
        model = Changelog()
        if include_optional:
            return Changelog(
                uid = '7131daad-e813-4b8f-8f42-be1e241e8cdb',
                status = 'ACTIVE',
                entity_uid = '7131daad-e813-4b8f-8f42-be1e241e8cdb',
                events = [
                    cdo_sdk_python.models.event.Event(
                        description = '', 
                        diff = '', 
                        username = 'myuser@cisco.com', 
                        date = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                        action = 'CREATE', 
                        change_request_uid = '7131daad-e813-4b8f-8f42-be1e241e8cdb', 
                        change_request_name = 'LH-12345', )
                    ]
            )
        else:
            return Changelog(
                uid = '7131daad-e813-4b8f-8f42-be1e241e8cdb',
        )
        """

    def testChangelog(self):
        """Test Changelog"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()