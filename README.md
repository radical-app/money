# GoLang Money

GoLang library to make working with money safer, easier, and fun!

|![golang money](./assets/radical-golang-money.png "Money") | - No dependencies<br>- No anaemic model<br>- JSON formatter<br>- SQL driver<br>- Localized formatter(s)| 
|     :---:      | :--- |
|  <strong>Money Value Object</strong> | <blockquote> "If I had a dime for every time I've seen someone use FLOAT to store currency, I'd have $999.997634"</blockquote> -- [Bill Karwin](https://twitter.com/billkarwin/status/347561901460447232) <br><br>In short: You shouldn't represent monetary values by a float.<br>Wherever you need to represent money, use this Money value object.    |

```go
package main

import (
    "fmt"
    "github.com/radical-app/money"
)

func main() {
    fiveEur := money.EUR(500) // see list of all currencies
    tenEur, err := fiveEur.Add(fiveEur)
    fmt.Print(err)
    
    zeroEur, err := tenEur.Subtract(tenEur)
    fmt.Print(err)
    
    zeroEur.IsZero() // true
    
    anotherFiveEur,err := zeroEur.Add(fiveEur)
    fmt.Print(err)
    
    fiveEur.IsEquals(anotherFiveEur) // true

    fmt.Print(fiveEur.String()) // EUR 500 for beautiful formatter see below 
}
```

## .Forge and .Parse 


<details><summary>
Money from Int and only if you really-really-really need ...Forge from Float<br><br>

```go
usd312 := money.USD(312)
usd312, err := money.Forge(312, "USD")
```

</summary>
<p>

```go
usd312 := money.FloatUSD(3.12)
usd312, err := money.ForgeFloat(3.12, "USD")
```

</p>
</details>
        
   
### More .Parse() from string

<details><summary>
.Parse() is the opposite of .String()<br><br>

```go
usd312, err := money.Parse("USD 312")
usd312.String() // "USD 312"
```
</summary>
<p>

```go
eur312, err := money.ParseWithFallback("312", "EUR")
 
// this uses EUR because the string has it   
eur312, err := money.ParseWithFallback("EUR 312", "JPY")

// not suggested solution use ParseWithFallback if you have to deal with multiple currencies
money.DefaultCurrencyCode="JPY"
jpy312, err := money.Parse("312")
```

</p>
</details>

[example at parse_test.go](./parse_test.go)
        
## Marshal/UnMarshal Custom Formatter 

<details><summary>
Custom Marshaller from and to Json<br><br>

```go
json.Marshal(money.EUR(123))
```

</summary>
<p>

will produce the simplified json for `money.DTO`:

```json
{"amount":123,"currency":"EUR","symbol":"€","cents":100}
```

and

```go
m := &Money{}
json.Unmarshal([]byte('{"amount":123,"currency":"EUR","symbol":"€","cents":100}'), m)
```

will produce the `money.EUR(123)`  

</p>
</details>

[example at marshal_test.go](./marshal_test.go)

## .String()

```go
money.EUR(123).String() // "EUR 123"
```

## .Display() beautiful money depending based on locale 

```go
import "github.com/radical-app/money/moneyfmt"
import "github.com/radical-app/money"

moneyfmt.Display(money.EUR(123400), "ru") // € 1 234
moneyfmt.Display(money.EUR(123456), "ru") // € 1 234,56

moneyfmt.Display(money.EUR(123456), "it") // € 1.234,56
moneyfmt.Display(money.EUR(123400), "it") // € 1.234

moneyfmt.Display(money.EUR(123456), "en") // € 1,234.56
moneyfmt.Display(money.EUR(123456), "jp") // € 1,234.56
moneyfmt.Display(money.EUR(123456), "zh") // € 1,234.56
```

[example at moneyfmt/moneyfmt_test.go](./moneyfmt/moneyfmt_test.go)
    
     
## SQL custom field support driver 

Is possible to use in mysql the field as `int` or `varchar` or if you really really need `decimal(13,4)`

```go
_, err := db.Exec("insert into blablabla int, string, decimal (?,?,?)",
  money.EUR(123).Int64(),
  money.EUR(123).String(),
  money.EUR(123).Float()
)  
```

and the Scan during a select is auto-magically done: 

```go
rows.Scan(&moneyStoredAsInt64,
  &moneyStoredAsString,
  &moneyStoredAsFloat
)
```

[Real example: driver_integration_test.go](./driver_integration_test.go)
   
## Limit

The biggest amount you can store in is `92.233.720.368.547.758,07` the `math.MaxInt64 / currency.cents`    