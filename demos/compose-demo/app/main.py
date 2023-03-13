#!/usr/bin/env python3
from flask import Flask
import socket
import mysql.connector as database
import sys

# Connect to MariaDB Platform
try:
    connection = database.connect(
        user="db_user",
        password="changeme_too",
        host="app_db",
        port=3306,
        database="app_db"
    )
except:
    print(f"Connecting to maria db failed")
    sys.exit(1)

# Get Cursor
cursor = connection.cursor()

app = Flask(__name__)

@app.route('/hit')
def hit():
	try:
		cursor.execute("UPDATE hit SET hit_count = hit_count + 1 WHERE id = 0;")
		return hits()
	except:
		return "Update failed hitcount failed", 500
	

@app.route('/hits')
def hits():
	try:
		cursor.execute("SELECT hit_count FROM hit;")
		hit_count = 0
		for (h) in cursor:
			hit_count = h
		return "host:{} - Current hit count is {}\n".format(
			socket.gethostname(),
			hit_count[0],
		), 200
	except:
		return "Retrieving hit count failed", 500

@app.route('/reset')
def reset():
	try:
		cursor.execute("UPDATE hit SET hit_count = 0 WHERE id = 0;")
		return hits()
	except:
		return "Reset hit count failed", 500

if __name__ == '__main__':
	app.run('0.0.0.0', debug=True)
