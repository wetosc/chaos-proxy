import requests
from pprint import pprint
from collections import Counter

API = 'http://localhost:8000/'
PROXY = 'http://localhost:8080/'

users_url = API + 'users'
users_url_proxy = PROXY + 'users'
wine_url = API + 'wines'
wine_url_proxy = PROXY + 'wines'


def check_json_response():
    pprint(requests.get(API+'users').json())
    print()
    pprint(requests.get(API+'wines').json())

def run_request(url, iterations):
    url = users_url
    l = []
    for i in range(iterations):
        res = requests.get(url)
        l += [res.status_code]
    print(dict(Counter(l)))

if __name__ == "__main__":
    iterations = 100
    check_json_response()
    
    run_request(users_url)
    run_request(users_url_proxy)

    run_request(wine_url)
    run_request(wine_url_proxy)
