#!/bin/bash
[[ ${BASH_VERSINFO[0]} -lt 5 ]] && { echo "bash 5 minimum"; exit 1; }
die() { echo "${1:-waaah}" >&2; exit "${2:-1}"; }
err() { echo "!  $1"; ((${2:-0}>0)) && exit "$2"; }
hash go 2>/dev/null || die "missing dep: go"


go build || die "cant build"


echo "help"
./kao &>/dev/null; e=$?
((e == 64)) || err "./kao : expected exit code 64, got $e"

echo "happy"
o=$(./kao happy); e=$?
((e > 0)) && err "./kao happy : expected exit code 0, got $e"
[[ "$o" == "" ]] && err "./kao happy : expected some output, got none"

echo "sad"
o=$(./kao sad); e=$?
((e > 0)) && err "./kao sad : expected exit code 0, got $e"
[[ "$o" == "" ]] && err "./kao sad : expected some output, got none"

echo "not-a-real-girl"
o=$(./kao not-a-real-girl); e=$?
((e == 0)) && err "./kao not-a-real-girl : expected non-zero exit code, didn't get it"

echo "all"
o=$(./kao -all angry); e=$?
((e > 0)) && err "./kao -all angry : expected exit code 0, got $e"
(($(wc -l <<<"$o") > 1)) || err "./kao -all angry : expected multi-line output"
