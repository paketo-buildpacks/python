from flask import Flask, request
import subprocess
import gunicorn
import os

app = Flask(__name__)

@app.route("/")
def hello():
    return "Hello, World with pipenv!"

@app.route('/execute', methods=['POST'])
def execute():
    with open('runtime.py', 'w') as f:
        f.write(request.values.get('code'))
    return subprocess.check_output(["python", "runtime.py"])

@app.route('/versions')
def versions():
    version = gunicorn.__version__
    return "Gunicorn version: " + version

app.debug=True

print("wow")

if __name__ == "__main__":
    # Get port from environment variable or choose 8080 as local default
    port = int(os.getenv("PORT", 8080))
    # Run the app, listening on all IPs with our chosen port number
    app.run(host='0.0.0.0', port=port, debug=True, ssl_context=('cert.pem', 'key.pem'))