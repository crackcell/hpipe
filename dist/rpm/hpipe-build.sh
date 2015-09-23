#!/bin/bash
##! @description: usage
##! @version: 1
##! @author: Menglong TAN <tanmenglong@gmail.com>
##! @date:   Wed Sep 23 21:49:37 2015

rpmbuild --buildroot $PWD/rpmbuild -bb ./hpipe.spec
