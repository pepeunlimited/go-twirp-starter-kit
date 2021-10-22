#!/usr/bin/env bash

set -euo pipefail

ORIGIN=$1
git remote set-url origin $ORIGIN
git push --set-upstream origin master
