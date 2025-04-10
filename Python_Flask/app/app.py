from flask import Flask, jsonify

# Initialize the Flask application
app = Flask(__name__)

# Define a route for the default endpoint
@app.route('/api', methods=['GET'])
def hello_world():
    return jsonify(message="Hello, World!")

@app.route('/api/user', methods=['GET'])
def get_user():
    user = {
        "id": 1,
        "name": "John Doe",
        "email": "johndoe@example.com"
    }
    return jsonify(user)

# Run the application
if __name__ == '__main__':
    app.run(debug=True)
