#!/usr/bin/env bash

echo  "\033[30;40m""\033[1m#############################################################################\033[0m"
echo ""

read -r -p "是否仍要继续 [Y/n] " continue
case $continue in
[yY][eE][sS] | [yY])
    echo ""
    echo  "\033[30;40m""\033[1m接下来继续执行code generator.....\033[0m"
    echo ""
    ;;
  *)
    echo ""
    echo  "\033[30;40m""\033[1m终止 cg.sh 流程，请移步>>>>>> \033[0m"  "\033[36;40m""\033[1m${wiki} \033[0m"
    echo ""
    exit
    ;;
esac
#######################################################################################################################

# 使用帮助
if [ $# -lt 1 ] || [ "-h" = "$1" ] || [ "--help" = "$1" ]
then
    echo "介绍: 初始化项目"
    echo "用法: sh $0 project"
    echo "说明: project: 项目名称"
    echo "示例: sh $0 godemo"
    echo "      那么初始化后的项目为: $(dirname "$PWD")/godemo"
    exit
fi
#######################################################################################################################

export LANGUAGE="utf-8"

project="$1"

CATALOG="$(dirname "$PWD")/$project"
mkdir -p "${CATALOG}"
echo  "\033[31;43m""\033[1m项目目录: \033[0m"   "\033[32;40m""\033[1m${CATALOG} \033[0m"

rsync -rv --exclude=.git --exclude=.idea --exclude=cg.sh --exclude=log ./ "${CATALOG}"

cd "${CATALOG}"
rm -rf goweb-docker-cg-cg

os=$(uname)
echo "os: $os"

if [[ "$os" == "Darwin" ]]; then
  grep -rl 'goweb-docker-cg-cg' ./  | xargs sed -i "" "s/goweb-docker-cg-cg/$project/g"
elif [[ "$(expr substr $os 1 5)" == "Linux" ]];then
  sed -i "s/goweb-docker-cg-cg/$project/g" `grep -rl 'goweb-docker-cg-cg' ./`
elif [[ "$(expr substr $os 1 10)" == "MINGW32_NT" ]];then
  sed -i "s/goweb-docker-cg-cg/$project/g" `grep -rl 'goweb-docker-cg-cg' ./`
fi


echo  "\033[30;40m""\033[1m#############################################################################\033[0m"
echo ""
echo  "\033[30;40m""\033[1m成功生成项目，请移步 ${CATALOG} 查看 😊\033[0m"
echo ""
echo  "\033[30;40m""\033[1m项目开始前，建议阅读\033[0m"  "\033[31;43m""\033[1m【goweb-docker-cg 项目引导】\033[0m" "\033[30;40m""\033[1m ，以便更快上手项目：\033[0m"
echo ""
echo  "\033[30;40m""\033[1m#############################################################################\033[0m"
