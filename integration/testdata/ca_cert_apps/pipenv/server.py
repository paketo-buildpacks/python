from flask import Flask, request
import subprocess
import gunicorn
import os, ssl

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
    port = int(os.getenv("PORT", 8080))

    context = ssl.SSLContext(ssl.PROTOCOL_TLS)
    context.verify_mode = ssl.CERT_REQUIRED
    context.load_default_certs()
    context.load_cert_chain(certfile='cert.pem', keyfile="key.pem")

    app.run(host='0.0.0.0', port=port, debug=True, ssl_context=context)