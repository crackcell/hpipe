#!/usr/bin/perl
##! @description: usage
##! @version: 1
##! @author: crackcell <tanmenglong@gmail.com>
##! @date:   Fri Jan 31 13:58:58 2014

use strict;

#--------------- global variable --------------


#------------------ function ------------------


#------------------- main -------------------

while (<STDIN>) {
    chomp;
    print $_, "\t1\n" if length($_) > 0
}
