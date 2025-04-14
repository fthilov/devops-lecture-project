# devops-lecture-project-2025

Unser Webshop für IT-Produkte bietet eine breite Auswahl an hochwertiger Hardware, leistungsstarker Software und praktischem Zubehör – ideal für Privatkunden, Unternehmen und Technikbegeisterte. Ob Sie nach aktuellen Laptops, leistungsstarken Prozessoren, professioneller Business-Software oder nützlichem IT-Zubehör suchen – bei uns finden Sie alles aus einer Hand, schnell und zuverlässig.

## Aufgabe 2

TODO: David

## Aufgabe 3

Der Webshop kann auch über Docker gestartet werden. Folgende Schritte sind dazu benötigt:

1. `docker pull ftvdb/devops-lecture-project`
2. `docker build . -t devops-lecture-project`
3. `docker run --name devops-lecture-project devops-lecture-project`

## Aufgabe 4

0. **Optional:** `minikube delete`
1. `minikube start`
2. `kubectl create namespace webshop-backend`
3. `kubectl apply -f manifests/deployment.yaml`
4. `kubectl apply -f manifests/services.yaml`
5. To check if the deployment was successfully applied, run: `kubectl get -n webshop-backend deployments`
6. To check if the services was successfully applied, run: `kubectl get -n webshop-backend services`

## Aufgabe 5

-   To run all microservices, run: `make all`
-   NOTE: valid service names are: auth, checkout or product
-   There are several other make commands :
    -   build-all
    -   build service=SERVICE_NAME
    -   test
    -   clean-all
    -   clean service=SERVICE_NAME
    -   start-all
    -   start service=SERVICE_NAME

## Aufgabe 6

1. `kubectl create namespace argocd`
2. `kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml`
3. `kubectl apply -f application.yaml`
4. To check, run: `kubectl get -n argocd deployments`

To check if the ArgoCD application is running, run the following commands:

0. **Optional:** `brew install argocd`
1. `kubectl patch svc argocd-server -n argocd -p '{"spec": {"type": "LoadBalancer"}}'`
2. `kubectl port-forward svc/argocd-server -n argocd 8080:443`
3. `argocd admin initial-password -n argocd`
4. Access localhost:8080 and login with **admin** as username and the provided password

## Aufgabe 7

TODO: David

## Aufgabe 8

1. kubectl apply -f monitoring/lgtm-app.yaml
2. Check on localhost:8080 to see if the lgtm stack is running
