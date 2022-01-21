import sys
import pika

def callback(ch, method, property, body):
    print("server got ping")
    ch.basic_ack(delivery_tag = method.delivery_tag)

msg = "ping"

# Set up connection
connection = pika.BlockingConnection(pika.ConnectionParameters(host='localhost'))
# get channel from connection
channel = connection.channel()

# declare queue
channel.queue_declare(queue='ping', durable=True)
channel.queue_declare(queue='pong', durable=True)
# send message 

channel.basic_consume(queue='ping', on_message_callback=callback)
channel.basic_publish(exchange='', routing_key='pong', body=msg)

channel.start_consuming()
