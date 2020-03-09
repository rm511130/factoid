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

## Using Factoid with Kubernetes

```
Mac $ kubectl cluster-info
Mac $ kubectl run factoid --image=rmeira/factoid
Mac $ kubectl expose deployment factoid --target-port=3000 --port=3000 --type=LoadBalancer

Mac $ kubectl get service factoid

NAME      TYPE           CLUSTER-IP      EXTERNAL-IP     PORT(S)          AGE
factoid   LoadBalancer   10.100.200.43   34.74.183.205   3000:31549/TCP   3m8s
```

```
Mac $ curl http://34.74.183.205:3000/5
[Factoid] Calculating Factorials: 5! = 120
```

```
Mac $ kubectl delete service factoid
service "factoid" deleted

Mac $ kubectl delete deployment factoid
deployment.extensions "factoid" deleted
```
