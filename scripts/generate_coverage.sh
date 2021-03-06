#!/bin/bash -e
# Requires local installation of: `github.com/wadey/gocovmerge`

FULL_COVERAGE_OUT=${1:-"full_coverage.out"}
COV_DIR=${2:-"./cov"}

cd ${GOPATH}/src/github.com/cosminrentea/gobbler

rm -rf ${COV_DIR}
mkdir ${COV_DIR}

for dir in $(find . -maxdepth 10 -not -path './.git*' -not -path './vendor/*' -not -path '*/_test.go' -type d);
do
    if ls ${dir}/*_test.go &> /dev/null; then
        COVERAGE_OUT=$(echo ${dir} | tr './' '-')
        echo "Generating test coverage for dir in file: ${dir} : ${COVERAGE_OUT}"
        GO_TEST_DISABLED=true go test -v -covermode=atomic -coverprofile=${COV_DIR}/${COVERAGE_OUT}.out ./${dir}
    fi
done

wait

gocovmerge ${COV_DIR}/*.out > ${FULL_COVERAGE_OUT}

rm -rf ${COV_DIR}
