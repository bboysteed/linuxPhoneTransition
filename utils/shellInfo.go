package utils

import (
	"bytes"
	"os/exec"
)

func GetShellPrefix()[]byte{
	who,_ := exec.Command("/bin/bash","-c","whoami").Output()
	hostname,_ := exec.Command("/bin/bash","-c","cat /etc/hostname").Output()
	var res bytes.Buffer
	res.Write(who)
	res.WriteString("@")
	res.Write(hostname)
	res.WriteString(":/#")
	return bytes.ReplaceAll(res.Bytes(),[]byte("\n"),[]byte(""))

}
