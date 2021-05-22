#!/bin/sh
rustc +nightly --crate-type=cdylib  --target wasm32-unknown-unknown -O ip_validator.rs && go build