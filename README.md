# container_from_scratch
Tested container from Liz Rice on Ubuntu 20.04 LTS, release 20.04

# Follow tutorial from: 
https://www.youtube.com/watch?v=8fi7uSYlOdc <br/>
fork on github from: https://github.com/lizrice/containers-from-scratch <br/>
you can easily edit main.go code on Ubuntu Text Editor <br/>

# Golang references for palackges and library function parameters: 
https://pkg.go.dev/syscall <br/>
https://man7.org/linux/man-pages/man2/clone.2.html <br/>


# Concepts: 
kernel space<br />
userspace<br /> 
processes<br /> 
system calls<br /> 
operating system internals and resources<br />
linux<br />

# Ran on ubuntu:
Distributor ID:	Ubuntu<br />
Description:	Ubuntu 20.04 LTS<br />
Release:	20.04<br />
Codename:	focal<br />

# Install these packages if not already present:
$sudo<br />
$snap install docker<br />
$apt install docker.io<br />
$apt install go<br />
$apt install golang-go<br />

# From container: successful run as root: 
root@ubuntu:/home/samira/cont# go build # main.go<br />
root@ubuntu:/home/samira/cont# ./cont run /bin/bash<br />
RunningP1 [/bin/bash] as 18431<br />
RunningP2 [/bin/bash] as 1<br />
root@CONTAINER:/# <br />

# Test max pids set from host running a fork bomb: 
root@CONTAINER:/# :(){ : | : & }; :   <br />                    
[1] 17<br />
root@CONTAINER:/# bash: fork: retry: No child processes<br />
bash: fork: retry: No child processes<br />
bash: fork: retry: No child processes<br />
bash: fork: retry: No child processes<br />
bash: fork: retry: No child processes<br />
bash: fork: retry: No child processes<br />
bash: fork: retry: No child processes<br />
bash: fork: retry: Resource temporarily unavailable<br />
bash: fork: retry: No child processes<br />
bash: fork: retry: No child processes<br />
bash: fork: retry: No child processes<br />
bash: fork: retry: No child processes<br />
bash: fork: retry: No child processes<br />
bash: fork: retry: No child processes<br />
bash: fork: retry: No child processes<br />
bash: fork: retry: No child processes<br />
bash: fork: retry: No child processes<br />
bash: fork: retry: Resource temporarily unavailable<br />
bash: fork: retry: No child processes<br />
bash: fork: retry: No child processes<br />
bash: fork: retry: No child processes<br />
bash: fork: retry: No child processes<br />
bash: fork: retry: No child processes<br />
bash: fork: retry: No child processes<br />
bash: fork: retry: No child processes<br />
bash: fork: retry: No child processes<br />
bash: fork: retry: Resource temporarily unavailable<br />
bash: fork: retry: No child processes<br />
bash: fork: retry: No child processes<br />
bash: fork: retry: No child processes<br />
bash: fork: retry: No child processes<br />
bash: fork: retry: No child processes<br />
bash: fork: retry: No child processes<br />
bash: fork: retry: No child processes<br />
bash: fork: retry: No child processes<br />
bash: fork: retry: Resource temporarily unavailable<br />
bash: fork: retry: No child processes<br />
bash: fork: retry: No child processes<br />
bash: fork: retry: No child processes<br />
bash: fork: retry: No child processes<br />
bash: fork: Resource temporarily unavailable<br />
bash: fork: Resource temporarily unavailable<br />
^C<br />
[1]+  Done                    : | :<br />
root@CONTAINER:/# <br />

# On host, $ps -fax will show no more than 20 defunct processes :)<br />
samira@ubuntu:/sys/fs/cgroup/pids$ cat /sys/fs/cgroup/pids/samira/pids.max<br />
20<br />
samira@ubuntu:/sys/fs/cgroup/pids$ cat /sys/fs/cgroup/pids/samira/notify_on_release<br />
1<br />
samira@ubuntu:/sys/fs/cgroup/pids$ cat /sys/fs/cgroup/pids/samira/cgroup.procs<br />
18734<br />
18738<br />
samira@ubuntu:/sys/fs/cgroup/pids$ <br />
samira@ubuntu:/sys/fs/cgroup/pids/samira$ cat pids.current<br />
samira@ubuntu:/sys/fs/cgroup/pids/samira$ cat tasks<br />
18734<br />
18735<br />
18736<br />
18737<br />
18738<br />
samira@ubuntu:/sys/fs/cgroup/pids/samira$ <br />

samira@ubuntu:/sys/fs/cgroup/pids/samira$ ps -fax <br />
10650 pts/0    Ss     0:00  |   \_ bash<br />
  16018 pts/0    S      0:00  |   |   \_ su<br />
  16019 pts/0    S      0:00  |   |       \_ bash<br />
  27424 pts/0    Sl     0:00  |   |           \_ ./cont run /bin/bash<br />
  27428 pts/0    Sl     0:00  |   |               \_ /proc/self/exe child /bin/bash<br />
  27433 pts/0    S+     0:00  |   |                   \_ /bin/bash<br />
  27463 pts/0    Z      0:00  |   |                   \_ [bash] <defunct><br />
  27464 pts/0    Z      0:00  |   |                   \_ [bash] <defunct><br />
  27466 pts/0    Z      0:00  |   |                   \_ [bash] <defunct><br />
  27467 pts/0    Z      0:00  |   |                   \_ [bash] <defunct><br />
  27468 pts/0    S      0:00  |   |                   \_ /bin/bash<br />
  27469 pts/0    S      0:00  |   |                   \_ /bin/bash<br />
  27470 pts/0    S      0:00  |   |                   \_ /bin/bash<br />
  27471 pts/0    S      0:00  |   |                   \_ /bin/bash<br />
  27472 pts/0    S      0:00  |   |                   \_ /bin/bash<br />
  27473 pts/0    S      0:00  |   |                   \_ /bin/bash<br />
  27474 pts/0    S      0:00  |   |                   \_ /bin/bash<br />
  27475 pts/0    S      0:00  |   |                   \_ /bin/bash<br />
  27482 pts/0    S      0:00  |   |                   |   \_ /bin/bash<br />
  27476 pts/0    S      0:00  |   |                   \_ /bin/bash<br />
  11046 pts/2    Ss     0:00  |   \_ bash<br />
  18819 pts/2    S      0:00  |   |   \_ su<br />
  18821 pts/2    S+     0:00  |   |       \_ bash<br />
  11378 pts/1    Ss+    0:00  |   \_ bash<br />
  24118 pts/3    Ss     0:00  |   \_ bash<br />
  27520 pts/3    R+     0:00  |       \_ ps -fax<br />


# after we run again the container: <br />
can still see info from container process using sleep: <br />
root@CONTAINER:/# sleep 100<br />
# from host:<br />
samira@ubuntu:/sys/fs/cgroup/pids/samira$ ps -C sleep<br />
    PID TTY          TIME CMD<br />
  27758 pts/0    00:00:00 sleep<br /><br />

Next, you can play with docker and kubernestes:<br />
On docker, you can follow this tutorial and exercise running priviledged, setting max pids, etc:<br />
ex: docker container run --pids-limit 20 <image><br />
I recommend:<br />
  1. go through entire tutorial of installing on desktop, then building image, running app, persistance of data across app updates etc.<br />
  https://www.docker.com/101-tutorial<br />
  2. https://kubernetes.io/docs/tutorials/kubernetes-basics/<br /><br />


