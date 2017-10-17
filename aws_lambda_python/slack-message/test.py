#!/usr/bin/env python
##-------------------------------------------------------------------
## @copyright 2017 DennyZhang.com
## Licensed under MIT
##   https://www.dennyzhang.com/wp-content/mit_license.txt
##
## File: test.py
## Author : Denny <contact@dennyzhang.com>
## Description :
## --
## Created : <2017-10-16>
## Updated: Time-stamp: <2017-10-16 19:52:58>
##-------------------------------------------------------------------
import os
import logging

logger = logging.getLogger()
logger.setLevel(logging.INFO)

def send_slack_message(msg):
    from slackclient import SlackClient
    slack_token = os.environ["SLACK_API_TOKEN"]
    slack_channel = os.environ["SLACK_CHANNEL"]
    sc = SlackClient(slack_token)
    sc.api_call("chat.postMessage", channel = slack_channel, text = msg)

def lambda_handler(event, context):
    logger.info("Event: " + str(event))
    logger.info("context: " + str(context))
    msg = event["msg"]
    send_slack_message(msg)
    return {
        'message' : "OK: %s" % (msg)
    }

# Test in command line
# export SLACK_API_TOKEN="XXX"
# export SLACK_CHANNEL="#denny-alerts"
if __name__ == '__main__':
    fake_event={}
    fake_event["msg"] = "hello, world"
    lambda_handler(fake_event, "fake_context")
## File: test.py ends
