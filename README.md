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

## Instalation
Clone this repository
```bash
git clone https://github.com/Chavis00/media-volume-fadeout.git
```
cd into directory
```bash
cd media-volume-fadeout
```

Build fadeout
```bash
go build fadeout
```

## Run
This is an example to run a 30 seconds fadeout
```bash
./fadeout -s 30
```
If not -s flag, time will be 15s by default

Now you can add a new shortcut in your operative sistem to execute
```bash
./path/to/fadeout -s {time-in-seconds}
``` 