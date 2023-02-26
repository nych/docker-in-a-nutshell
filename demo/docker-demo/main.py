#!/usr/bin/env python3

import datetime
from flask import Flask

app = Flask('My first flask based WebApp')

@app.route('/')
def index():
    return 'Hey you! The time is: ' + str(datetime.datetime.now().time())

if __name__ == '__main__':
    app.run(host='0.0.0.0')