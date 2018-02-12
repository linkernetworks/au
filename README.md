# AURORA Console tool

Note: For localhost support only.

## Install Aurora commands:

```
go install bitbucket.org/linkernetworks/aurora/src/sdk/cmd/au
```

(you need have bitbucket permission)

## Support commamds:

- Get all worksapce: `au get ws`
    - `-p`: Page number (default: 1)
    - `-f`: Filter string

- Create workspace: `au create ws datasetID1, datasetID2 ...`
    - `-t`: Specific workspace type (defult: general)

- Upload file to workspace: `au upload WORKSPACE_ID FILE_PATH1, FILE_PATH2 ....`
    - `FILE_PATH` could be file or path.

- Get application version number: `au version`

#### Other options:

- Enablbe SSL connection (default is disable) 
    - `-s`: "true"
