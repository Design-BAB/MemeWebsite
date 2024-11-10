#!/bin/python
from flask import Flask, render_template
import requests
import json

app = Flask(__name__)

def get_meme():
    url = "https://meme-api.com/gimme/wholesomememes"
    response = requests.get(url)
    print(response.text)  # Print the response content
    meme_data = response.json() # this graps the JS info from the website
    return meme_data #this returns the JS stuffz to whomever calls this function

@app.route("/")
def index():
    try:
        meme_data = get_meme() #this calls the funtion
        meme_pic = meme_data["url"]  # Correct key to fetch the meme image from the URL variable
        return render_template("meme_index.html", meme_pic=meme_pic) #Starts the .HTML file and passes memes. Rest is Error handeling
    except json.JSONDecodeError as e:
        return f"JSON Decode Error: {e}"
    except Exception as e:
        return f"An error occurred: {e}"

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=80)

