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
fmt.Println(c)
```

</details>

</details>
