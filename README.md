# How to run?


  Create **config.yaml** with help of the given example,then to run
```  
  go run .
```

By default the application looks at the current directory for the configuration file, but any file can be passed in by using the **config-path** flag
```  
  go run .  --config-path=/path/to/conig.yaml
```

For social login, the application has to be registered in the corresponding provider site. Only Github has been implemented for the time being
