FROM ubuntu:14.04
MAINTAINER haibzhou <haibzhou@ebay.com>

ADD downward /usr/bin/downward
ENTRYPOINT ["downward"]
