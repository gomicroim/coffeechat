#!/bin/sh

read -r -p "此操作会清理除了pkg外所有文件，请确认所有修改已提交后继续? [Y/n] " input

case $input in
    [yY][eE][sS]|[yY])
		echo "Yes"
		;;

    [nN][oO]|[nN])
		echo "No"
                exit 1
       	        ;;
    *)
		echo "Invalid input..."
		exit 1
		;;
esac

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

git rm -rf ../docs
git rm -rf ../*.yml
git rm -rf ../*.md

# add all files
git add .

# commit
git commit -a -m "Update pkg"

# push to the origin
git push origin gh-pages

# checkout to the master branch
git checkout dev