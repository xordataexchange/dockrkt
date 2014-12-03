DockRkt
=======

Because Orange is the new Black.

Usage
=====

curl -O http://dockrkt.com/{publicrepository}/{username}/{image}:{tag}

DockRkt will return an ACI file that represents the docker container converted to Rocket format.

Advanced Usage
==============
$ cd myproject
$ docker build -t dockrkt.com/docker/joecool/etcd:latest .
$ docker push dockrkt.com/docker/joecool/etcd:latest
(after converston redirects to)
http://dockrkt.com/joecool-etcd-latest.aci

Conventions
===========

Repository is required in the URL.  For hub.docker.com use "docker".  So to convert bketelsen/nsqd you would
request http://dockrkt.com/docker/bketelsen/nsqd:latest

Tags are optional, latest is assumed if none given.

We aren't in the business of authentication, so won't be supporting private repositories.  

Contributing
============

Fork it, rock it, PR it.  Include tests.


Hipsters, you've found your home.

