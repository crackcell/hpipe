#!/usr/bin/perl
##! @description: usage
##! @version: 1
##! @author: crackcell <tanmenglong@gmail.com>
##! @date:   Fri Jan 31 14:16:57 2014

use strict;

#--------------- global variable --------------


#------------------ function ------------------


#------------------- main -------------------

my %wordcount = ();
while (<STDIN>) {
    chomp;
    my @tokens = split(/\t/, $_);
    my $word = $tokens[0];
    my $count = $tokens[1];
    if (! exists $wordcount{$word}) {
        $wordcount{$word} = 1;
    } else {
        $wordcount{$word} = $wordcount{$word} + 1
    }
}

for my $word (sort keys %wordcount) {
    print $word, "\t", $wordcount{$word}, "\n"
}
