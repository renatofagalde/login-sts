#!/bin/sh

set -e

echo " loading environment"
source /app/app.env


echo "start the app"
exec "$@"