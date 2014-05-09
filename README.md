gvolinfojson
============

A utility to convert xml output of gluster volume info to json.

## Install

Download binary from [here](https://github.com/aravindavk/gvolinfojson/releases/download/1.0/gvolinfojson) and copy to `/usr/local/bin`(or any other directory, which is available in PATH).

    wget https://github.com/aravindavk/gvolinfojson/releases/download/1.0/gvolinfojson
    sudo cp gvolinfojson /usr/local/bin/
    chmod +x /usr/local/bin/gvolinfojson

If you have golang installed, then(make sure $GOPATH/bin is available in PATH)

    go get github.com/aravindavk/gvolinfojson

## Usage

Run gluster volume info command and pipe its output to gvolinfojson.

    sudo gluster volume info --xml | gvolinfojson

If you need to print pretty JSON output, then

    sudo gluster volume info --xml | gvolinfojson --pretty

## Source

    git clone https://github.com/aravindavk/gvolinfojson.git

## Blogs

1. [gvolinfojson - A utility to convert xml output of gluster volume info to json](http://aravindavk.in/blog/gvolinfojson)
