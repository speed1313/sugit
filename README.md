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
- git commit
```
$ go run main.go git commit message_hogehoge
// first commit object doesn't point parent commmit object

$ cat .sugit/objects/79/4fea08c1d1797fe85067460d9733c393890566
tree 0a448e7835621fa31ac9503d8cd31bacbc7e8f00
author speed1313
committer speed1313

message message_hogehoge%

$ go run main.go git commit hoge
// commit object points the parent commit object.
$ cat .sugit/objects/90/543b755444e4a23b0bbcff28733fa5d7ce90e0
tree 0a448e7835621fa31ac9503d8cd31bacbc7e8f00
parent 794fea08c1d1797fe85067460d9733c393890566
author speed1313
committer speed1313

message hoge%
```

- git log
```
$ go run main.go git log
90543b755444e4a23b0bbcff28733fa5d7ce90e0
* 90543b7 hoge
* 794fea0 message_hogehoge
```