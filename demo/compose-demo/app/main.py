#!/usr/bin/env python3

from flask import Flask
from redis import Redis
import socket

app = Flask(__name__)
redis = Redis(host='redis', port=6379)

@app.route('/whoami')
def whoami():
	return socket.gethostname(), 200

@app.route('/hit')
def hit():
	redis.incr('hits')
	return 'host:{} - Current hit count is {}'.format(
		socket.gethostname(),
		redis.get('hits'),
	), 200

@app.route('/reset')
def reset():
	redis.delete('hits')
	return hit()

if __name__ == '__main__':
	app.run('0.0.0.0', debug=True)
