from flask import Flask, render_template
import random

app = Flask(__name__)

# list of cat images
images = [
        "https://www.google.com/imgres?imgurl=https%3A%2F%2Fmedia2.giphy.com%2Fmedia%2Fl3q2RJBdaqJY2SV3O%2Fgiphy.gif&imgrefurl=https%3A%2F%2Fgiphy.com%2Fgifs%2Fjerseydemic-l3q2RJBdaqJY2SV3O&tbnid=KdSmtX9vVpmnlM&vet=12ahUKEwjIoNjw6KjsAhXWAaYKHYrwBBcQMygNegUIARClAg..i&docid=YAoBa4x3mct13M&w=400&h=332&q=cat%20gif&ved=2ahUKEwjIoNjw6KjsAhXWAaYKHYrwBBcQMygNegUIARClAg",
        "https://www.google.com/imgres?imgurl=https%3A%2F%2Fmedia4.giphy.com%2Fmedia%2Ff7kUyaVVXn4JqKovE1%2Fgiphy.gif&imgrefurl=https%3A%2F%2Fgiphy.com%2Fgifs%2Ftiktok-aww-f7kUyaVVXn4JqKovE1&tbnid=AWhc_Ta6p3jNVM&vet=12ahUKEwjIoNjw6KjsAhXWAaYKHYrwBBcQMygZegUIARDNAg..i&docid=y_LyaTp5NKI2QM&w=475&h=329&q=cat%20gif&ved=2ahUKEwjIoNjw6KjsAhXWAaYKHYrwBBcQMygZegUIARDNAg",
        "https://www.google.com/imgres?imgurl=https%3A%2F%2Fmedia3.giphy.com%2Fmedia%2FjUKBVRKJwoB9fC8g8p%2Fgiphy.gif&imgrefurl=https%3A%2F%2Fgiphy.com%2Fgifs%2Ftiktok-funny-scaredy-cat-scared-jUKBVRKJwoB9fC8g8p&tbnid=MvECeG9IShAU0M&vet=12ahUKEwjIoNjw6KjsAhXWAaYKHYrwBBcQMygtegUIARCJAw..i&docid=T4Dv6wVdVl3yaM&w=308&h=368&q=cat%20gif&ved=2ahUKEwjIoNjw6KjsAhXWAaYKHYrwBBcQMygtegUIARCJAw",
        "https://www.google.com/imgres?imgurl=https%3A%2F%2Fmedia4.giphy.com%2Fmedia%2FWTL02R1L7YCGUEunFy%2Fgiphy.gif&imgrefurl=https%3A%2F%2Fgiphy.com%2Fexplore%2Fwoman-with-cat&tbnid=yHGrn50dGYjmeM&vet=12ahUKEwjIoNjw6KjsAhXWAaYKHYrwBBcQMyg2egQIARBf..i&docid=O1Uy2sUlUX-c1M&w=498&h=498&q=cat%20gif&ved=2ahUKEwjIoNjw6KjsAhXWAaYKHYrwBBcQMyg2egQIARBf"
    ]

@app.route('/')
def index():
    url = random.choice(images)
    return render_template('index.html', url=url)

if __name__ == "__main__":
    app.run(host="0.0.0.0")
