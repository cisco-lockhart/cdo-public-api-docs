# coding: utf-8

"""
    CDO API

    Use the documentation to explore the endpoints CDO has to offer

    The version of the OpenAPI document: 1.4.0
    Contact: cdo.tac@cisco.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest

from cdo_sdk_python.models.environment import Environment

class TestEnvironment(unittest.TestCase):
    """Environment unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> Environment:
        """Test Environment
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `Environment`
        """
        model = Environment()
        if include_optional:
            return Environment(
                active_profiles = [
                    ''
                    ],
                default_profiles = [
                    ''
                    ]
            )
        else:
            return Environment(
        )
        """

    def testEnvironment(self):
        """Test Environment"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()