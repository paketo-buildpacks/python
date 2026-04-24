import os
import ssl
import sys

from flask import Flask

app = Flask(__name__)


@app.route('/')
def root():
    python_version = sys.version
    return "Hello, world!\nUsing python: " + python_version + "\n"


if __name__ == '__main__':
    # Get port from environment variable or choose 9099 as local default
    port = int(os.getenv("PORT", 8080))

    context = ssl.SSLContext(ssl.PROTOCOL_TLS)
    context.verify_mode = ssl.CERT_REQUIRED
    context.load_default_certs()
    context.load_cert_chain(certfile='cert.pem', keyfile="key.pem")

    # Run the app, listening on all IPs with our chosen port number
    app.run(host='0.0.0.0', port=port, debug=True, ssl_context=context)
