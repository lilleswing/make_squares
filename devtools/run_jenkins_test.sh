#!/bin/bash
export PATH="`pwd`/anaconda/bin:$PATH"
source activate make_squares
export gcc=`which gcc`
export CC=`which gcc`

go test -count=1 ./cmd/... ./pkg/...
