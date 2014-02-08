#!/bin/bash
##! @description: usage
##! @version: 1
##! @author: crackcell <tanmenglong@gmail.com>
##! @date:   Tue Jan 28 11:23:34 2014

#------------------- library ------------------



#--------------- global variable --------------

typeset -r VERSION="1.0"

#------------------ function ------------------


#------------------- main -------------------

set -eu
workroot=$(pushd $(dirname $0)/.. >/dev/null && pwd -P && popd >/dev/null)
source $workroot/conf/hpipe.env

function usage() {
    echo "Usage: $0 [-t today] [-h] FLOW" >&2
    exit 1
}

export today=$(date +%Y%m%d)
export hpipe_workroot=$workroot

properties=()
while getopts "t:d:h" opt
do
    case $opt in
        t) today=$OPTARG;;
        d) workroot=$OPTARG;;
        h) usage 0;;
        \?) echo "Invalid option: -$OPTARG" >&2; usage 1;;
    esac
done
shift $((OPTIND-1))

[[ $# == 1 ]] || usage 1
flow=$1

PYTHONPATH=$workroot \
    python $workroot/bin/hpipe_run.py $flow
