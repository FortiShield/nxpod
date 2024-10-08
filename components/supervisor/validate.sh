#!/bin/bash
# Copyright (c) 2023 Nxpod GmbH. All rights reserved.
# Licensed under the GNU Affero General Public License (AGPL).
# See License.AGPL.txt in the project root for license information.

set -Eeuo pipefail

ROOT_DIR="$(realpath "$(dirname "$0")/../..")"
bash "$ROOT_DIR/components/nxpod-cli/rebuild.sh"
bash "$ROOT_DIR/components/ide/code/codehelper/rebuild.sh"

DIR="$(dirname "$(realpath "$0")")"
COMPONENT="$(basename "$DIR")"
cd "$DIR"

# build
go build -gcflags=all="-N -l" .
echo "$COMPONENT built"

sudo rm -rf "/.supervisor/$COMPONENT" && true
sudo mv ./"$COMPONENT" /.supervisor
echo "$COMPONENT in /.supervisor replaced"

yarn --cwd "$DIR/frontend" run build

sudo rm -rf /.supervisor/frontend && true
sudo ln -s "$DIR/frontend/dist" /.supervisor/frontend
echo "$DIR/frontend/dist linked in /.supervisor/frontend"

gp validate --workspace-folder="$ROOT_DIR/dev/ide/example/workspace" --nxpod-env "NXPOD_ANALYTICS_SEGMENT_KEY=YErmvd89wPsrCuGcVnF2XAl846W9WIGl" --nxpod-env "GP_OPEN_EDITOR=" --nxpod-env "GP_PREVIEW_BROWSER=" --nxpod-env "GP_EXTERNAL_BROWSER=" "$@"
