def yes():
    """
    >>> yes_man = yes()
    >>> yes_man.__class__.__name__
    generator
    >>> print(next(yes_man))
    yes
    >>> print(next(yes_man))
    yes
    >>> print(next(yes_man))
    yes
    """
    while True:
        yield 'yes'

yes_man = yes()

print(next(yes_man))
print(next(yes_man))
print(next(yes_man))
print(next(yes_man))
print(next(yes_man))