def yes():
    """
    >>> yes_man = yes()
    >>> print(next(yes_man))
    A
    B
    yes
    >>> print(next(yes_man))
    C
    B
    yes
    >>> print(next(yes_man))
    C
    B
    yes
    """
    print('A')
    while True:
        print('B')
        yield 'yes'
        print('C')

yes_man=yes()
# next関数を利用すると、yieldまで処理をすすめる（yieldで処理がとまる。）
print(next(yes_man))
print(next(yes_man))
