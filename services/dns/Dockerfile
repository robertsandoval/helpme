FROM centos:centos8
RUN dnf -y install bind bind-utils
ADD named.conf /etc/named.conf
ADD zonefile.db /var/named/zonefile.db
ADD reverse.db /var/named/reverse.db
RUN chown named -R /var/named
EXPOSE 53/tcp 53/udp
CMD ["/usr/sbin/named", "-u", "named", "-g"]
