#!/usr/bin/env bash

ls ../**/*.go | entr -c sh -c "cd tests && go test"
