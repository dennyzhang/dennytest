#!/usr/bin/env python
##-------------------------------------------------------------------
## @copyright 2017 DennyZhang.com
## Licensed under MIT
##   https://www.dennyzhang.com/wp-content/mit_license.txt
##
## File: lambda_function.py
## Author : Denny <https://www.dennyzhang.com/contact>
## Description :
## --
## Created : <2017-10-16>
## Updated: Time-stamp: <2017-11-13 11:00:54>
##-------------------------------------------------------------------
import requests, json
import os
import logging

logger = logging.getLogger()
logger.setLevel(logging.INFO)

def send_slack_message(msg):
    slack_token = os.environ["SLACK_API_TOKEN"]
    slack_channel = os.environ["SLACK_CHANNEL"]
    slack_username = 'denny-test'
    url = 'https://slack.com/api/chat.postMessage'
    payload = {'username': slack_username, \
               'token': slack_token, \
               'channel': slack_channel, \
               'text': msg}
    print("payload: %s" % payload)
    r = requests.post(url, data=payload)
    return r

def lambda_handler(event, context):
    logger.info("Event: " + str(event))
    logger.info("context: " + str(context))
    msg = event["msg"]
    r = send_slack_message(msg)
    if r.status_code == 200:
        return_msg = "OK: HTTP 200"
    else:
        return_msg = "ERROR: %s" % (r.status_code)
    print(return_msg)

    # TODO: error handling
    return {
        'message' : return_msg
    }

# Test in command line
# export SLACK_API_TOKEN="XXX"
# export SLACK_CHANNEL="#denny-alerts"
if __name__ == '__main__':
    fake_event={}
    fake_event["msg"] = "hello, world"
    lambda_handler(fake_event, "fake_context")
## File: lambda_function.py ends
