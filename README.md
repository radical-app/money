# GoLang Money

GoLang library to make working with money safer, easier, and fun!

> "If I had a dime for every time I've seen someone use FLOAT to store currency, I'd have $999.997634" -- [Bill Karwin](https://twitter.com/billkarwin/status/347561901460447232)

In short: You shouldn't represent monetary values by a float. Wherever
you need to represent money, use this Money value object.

### This library doesn't have dependencies.

```go
import "github.com/radicalcompany/money"

fiveEur := money.EUR(500) // see list of all currencies
tenEur := fiveEur.Add(fiveEur)
zeroEur := tenEur.Subtract(tenEur)

zeroEur.IsZero() // true

fiveEur.IsEquals(zeroEur.Add(FiveEur)) // true
```

## .Forge() Money from Int, String and if you really-really-really need ... Float 

<details><summary><code>usd312, err := money.Forge(312, "USD")</code><br>more</summary>
<p>

```go
usd312 := money.USD(312)
usd312, err := money.Forge(312, "USD")

usd312 := money.Parse("USD 312")

usd312 := money.FloatUSD(3.12)
usd312, err := money.ForgeFloat(3.12, "USD")
```

</p>
</details>
        
   
### More .Parse() from string example [example at parse_test.go](./parse_test.go)

<details><summary><code>usd312, err := money.Parse("USD 312")</code><br>more</summary>
<p>

```go
usd312, err := money.Parse("USD 312")

eur312, err := money.ParseWithFallback("312", "EUR")
 
// this uses EUR because the string has it   
eur312, err := money.ParseWithFallback("EUR 312", "JPY")

// not suggested solution use ParseWithFallback if you have to deal with multiple currencies
money.DefaultCurrencyCode="JPY"
jpy312, err := money.Parse("312")
```

</p>
</details>
        
    
## Marshal/UnMarshal Custom Formatter [example at marshal_test.go](./marshal_test.go)

<details><summary><code>json.Marshal(money.EUR(123))</code><br>more</summary>
<p>

    json.Marshal(money.EUR(123))

will produce the simplified json for `money.DTO`:
    
    {"amount":123,"currency":"EUR"}

and the
 
    m := &Money{}
    json.Unmarshal([]byte('{"amount":123,"currency":"EUR"}'), m)

will produce the `money.EUR(123)`  

</p>
</details>

## .String()
 
    money.EUR(123).String() // "EUR 123"
 
## .Display() beautiful money depending based on locale [example at moneyfmt/moneyfmt_test.go](./moneyfmt/moneyfmt_test.go)
    
```go
import "github.com/radicalcompany/money/moneyfmt"
import "github.com/radicalcompany/money"

moneyfmt.Display(money.EUR(123400), "ru") // € 1 234
moneyfmt.Display(money.EUR(123456), "ru") // € 1 234,56

moneyfmt.Display(money.EUR(123456), "it") // € 1.234,56
moneyfmt.Display(money.EUR(123400), "it") // € 1.234

moneyfmt.Display(money.EUR(123456), "en") // € 1,234.56
moneyfmt.Display(money.EUR(123456), "jp") // € 1,234.56
moneyfmt.Display(money.EUR(123456), "zh") // € 1,234.56
```
      
## Mysql custom field support driver 

Is possible to use in mysql the field as `int` or `varchar` or if you really really need `decimal(13,4)`

    _, err := db.Exec("insert into blablabla int, string, decimal (?,?,?)",
      money.EUR(123).Int64(),
      money.EUR(123).String(),
      money.EUR(123).Float()
    )  

and the Scan during a select is auto-magically done: 
    
    rows.Scan(&moneyStoredAsInt64,
    		  &moneyStoredAsString,
    		  &moneyStoredAsFloat)
    		  
[Real example: driver_integration_test.go](./driver_integration_test.go)
   
## Limit

The biggest amount you can store in is `92.233.720.368.547.758,07` the `math.MaxInt64 / currency.cents`    
