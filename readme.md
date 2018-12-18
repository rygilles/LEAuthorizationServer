# LEAuthorizationServer

A simple HTTP Api for Let's Encrypt challenge validation with `http-01` type, written in GO.

**For tests purpose only, Api HTTP routes are not secured !**

## OpenAPI / Swagger specs ##

Check this [file](swagger-api/swagger.json) for API specifications.

## Windows helpers ##

Initially codded on Windows environment, a few `.bat` helpers files are located in `cmd/app` directory.

- `run_main.bat` Run the server.
- `serve_swagger.bat` Run a go-swagger server base on `swagger-api/swagger.json` specifications file.
- `generate_swagger_spec.bat` Re-generate `swagger-api/swagger.json` specifications file.
- `generate_swagger_spec.bat` Re-validate `swagger-api/swagger.json` specifications file.

## Routes ##

### Get all challenges ###

```bash
curl -X GET "http://domain.tld/challenges" \
     -H "accept: application/json" 
```

### Get a challenge ###

```bash
curl -X GET "http://domain.tld/challenge/{token}" \
     -H "accept: application/json" 
```

### Create/Update a challenge ###

This will add a challenge record.

```bash
curl -X POST "http://domain.tld/challenge" \
     -H "accept: application/json" \
     -H "Content-Type: application/json" \
     -d "{ \"token\": \"{token}\", \"value\": \"{value}\"}"
```

### ACME validation ### 

This route is for ACME validation purpose.

```bash
curl -X GET "http://domain.tld/.well-known/acme-challenge/{token}" \
     -H "accept: text/plain"
```

## *TODO* ##

- Builds for Linux and Windows.
- Handle certificates to serve with TLS HTTPS end-point.