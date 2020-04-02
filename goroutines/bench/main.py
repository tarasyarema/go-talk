from threading import Lock, Thread
from sys import argv

MAX_T, LOOPS = 1000, 1000
counter = 0


def task(lock):
    for _ in range(LOOPS):
        continue

    lock.acquire()
    global counter
    counter += 1
    lock.release()


def main():
    global counter
    counter = 0

    lock = Lock()

    threads = [Thread(target=task, args=(lock,)) for _ in range(MAX_T)]

    for t in threads:
        t.start()

    for t in threads:
        t.join()


if __name__ == "__main__":
    if len(argv) > 1:
        MAX_T = int(argv[1])
    
    main()
    print(f"> the counter is {counter}")
