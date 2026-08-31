package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"
	"time"

	"github.com/hugolgst/rich-go/client"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/net"
	"golang.org/x/sys/windows/registry"
)

const (
	ClientID    = "1543382496031281304"
	LargeImage  = "https://i.imgur.com/zyPZ9Gw.gif"
	SmallImage  = "https://i.imgur.com/FP0hM27.png"
	UpdateEvery = 10 * time.Second
	ShowOS      = true
	AppName     = "SystemMonitorRPC"
)

func humanReadable(bytes uint64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := uint64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

func getOSName() string {
	switch runtime.GOOS {
	case "windows":
		return "Windows"
	case "darwin":
		return "macOS"
	default:
		return "Linux"
	}
}

func getPresence() (details, state string) {
	cpuPercent, _ := cpu.Percent(0, false)
	cpuVal := 0.0
	if len(cpuPercent) > 0 {
		cpuVal = cpuPercent[0]
	}

	vm, _ := mem.VirtualMemory()

	diskIO, _ := disk.IOCounters()
	var readBytes, writeBytes uint64
	for _, v := range diskIO {
		readBytes += v.ReadBytes
		writeBytes += v.WriteBytes
	}

	netIO, _ := net.IOCounters(false)
	var sent, recv uint64
	if len(netIO) > 0 {
		sent = netIO[0].BytesSent
		recv = netIO[0].BytesRecv
	}

	osName := getOSName()
	osSuffix := ""
	if ShowOS {
		osSuffix = " | " + osName
	}

	variants := [][2]string{
		{
			fmt.Sprintf("CPU: %.1f%%%s", cpuVal, osSuffix),
			fmt.Sprintf("RAM: %.1f%% (%s/%s)", vm.UsedPercent, humanReadable(vm.Used), humanReadable(vm.Total)),
		},
		{
			fmt.Sprintf("Disk R/W: %s/%s%s", humanReadable(readBytes), humanReadable(writeBytes), osSuffix),
			fmt.Sprintf("Net S/R: %s/%s", humanReadable(sent), humanReadable(recv)),
		},
		{
			fmt.Sprintf("CPU: %.1f%%%s", cpuVal, osSuffix),
			fmt.Sprintf("Disk R/W: %s/%s", humanReadable(readBytes), humanReadable(writeBytes)),
		},
		{
			fmt.Sprintf("RAM: %.1f%% (%s/%s)%s", vm.UsedPercent, humanReadable(vm.Used), humanReadable(vm.Total), osSuffix),
			fmt.Sprintf("Net S/R: %s/%s", humanReadable(sent), humanReadable(recv)),
		},
	}

	choice := variants[rand.Intn(len(variants))]
	return choice[0], choice[1]
}

func updatePresence() {
	details, state := getPresence()
	osName := getOSName()

	now := time.Now()

	_ = client.SetActivity(client.Activity{
		Details:    details,
		State:      state,
		LargeImage: LargeImage,
		LargeText:  fmt.Sprintf("System Performance Monitor - %s", osName),
		SmallImage: SmallImage,
		SmallText:  fmt.Sprintf("Monitoring %s System...", osName),
		Timestamps: &client.Timestamps{
			Start: &now,
		},
	})
}

func getExePath() (string, error) {
	exe, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Abs(exe)
}

func addToStartup() error {
	if runtime.GOOS != "windows" {
		return nil
	}

	exePath, err := getExePath()
	if err != nil {
		return err
	}

	key, err := registry.OpenKey(registry.CURRENT_USER,
		`Software\Microsoft\Windows\CurrentVersion\Run`,
		registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer key.Close()

	return key.SetStringValue(AppName, exePath)
}

func removeFromStartup() error {
	if runtime.GOOS != "windows" {
		return nil
	}

	key, err := registry.OpenKey(registry.CURRENT_USER,
		`Software\Microsoft\Windows\CurrentVersion\Run`,
		registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer key.Close()

	return key.DeleteValue(AppName)
}

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "--install":
			err := addToStartup()
			if err != nil {
				fmt.Println("Failed to add to startup:", err)
			} else {
				fmt.Println("Successfully added to Windows startup!")
			}
			return
		case "--uninstall":
			err := removeFromStartup()
			if err != nil {
				fmt.Println("Failed to remove from startup:", err)
			} else {
				fmt.Println("Successfully removed from Windows startup!")
			}
			return
		}
	}

	_ = addToStartup()

	err := client.Login(ClientID)
	if err != nil {
		os.Exit(0)
	}
	defer client.Logout()

	updatePresence()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	ticker := time.NewTicker(UpdateEvery)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			updatePresence()
		case <-c:
			return
		}
	}
}