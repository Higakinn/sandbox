def delegator():
    """
    >>> g = delegator()
    >>> next(g)
    >>> g.send('A')
    A
    >>> g.send('B')
    B
    >>> g.send(None)
    fin.
    """
    while True:
        res = yield from subgenerator()
        print(res)

def subgenerator():
    while True:
        print("hoge")
        recv = yield
        if recv is None:
            break
        print(recv)
    return 'fin.'

g = delegator()
next(g)
g.send('A')
# g.send('B')
# g.send(None)