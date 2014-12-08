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

.PHONY : all output clean help

all : output

output :
	mkdir -p output/bin
	cd main; go build
	mv main/main output/bin/hpipe
	cp scripts/* output/bin/

clean :
	rm -rf output

help :
	@echo 'Usage: make [TARGET]'
	@echo 'TARGETS:'
	@echo '  all       (=make) compile and link.'
	@echo '  clean     clean objects and the executable file.'
	@echo '  help      print this message.'
	@echo
