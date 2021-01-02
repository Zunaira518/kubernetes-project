package models

import (
	"bapi/constants"
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	"net"
	"runtime"
	"strconv"
	"time"
)


type VM struct {
	Id       		string		`json:"-" bson:"_id,omitempty"`
	Name			string		`json:"name" bson:"name" valid:"required" description:"User name of the VM [required]"`
	IpAddress 		string		`json:"ip_address" bson:"ip_address" valid:"required" description:"Ip address of the VM [required]"`
	SSHPassword	 	string      `json:"ssh_key" bson:"ssh_key" valid:"required" description:"SSH password for the VM [required]"`
}

func AddWorkerNode(vm VM) string {

	os := runtime.GOOS
	version := runtime.GOARCH
	switch os {
		case "linux":
			fmt.Println("Linux")
			switch version {
				case "amd64":
					runScripts(vm, "Linux")
				default:
					fmt.Printf("%s.\n", os+" is not supported for this application. Run it on Linux centos 7.8 and ubuntu 18.04.")
	}
	default:
		fmt.Printf("%s.\n", os+" is not supported for this application. Run it on Linux centos 7.8 and ubuntu 18.04.")
	}

	vm.Id = "vm_" + strconv.FormatInt(time.Now().UnixNano(), 10)

	return vm.Id

}

func setScriptsPermission(vm VM ,fileName string ) error{

	cmd := "chmod 700 " +fileName
	_ ,err := remoteRun(vm.Name,vm.IpAddress,vm.SSHPassword,cmd)
	if err !=nil{
		fmt.Print("Error in setting file permission")
		return err
	}
	return nil
}

func runScripts(vm VM, os string) error{

	filename := "/home/" + vm.Name + "/scripts/"+ os +  constants.CENTOS_K8S_CONTAINERD

	err := setScriptsPermission(vm ,filename)
	if err !=nil {
		return err
	}
	_ ,err = remoteRun(vm.Name,vm.IpAddress,vm.SSHPassword,filename)
	if err !=nil{
		fmt.Print(err)
		return err
	}

	filename = "/home/" + vm.Name + "/scripts/"+ os +  constants.CENTOS_K8S_CRIO

	err = setScriptsPermission(vm ,filename)
	if err !=nil {
		return err
	}

	_ ,err = remoteRun(vm.Name,vm.IpAddress,vm.SSHPassword,filename)
	if err !=nil{
		fmt.Print(err)
		return err
	}

	filename = "/home/" + vm.Name + "/scripts/"+ os + constants.CENTOS_K8S_DOCKER

	err = setScriptsPermission(vm, filename)
	if err !=nil {
		return err
	}

	_ ,err = remoteRun(vm.Name,vm.IpAddress,vm.SSHPassword,filename)
	if err !=nil{
		fmt.Print(err)
		return err
	}
	return nil
}
func remoteRun(user string, addr string, password string, cmd string) (string, error) {
	config := &ssh.ClientConfig{
		User: user,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		   Auth: []ssh.AuthMethod{
		       ssh.Password(password),
		   },
	}
	// Connect
	client, err := ssh.Dial("tcp", net.JoinHostPort(addr, "22"), config)
	if err != nil {
		return "", err
	}
	// Create a session. It is one session per command.
	session, err := client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()
	var b bytes.Buffer  // import "bytes"
	session.Stdout = &b // get output
	// Run the command
	err = session.Run(cmd)
	return b.String(), err
}