from flask import Flask, jsonify, request
from flask_cors import CORS
import mysql.connector
from mysql.connector import Error
import os
from dotenv import load_dotenv

# Load environment variables
load_dotenv()

app = Flask(__name__)
CORS(app)

# Create database connection
def create_connection():
    try:
        connection = mysql.connector.connect(
            host=os.getenv("DB_HOST"),
            user=os.getenv("DB_USER"),
            password=os.getenv("DB_PASSWORD"),
            database=os.getenv("DB_NAME")
        )
        return connection
    except Error as e:
        print("Error connecting to MySQL:", e)
        return None

# Route: Hello World
@app.route('/api', methods=['GET'])
def hello_world():
    return jsonify(message="Hello, World!")

# Route: Get all users
@app.route('/api/users', methods=['GET'])
def get_users():
    connection = create_connection()
    if connection is None:
        return jsonify(error="Database connection failed"), 500

    cursor = connection.cursor(dictionary=True)
    cursor.execute("SELECT * FROM users")
    users = cursor.fetchall()
    cursor.close()
    connection.close()
    return jsonify(users)

# Route: Add a user
@app.route('/api/addUser', methods=['POST'])
def add_user():
    data = request.get_json()
    username = data.get("username")
    name = data.get("name")

    if not username or not name:
        return jsonify(error="Missing username or name"), 400

    connection = create_connection()
    if connection is None:
        return jsonify(error="Database connection failed"), 500

    cursor = connection.cursor()
    query = "INSERT INTO users (username, name) VALUES (%s, %s)"
    cursor.execute(query, (username, name))
    connection.commit()
    cursor.close()
    connection.close()

    return jsonify(message="User added successfully", username=username), 201

# Run the server
if __name__ == '__main__':
    port = int(os.getenv("PORT", os.getenv("DB_PORT") or 5000))
    app.run(debug=True, port=port)
