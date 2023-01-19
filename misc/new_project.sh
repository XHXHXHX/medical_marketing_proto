#!/bin/bash

cur_pwd=`pwd`
PROJECT_NAME=`basename $cur_pwd`

DEMO_GIT="http://gitlab.aiforward.cn/proto/demo.git"
DEMO_PROJECT="demo"

echo "=> Git clone demo..."
rm -rf $DEMO_PROJECT
git clone $DEMO_GIT
rm -rf $DEMO_PROJECT/.git
cp -rf $DEMO_PROJECT/* .
rm -rf $DEMO_PROJECT
mv proto/demo proto/$PROJECT_NAME

echo "=> Replace go.mod module"
sed -i "" 's@module gitlab.aiforward.cn/proto/demo@module gitlab.aiforward.cn/proto/'$PROJECT_NAME'@g' go.mod
rm -rf go.sum

echo "=> Replace proto files"
find  ./proto -type f  | xargs -I {} sed -i "" 's@gitlab.aiforward.cn/proto/demo/gen/go/proto/demo@gitlab.aiforward.cn/proto/'$PROJECT_NAME'/gen/go/proto/'$PROJECT_NAME'@g' {}
find  ./proto -type f  | xargs -I {} sed -i "" 's@gitlab.aiforward.cn/proto/demo@gitlab.aiforward.cn/proto/'$PROJECT_NAME'@g' {}
find  ./proto -type f  | xargs -I {} sed -i "" 's@proto/demo@proto/'$PROJECT_NAME'@g' {}
find  ./proto -type f  | xargs -I {} sed -i "" 's@demo\.@'$PROJECT_NAME'\.@g' {}
