import sys
import pika

def callback(ch, method, property, body):
    print("client got pong")
    print(" [x] Done")

    ch.basic_ack(delivery_tag = method.delivery_tag)

msg = "ping"

# Set up connection
connection = pika.BlockingConnection()
# get channel from connection
channel = connection.channel()
# send message 
channel.basic_publish(exchange='', routing_key='ping', body=msg)
print("ping sent from client")

# block get message
channel.basic_consume(queue='pong', on_message_callback=callback)
channel.start_consuming()
