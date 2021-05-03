from bs4 import BeautifulSoup

VULNERABLE_ENDPOINT = 'http://127.0.0.1:4000/secret'


def generate_api_calls(base_url, param):
    return [(VULNERABLE_ENDPOINT, {'name': param}, 'post')]


def parse_response(response_text):
    soup = BeautifulSoup(response_text, 'html.parser')
    alert = soup.find("div", class_="alert alert-success col-12")

    return alert.contents[0].strip().split(' ')[3]