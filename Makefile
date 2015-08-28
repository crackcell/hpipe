# Author:  Tan Menglong <tanmenglong@gmail.com>
# Date:    Mon Dec  8 12:25:53 2014
#
# Make Target:
# ------------
# The Makefile provides the following targets to make:
#   $ make           compile and link
#   $ make clean     clean objects and the executable file
#
#===========================================================================

.PHONY : all deps output clean help hpipe-run

all : output

deps :
	go get github.com/colinmarc/hdfs

output : hpipe-run
	mkdir -p output/bin
	mv hpipe-run/hpipe-run output/bin/
#	cp scripts/* output/bin/

hpipe-run : deps
	cd hpipe-run; go build

clean :
	rm -rf output

help :
	@echo 'Usage: make [TARGET]'
	@echo 'TARGETS:'
	@echo '  all       (=make) compile and link.'
	@echo '  clean     clean objects and the executable file.'
	@echo '  help      print this message.'
	@echo
