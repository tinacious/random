#!/usr/bin/env sh
VERSION="0.0.0-$(date +%s)";

git tag -a $VERSION -m "'$VERSION'"
echo "✅ Created version $VERSION";
