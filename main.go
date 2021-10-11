package main

//what is a container really? 
//nice tutorial from Liz Rice: build a container from scratch in Go.
//docker      run image <cmd> <params>
//go run main.go run    <cmd> <params>

//container is created from namespace and limited by control groups
//image/fs works by pointing container to subset of filesystem on host.

import (
    "fmt"
    "os"
    "os/exec"
    "syscall"
    "strconv" 
    "path/filepath"
    "io/ioutil"
)

//control groups let us limit resources to container (mem, cpu, i/o, process #s, etc.)
func main(){

    switch os.Args[1]{
    case "run": 
        run()
    case "child":
        child()
    default:
	panic("bad command")	
    }
}

func run(){

    fmt.Printf("RunningP1 %v as %d\n", os.Args[2:], os.Getpid())
    cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...) //run itself
    cmd.Stdin = os.Stdin
    cmd.Stderr = os.Stderr
    cmd.Stdout = os.Stdout
    cmd.SysProcAttr = &syscall.SysProcAttr {
    //creates a namespace
    	Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
    	Unshareflags: syscall.CLONE_NEWNS,//kernel:don't share mount ns with host.
    } 
    must(cmd.Run())
} 

func child(){

    fmt.Printf("RunningP2 %v as %d\n", os.Args[2:], os.Getpid())
    cg()
    syscall.Sethostname([]byte("CONTAINER"))
    //change root directory so container sees only our new filesystem:
    //used a lite ubuntu base fs:
    syscall.Chroot("ubuntu-base-14.04-core-i386")
    //change dir to root else you end up in undefined location:
    syscall.Chdir("/.")
    //proc is a pseudofilesys: interface between user and kernal space for communication
    syscall.Mount("proc", "proc", "proc", 0, "")
    cmd := exec.Command(os.Args[2], os.Args[3:]...)//run arbitrary command
    cmd.Stdin = os.Stdin
    cmd.Stderr = os.Stderr
    cmd.Stdout = os.Stdout
    must(cmd.Run())
    must(syscall.Unmount("/proc", 0))
} 

func cg(){
	cgroups := "/sys/fs/cgroup"
	pids := filepath.Join(cgroups, "pids")
	err := os.Mkdir(filepath.Join(pids, "samira"), 0755)
	if(err != nil && !os.IsExist(err)){
		panic(err)  
	}
	//set max processes to 20
	must(ioutil.WriteFile(filepath.Join(pids, "samira/pids.max"), []byte("20"), 0700))
	//remove new cgroup in place after the container exists
	must(ioutil.WriteFile(filepath.Join(pids, "samira/notify_on_release"), []byte("1"), 0700))
       must(ioutil.WriteFile(filepath.Join(pids, "samira/cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700)) 	

}

func must(err error) {

	if err != nil {
	    panic(err)
	}

}
