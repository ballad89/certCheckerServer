**CertCheckerServer**

This is a real time web application server that monitors the status of certificates across a number of servers

***How to run it***
The application runs in a docker container

1. Clone this repository locally
2. Build the docker image ```docker build -t certchecker . ```
3. Run it ```docker run -p 3000:3000 -d certchecker```
4. Point your browser to ```http://localhost:3000/```
5. Simulate the posting of information to the application
```
curl -X POST -H "Cache-Control: no-cache" -H "Content-Type: application/x-www-form-urlencoded" -d 'cname=cert01.com&SigningAlgorithm=sha256&issuer=godaddy¬after=25-02-2018¬before=25-02-2014&servername=server-pc' "http://localhost:3000/certificates"
```
This should update the browser in realtime