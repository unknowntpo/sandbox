import sys
import pika
import threading

def callback(ch, method, property, body):
    print("server got ping")
    ch.basic_ack(delivery_tag = method.delivery_tag)
    channel.basic_publish(exchange='', routing_key='pong', body="pong")

def declareQueue(ch):
    channel.queue_declare(queue='ping', durable=True)
    channel.queue_declare(queue='pong', durable=True)

def getChannel():
    # Set up connection
    connection = pika.BlockingConnection(pika.ConnectionParameters(host='localhost'))
    # get channel from connection
    return connection.channel()

def serverFn():
    channel.basic_consume(queue='ping', on_message_callback=callback)
    channel.start_consuming()

def clientFn():
    channel.basic_publish(exchange='', routing_key='ping', body="ping")



channel = getChannel()
declareQueue(channel)
# send message 

server = threading.Thread(name="serverThread", target=serverFn)
server.start()

client = threading.Thread(name="clientThread", target=clientFn)
client.start()

client.join()
server.join()
