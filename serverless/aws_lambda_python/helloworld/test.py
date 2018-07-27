import logging

logger = logging.getLogger()
logger.setLevel(logging.INFO)

def lambda_handler(event, context):
    logger.info("Event: " + str(event))
    logger.info("context: " + str(context))
    message = 'Hello {} {}!'.format(str(event), str(context))
    return {
        'message' : message
    }
