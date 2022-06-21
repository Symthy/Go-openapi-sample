#!/bin/bash
VSCODE_SETTINGS=".vscode/settings.json"
TAGS_WIN="\"-tags=windows\""
TAGS_LINUX="\"-tags=linux\""
REGEX_WIN=".*${TAGS_WIN}.*"
REGEX_LINUX=".*${TAGS_LINUX}.*"

build_flag_line=`grep "build.buildFlags" ${VSCODE_SETTINGS}`

if [[ ${build_flag_line} =~ ${REGEX_WIN} ]]; then
    echo "chage windows to linux"
    sed -i -e "s/${TAGS_WIN}/${TAGS_LINUX}/g" ${VSCODE_SETTINGS}
elif [[ ${build_flag_line} =~ ${REGEX_LINUX} ]]; then
    echo "change linux to windows"
    sed -i -e "s/${TAGS_LINUX}/${TAGS_WIN}/g" ${VSCODE_SETTINGS}
else
    echo "no change"
fi

cat ${VSCODE_SETTINGS}