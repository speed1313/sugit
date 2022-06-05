package cmd


// git commit creates commit object and HEAD file points it.
// commit object contains

/* example of commit object
tree 83de101e54e9d5b9ce77d0cf0870b93f4af7358b
author speed1313
committer speed1313

add tmp.txt
*/
/* example of HEAD and heads/main
$ cat .git/refs/heads/main
27ad1210df49123d0aac5caa8eea9237c24e3592

$ cat .git/HEAD
ref: refs/heads/main
*/

func Git_commit(message []string){

}