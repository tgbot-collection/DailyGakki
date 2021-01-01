#!/usr/local/bin/python3
# coding: utf-8

# untitled - gakki_data_convertor.py
# 1/1/21 13:07
#

__author__ = "Benny <benny.think@gmail.com>"

import json

old = json.load(open("gakki.json"))

converted = {}

for sub in old:
    uid = sub["chat_id"]
    template = {
        "chat_id": uid,
        "time": [
            "18:11"
        ]
    }
    converted[uid] = template

print(json.dumps(converted, indent=4))

if input("Are you sure you want to dump data? yes/no") == "yes":
    with open("gakki.json", "w")as f:
        json.dump(converted, f, ensure_ascii=False, indent=4)
