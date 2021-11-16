
from .celery import app

import time

@app.task
def add(x, y):

    return x + y


@app.task
def mul(x, y):
    return x * y


@app.task
def xsum(numbers):
    return sum(numbers)