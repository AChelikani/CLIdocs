import sqlite3

def populate_docs():
    conn = sqlite3.connect('documentation.db')
    c = conn.cursor()

    c.execute('DROP TABLE IF EXISTS docs')

    c.execute('''CREATE TABLE docs
                 (q_id INTEGER PRIMARY KEY NOT NULL, language text, name text, dataType text, declaration text, description text, exampleUsage text)''')

    language = "python"
    name = "abs"
    dataType = "function"
    declaration = "abs(number)"
    description = "Return the absolute value of a number. The argument may be a plain or long integer or a floating point number. If the argument is a complex number, its magnitude is returned."
    exampleUsage = "abs(-1) -> 1"

    c.execute('INSERT INTO docs(language, name, dataType, declaration, description, exampleUsage) VALUES (?, ?, ?, ?, ?, ?)',
             (language, name, dataType, declaration, description, exampleUsage))

    conn.commit()
    conn.close()

populate_docs()
