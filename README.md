@UTF-8

Xamboo for GO v1
=============================

This is the main project environment to install a working xamboo example site.


INSTALATION AND COMPILATION
=============================

Create a new directory for Xamboo, for instance /home/sites/xamboo

$ mkdir /home/sites/xamboo

Go in xamboo directory
create a git:
$ git init

Pull the last verion of Xamboo

$ git pull https://github.com/webability-go/xamboo.git

You need to edit each .json files to adapt it to your own IP and ports

Set the Listeners IP and Port so the service will work on your machine.
Set the Hosts domains so the service will resolve. Do not forget to add those domains to your DNS too.

Run the xamboo with master and examples

$ start.sh


To build your own server:
Edit start.sh, json config files and change the config file path.

You can copy the example directory and change anything you need.

The master site is not necessary to make the CMS work. It's a helpfull tool to configure and install anything easier.
Install the master site and install contexts with XModules for any site you need.

You can compile xamboo to an executable with
go build xamboo.go
You do not need to recompile any app and page any time you restart the server. The system compile things as needed. You may recompile anything before launching on a production site, for velocity.
You will need the original code so the compiler is able to compile pages and libraries without problem at anytime. It will use the go.mod and go.sum retrieved with the Xamboo.

You may attach the xamboo as a OS/service, calling the start.sh


Version Changes Control
=======================

V0.0.3 - 2021-02-12
-----------------------
- Adaptation to Xamboo v1.5.0 not compatible with previous versions

V0.0.2 - 2020-07-29
-----------------------
- Correct some config errors
- Adjust go.mod with last version

V0.0.1 - 2018-11-06
-----------------------
- First commit
