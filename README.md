# container_from_scratch
Tested container from Liz Rice on Ubuntu 20.04 LTS, release 20.04

# Follow tutorial from: https://www.youtube.com/watch?v=8fi7uSYlOdc
# fork on github from: https://github.com/lizrice/containers-from-scratch

# Ran on ubuntu:

Distributor ID:	Ubuntu
Description:	Ubuntu 20.04 LTS
Release:	20.04
Codename:	focal

# On container:
root@ubuntu:/home/samira/cont# go build # main.go
root@ubuntu:/home/samira/cont# ./cont run /bin/bash
RunningP1 [/bin/bash] as 27424
RunningP2 [/bin/bash] as 1

# Install these packages if not already present:
$sudo
$snap install docker
$apt install docker.io
$apt install go
$apt install golang-go

# From container: successful run as root: 
root@ubuntu:/home/samira/cont# ./cont run /bin/bash
RunningP1 [/bin/bash] as 18431
RunningP2 [/bin/bash] as 1
root@CONTAINER:/# 

# Test max pids set from host running a fork bomb: 
root@CONTAINER:/# :(){ : | : & }; :                       
[1] 17
root@CONTAINER:/# bash: fork: retry: No child processes
bash: fork: retry: No child processes
bash: fork: retry: No child processes
bash: fork: retry: No child processes
bash: fork: retry: No child processes
bash: fork: retry: No child processes
bash: fork: retry: No child processes
bash: fork: retry: Resource temporarily unavailable
bash: fork: retry: No child processes
bash: fork: retry: No child processes
bash: fork: retry: No child processes
bash: fork: retry: No child processes
bash: fork: retry: No child processes
bash: fork: retry: No child processes
bash: fork: retry: No child processes
bash: fork: retry: No child processes
bash: fork: retry: No child processes
bash: fork: retry: Resource temporarily unavailable
bash: fork: retry: No child processes
bash: fork: retry: No child processes
bash: fork: retry: No child processes
bash: fork: retry: No child processes
bash: fork: retry: No child processes
bash: fork: retry: No child processes
bash: fork: retry: No child processes
bash: fork: retry: No child processes
bash: fork: retry: Resource temporarily unavailable
bash: fork: retry: No child processes
bash: fork: retry: No child processes
bash: fork: retry: No child processes
bash: fork: retry: No child processes
bash: fork: retry: No child processes
bash: fork: retry: No child processes
bash: fork: retry: No child processes
bash: fork: retry: No child processes
bash: fork: retry: Resource temporarily unavailable
bash: fork: retry: No child processes
bash: fork: retry: No child processes
bash: fork: retry: No child processes
bash: fork: retry: No child processes
bash: fork: Resource temporarily unavailable
bash: fork: Resource temporarily unavailable
^C
[1]+  Done                    : | :
root@CONTAINER:/# 

# On host, $ps -fax will show no more than 20 defunct processes :)
samira@ubuntu:/sys/fs/cgroup/pids$ cat /sys/fs/cgroup/pids/samira/pids.max
20
samira@ubuntu:/sys/fs/cgroup/pids$ cat /sys/fs/cgroup/pids/samira/notify_on_release
1
samira@ubuntu:/sys/fs/cgroup/pids$ cat /sys/fs/cgroup/pids/samira/cgroup.procs
18734
18738
samira@ubuntu:/sys/fs/cgroup/pids$ 
samira@ubuntu:/sys/fs/cgroup/pids/samira$ cat pids.current
samira@ubuntu:/sys/fs/cgroup/pids/samira$ cat tasks
18734
18735
18736
18737
18738
samira@ubuntu:/sys/fs/cgroup/pids/samira$ 

samira@ubuntu:/sys/fs/cgroup/pids/samira$ ps -fax 
10650 pts/0    Ss     0:00  |   \_ bash
  16018 pts/0    S      0:00  |   |   \_ su
  16019 pts/0    S      0:00  |   |       \_ bash
  27424 pts/0    Sl     0:00  |   |           \_ ./cont run /bin/bash
  27428 pts/0    Sl     0:00  |   |               \_ /proc/self/exe child /bin/bash
  27433 pts/0    S+     0:00  |   |                   \_ /bin/bash
  27463 pts/0    Z      0:00  |   |                   \_ [bash] <defunct>
  27464 pts/0    Z      0:00  |   |                   \_ [bash] <defunct>
  27466 pts/0    Z      0:00  |   |                   \_ [bash] <defunct>
  27467 pts/0    Z      0:00  |   |                   \_ [bash] <defunct>
  27468 pts/0    S      0:00  |   |                   \_ /bin/bash
  27469 pts/0    S      0:00  |   |                   \_ /bin/bash
  27470 pts/0    S      0:00  |   |                   \_ /bin/bash
  27471 pts/0    S      0:00  |   |                   \_ /bin/bash
  27472 pts/0    S      0:00  |   |                   \_ /bin/bash
  27473 pts/0    S      0:00  |   |                   \_ /bin/bash
  27474 pts/0    S      0:00  |   |                   \_ /bin/bash
  27475 pts/0    S      0:00  |   |                   \_ /bin/bash
  27482 pts/0    S      0:00  |   |                   |   \_ /bin/bash
  27476 pts/0    S      0:00  |   |                   \_ /bin/bash
  11046 pts/2    Ss     0:00  |   \_ bash
  18819 pts/2    S      0:00  |   |   \_ su
  18821 pts/2    S+     0:00  |   |       \_ bash
  11378 pts/1    Ss+    0:00  |   \_ bash
  24118 pts/3    Ss     0:00  |   \_ bash
  27520 pts/3    R+     0:00  |       \_ ps -fax


# after we run again the container: 
# can still see info from container process using sleep: 
# from container
root@CONTAINER:/# sleep 100
# from host:
samira@ubuntu:/sys/fs/cgroup/pids/samira$ ps -C sleep
    PID TTY          TIME CMD
  27758 pts/0    00:00:00 sleep





