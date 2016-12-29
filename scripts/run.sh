#!/bin/sh

echo "Running local API with reload enabled"
gin --port 3000 --appPort 8000
