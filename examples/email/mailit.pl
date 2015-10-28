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

my $mail_recv = 'tanmenglong@gmail.com';

my @finished = ();
my @failed = ();

my $obj = from_json($ARGV[1]);
foreach my $key (keys %{$obj}) {
	if ($obj->{$key} eq "failed") {
		push(@failed, $key)
	} elsif ($obj->{$key} eq "finished") {
		push(@finished, $key)
	}
}

my $mail_body = "";

$mail_body = $mail_body . "finished: ";
foreach my $job (@finished) {
	$mail_body = $mail_body . $job . ", ";
}
$mail_body = $mail_body . '\\n';

$mail_body = $mail_body . "failed: ";
foreach my $job (@failed) {
	$mail_body = $mail_body . $job . ", ";
}
$mail_body = $mail_body . '\\n';

print "failed jobs: ", $mail_body, "\n";
print "echo `$mail_body` | mail -s 'Task Report' $mail_recv\n";
#system("echo `$mail_body` | mail -s 'Task Report' $mail_recv");
