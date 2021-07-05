Generate swagger specification from Go source code
--------------------------------------------------

# Step 1

```bash

go run .

```

# Step 2

Perform this step in another window

```bash
curl -u foo:bar   -X 'POST' \
  'http://localhost:8000/foobar' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "bar": [
    0
  ],
  "foo": "string"
}'

```

# Step 3

```bash
curl -u foo:bar -X 'POST' \
  'http://localhost:8000/' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "msg": "string"
}'
```

# Step 4

```bash

make swagger-json

docker run --rm  \
-v ${PWD}:/local \
swaggerapi/swagger-codegen-cli:2.4.21 \
generate -i /local/swagger.json -l \
go --http-user-agent \'BOZO-SDK/v0/go\' -o /local/sdk \
--additional-properties packageName=bozosdk --additional-properties packageVersion='v0.1.0'

```





This is the source code accompanying my blog post on
[Generating swagger specification from Go source code](https://medium.com/@pedram.esmaeeli/generate-swagger-specification-from-go-source-code-648615f7b9d9)
