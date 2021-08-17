types_of_people = 10
x = f"There are {types_of_people} types of people."

binary = "binary"
do_not = "don't"
y = f"Those who know {binary} and those who {do_not}."

print(x)
print(y)

print(f"I said: {x}")
print(f"I also said: '{y}'")

joke_evaluation = "Isn't that joke so funny?! {}"

hilarious = False

print(joke_evaluation.format(hilarious))


w = "This is the left side of..."
e = "a string with a right side."

print(w + e)

# python way to send formatted output:
# Ref: https://pyformat.info/
print("current temp: {:>10f}".format(123))

