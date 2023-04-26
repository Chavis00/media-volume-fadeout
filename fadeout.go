package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var (
	fadeOutInSeconds int
)

func init() {
	flag.IntVar(&fadeOutInSeconds, "s", 15, "Fade out time in seconds") // 15 seconds by default
	flag.Parse()
}

func main() {
	fadeOutTime := float64(fadeOutInSeconds)
	startTime := time.Now()
	initialVolume, err := getCurrentVolume()

	if err != nil {
		log.Fatal(err)
		return
	}

	tau := getTau(initialVolume, fadeOutTime)
	volume := initialVolume

	var elapsed time.Duration

	for {
		elapsed = time.Since(startTime)
		if elapsed.Seconds() >= float64(fadeOutTime) {
			break
		}
		fadeOut(volume, initialVolume, tau, elapsed)

	}

	endFadeOut(initialVolume)

}

func getCurrentVolume() (float64, error) {
	// Run the "amixer" command with the "-D pulse" flag to specify the default pulseaudio sound card
	cmd := exec.Command("amixer", "-D", "pulse", "get", "Master")
	output, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	// Extract the volume level from the output using string manipulation
	volumeStr := strings.Split(string(output), "[")[1]
	volumeStr = strings.Split(volumeStr, "]")[0]
	volumeStr = strings.TrimSpace(volumeStr)
	volumeStr = strings.TrimRight(volumeStr, "%")

	volume, err := strconv.ParseFloat(volumeStr, 64)
	if err != nil {
		return 0, err
	}

	return volume, nil
}

func setVolume(volume float64) error {
	// Set current volume of OS
	volumeStr := fmt.Sprintf("%d%%", int(volume))
	cmd := exec.Command("amixer", "-D", "pulse", "sset", "Master", volumeStr)
	err := cmd.Run()

	if err != nil {
		return err
	}

	return nil
}
func exponentialFadeOut(initialVolume, elapsed, tau float64) float64 {
	/*
		v(t) = v0 * e^(-t/tau)
		where:
			v(t): is the specific value of volume in t .
			v0: is the volume at the beggining of the fadeout.
			t: is the time elapsed from the beginning of the fade out to the current moment
			tau: is a time constant that determines the rate at which the quantity exponentially decreases over time.
	*/
	return initialVolume * math.Exp(-elapsed/tau)
}

func getTau(initialVolume, totalTime float64) float64 {
	// tau depends on total time expected and the initial volume
	return totalTime / math.Log(initialVolume)
}
func fadeOut(volume, initialVolume, tau float64, elapsed time.Duration) {
	// Calculate an set new volume
	volume = exponentialFadeOut(initialVolume, elapsed.Seconds(), tau)
	setVolume(volume)
}

func toggleMediaReproduction() {
	// Pause media using XF86AudioPlay key and wait 1 sec to return
	pause := exec.Command("xdotool", "key", "XF86AudioPlay")
	err := pause.Run()
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Second)
}

func endFadeOut(initialVolume float64) {
	// Toggle to pause and set initial volume
	toggleMediaReproduction()
	setVolume(initialVolume)
}
