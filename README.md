# pwgen
Password Generator


# Container image

To build the container image and tag with a version number as well as the de facto standard `latest` tag:

```
version=0.1
podman build -f Containerfile --label pwgen=1 --build-arg VERSION=$version -t ploregistry01.westernpower.co.uk:5050/wpd/pwgen:$version .
podman build -f Containerfile --label pwgen=1 --build-arg VERSION=latest -t ploregistry01.westernpower.co.uk:5050/wpd/pwgen:latest .
```

If you want to push those to ploregistry01 (you probably do eventually):

```
podman push ploregistry01.westernpower.co.uk:5050/wpd/pwgen:${version}
podman push ploregistry01.westernpower.co.uk:5050/wpd/pwgen:latest
```

