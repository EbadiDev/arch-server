package cmd

import (
	"fmt"
	"time"

	"github.com/EbadiDev/Arch-Server/common/exec"
	"github.com/spf13/cobra"
)

var (
	startCommand = cobra.Command{
		Use:   "start",
		Short: "Start Arch-Server service",
		Run:   startHandle,
	}
	stopCommand = cobra.Command{
		Use:   "stop",
		Short: "Stop Arch-Server service",
		Run:   stopHandle,
	}
	restartCommand = cobra.Command{
		Use:   "restart",
		Short: "Restart Arch-Server service",
		Run:   restartHandle,
	}
	logCommand = cobra.Command{
		Use:   "log",
		Short: "Output Arch-Server log",
		Run: func(_ *cobra.Command, _ []string) {
			exec.RunCommandStd("journalctl", "-u", "Arch-Server.service", "-e", "--no-pager", "-f")
		},
	}
)

func init() {
	command.AddCommand(&startCommand)
	command.AddCommand(&stopCommand)
	command.AddCommand(&restartCommand)
	command.AddCommand(&logCommand)
}

func startHandle(_ *cobra.Command, _ []string) {
	r, err := checkRunning()
	if err != nil {
		fmt.Println(Err("check status error: ", err))
		fmt.Println(Err("Failed to start Arch-Server"))
		return
	}
	if r {
		fmt.Println(Ok("Arch-Server is already running, no need to start again. If you want to restart, please use the restart command"))
	}
	_, err = exec.RunCommandByShell("systemctl start Arch-Server.service")
	if err != nil {
		fmt.Println(Err("exec start cmd error: ", err))
		fmt.Println(Err("Failed to start Arch-Server"))
		return
	}
	time.Sleep(time.Second * 3)
	r, err = checkRunning()
	if err != nil {
		fmt.Println(Err("check status error: ", err))
		fmt.Println(Err("Failed to start Arch-Server"))
	}
	if !r {
		fmt.Println(Err("Arch-Server may have failed to start, please use the Arch-Server log command to view the log information later"))
		return
	}
	fmt.Println(Ok("Arch-Server started successfully, please use the Arch-Server log command to view the running log"))
}

func stopHandle(_ *cobra.Command, _ []string) {
	_, err := exec.RunCommandByShell("systemctl stop Arch-Server.service")
	if err != nil {
		fmt.Println(Err("exec stop cmd error: ", err))
		fmt.Println(Err("Failed to stop Arch-Server"))
		return
	}
	time.Sleep(2 * time.Second)
	r, err := checkRunning()
	if err != nil {
		fmt.Println(Err("check status error:", err))
		fmt.Println(Err("Failed to stop Arch-Server"))
		return
	}
	if r {
		fmt.Println(Err("Failed to stop Arch-Server, it may be because the stop time exceeded two seconds, please check the log information later"))
		return
	}
	fmt.Println(Ok("Arch-Server stopped successfully"))
}

func restartHandle(_ *cobra.Command, _ []string) {
	_, err := exec.RunCommandByShell("systemctl restart Arch-Server.service")
	if err != nil {
		fmt.Println(Err("exec restart cmd error: ", err))
		fmt.Println(Err("Failed to restart Arch-Server"))
		return
	}
	r, err := checkRunning()
	if err != nil {
		fmt.Println(Err("check status error: ", err))
		fmt.Println(Err("Failed to restart Arch-Server"))
		return
	}
	if !r {
		fmt.Println(Err("Arch-Server may have failed to start, please use the Arch-Server log command to view the log information later"))
		return
	}
	fmt.Println(Ok("Arch-Server restarted successfully"))
}
