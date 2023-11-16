root_dir=$(dirname $(dirname $(readlink -f "$0")))
filename=openapi.yaml
#url="https://edge.staging.cdo.cisco.com/api/platform/public-api/v3/api-docs.yaml"
url="http://localhost:4033/v3/api-docs.yaml"
echo -n "$(yellow Downloading file from) $(yellow_underlined ${url}) to $(blue_bold ${root_dir}/${filename})... "

curl -X GET --silent --url  "${url}" -o $root_dir/$filename

if [[ -z ${args[--do-not-commit]} ]]; then
    cd ${root_dir}
    git checkout main
    git pull origin main
    git add .
    git commit -am "Update OpenAPI.yaml from staging"
    git push origin main
    cd -
fi