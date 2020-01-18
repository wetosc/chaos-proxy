from flask import Flask, jsonify
from flask_cors import CORS

app = Flask(__name__)

CORS(app)

users = [dict(id=1,name='User 1'), dict(id=2,name='User 2')]
wines = [
    dict(id=1,name='Chateau Rieussec', rating=9.5), 
    dict(id=2,name='Domaine Bousques', rating=8.5), 
    dict(id=2,name='Vina Garces Silva', rating=7.8)
    ]

@app.route('/users', methods=['GET'])
def get_users():
    return jsonify(dict(data=users)), 200

@app.route('/wines', methods=['GET'])
def get_wines():
    return jsonify(dict(data=wines)), 200

if __name__ == '__main__':
    app.run(port=3000)
