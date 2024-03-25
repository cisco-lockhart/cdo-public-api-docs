secret_arn=arn:aws:secretsmanager:us-west-2:107042026245:secret:jenkins-pypi-credentials-mD4NdK
echo -n "$(yellow Retrieving the secret value from AWS Secrets Manager)... "
secret_value=$(aws secretsmanager get-secret-value --secret-id $secret_arn --query SecretString --output text)

package_name=$(grep -Eo 'NAME = "[^"]+"' cdo-sdk/python/setup.py | cut -d'"' -f2)
current_version=$(grep -Eo "VERSION = \"[0-9]+\.[0-9]+\.[0-9]+\"" cdo-sdk/python/setup.py | cut -d'"' -f2)

cd cdo-sdk/python
pip3 install wheel twine
echo "PATH: $PATH"
echo "Wheel installation directory: $(which wheel)"
echo "Twine installation directory: $(which twine)"

echo -n "$(yellow Creating the Wheel and Source)... "
python3 setup.py sdist bdist_wheel

echo -n "$(yellow Publishing to PyPI)... "
twine upload --username __token__ --password $secret_value dist/${package_name}-${current_version}.tar.gz dist/${package_name//-/_}-${current_version}-py3-none-any.whl

unset secret_value