# tf-release-info

Write release-info JSON for Terraform provider plugins.

## Usage

```bash
NAME:
   tf-release-info - Write terraform plugin release-info JSON

USAGE:
   tf-release-info PLUGIN-VERSION TERRAFORM-VERSION JSON-FILE-PATH

ARGS:
   PLUGIN-VERSION    : terraform plugin version via "v1.0.0"
   TERRAFORM-VERSION : terraform version via "v1.0.0"
   JSON-FILE-PATH    : file path of target JSON

COPYRIGHT:
   Copyright (C) 2017 Kazumichi Yamamoto.
```

## Usage(Docker)

```bash
docker run -it --rm -v $PWD:/workdir sacloud/tf-release-info v1.0.0 v0.11.0 releases/versions.json
```


## License

 `tf-release-info` Copyright (C) 2017 Kazumichi Yamamoto.

  This project is published under [Apache 2.0 License](LICENSE.txt).
  
## Author

  * Kazumichi Yamamoto ([@yamamoto-febc](https://github.com/yamamoto-febc))
