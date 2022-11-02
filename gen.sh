npx @openapitools/openapi-generator-cli generate -i api_v1.yaml -g go-gin-server -o server_stub
npx @openapitools/openapi-generator-cli generate -i api_v1.yaml -g python-prior -o control_hub/api
npx @openapitools/openapi-generator-cli generate -i api_v1.yaml -g dart-dio -o ui/api