from flask import Flask
import sqlite3
import json

app = Flask(__name__)

def get_info(lang, query):
    conn = sqlite3.connect('documentation.db')
    c = conn.cursor()

    c.execute('SELECT * FROM docs WHERE language=? AND name=?', (lang, query))
    res = c.fetchone()
    return res

    conn.close()

def to_json_string(query_result):
    print query_result
    fieldNames = ["Name", "DataType", "Declaration", "Description", "ExampleUsage"]
    d = {}
    for x in range(len(fieldNames)):
        d[fieldNames[x]] = query_result[x+2]
    return json.dumps(d)

@app.route('/python/<path:command>')
def index(command):
    resp = get_info("python", command)
    if (resp):
        return to_json_string(resp)
if __name__ == '__main__':
    app.run(debug=True, port=8080)
