# factoid

Old style, Martini-based, factorial progra

## How to use it:

```
Mac $ cd /work
Mac $ git clone https://github.com/rm511130/factoid
Mac $ cd factoid
Mac $ docker build . -t factoid
Mac $ docker run -d --rm -p 3000:3000 factoid
```

You can now access the factorial program at:  http://127.0.0.1:3000/5

## I then Docker Pushed the image to the Docker Hub 

```
Mac $ docker login
Mac $ docker tag factoid rmeira/factoid
Mac $ docker push rmeira/factoid
```

And voilà. Now anyone can use the `rmeira/factoid` image
