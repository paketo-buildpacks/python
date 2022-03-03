import os

from flask import Flask
app = Flask(__name__)

@app.route('/')
def hello_world():
    return 'Hello, World!'

if __name__ == "__main__":
    app.run()

def run():
    port = int(os.getenv("PORT"))
    app.run(host='0.0.0.0', port=port)
