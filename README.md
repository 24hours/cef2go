Chrome - Go binding to CEF [![Build Status](https://travis-ci.org/24hours/chrome.svg?branch=master)](https://travis-ci.org/24hours/chrome) [![Coverage Status](https://coveralls.io/repos/24hours/chrome/badge.svg)](https://coveralls.io/r/24hours/chrome)
======
Table of contents:
 * [Introduction](#introduction)
 * [Compatibility](#compatibility)
 * [Getting started on Windows `not yet supported`](#getting-started-on-windows)
 * [Getting started on Linux](#getting-started-on-linux)
 * [Getting started on Mac OS X](#getting-started-on-mac-os-x)


Introduction
------------

Chrome is a project forked from CEF2go by [Czarek Tomczak](http://www.linkedin.com/in/czarektomczak) and [FromKeith](https://github.com/fromkeith). 
This fork aim to have similar role with [PhantomJs](http://phantomjs.org/) 

Currently the Chrome only allow simple programming such as opening a browser. 

CEF2go is licensed under the BSD 3-clause license, see the LICENSE
file.

Compatibility
-------------
Supported platforms: Windows, Linux, Mac OSX.

Build
------------
Getting started on Windows
--------------------------


Getting started on Linux
------------------------
1. These instructions work fine with Ubuntu 12.04 64-bit. 
   On Ubuntu 13/14 libudev.so.0 is missing and it is required to 
   create a symbolic link to libudev.so.1. For example on 
   Ubuntu 14.04 64-bit run this command: 
  `cd /lib/x86_64-linux-gnu/ && sudo ln -sf libudev.so.1 libudev.so.0`.

2. Install CEF dependencies:  
   `sudo apt-get install libatk1.0-0 libc6 libasound2 libcairo2 libcap2 libcups2 libexpat1
  libexif12 libfontconfig1 libfreetype6 libglib2.0-0 libgnome-keyring0 libgtk2.0-0
  libpam0g libpango1.0-0 libpci3 libpcre3 libpixman-1-0 libpng12-0 libspeechd2 libstdc++6
  libsqlite3-0 libx11-6 libxau6 libxcb1 libxcomposite1 libxcursor1 libxdamage1 libxdmcp6
  libxext6 libxfixes3 libxi6 libxinerama1 libxrandr2 libxrender1 libxtst6 zlib1g
  libpulse0 libbz2-1.0 libnss3-dev libgconf2-dev`

3. Download CEF 3 Branch 2171 binaries:
   [release_linux.zip](https://github.com/24hours/chrome/releases/download/v0.13/Release_linux.zip)  
   Extract the file to `GOPATH/release/*`  

4. Copy `GOPATH/github.com/24hours/chrome/doc/Makefile` to `GOPATH/Makefile`  
   Copy `GOPATH/github.com/24hours/chrome/doc/minimum/main_linux.go` to `GOPATH/main_linux.go`  
   
5. Run "make" command.


Getting started on Mac OS X
---------------------------
1. These instructions work fine with OS X 10.10 Yosemite.

2. Install command line tools (make is required) from:  
   https://developer.apple.com/downloads/  
   (In my case command line tools for Mountain Yosemite from 2 Dec 2014)

3. Install XCode (gcc that comes with XCode is required).   
   My version is Version 6.1.1  

5. Download CEF 3 Branch 2171 binaries:
   [release_mac.zip](https://github.com/24hours/chrome/releases/download/v0.13/release_mac.zip)   
   Extract the file to `GOPATH/release/*`  

6. Copy `GOPATH/github.com/24hours/chrome/doc/Makefile` to `GOPATH/Makefile`  
   Copy `GOPATH/github.com/24hours/chrome/doc/minimum/main_mac.go` to `GOPATH/main_mac.go`  
   
7. Run "make" command.