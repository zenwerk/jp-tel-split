# jptel 電話番号ユーティリティ

jptel は日本の電話番号を市外局番・市内局番・加入者番号に分割して返します。

This package is utility for japaneses telephone number.

## インストール

```
$ go get github.com/zenwerk/jptel
```

## 使い方

```go
package main

import (
	"fmt"

	"github.com/zenwerk/jptel"
)

func main() {
	result, err := jptel.Split("0312345678")
	if err != nil {
		panic(err)
	}
	fmt.Println(result.AreaCode)       // 03
	fmt.Println(result.CityCode)       // 1234
	fmt.Println(result.SubscriberCode) // 5678

	result, err = jptel.Split("０９０９８７６５４３２")
	if err != nil {
		panic(err)
	}
	fmt.Println(result.AreaCode)       // 090
	fmt.Println(result.CityCode)       // 9876
	fmt.Println(result.SubscriberCode) // 5432
	
	fmt.Println(jptel.Normalize("0123456789"))           // 0123-45-6789
	fmt.Println(jptel.Normalize("０１ー２ー３４５６７８９")) // 0123-45-6789
	
	err = jptel.Validate("090-1234-5678") // nil
	if err != nil {
		panic(err)
	}
	
	err = jptel.Validate("0-90-12345678") // error
	if err != nil {
		panic(err)
	}
}
```

## その他
固定電話の市外局番データは[総務省のサイト](http://www.soumu.go.jp/main_sosiki/joho_tsusin/top/tel_number/number_shitei.html#kotei-denwa)からダウンロードできるExcelから生成しています。
再生成する場合は以下の手順で行って下さい。

```
$ pip install -r freeze.txt
$ python generate_master_data.py
```

### thanks
- https://gist.github.com/kennyj/4966002