Name:           opendocking
Version:        2.0.0
Release:        0
License:        Zlib
Summary:        A lightweight docker management UI
Url:            https://opendocking.io
Group:          BLAH
Source0:        https://github.com/opendocking/opendocking/releases/download/%{version}/opendocking-%{version}-linux-amd64.tar.gz
Source1:        opendocking.service
BuildRoot:      %{_tmppath}/%{name}-%{version}-build
%if 0%{?suse_version}
BuildRequires:  help2man
%endif
Requires:       docker
%{?systemd_requires}
BuildRequires: systemd

## HowTo build ## 
# You can use spectool to fetch sources
# spectool -g -R distribution/opendocking.spec 
# Then build with 'rpmbuild -ba distribution/opendocking.spec' 


%description
OpenDocking is a lightweight management UI which allows you to easily manage
your different container environments (Docker hosts, Kubernetes, or Swarm clusters).
OpenDocking is meant to be as simple to deploy as it is to use.
It consists of a single container that can run on any container engine
(can be deployed as Linux container or a Windows native container).
OpenDocking allows you to manage your containers, images, volumes,
networks and more! It is compatible with the standalone Docker engine, Kubernetes, and Docker Swarm mode.

%prep
%setup -qn opendocking

%build
%if 0%{?suse_version}
help2man -N --no-discard-stderr ./opendocking  > opendocking.1
%endif

%install
# Create directory structure
install -D -m 0755 opendocking %{buildroot}%{_sbindir}/opendocking
install -d -m 0755 %{buildroot}%{_datadir}/opendocking/public
install -d -m 0755 %{buildroot}%{_localstatedir}/lib/opendocking
install -D -m 0644 %{S:1} %{buildroot}%{_unitdir}/opendocking.service
%if 0%{?suse_version}
install -D -m 0644 opendocking.1 %{buildroot}%{_mandir}/man1/opendocking.1
( cd  %{buildroot}%{_sbindir} ; ln -s service rcopendocking )
%endif
# populate
# don't install docker binary with package use system wide installed one
cp -ra public/ %{buildroot}%{_datadir}/opendocking/

%pre
%if 0%{?suse_version}
%service_add_pre opendocking.service
#%%else # this does not work on rhel 7?
#%%systemd_pre opendocking.service
true
%endif

%preun
%if 0%{?suse_version}
%service_del_preun opendocking.service
%else
%systemd_preun opendocking.service
%endif

%post
%if 0%{?suse_version}
%service_add_post opendocking.service
%else
%systemd_post opendocking.service
%endif

%postun
%if 0%{?suse_version}
%service_del_postun opendocking.service
%else
%systemd_postun_with_restart opendocking.service
%endif


%files
%defattr(-,root,root)
%{_sbindir}/opendocking
%{_datadir}/opendocking/public
%dir %{_localstatedir}/lib/opendocking/
%{_unitdir}/opendocking.service
%if 0%{?suse_version}
%{_mandir}/man1/opendocking.1*
%{_sbindir}/rcopendocking
%endif
