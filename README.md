# Crashlytics from Fabric upload DSYM tool

This is a simple command line tool made with Golang that upload any zipped DSYM to **Crashlytics** from **Fabric.io**

# Usage

```
fabric-upload-dsym --bundleid=[YOUR-APP-BUNDLE] --fabricapikey=[YOUR-FABRIC-API-KEY] --file=[ZIPPED-DSYM-FILE]
```

**Replace in command with your info, example:**  

- [YOUR-APP-BUNDLE] = com.prsolucoes.myapp
- [YOUR-FABRIC-API-KEY] = 12xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxbf
- [ZIPPED-DSYM-FILE] = /tmp/yourapp/dsym.zip

# Prebuilt

You can find all prebuilt files by operacional system inside folder **"dist"**.  

With this file you don't need compile, only download and use binary.  

# Build and install

```
make deps
make install
```

# Available commands

To see all available commands on Makefile, type:
> make  

# Help

```
fabric-upload-dsym -help
```

# Support with donation
[![Support with donation](http://donation.pcoutinho.com/images/donate-button.png)](http://donation.pcoutinho.com/)