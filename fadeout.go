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
	/*
		y = A * ln(1 + (B/x))
		Where:

			y: is the amplitude value of the signal at a given time.
			A: is the maximum amplitude value of the signal (i.e., the amplitude value before the fade out begins).
			B: is a constant that controls how quickly the amplitude decreases. A higher value of B will produce a faster fade out.
			x: is the time elapsed from the beginning of the fade out to the current moment.

	*/
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

	// Retrieve the output of the command
	output, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	// Extract the volume level from the output using string manipulation
	volumeStr := strings.Split(string(output), "[")[1]
	volumeStr = strings.Split(volumeStr, "]")[0]
	volumeStr = strings.TrimSpace(volumeStr)
	volumeStr = strings.TrimRight(volumeStr, "%")

	// Parse the volume level as an integer
	volume, err := strconv.ParseFloat(volumeStr, 64)
	if err != nil {
		return 0, err
	}

	return volume, nil
}

func setVolume(volume float64) error {
	volumeStr := fmt.Sprintf("%d%%", int(volume))
	cmd := exec.Command("amixer", "-D", "pulse", "sset", "Master", volumeStr)
	err := cmd.Run()

	if err != nil {
		return err
	}

	return nil
}
func exponentialFadeOut(initialVolume, elapsed, tau float64) float64 {
	return initialVolume * math.Exp(-elapsed/tau)
}

func getTau(initialVolume, totalTime float64) float64 {
	return totalTime / math.Log(initialVolume)
}
func fadeOut(volume, initialVolume, tau float64, elapsed time.Duration) {
	volume = exponentialFadeOut(initialVolume, elapsed.Seconds(), tau)
	setVolume(volume)
}

func toggleMediaReproduction() {
	pause := exec.Command("xdotool", "key", "XF86AudioPlay")
	err := pause.Run()
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Second)
}

func endFadeOut(initialVolume float64) {
	toggleMediaReproduction()
	setVolume(initialVolume)
}
