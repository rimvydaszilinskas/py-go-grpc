## Render GO code

```
protoc --go_out=plugins=grpc:server/countries countries.proto
```

## Render Python code

Source env before running

```
python -m grpc_tools.protoc -I. --python_out=client --grpc_python_out=client countries.proto
```
