# Introduction
A simple fadeout for pause media in linux, using exponential fadeout 
```math
v(t) = v0 * e^(-t/tau)

```
Where:
- v(t): is the specific value of volume in t .
- v0: is the volume at the beggining of the fadeout.
- t: is the time elapsed from the beginning of the fade out to the current moment
- tau: is a time constant that determines the rate at which the quantity exponentially decreases over time.

## Installation
Clone this repository
```bash
git clone https://github.com/Chavis00/media-volume-fadeout.git
```
cd into the directory
```bash
cd media-volume-fadeout
```

Build fadeout
```bash
go build fadeout
```

## Run
Here is an example of how to run a 30-second fadeout:
```bash
./fadeout -s 30
```
If the -s flag is not specified, the time will default to 15 seconds.

Now, you can add a new shortcut in your operating system to execute:

```bash
./path/to/fadeout -s {time-in-seconds}
``` 