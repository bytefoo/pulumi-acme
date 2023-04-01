#!/bin/bash

VERSION=$1
TAG_VERSION=v$VERSION

echo $TAG_VERSION

git tag -a $TAG_VERSION -m "wip $TAG_VERSION"
git push origin $TAG_VERSION