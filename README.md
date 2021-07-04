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






This is the source code accompanying my blog post on
[Generating swagger specification from Go source code](https://medium.com/@pedram.esmaeeli/generate-swagger-specification-from-go-source-code-648615f7b9d9)
