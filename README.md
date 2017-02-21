golang utility to marshal/unmarshal url.Values to/from struct

```go
import "query"

...

type TestType struct {
    GUID      string
    Name      string `query:"nom"`
    lowercase string `query:"lc"`
    Number    int    `query:"mispar"`
}

   
obj := TestType{"abc", "def", "hello", 31}
    
// v is url.Values
v := query.Marshal(obj)
obj2 := TestType{}
err := query.Unmarshal(v, &obj2)
```
