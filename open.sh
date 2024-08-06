#!/bin/sh
#### Julio S. Zaima ( zaima@br.ibm.com )
###  release: 2019/08/21
#### usage: sudo sh sysinfo.sh
clear
echo .
echo ..OS
head -n1 /etc/issue
# OS release REDHAT
if [ -f /etc/redhat-release ]; then
        cat /etc/redhat-release
fi
# OS release SUSE
if [ -f /etc/os-release ]; then
        grep PRETTY /etc/os-release
fi
echo .
echo ..HOST
export HNAME=`hostname -f`
export HADDR=`ifconfig -a | grep -i inet | grep -i cast | awk '{print "("$1": " $2" " $3": " $4 ")"}'`
echo " Server: "$HNAME" "$HADDR
hostnamectl | grep -v ID | grep -v Icon | grep -v Chassis | grep -v CPE | grep -v Static
echo .
echo ..MEMORIA
grep MemTotal /proc/meminfo
echo .
echo ..CPU
echo "TOTAL CPU :" `grep "model name" /proc/cpuinfo | wc -l` " Core"
echo `grep  "# processors" /proc/cpuinfo` " Core"
grep -m1 "model name" /proc/cpuinfo
grep "vendor_id" /proc/cpuinfo
echo .
echo ..DISKS
# cat /proc/partitions
# fdisk -l |grep '^Disk /dev/' |grep -v '/dev/mapper'
# echo
echo Show all partitions registered on o system
lsblk
# echo ..# df -h
echo .
df -h
# hdparm -i /dev/sda
echo .
echo ..NETWORKs
ip a | grep inet
echo ...
netstat -ar
echo .
echo ..NETPORTS
netstat -tnlp
echo .
echo : /etc/resolv.conf // DNS configuration
cat /etc/resolv.conf | grep -v \#
# echo .
# echo : /etc/hosts // File-based network name resolution
# cat /etc/hosts
echo .
echo ..TIMESYNC
echo .Checking Time Synchronization
timedatectl
echo .
echo ..SELINUX
if [ -f /etc/selinux/config ]; then
        cat /etc/selinux/config
fi
cat /etc/selinux/config
echo .
echo ..USERLIMITS !
cat /etc/security/limits.conf | grep -v \#
echo .
ulimit -a
echo .
# echo .IWS USERS
# grep -i iws /etc/passwd

echo .
echo DONE !!
## fim do arquivo







sudo ./sysinfo.sh  > $(hostname -f)_info.out

# mkdir -p /opt/CP4D/binarios/openshift
# mkdir -p /opt/CP4D/config/prd 
# mkdir -p /opt/CP4D/config/bin










[3J[H[2J.
..OS
\S
Red Hat Enterprise Linux release 8.10 (Ootpa)
PRETTY_NAME="Oracle Linux Server 8.10"
.
..HOST
 Server: melponeme-01.fazenda.mg.gov.br (inet: 172.23.218.25 netmask: 255.255.255.0) (inet: 192.168.122.1 netmask: 255.255.255.0)
    Virtualization: vmware
  Operating System: Oracle Linux Server 8.10
            Kernel: Linux 5.15.0-206.153.7.1.el8uek.x86_64
      Architecture: x86-64
.
..MEMORIA
MemTotal:        7837560 kB
.
..CPU
TOTAL CPU : 4  Core
 Core
model name	: Intel(R) Xeon(R) Gold 6238R CPU @ 2.20GHz
vendor_id	: GenuineIntel
vendor_id	: GenuineIntel
vendor_id	: GenuineIntel
vendor_id	: GenuineIntel
.
..DISKS
Show all partitions registered on o system
NAME                      MAJ:MIN RM  SIZE RO TYPE MOUNTPOINT
sda                         8:0    0   80G  0 disk 
├─sda1                      8:1    0   20G  0 part /
└─sda2                      8:2    0   59G  0 part 
  ├─ol_melponeme--01-opt  252:0    0   44G  0 lvm  /opt
  ├─ol_melponeme--01-home 252:1    0    5G  0 lvm  /home
  └─ol_melponeme--01-var  252:2    0   10G  0 lvm  /var
sdb                         8:16   0   80G  0 disk 
sr0                        11:0    1 53,6M  0 rom  
.
Sist. Arq.                         Tam. Usado Disp. Uso% Montado em
devtmpfs                           3,8G     0  3,8G   0% /dev
tmpfs                              3,8G     0  3,8G   0% /dev/shm
tmpfs                              3,8G  9,4M  3,8G   1% /run
tmpfs                              3,8G     0  3,8G   0% /sys/fs/cgroup
/dev/sda1                           20G  6,2G   14G  31% /
/dev/mapper/ol_melponeme--01-opt    44G  347M   44G   1% /opt
/dev/mapper/ol_melponeme--01-var    10G  3,8G  6,3G  38% /var
/dev/mapper/ol_melponeme--01-home  5,0G  409M  4,6G   9% /home
tmpfs                              766M   56K  766M   1% /run/user/208
.
..NETWORKs
    inet 127.0.0.1/8 scope host lo
    inet6 ::1/128 scope host 
    inet 172.23.218.25/24 brd 172.23.218.255 scope global noprefixroute ens192
    inet6 fe80::250:56ff:feab:f535/64 scope link noprefixroute 
    inet 192.168.122.1/24 brd 192.168.122.255 scope global virbr0
...
Tabela de Roteamento IP do Kernel
Destino         Roteador        MáscaraGen.    Opções   MSS Janela  irtt Iface
default         _gateway        0.0.0.0         UG        0 0          0 ens192
172.23.218.0    0.0.0.0         255.255.255.0   U         0 0          0 ens192
192.168.122.0   0.0.0.0         255.255.255.0   U         0 0          0 virbr0
.
..NETPORTS
Conexões Internet Ativas (sem os servidores)
Proto Recv-Q Send-Q Endereço Local          Endereço Remoto         Estado      PID/Program name    
tcp        0      0 127.0.0.1:6010          0.0.0.0:*               OUÇA       7717/sshd: geisler@ 
tcp        0      0 127.0.0.1:6011          0.0.0.0:*               OUÇA       8124/sshd: geisler@ 
tcp        0      0 127.0.0.1:631           0.0.0.0:*               OUÇA       950/cupsd           
tcp        0      0 192.168.122.1:53        0.0.0.0:*               OUÇA       1870/dnsmasq        
tcp        0      0 0.0.0.0:22              0.0.0.0:*               OUÇA       952/sshd            
tcp        0      0 0.0.0.0:111             0.0.0.0:*               OUÇA       1/systemd           
tcp6       0      0 :::9100                 :::*                    OUÇA       3763/node_exporter  
tcp6       0      0 :::22                   :::*                    OUÇA       952/sshd            
tcp6       0      0 :::111                  :::*                    OUÇA       1/systemd           
tcp6       0      0 ::1:631                 :::*                    OUÇA       950/cupsd           
tcp6       0      0 ::1:6011                :::*                    OUÇA       8124/sshd: geisler@ 
tcp6       0      0 ::1:6010                :::*                    OUÇA       7717/sshd: geisler@ 
.
: /etc/resolv.conf // DNS configuration
search fazenda.mg.gov.br fazenda.mg
nameserver 172.23.212.226
nameserver 172.23.212.216
.
..TIMESYNC
.Checking Time Synchronization
               Local time: ter 2024-07-30 13:10:18 EDT
           Universal time: ter 2024-07-30 17:10:18 UTC
                 RTC time: ter 2024-07-30 17:10:18
                Time zone: America/New_York (EDT, -0400)
System clock synchronized: no
              NTP service: active
          RTC in local TZ: no
.
..SELINUX

# This file controls the state of SELinux on the system.
# SELINUX= can take one of these three values:
#     enforcing - SELinux security policy is enforced.
#     permissive - SELinux prints warnings instead of enforcing.
#     disabled - No SELinux policy is loaded.
SELINUX=disabled
# SELINUXTYPE= can take one of these three values:
#     targeted - Targeted processes are protected,
#     minimum - Modification of targeted policy. Only selected processes are protected. 
#     mls - Multi Level Security protection.
SELINUXTYPE=targeted



# This file controls the state of SELinux on the system.
# SELINUX= can take one of these three values:
#     enforcing - SELinux security policy is enforced.
#     permissive - SELinux prints warnings instead of enforcing.
#     disabled - No SELinux policy is loaded.
SELINUX=disabled
# SELINUXTYPE= can take one of these three values:
#     targeted - Targeted processes are protected,
#     minimum - Modification of targeted policy. Only selected processes are protected. 
#     mls - Multi Level Security protection.
SELINUXTYPE=targeted


.
..USERLIMITS !


.
core file size          (blocks, -c) 0
data seg size           (kbytes, -d) unlimited
scheduling priority             (-e) 0
file size               (blocks, -f) unlimited
pending signals                 (-i) 30378
max locked memory       (kbytes, -l) 64
max memory size         (kbytes, -m) unlimited
open files                      (-n) 1024
pipe size            (512 bytes, -p) 8
POSIX message queues     (bytes, -q) 819200
real-time priority              (-r) 0
stack size              (kbytes, -s) 8192
cpu time               (seconds, -t) unlimited
max user processes              (-u) 30378
virtual memory          (kbytes, -v) unlimited
file locks                      (-x) unlimited
.
.
DONE !!




sh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDMkmGU7+tpnJzs+ui+jWjm3898AiXNabR/OOx02vZVeOIQA5Wc/zPX+zKP5BhtWzIs8PTu14tSvCwtH3XAD42FiGHzWzy54FNIp5JEZrbZoAZxGGMviOXlSZTF/mOOIYDV/jXau7PMyXsvPcBS8Np1
+trB/NuhAS0j0aqTRwBQ4qWLTnUFxvTf0zWmoQHjbp9jePNWthGZntIiwgJW+rZZ8TjQSZNfxyGsht42AxoSDWmgo6FwoYl8ULvakVYtqRWmtzb6xL+2byBsq1GB07hFfDlLlHgo//Hr8uj9ASCqc33tNLSvBNtxtZTDzdECOG5HU28yqIdJ7qNkpbkI
MR2WbG4MyFgL2IsiIXTKY8bGi1zqKQd7Q5iMqxPGl5Ax2Bx4/3PqJfkdT7Hkwq/gZ5cAHPkt3WYIZ5pxfIjllPeEdw387Gu5gmQIw+qtYP5srXv/s7/cB6kN/BGkPxLjguvwKTng19aw36TfE9yB6mXoT/ahfvsK1CXIvOVE6S+k5U1FescZ/+U8i0bS
Ped9Kz6C+kYr5dwS4cCbZ0WSh3CRhABnCDLzVBtV/l5NMYrU95ADxn7PP6VhtF931qPIu2bHatyZWpL0ALKO19YIzHIbrljYSHJde1wl1uWWmVBOpklfKr7AAKS/euxngLZ21zPW6rFUaiKtz3DVhRpLOTGQTQ== k8s@melponeme-01.fazenda.mg
.gov.br