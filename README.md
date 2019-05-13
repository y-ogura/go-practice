# go-practice

## Installation
```
mkdir $GOPATH/src/github.com/y-ogura
cd $GOPATH/src/gihub.com/y-ogura
git clone https://github.com/y-ogura/go-practice.git
```

## Overview
/answers/{sectionName}に各問題のファイルがあります。    
こちらのファイルに回答を記述してください。

下に記載されている回答はあくまで1つの記述方法です。

プルリク歓迎です。

## 回答確認

```
cd answers/{sectionName}
go test -run {functionName}
```

## 問題
<details><summary>問1~50</summary>

**問1**    
`a := map[string]string{"a": "a"}`と`b := map[string]string{"b": "b"}`をマージしたｃを出力してください
> eg map[a:a b:b]

<details><summary>回答</summary>

```go
a := map[string]string{"a": "a"}
b := map[string]string{"b": "b"}
c := map[string]string{}

for key, val := range a {
    c[key] = val
}
for key, val := range b {
    c[key] = val
}

return c
```

**問2**    
```go
arry := []string{"aa", "bb", "cc", "dd", "ee", "ff", "gg"}
```
のdd,ee,ffを新たな配列として返してください
> e.g `[dd ee ff]`

<details><summary>回答</summary>

```go
arry := []string{"aa", "bb", "cc", "dd", "ee", "ff", "gg"}
return arry[3:6]
```


</details>

</details>
