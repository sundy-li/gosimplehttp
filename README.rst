gosimplehttp
=========


What's it
-------------
Gosimplehttp is a Golang implementation to replace "python -m SimpleHTTPServer"



How to install
-------------
.. code-block:: shell

    go get -u github.com/sundy-li/gosimplehttp

Usage
-------------
.. code-block:: shell

    ➜  ~   gosimplehttp -p 8000
            2017/11/25 gosimplehttp.go:25: Serving HTTP on 0.0.0.0 port 8000 ...


Flag
------------
.. code-block:: shell

    ➜  ~   gosimplehttp -h
            Usage of ./gosimplehttp:
            -d string
                  directory of the files (default ".")
            -debug
                  show debug log or not
            -p string
                  port to bind (default "8000")
            -proxy string
                  reverse proxy


