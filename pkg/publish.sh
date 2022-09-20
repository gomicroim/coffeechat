#!/bin/sh

# checkout to the gh-pages branch
git checkout gh-pages

# pull the latest updates
git pull origin gh-pages --rebase

# remove 'node_modules' and '_book' directory
git clean -fx .idea

git rm -rf ../api
git rm -rf ../deploy
git rm -rf ../rpc
git rm -rf ../third_party

git rm -rf ../../server
git rm -rf ../../pb
git rm -rf ../../docs

# add all files
git add .

# commit
git commit -a -m "Update pkg"

# push to the origin
git push origin gh-pages

# checkout to the master branch
git checkout dev