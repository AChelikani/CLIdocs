from flask import Flask

app = Flask(__name__)

@app.route('/python/<path:command>')
def index(command):
    if command == "print":
        return '{"Name": "print", "DataType": "function", "Declaration": "print [string]", "Description": "To print something", "ExampleUsage": "print"}'

if __name__ == '__main__':
    app.run(debug=True, port=8080)
