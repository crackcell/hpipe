Name: hpipe
Version: 1.0
Release: 1
Summary: A workflow engine.

Group: Development/Libraries
License: GPL
URL: https://github.com/crackcell/hpipe
Packager: Menglong TAN <tanmenglong@gmail.com>
Source: https://github.com/crackcell/hpipe/archive/master.zip
BuildRoot: %{_tmppath}/%{name}-%{version}-%{release}-root-%(%{__id_u} -n)

#BuildRequires: golang sqlite-devel
#Requires: sqlite

%define bin_dir %{_prefix}/share/hpipe/bin

%description
Hpipe is a workflow engine supporting hybrid workflows with built-in support for Hadoop Streaming and Hive.

%prep

%build
pushd $OLDPWD/../..
make
popd

%install
rm -rf $RPM_BUILD_ROOT
mkdir -p ${RPM_BUILD_ROOT}%{bin_dir}
pushd $OLDPWD/../..
cp output/bin/* ${RPM_BUILD_ROOT}%{bin_dir}
popd

%clean
rm -rf $RPM_BUILD_ROOT

%files
%{bin_dir}
