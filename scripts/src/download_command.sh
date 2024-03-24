# Function to increment version number
increment_version() {
    local version=$1
    local parts=( ${version//./ } )
    ((parts[2]++))
    new_version="${parts[0]}.${parts[1]}.${parts[2]}"
    echo "$new_version"
}

root_dir=$(dirname $(dirname $(readlink -f "$0")))
filename=openapi.yaml
declare -A urls=(
    ["public-api"]="https://edge.staging.cdo.cisco.com/api/platform/public-api/v3/api-docs.yaml"
    ["object-service"]="https://edge.staging.cdo.cisco.com/api/platform/object-service/v3/api-docs.yaml"
)

scripts/cli transform-fmc-oas

filenames=()
for service in "${!urls[@]}"; do
    url=${urls[${service}]}
    filename="${service}-openapi.yaml"
    echo -n "$(yellow Downloading file from) $(yellow_underlined ${url}) to $(blue_bold ${root_dir}/${filename})... "
    curl -X GET --silent --url  "${url}" -o $root_dir/$filename
    filenames+=("${root_dir}/${filename}")
    echo "✅︎"
done

echo "$(yellow Installing redocly from npm)... "
npm i
echo "Installed ✅︎"

echo -n "$(yellow Combining all OpenAPI YAMLs into one)... "
 ./node_modules/.bin/redocly join ${filenames[*]} -o openapi.yaml
echo "Combined ✅︎"

echo -n "$(yellow Generating Postman collection from combined OpenAPI YAMLs)... "
./node_modules/.bin/openapi2postmanv2 -s openapi.yaml -o postman-collection.json -O folderStrategy=Tags
echo "Postman collection Generated ✅︎"

echo -n "$(yellow Generating Python SDK from combined OpenAPI YAMLs)... "
current_version=$(grep -Eo "VERSION = \"[0-9]+\.[0-9]+\.[0-9]+\"" cdo-sdk/python/setup.py | cut -d'"' -f2)
new_version=$(increment_version "$current_version")
openapi-generator generate -i openapi.yaml -g python -o ./cdo-sdk/python \
--additional-properties packageName=cdo_sdk_python,packageVersion="$new_version" \
--openapi-generator-ignore-list ".github/workflows/python.yml"
echo "Python SDK Generated ✅︎"

if [[ -z ${args[--do-not-commit]} ]]; then
    cd ${root_dir}
    git checkout main
    git pull origin main
    git add .
    git commit -am "Update OpenAPI.yaml from staging"
    git push origin main
    cd -
fi