import requests

r = requests.post("https://www.youtube.com/channel/UCDopCzR2JKYhpkwpdf7atUA")

print(r.text)