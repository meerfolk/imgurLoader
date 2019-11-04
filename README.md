# imgurLoader
small command to load images to imgur.com

## Installation
For installation run command ```go build``` inside the directory. You should have golang installed on your computer

## Configuration
After first run command will create ```.imgurLoader``` file inside your home directory. This file contains several fields:
- path - path where application will look for changes
- file - regexp for filter unnecessary changes

## Default Configuration
```
  {
    path: your_home_directory/Pictures,
    file: /.png
  }
```

## Beware
By default command will send all ```*.png``` files (new or changed) in default directory to imgur.com