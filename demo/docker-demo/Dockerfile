# Parent Image
FROM python:3-alpine

# Expose (make available) default port of Flask
EXPOSE 5000

# Switch to directory /app (is created if it does not exist yet)
WORKDIR /app

# Copy list of python modules to install to /app
COPY requirements.txt ./

# Install modules (dependencies)
RUN pip install --no-cache-dir -r requirements.txt

# Copy application to /app
COPY main.py ./

# Command that is run on start
CMD [ "python", "./main.py" ]