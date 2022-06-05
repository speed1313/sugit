# sugit
sugit is a reimplementation of Git.


# How to use

- go run main.go git init
```
$ go run main.go git init
$ tree .sugit
.sugit
├── HEAD
├── objects
└── refs
    └── heads
```
- git add
```
$ go run main.go git add test.txt
.sugit
├── HEAD
├── objects
│   └── ff
│       └── 49900e6b0daa37cf14f334e2fd0ba33a4d81d9
└── refs
    └── heads
```
- git cat-file
```
$ go run main.go git cat-file ff49900e6b0daa37cf14f334e2fd0ba33a4d81d9
blob 18This is test file.% 
```
