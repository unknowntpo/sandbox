# Example Package

This is a simple example package. You can use
[Github-flavored Markdown](https://guides.github.com/features/mastering-markdown/)
to write your content.

## Steps
```
# New venv
python3 -m venv ~/repo/unknowntpo/sandbox/python/packaging_tutorial
```

```
# Activate new venv
# We are at ./package_tutorial
source ./bin/activate
```

```
# [Generating distribution archives](https://packaging.python.org/en/latest/tutorials/packaging-projects/#generating-distribution-archives)
python3 -m pip install --upgrade build
```

```
# Now run this command from the same directory where pyproject.toml is located:
python3 -m build
```

```
# [Uploading the distribution archives](https://packaging.python.org/en/latest/tutorials/packaging-projects/#uploading-the-distribution-archives)
python3 -m pip install --upgrade twine
python3 -m twine upload --repository testpypi dist/*
```

