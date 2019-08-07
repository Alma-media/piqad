#!/usr/bin/env bash
docker run --rm -it $(docker build -q .) /app/bin/piqad $@
