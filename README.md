# kthulu-packer

Simple packer that packs a jar with jre into one single executable binary.

## Getting started
provided you have a shadowJar or uberJar

execute the following command from the root of the packer folder.
```
./kthulu-pack.sh /path/to/app/app-all.jar my_executable_binary
```

## Limitions
currently only supports Mac OS and only tested on ARM architecture.
PR's welcome for Windows and Linux
