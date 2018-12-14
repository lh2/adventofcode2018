#!/bin/sh
set -e
if [ ! -e session.txt ]; then
	echo 1>&2 Please create a \"session.txt\" file containing your advent of code HTTP cookie
	exit 1
fi
if [ "x$1" != "x" ]; then
	day=$1
else
	day=`date +%-d`
fi
if [ ! -d input ]; then
	mkdir input
fi

whydoihavetodothis() {
	case "$day" in
	11 | 14)
		sed 's/.*puzzle-input"\>\s*(\d+)\s*/\1/'
		;;
	*)
		cat
		;;
	esac
}

curl -b "session=`cat session.txt`" https://adventofcode.com/2018/day/$day/input | whydoihavetodothis > input/$day.txt
