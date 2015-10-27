#!/usr/bin/perl
##! @description: usage
##! @version: 1
##! @author: crackcell <tanmenglong@gmail.com>
##! @date:   Mon Oct 26 18:44:40 2015

use strict;
use JSON;

#--------------- global variable --------------


#------------------ function ------------------


#------------------- main -------------------

print "args: ", $ARGV[0], " ", $ARGV[1], "\n";
my $obj = from_json($ARGV[1]);
foreach my $key (keys %{$obj}) {
	print $key, ": ", $obj->{$key}, "\n";
}
