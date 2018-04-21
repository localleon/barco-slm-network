#!/usr/bin/python3

import datetime
import sacn
from multiprocessing import Process
from flask import Flask, render_template, jsonify, request, redirect, Response
from flask_cors import CORS
from flask_httpauth import HTTPBasicAuth

import logging
from logging.handlers import RotatingFileHandler
logger = logging.getLogger('werkzeug')  # Gets Flasks main Logger

# ============= SERIAL STuff
import serial   # Importing pyserial Libary
ser = serial.Serial('/dev/ttyUSB0')  # open serial port with baudrate of 9600 and standard settings

def send_serial(tosend):
    # Sends @tosend to the serial port
    ser.write(tosend)

#============  CONFIG  ============
from apiconfig import *        # Imports Config Values
app = Flask(__name__)          # Inits Flask APP
CORS(app)                      # Adds a Cors Header to the Responses
auth = HTTPBasicAuth()         # HTTP Basic Authentication
# =================== Logging ===================================0

def file_logger():
    """
    Creates a log file in the zeus directory named zeus.log, it will contain 
    everything the api logs.
    """
    logger.setLevel(logging.DEBUG)
    logFile = "logs/app.log"
    handler = RotatingFileHandler(logFile, mode='a', maxBytes=5 * 1024 * 1024,
                                  backupCount=2, encoding=None, delay=0)
    logger.addHandler(handler)

#============================= APP Routes for GUIS ===================================0
@app.route('/help')
@auth.login_required
def send_help():
    # Displays Help/Documentary Site
    return render_template('help.html'), 200

@app.route('/')
@auth.login_required
def send_app():
    # Returns the Frontend
    return render_template('app.html'), 200

#================= API Function Route ===================================


@app.route('/app/send/<string:cmd>', methods=['GET', 'POST', 'PUT'])
@auth.login_required
def tasks_control(cmd):
    # Matches Api Keys in serial_commands with gotten cmd and sending it out to serial port
    if cmd in serial_commands.keys():
        send_serial(serial_commands[cmd])
        return jsonify(success_response), 200
    else:
        return "", 404


@app.route('/ping', methods=['GET', 'POST', 'PUT'])
@auth.login_required
def ping():
    """
    A simple ping function that returns pong when executed. 
    Can be used for heartbeats and similar things.
    This is easier than always get the index.html. 
    Also know as Heartbeat !
    """
    return Response("pong", mimetype='text/plain')

#=============HTTP Basic  AUTHENTICATION ============================
@auth.get_password
def get_pw(username):
    if username in users:
        return users[username]
    return None

#================ Redirect on Port 80 ==================
redirectApp = Flask("redirectApp")
# redirect http traffic to https:

@redirectApp.before_request
def before_request():
    if request.url.startswith('http://'):
        return redirect(request.url.replace('http://', 'https://', 1), code=301)

def start_redirect():
    redirectApp.run(port=80, host="0.0.0.0")

#==================== MAIN =================================
if __name__ == '__main__':
    # file_logger()

    # sACN
    sacn_rcv = sacn.sACNreceiver()
    sacn_process = Process(target=sacn_rcv.start())
    sacn_process.start()
    # list for storing info wether a channel was flashed or not so they dont get flashed multiple times
    flashed = [False] * 12
    # listens on universe 99 for DMX data
    @sacn_rcv.listen_on('universe', universe=99)
    def callback(packet):  # packet type: sacn.DataPacket
        # print(packet.dmxData)  # print the received DMX data for Debugging
        for i, channel_value in enumerate(packet.dmxData):
            if i in sacn_mapping.keys():
                if channel_value < 255:
                    flashed[i] = False
                elif channel_value >= 255 and flashed[i] == False:
                    send_serial(sacn_mapping[i])
                    flashed[i] = True
            else:
                break

    sacn_rcv.join_multicast(100)

    # Runs FLASK Server
    process = Process(target=start_redirect)
    process.start()
    app.run(host="0.0.0.0", port=443, debug=False,
            ssl_context=('ssl/cert.pem', 'ssl/key.pem'))
    process.terminate()
    process.join()
    sacn_rcv.stop()
