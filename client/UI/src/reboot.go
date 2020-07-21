package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"

	//"io/ioutil"

	"golang.org/x/sys/windows"
)

// __________________________________________________________________________main________________________________________________________________
func main() {
	setup := getInfo()
	cameras := getCameras()
	fmt.Println("Num of Cameras: ", len(cameras)-1)

	dt := time.Now()
	input := setup.Status
	cur := setup.CurrentReboots
	max := setup.MaxReboots
	isPassed := setup.IsPassed

	if input != "noReboot" {
		for setup.CurrentReboots < setup.MaxReboots {
			fmt.Println("Status: ", input)
			fmt.Println("currentReboots: ", cur)
			fmt.Println("maxReboots: ", max)
			fmt.Println("isPassed: ", isPassed)
			if isPassed {

				if input == "SRX-Pro" { // reset srx-pro through date/time change
					// get admin permissions if not administrator
					if !amAdmin() {
						runMeElevated()
					}
					time.Sleep(10 * time.Second)

					tm := setup.Date

					time.Sleep(10 * time.Second) // we can replace this with time interval later on

					//dt := time.Date(2020, 06, 17, 20, 34, 58, 65, time.UTC) // yyyy, mm, dd, hh, min, ss, ms
					dt, _ = time.Parse("2006-01-02T15:04:05Z", tm)
					fmt.Println(dt)

					err1 := SetSystemDate(dt)
					if err1 != nil {
						fmt.Printf("Error: %s", err1.Error())
					}
					err2 := SetSystemTime(dt)
					if err2 != nil {
						fmt.Printf("Error: %s", err2.Error())
					}
					// With short weekday (Mon)
					// fmt.Println((dt).Format("01-02-2006 15:04:05.00 Mon"))

				} else if input == "Switch" { // reset switches through http request

					ip := setup.SwitchIP
					user := setup.User
					pass := setup.Pass

					rebootSwitch(ip, user, pass)
					// err := rebootSwitch(ip, user, pass)
					// if err != nil {
					// 	fmt.Printf("Error: %s", err.Error())
					// }
				} else if input == "UIO8" {

				}
			}
			setup.CurrentReboots++
			if setup.CurrentReboots == setup.MaxReboots {
				setup.Status = "noReboot"
			}
			// also call results function here and redefine isPassed
			postInfo(setup)

		}
	}
	ip := cameras[1].Ip             // placeholder
	setup.IsPassed = pingCamera(ip) // placeholder
	fmt.Println("Passed:", setup.IsPassed)

}

// ______________________________________________________________________setSystemDate && setSystemTime__________________________________________

func SetSystemDate(newTime time.Time) error { // Set the system date to desired reboot date
	// _, lookErr := exec.LookPath("date")
	// if lookErr != nil {
	// 	fmt.Printf("Date binary not found, cannot set system date: %s\n", lookErr.Error())
	// 	return lookErr
	// } else {
	dateString := newTime.Format("2006-01-02") // convert date to string
	fmt.Printf("Setting system date to: %s\n", dateString)
	args := []string{"/C", "date", dateString}
	cmd := exec.Command("cmd.exe", args...) // cmd
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	return cmd.Run()
	// if err != nil {
	// 	fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
	// }
	// fmt.Println("Result: " + out.String())
	// return err
}

// Set the system time to desired reboot time
func SetSystemTime(newTime time.Time) error {
	timeString := newTime.Format("15:04:05.00") // convert time to string
	// converting timezone
	a := []rune(timeString)
	hour := string(a[0:2])
	rest := string(a[len(a)-9 : len(a)])
	i, err := strconv.Atoi(hour)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
	if i >= 4 {
		i -= 4
	} else {
		i = i + 20
	}
	hour = strconv.Itoa(i)
	timeString = hour + rest

	fmt.Printf("Setting system time to: %s\n", timeString)
	args := []string{"/C", "time", timeString}
	cmd := exec.Command("cmd.exe", args...) // cmd
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	return cmd.Run()
}

//________________________________________________________________________rebootSwitch________________________________________________________________________________
func rebootSwitch(ip string, user string, pass string) { // Reboot switches through http
	uri := strings.Join([]string{"http://", ip, "/rb-cgi/reboot.cgi?user=", user, "&pwd=", pass}, "")
	res, err := http.Get(uri)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}

	// data, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()

}

//______________________________________________________________________admin______________________________________________________________________________________
// get admin permissions
func runMeElevated() {
	verb := "runas"
	exe, _ := os.Executable()
	cwd, _ := os.Getwd()
	args := strings.Join(os.Args[1:], " ")

	verbPtr, _ := syscall.UTF16PtrFromString(verb)
	exePtr, _ := syscall.UTF16PtrFromString(exe)
	cwdPtr, _ := syscall.UTF16PtrFromString(cwd)
	argPtr, _ := syscall.UTF16PtrFromString(args)

	var showCmd int32 = 1 //SW_NORMAL

	err := windows.ShellExecute(0, verbPtr, exePtr, argPtr, cwdPtr, showCmd)
	if err != nil {
		fmt.Println(err)
	}
}

// determines if the current user is an administrator
func amAdmin() bool {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	if err != nil {
		fmt.Println("You do not have administrator permissions.")
		return false
	}
	fmt.Println("You are currently in administrator mode.")
	return true
}
