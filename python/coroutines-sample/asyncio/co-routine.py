def recv():
    """
    >>> receiver = recv()
    >>> next(receiver)
    Started.
    >>> receiver.send(1)
    Receive: 1
    >>> receiver.send(2)
    Receive: 2
    >>> receiver.send(3)
    Receive: 3
    """
    print('Started.')
    while True:
        v = yield
        print(f'Receive: {v}')

receiver = recv()
next(receiver)
receiver.send(1)
receiver.send(2)
receiver.send(3)