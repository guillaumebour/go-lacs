# go-lacs

A simple **La**Tex **C**ompilation **S**ervice, exposed over a web API.
This project can be used as such as long as there is a LaTeX distribution 
installed on the machine and that the `pdflatex` binary is available in the PATH.

The envisioned used however, is in a docker container, to have *LaTeX as a Service*. 

**Important**: Do not use with untrusted inputs.

## Requirements

If run without Docker, a LaTeX distribution with `pdflatex` in the PATH.

## Run with Docker

You first need to build the image, from the project folder:

```
docker build -t lacs:latest .
```

NB: depending on your machine and network, this might take some time.

Once you have the image, you can start the service:

```
docker run -p 8081:8081 lacs:latest
```
