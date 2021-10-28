#!/usr/bin/env bash
set -euo pipefail

# Remove the submodule entry from .git/config
git submodule deinit -f third_party/temporal-server-docker-compose

# Remove the submodule directory from the superproject's .git/modules directory
rm -rf .git/modules/third_party/temporal-server-docker-compose

# Remove the entry in .gitmodules and remove the submodule directory located at path/to/submodule
git rm -f third_party/temporal-server-docker-compose

# Undo delete
# cd $PROJECT_ROOT/third_party
# git submodule add https://github.com/temporalio/docker-compose.git temporal-server-docker-compose