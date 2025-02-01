import requests 

def send_http_request(url):
    try:
        response = requests.get(url)  
        if response.status_code == 200:
            return True
        else:
            return False
    except:
        return False

def main():
    if send_http_request("http://127.0.0.1:8000/"):
        print("3alamya")
    else:
        print("wihdanya")

if __name__ == "__main__":
    main()
