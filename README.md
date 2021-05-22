# wasm-validators

## Building

    build.sh

## Running

    ./wasm-validators [ip]

## Example output

```bash
    ./wasm-validators ::1
    ::1 is a valid IPv6 or IPv4 address
```

```bash
./wasm-validators 127.0.0.1
127.0.0.1 is a valid IPv6 or IPv4 address
```

```bash
./wasm-validators definitely-not-valid
definitely-not-valid is not a valid IPv6 or IPv4 address
```