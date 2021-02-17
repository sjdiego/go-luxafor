# goluxafor
A little script to manage colors of Luxafor LED devices

## How to use
You must compile for the platform where is going to be executed.

The available parameters are:
* `-d (1-n)` device where command will be performed
* `-r (0-255)` intensity of red color
* `-g (0-255)` intensity of green color
* `-b (0-255)` intensity of blue color

To turn on a color you have to know the RGB codes. For example, if you want to turn on the leds of your device with
a green color, you could use the following command:
* `./luxafor -g 255`

Maybe you have multiple Luxafor LED devices. You can select which device will receive the commando using `-d` parameter.
For example, you can send the command to another device with this command:
* `./luxafor -d 4 -r 255 -b 255`

If you want to turn off the device, execute the command without color parameters:
* `./luxafor`

## Real use example
I have a Raspberry Pi 3B+. I can use it combined with a bash script to monitor RPi CPU temperature and switch colors when 
it's too high. The script was compiled for arm5 and copied to `/usr/bin/`to be able to execute simply using `luxafor` command.

Then a bash script is executed as cron job each minute to check the temperature and send the proper command to the device:

```
#!/bin/bash

MIN_TEMP=52
MAX_TEMP=56

cpu=$(</sys/class/thermal/thermal_zone0/temp)
cpu=$((cpu/1000))


echo "CPU temperature is: $cpu ºC"


if [ "$cpu" -lt $MIN_TEMP ]; then
   luxafor -g 20
elif [ "$cpu" -ge $MIN_TEMP ] && [ "$cpu" -lt $MAX_TEMP ]; then
   luxafor -r 20 -g 20
elif [ "$cpu" -ge $MAX_TEMP ]; then
   luxafor -r 60
fi
```


The possibilities are multiple. Use your own imagination :-)