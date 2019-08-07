# app

This is Vue.js project(frontend).  
This sample is based on [`基礎から学ぶ Vue.js`](https://cr-vue.mio3io.com/).

# Setup
This package is build by `Makefile` in parent directory.  
Plese create configuration files for every environment before create image.

```
NODE_ENV = <YOUE ENVIRONMENT>
VUE_APP_API_ORIGIN = <URI>
```

These configuration file must follow the `.env.<environment name>` format (Note: in production environment, you must create a `.env` file).  
For example, if you want to run in `local` environment.  
You have to create `.env.local` file like this.

```.env.local
NODE_ENV='local'
VUE_APP_API_ORIGIN='http://localhost:38080/monsters'
```
