


check-swagger:
	which swagger || (GO111MODULE=on go get -u github.com/go-swagger/go-swagger/cmd/swagger)

swagger: check-swagger
	GO111MODULE=on swagger generate spec -o ./swagger.yaml --scan-models

sdk: swagger-json
	docker run --rm  \
		-v ${PWD}:/local \
		swaggerapi/swagger-codegen-cli:2.4.21 \
		generate -i /local/swagger.json -l \
		go --http-user-agent \'BOZO-SDK/v0/go\' -o /local/sdk \
		--additional-properties packageName=bozosdk --additional-properties packageVersion='v0.1.0'


swagger-json: check-swagger
	GO111MODULE=on swagger generate spec -o ./swagger.json --scan-models


serve-swagger: check-swagger
	swagger serve -F=swagger swagger.yaml
