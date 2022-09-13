#!/bin/sh

# checkout to the gh-pages branch
git checkout gh-pages

# pull the latest updates
git pull origin gh-pages --rebase

# remove 'node_modules' and '_book' directory
git clean -fx api
git clean -fx deploy
git clean -fx rpc
git clean -fx third_party

# add all files
git add .

# commit
git commit -a -m "Update pkg"

# push to the origin
git push origin gh-pages

# checkout to the master branch
git checkout dev