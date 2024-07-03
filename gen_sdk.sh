# cd ./mist_openapi
# git pull
# cd ..

# python3 ./gen_sdk_pre.py

openapi-generator generate \
    -i ./mist.sdk.yml \
    -g go \
    -o ./sdk \
    --api-package mistapi-go \
    --additional-properties=packageName=mistapigo \
    --additional-properties=disallowAdditionalPropertiesIfNotPresent=false \
    --additional-properties=enumClassPrefix=true \
    --additional-properties=generateInterfaces=true \
    --additional-properties=packageVersion=1.0.0 \
    --git-repo-id msitapi-go \
    --git-user-id tmunzer


cp ./fix_sdk/* ./sdk/