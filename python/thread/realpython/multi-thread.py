import logging
import threading
import time

def thread_funtion(name):
    logging.info("thread %s starting", name)
    logging.info("thread %s finishing", name)

if __name__ == "__main__":
    format = "%(asctime)s: %(message)s"
    logging.basicConfig(format=format, level=logging.INFO, datefmt="%H:%M:%S")

    threads = list()
    for i in range(3):
        x = threading.Thread(target=thread_funtion, args=(i,))
        threads.append(x)
        x.start()

    for i, t in enumerate(threads):
        logging.info("Main    : before joining thread %d.", i)
        t.join()
        logging.info("Main    : thread %d done", i)

