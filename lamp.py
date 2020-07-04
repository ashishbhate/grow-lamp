#!/usr/bin/python3
import RPi.GPIO as GPIO # Import Raspberry Pi GPIO library
import datetime
import sys
from time import sleep, localtime

def turn_on(pin):
    GPIO.output(pin, GPIO.LOW)       # set port/pin value to 0/GPIO.LOW/False  

def turn_off(pin):
    GPIO.output(pin, GPIO.HIGH)       # set port/pin value to 1/GPIO.HIGH/True  

def is_on(pin):
    return not GPIO.input(pin)

def set_by_schedule(pin):
    tt = datetime.datetime.now().time()
    if tt >= TIME_ON and tt <= TIME_OFF:
        turn_on(pin)
        print("turned on")
        return
    turn_off(pin)
    print("turned off")

TIME_ON = datetime.time(9) 
TIME_OFF = datetime.time(19) 
PIN = 11

GPIO.setwarnings(False) # Ignore warning for now
GPIO.setmode(GPIO.BOARD) # Use physical pin numbering
GPIO.setup(PIN, GPIO.OUT) # set a port/pin as an output   

if len(sys.argv) < 2:
    print("on" if is_on(PIN) else "off")
    sys.exit(0)

if sys.argv[1] == "on":
    turn_on(PIN)
elif sys.argv[1] == "off": 
    turn_off(PIN)
elif sys.argv[1] == "set-by-schedule": 
    set_by_schedule(PIN)
else:
    print("on" if is_on(PIN) else "off")
