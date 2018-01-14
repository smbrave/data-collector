## 身份证信息解析

```bash
    info, err = idcard.Parse("xxxxx")
	if err !=nil {
		panic(err)
	}
	fmt.Printf("%+v\n", info)

```

```bash
&{Area:天津市蓟县 Sex:男 Birth:1987-03-26 IdNo:xxxxx}
```