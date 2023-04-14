#!/bin/bash

for file in *.proto; do
  protoc --go_out=. --go_opt=paths=source_relative $file
done

