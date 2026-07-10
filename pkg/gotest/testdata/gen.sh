#!/usr/bin/env bash
# Regenerate golden test2json streams from the fixture modules in src/.
# Streams are INPUTS to parser tests (timestamps vary per regen — that's fine).
# Usage: ./gen.sh   (from this directory)
set -uo pipefail
cd "$(dirname "$0")"

capture() {
  local name="$1"; shift
  echo "==> $name"
  (cd "src/$name" && GOFLAGS= go test -json "$@" ./... 2>/dev/null) > "$name.jsonl"
  # go test exits non-zero on failures/build errors — expected for fixtures.
  true
}

capture allpass
capture failing
capture panic_direct
capture panic_goroutine
capture count2 -count=2
capture parallel -parallel 4
capture notests

# buildfail: capture stdout AND a merged-stream variant (stderr interleaved),
# which is how pre-Go-1.24 toolchains surface compiler errors as raw text.
echo "==> buildfail"
(cd src/buildfail && GOFLAGS= go test -json ./... 2>/dev/null) > buildfail.jsonl; true
(cd src/buildfail && GOFLAGS= go test -json ./... 2>&1) > buildfail_merged.jsonl; true

echo "done"
