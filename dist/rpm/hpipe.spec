Name: hpipe
Version: 1.0
Release: 1
Summary: A workflow engine.

Group: Development/Libraries
Licence: GPL
URL: https://github.com/crackcell/hpipe
Packager: Menglong TAN
Source0: https://github.com/crackcell/hpipe/archive/master.zip
BuildRoot: %{_tmppath}/%{name}-%{version}-%{release}-root-%(%{__id_u} -n)

BuildRequires:
Requires:

%description
Hpipe is a workflow engine supporting hybrid workflows with built-in support for Hadoop Streaming and Hive.

%prep

%build
make

%install
rm -rf $RPM_BUILD_ROOT
make install DESTDIR=$RPM_BUILD_ROOT

%clean
rm -rf $RPM_BUILD_ROOT

%files
