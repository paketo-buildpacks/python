import http.server, ssl
from http import HTTPStatus

class Handler(http.server.SimpleHTTPRequestHandler):
    def do_GET(self):
        self.send_response(HTTPStatus.OK)
        self.end_headers()
        self.wfile.write(b'Powered by Paketo Buildpacks')

# SSL
context = ssl.SSLContext(ssl.PROTOCOL_TLS)
context.verify_mode = ssl.CERT_REQUIRED
context.load_default_certs(ssl.Purpose.CLIENT_AUTH)
context.load_cert_chain(certfile='cert.pem', keyfile="key.pem")
# Wrap http
server_address = ('0.0.0.0', 8080)
httpd = http.server.HTTPServer(server_address, Handler)
httpd.socket = context.wrap_socket(httpd.socket, server_side=True)

httpd.serve_forever()